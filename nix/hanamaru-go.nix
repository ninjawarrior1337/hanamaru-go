{
  buildGoModule,
  go_1_27,
  hanamaru-lib,
}:
(buildGoModule.override {go = go_1_27;}) {
  pname = "hanamaru-go";
  version = builtins.readFile ../version/version.txt;
  src = ../.;

  ldflags = [
    "-s"
    "-w"
  ];

  tags = [
    "jp"
    "ij"
  ];

  preBuild = [
    "go generate"
  ];

  buildInputs = [
    hanamaru-lib
  ];

  doCheck = false;

  vendorHash = "sha256-NMb6V5d45CZz+HSJioNjmXb/J0rJfYfSqt9YV8Hmmmk=";
  proxyVendor = true;
}
