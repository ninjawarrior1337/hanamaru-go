use std::{
    ffi::{CStr, CString},
    os::raw::c_void,
};

use libc::{c_char, free, uintptr_t};
use uiua::{
    format::{format_str, FormatConfig}, Uiua, UiuaError, UiuaErrorKind,
};

extern "C" {
    fn goSendMessage(ctx: uintptr_t, msg: *const c_char);
    fn goReferencedMessage(ctx: uintptr_t) -> *const c_char;
    fn goAssignRole(ctx: uintptr_t, target_id: *const c_char, role_id: *const c_char);
}

fn referenced_message(ctx: uintptr_t) -> Option<String> {
    unsafe {
        let m = goReferencedMessage(ctx);
        if m.is_null() {
            return None;
        }
        let s = CStr::from_ptr(m).to_str().unwrap().to_owned();
        free(m as *mut c_void);
        return Some(s);
    }
}

fn assign_role(ctx: uintptr_t, target_id: &str, role_id: &str) {
    let target_c = CString::new(target_id).unwrap();
    let role_c = CString::new(role_id).unwrap();
    unsafe { goAssignRole(ctx, target_c.into_raw(), role_c.into_raw()) }
}

fn send_message(ctx: uintptr_t, msg: &str) {
    unsafe {
        let s = CString::new(msg).unwrap();
        goSendMessage(ctx, s.into_raw())
    }
}

#[no_mangle]
pub extern "C" fn run_uiua(ctx: uintptr_t, input_ptr: *const c_char) -> *const c_char {
    let mut uiua = Uiua::with_safe_sys();

    let input_cstr = unsafe { CStr::from_ptr(input_ptr) };
    let input = input_cstr.to_str().unwrap();

    let result = uiua.compile_run(|comp| {
        comp.print_diagnostics(true);

        // Pushes the content of the referenced discord message onto the stack
        comp.create_bind_function("&dr", (0, 1), move |ua| {
            let ref_msg = referenced_message(ctx);
            match ref_msg {
                Some(m) => {
                    ua.push(m);
                    Ok(())
                }
                None => Err(UiuaError::from(UiuaErrorKind::Run(
                    ua.span().sp("no referenced message".to_owned()),
                    ua.inputs().clone().into(),
                ))),
            }
        })
        .unwrap();
        // Pops the top of the stack and sends a discord message
        comp.create_bind_function("&ds", (1, 0), move |ua| {
            let s = ua.pop_string()?;
            send_message(ctx, &s);
            Ok(())
        })
        .unwrap();

        // Pops a target and a role id off the stack and tells discord to give role to target
        comp.create_bind_function("&dar", (2, 0), move |ua| {
            let target_id = ua.pop_string()?;
            let role_id = ua.pop_string()?;
            assign_role(ctx, &target_id, &role_id);
            Ok(())
        })
        .unwrap();

        comp.load_str(input)
    });

    let result_str = if let Err(e) = result {
        e.to_string()
    } else {
        let config = FormatConfig::default().with_trailing_newline(false);
        let mut formatted = "Formatted: ".to_owned() + &format_str(input, &config).unwrap().output;
        let v: Vec<String> = uiua.stack().iter().map(|e| e.to_string()).collect();
        formatted += "\n";
        formatted += &v.join("\n");
        formatted
    };

    return CString::new(result_str).unwrap().into_raw();
}

#[no_mangle]
pub unsafe extern "C" fn drop_string(ptr: *mut c_char) {
    let _ = CString::from_raw(ptr);
}
