#! /bin/bash

systemctl --user enable --now podman.socket
systemctl --user disable --now podman.socket
systemctl --user enable --now podman.socket


XDG_RUNTIME_DIR=${XDG_RUNTIME_DIR:-/run/user/$(id -u)}
export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/podman/podman.sock
export DOCKER_SOCK=$XDG_RUNTIME_DIR/podman/podman.sock

export PATH="$PATH:$(go env GOPATH)/bin"
