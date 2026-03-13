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
  }: (
    flake-utils.lib.eachDefaultSystem
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
    })
  );
}
