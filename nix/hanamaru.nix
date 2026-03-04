{
  writeShellScriptBin,
  hanamaru-go,
}: (writeShellScriptBin "hanamaru" ''
  exec ${hanamaru-go}/bin/hanamaru-go "$@"
'')
