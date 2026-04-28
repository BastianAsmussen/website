{
  pkgs ? (
    let
      inherit (builtins) fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
      import (fetchTree nixpkgs.locked) {
        overlays = [
          (import "${fetchTree gomod2nix.locked}/overlay.nix")
        ];
      }
  ),
  buildGoApplication ? pkgs.buildGoApplication,
}:
buildGoApplication {
  pname = "website";
  version = "0.1.2";
  pwd = ./.;
  src = ./.;
  modules = ./gomod2nix.toml;

  # Bundle runtime assets (templates, static files, content) alongside the
  # binary and wrap the executable so it runs from that data directory.
  nativeBuildInputs = [pkgs.makeWrapper];

  postInstall = ''
    mkdir -p $out/share/website
    cp -r templates static content $out/share/website/

    wrapProgram $out/bin/website \
      --chdir "$out/share/website"
  '';

  meta = {
    description = "Personal website built with Go and HTMX";
    mainProgram = "website";
  };
}
