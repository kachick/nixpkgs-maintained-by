rec {
  lib = {
    attrNames = builtins.attrNames;
    filterAttrs =
      pred: set:
      builtins.listToAttrs (
        builtins.concatMap (
          n:
          if pred n set.${n} then
            [
              {
                name = n;
                value = set.${n};
              }
            ]
          else
            [ ]
        ) (builtins.attrNames set)
      );
    isDerivation = x: x.type or "" == "derivation";
    maintainers = {
      user1 = "user1";
      user2 = "user2";
    };
  };

  mkPkg =
    name: maintainers:
    builtins.derivation {
      inherit name;
      system = "unused";
      builder = "unused";
      meta = { inherit maintainers; };
    };

  packageSingleMaintainer = mkPkg "packageSingleMaintainer" [ "user1" ];
  packageSharedMaintainers = mkPkg "packageSharedMaintainers" [
    "user1"
    "user2"
  ];
  packageNoMaintainers = mkPkg "packageNoMaintainers" [ ];

  brokenMeta = builtins.derivation {
    name = "brokenMeta";
    system = "unused";
    builder = "unused";
    meta = builtins.throw "error";
  };

  notAPkg = { };
}
