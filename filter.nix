{
  maintainerId,
  pkgs ? import <nixpkgs> { },
}:
let
  lib = pkgs.lib;
  targetMaintainer = lib.maintainers."${maintainerId}" or null;
  isMatch = m: targetMaintainer != null && m == targetMaintainer;
  hasMaintainer =
    pkg:
    let
      meta = (builtins.tryEval (pkg.meta or { })).value or { };
      maintainers = meta.maintainers or [ ];
    in
    builtins.any isMatch maintainers;
  check =
    n: v:
    let
      res = builtins.tryEval (lib.isDerivation v && hasMaintainer v);
    in
    res.success && res.value;
in
if targetMaintainer == null then [ ] else lib.attrNames (lib.filterAttrs check pkgs)
