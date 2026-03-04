{
  description = "Hanamaru Discord Bot";

  # Nixpkgs / NixOS version to use.
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }: let
    # to work with older version of flakes
    lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";

    # Generate a user-friendly version number.
    version = builtins.substring 0 8 lastModifiedDate;
  in
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
