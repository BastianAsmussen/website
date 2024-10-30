{pkgs, ...}: {
  packages = with pkgs; [
    git
    nodejs
    vite
  ];

  languages.typescript.enable = true;

  processes = {
    install.exec = "npm install";
    build.exec = "npm run build";
  };

  pre-commit.hooks.prettier.enable = true;
}
