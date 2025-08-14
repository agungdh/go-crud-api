{ pkgs, ... }: {
  # Which nixpkgs channel to use.
  channel = "stable-25.05"; # or "unstable"

  # Packages yang kamu butuhkan
  packages = [
    pkgs.go
    pkgs.air
    pkgs.btop
    pkgs.htop
    pkgs.fastfetch
    pkgs.nano
    pkgs.wget
    pkgs.docker
    pkgs.docker-buildx
    pkgs.docker-compose
    pkgs.flyway
    pkgs.bash-completion
    pkgs.zram-generator
    pkgs.util-linux
    
  ];

  # Enable Docker (rootless) sebagai service supaya otomatis jalan saat start
  services.docker.enable = true;

  # Sets environment variables in the workspace
  env = {};

  idx = {
    # VS Code extensions
    extensions = [
      "golang.go"
    ];

    workspace = {
      onCreate = {
        # Open editors for the following files by default, if they exist:
        default.openFiles = ["main.go"];
      };

      # Runs when a workspace is (re)started
      onStart = {
        # Jalankan hot-reload Go
        run-server = "air";
      };
    };
  };
}
