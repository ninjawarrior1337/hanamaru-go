{
  description = "Hanamaru Discord Bot";

  # Nixpkgs / NixOS version to use.
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
    in {
      # Provide some binary packages for selected system types.
      packages = import ./nix {inherit pkgs;};

      # Add dependencies that are only needed for development
      devShells = {
        default = pkgs.mkShell {
          buildInputs = with pkgs; [go gopls gotools go-tools just];
        };
      };
    });
}
