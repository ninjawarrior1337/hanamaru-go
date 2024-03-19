use std::{ffi::{CStr, CString}, os::raw::c_void, ptr::null, result};

use libc::{c_char, free, uintptr_t};
use uiua::{Compiler, Uiua, UiuaError, UiuaResult};

extern "C" {
    fn goSendMessage(ctx: uintptr_t, msg: *const c_char);
    fn goReferencedMessage(ctx: uintptr_t) -> *const c_char;
}

fn referenced_message(ctx: uintptr_t) -> UiuaResult<String> {
    unsafe {
        let m = goReferencedMessage(ctx);
        if m.is_null() {
            return Err(UiuaError::Panic("no referenced message".to_owned()));
        }
        let s = CStr::from_ptr(m).to_str().unwrap().to_owned();
        free(m as *mut c_void);
        return Ok(s)
    }
}

fn send_message(ctx: uintptr_t, msg: &str) {
    unsafe {
        let s = CString::new(msg).unwrap();
        goSendMessage(ctx, s.into_raw())
    }
}

#[no_mangle]
pub extern "C" fn run_uiua(ctx: uintptr_t, input: *const c_char) {
    let mut uiua = Uiua::with_safe_sys();

    let input_cstr = unsafe { CStr::from_ptr(input) };

    let result = uiua.compile_run(|comp| {
        comp.print_diagnostics(true);

        comp.create_bind_function("&dr", (0, 1), move |ua| {
            ua.push(referenced_message(ctx)?);
            Ok(())
        }).unwrap();
        comp.create_bind_function("&ds", (1, 0), move |ua| {
            let s = ua.pop_string()?;
            send_message(ctx, &s);
            Ok(())
        }).unwrap();

        comp.load_str(input_cstr.to_str().unwrap())
    });

    if let Err(e) = result {
        send_message(ctx, &e.to_string())
    } else {
        let v: Vec<String> = uiua.stack().iter().map(|e| e.to_string()).collect();
        send_message(ctx, &v.join("\n"))
    }
}

#[no_mangle]
pub unsafe extern "C" fn drop_string(ptr: *mut c_char) {
    let _ = CString::from_raw(ptr);
}