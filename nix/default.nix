{pkgs}: let
  callPackage = pkgs.lib.callPackageWith (pkgs // packages);
  packages = rec {
    hanamaru-go = callPackage ./hanamaru-go.nix {};
    hanamaru-lib = callPackage ./hanamaru-lib.nix {};

    hanamaru = callPackage ./hanamaru.nix {};

    dockerImages.latest = callPackage ./docker.nix {};

    default = hanamaru;
  };
in
  packages
