{
  lib,
  buildGoModule,
  versionCheckHook,
}:

buildGoModule (finalAttrs: {
  pname = "nixpkgs-maintained-by";
  version = "0.1.0"; # TODO: Load the actual version
  src = lib.fileset.toSource {
    root = ./.;
    # - Don't just use `fileset.gitTracked root`, then always rebuild even if just changed the README.md
    # - Don't use gitTracked for now, even if filtering with intersection, the feature is not supported in nix-update. See https://github.com/Mic92/nix-update/issues/335
    fileset = lib.fileset.unions [
      ./go.mod
      # ./go.sum
      ./main.go
      ./filter.nix
    ];
  };
  # src = lib.cleanSource self; # Requires this old style if I use nix-update

  # https://github.com/NixOS/nixpkgs/issues/346380
  ldflags = [
    "-s"
  ];

  vendorHash = null;

  # https://github.com/kachick/times_kachick/issues/316
  env.CGO_ENABLED = "0";

  nativeInstallCheckInputs = [
    versionCheckHook
  ];
  doInstallCheck = true;

  meta = {
    description = "Filter nixpkgs by the maintainer ID";
    homepage = "https://github.com/kachick/nixpkgs-maintained-by";
    license = lib.licenses.mit;
    mainProgram = "nixpkgs-maintained-by";
  };
})
