{
  description = "theatre devshell and package";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          name = "theatre-devshell";

          packages = with pkgs; [
            go
            gopls
            gotools
            delve
          ];
        };

        packages.theatre = pkgs.buildGoModule {
          pname = "theatre";
          version = "2026.02.24-a";

          src = self;

          vendorHash = "sha256-ny8EzevXK264YANZcJLHx6C16BMRVIs+XiZzY+Ncex0=";

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "A TUI presentation program using Markdown";
            license = licenses.mit;
            platforms = platforms.linux;
          };
        };

        apps.theatre = {
          type = "app";
          program = "${self.packages.${system}.theatre}/bin/theatre";
        };
      });
}
