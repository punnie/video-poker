{
  description = "A simple Go project";

  # Inputs include the Nix packages and any other dependencies.
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  # Outputs use inputs to define packages and development environments.
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        # Setting up the package environment with the required dependencies.
        pkgs = import nixpkgs {
          inherit system;
          config = { 
            allowUnfree = true; 
          }; # Modify this line based on your project's license requirements.
        };

        # Defining the development shell.
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            semgrep

            # Copilot
            nodejs

            llm
          ];
        };

      in {
        # The devShell provided by this flake.
        devShells.default = devShell;
      }
    );
}

