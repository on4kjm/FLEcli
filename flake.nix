{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    nixpkgs-24-05.url = "github:nixos/nixpkgs?ref=nixos-24.05";
  };

  outputs = { self, nixpkgs, nixpkgs-24-05 }:
    let
      pkgs = import nixpkgs {
          system = "x86_64-linux";
      };
      pkgs-24-05 = import nixpkgs-24-05 {
          system = "x86_64-linux";
      };
    in {

      devShell.x86_64-linux = pkgs.mkShell {
        buildInputs = with pkgs; [
          go
          gopls
          golangci-lint
          bats
          pkgs-24-05.goreleaser
        ];
      };

  };
}
