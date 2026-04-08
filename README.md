# batchjob-agent-kit

Public distribution repository for:

- `batchjob-cli`
- agent skill packs for Codex / Claude
- install scripts
- examples
- release workflows

This repository is the public delivery surface for developers and agents to use hosted BatchJob skills.

## Current MVP

The first public CLI is HTTP-based and focuses on:

- `batchjob-cli doctor`
- `batchjob-cli template list`
- `batchjob-cli template schema <template-id>`
- `batchjob-cli run submit <template-id> -f rows.jsonl`
- `batchjob-cli run watch <run-id>`
- `batchjob-cli artifact list <run-id>`
- `batchjob-cli artifact download <run-id>`

Authentication is environment-variable based:

```bash
export BATCHJOB_SERVER="https://batchjob-test.shengsuanyun.com/batch"
export BATCHJOB_TOKEN="your-token"
```

## Install From GitHub Release

```bash
curl -fsSL https://raw.githubusercontent.com/cocovs/batchjob-agent-kit/main/install.sh | bash
```

By default the installer downloads the latest release, installs `batchjob-cli` into `~/.local/bin`, and installs the Codex skill into `~/.codex/skills/batchjob/SKILL.md`.

Useful flags:

```bash
curl -fsSL https://raw.githubusercontent.com/cocovs/batchjob-agent-kit/main/install.sh | bash -s -- --agent claude
curl -fsSL https://raw.githubusercontent.com/cocovs/batchjob-agent-kit/main/install.sh | bash -s -- --no-skill
curl -fsSL https://raw.githubusercontent.com/cocovs/batchjob-agent-kit/main/install.sh | bash -s -- --version v0.1.0
```

## Local Build

```bash
cd cli
GOWORK=off go build ./cmd/batchjob-cli
```

## Quick Start

```bash
export BATCHJOB_SERVER="https://batchjob-test.shengsuanyun.com/batch"
export BATCHJOB_TOKEN="your-token"

./cli/batchjob-cli doctor
./cli/batchjob-cli template list
./cli/batchjob-cli template schema text-image-v1
./cli/batchjob-cli run submit text-image-v1 -f examples/text-image-v1.input.jsonl
./cli/batchjob-cli run watch <run-id>
./cli/batchjob-cli artifact list <run-id>
./cli/batchjob-cli artifact download <run-id> --output-dir ./downloads
```

If `template list` returns `no templates`, the target environment likely has not imported official template seed data yet.

## Input File Format

`run submit` accepts:

- JSONL with one flat object per line
- JSON array of flat objects

The field names must match the template schema. Starter files live under `examples/`.
