{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  hardeningDisable = [ "fortify" ];
  buildInputs = with pkgs; [
    go
  ];
  shellHook = ''
    getrootdir() {
      local dir=""
      local os=`uname`
      case $os in
        Darwin) dir="/Users/noor/work/Vivasoft/Blogs" ;;
        Linux)  dir="/home/noor/work/vivasoft-blogs" ;;
	* ) ;;
      esac
      printf "%s" "$dir"
    }
  '';
}
