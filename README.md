# website

Personal website built with Go and HTMX.

This project is built using [Nix](https://nixos.org).

### Build

```sh
nix build
```

The resulting binary is at `./result/bin/website`. It bundles `templates/`,
`static/`, and `content/` at build time, so it is self-contained.

### Run

```sh
nix run
```

### Test / CI

```sh
nix flake check --all-systems
```

### Updating Go dependencies

After editing `go.mod` / running `go get`, regenerate the gomod2nix lockfile:

```sh
nix develop
gomod2nix import
```

## Adding content

### Posts

Drop a Markdown file into `content/posts/{school,work,home}/`:

```
content/posts/home/my-new-post.md
```

Frontmatter at the top of the file:

```yaml
---
title: "My New Post"
date: 2026-03-13
description: "One-line summary shown in the posts list."
tags: [rust, linux]
---
Markdown body here.
```

The category is inferred from the directory name (`school` | `work` | `home`).

### Projects / About

Edit `content/projects.md` and `content/about.md` directly.
