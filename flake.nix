{
  description = "A simple Go package";

  # Nixpkgs / NixOS version to use.
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    let

      # to work with older version of flakes
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";

      # Generate a user-friendly version number.
      version = builtins.substring 0 8 lastModifiedDate;

    in
    flake-utils.lib.eachDefaultSystem (system: 
    let 
      pkgs = import nixpkgs {
        inherit system;
      };
    in {

      # Provide some binary packages for selected system types.
      packages = 
        rec {
          hanamaru-go = pkgs.buildGoModule {
            pname = "hanamaru-go";
            inherit version;
            src = ./.;

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

            vendorHash = "sha256-zmVCsmy/sDo9WCI3n/w8dGeQauDGK96gRhDjXtGZtXo=";
          };

          hanamaru-lib = pkgs.rustPlatform.buildRustPackage {
            pname = "hanamaru-lib";
            inherit version;
            src = ./lib;
            cargoHash = "sha256-U13p/7temhpIVKcadfzKH0YUMChi+oNBiTLPpVLbl6A=";
          };

          hanamaru = pkgs.writeShellScriptBin "hanamaru.sh" ''
              exec ${hanamaru-go}/bin/hanamaru-go "$@"
            '';
        };

      # Add dependencies that are only needed for development
      devShells = 
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [ go gopls gotools go-tools ];
          };
        };

      # The default package for 'nix build'. This makes sense if the
      # flake provides only one package or there is a clear "main"
      # package.
      defaultPackage = self.packages.${system}.hanamaru;
    });
}