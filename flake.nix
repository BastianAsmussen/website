{
  description = "Personal website built with Go and HTMX.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    gomod2nix,
  }:
    (flake-utils.lib.eachDefaultSystem
      (system: let
        pkgs = nixpkgs.legacyPackages.${system};
        application = pkgs.callPackage ./. {
          inherit (gomod2nix.legacyPackages.${system}) buildGoApplication;
        };
      in {
        checks.build = application;
        packages.default = application;
        devShells.default = pkgs.callPackage ./shell.nix {
          inherit (gomod2nix.legacyPackages.${system}) mkGoEnv gomod2nix;
        };
      }))
    // {
      nixosModules.default = {
        config,
        lib,
        pkgs,
        ...
      }: {
        options.services.website = {
          enable = lib.mkEnableOption "personal website";
          port = lib.mkOption {
            type = lib.types.port;
            default = 8080;
            description = "TCP port the server listens on.";
          };
        };

        config = lib.mkIf config.services.website.enable {
          systemd.services.website = {
            description = "Personal website";
            wantedBy = ["multi-user.target"];
            after = ["network.target"];
            serviceConfig = {
              ExecStart = "${self.packages.${pkgs.system}.default}/bin/website";
              Environment = "PORT=${toString config.services.website.port}";
              DynamicUser = true;
              Restart = "on-failure";
              NoNewPrivileges = true;
              PrivateTmp = true;
              ProtectSystem = "strict";
              ProtectHome = true;
            };
          };
        };
      };
    };
}
