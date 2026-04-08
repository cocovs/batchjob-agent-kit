#!/usr/bin/env bash
set -euo pipefail

REPO="cocovs/batchjob-agent-kit"
VERSION="${VERSION:-latest}"
AGENT="codex"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --agent)
      AGENT="${2:-codex}"
      shift 2
      ;;
    --install-dir)
      INSTALL_DIR="${2:-$HOME/.local/bin}"
      shift 2
      ;;
    *)
      echo "unknown argument: $1" >&2
      exit 1
      ;;
  esac
done

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
case "$ARCH" in
  arm64|aarch64) ARCH="arm64" ;;
  x86_64|amd64) ARCH="amd64" ;;
  *)
    echo "unsupported architecture: $ARCH" >&2
    exit 1
    ;;
esac

mkdir -p "$INSTALL_DIR"

echo "batchjob-agent-kit installer"
echo "repo: $REPO"
echo "version: $VERSION"
echo "agent: $AGENT"
echo "install dir: $INSTALL_DIR"
echo
echo "MVP note:"
echo "- CLI release packaging is not wired yet."
echo "- For now, build from source inside the public repo:"
echo "    cd cli && go build -o \"$INSTALL_DIR/batchjob-cli\" ./cmd/batchjob-cli"
echo "- Then install the skill file for your agent from ./skills/$AGENT/batchjob/SKILL.md"
echo
echo "Expected environment:"
echo "  export BATCHJOB_SERVER=https://batchjob-test.shengsuanyun.com/batch"
echo "  export BATCHJOB_TOKEN=your-token"
