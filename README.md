# nixpkgs-maintained-by

## Usage

```bash
# Find from your system's nixpkgs
go run main.go -id kachick -json

# Find from your specified nixpkgs
NIX_PATH=nixpkgs=/THE_PATH go run main.go -id kachick

# Run via flake
nix run github:kachick/nixpkgs-maintained-by -- -id kachick
```

## Motivation

I often need to automatically get a list of `pname`s for packages where I am listed as a maintainer in `nixpkgs`.

This is because several tools use `pname` as a key, such as:

- [https://github.com/nix-community/hydra-check](https://github.com/nix-community/hydra-check)
- [https://github.com/kachick/nixpkgs-update-log-checker](https://github.com/kachick/nixpkgs-update-log-checker)

While it’s possible to get similar data from [Repology](https://repology.org/), it has some downsides:

- Some package names differ from `nixpkgs` `pname`s (e.g., `hugo` appears as `hugo-sitegen`).
- It may take time to sync the list with nixpkgs.

This CLI helps extract the actual `pname`s directly from the `nixpkgs` source.

## Links

When looking for a way to solve this, I found [this helpful snippet](https://discourse.nixos.org/t/how-to-get-a-list-of-packages-maintained-by-someone/29963/3).
