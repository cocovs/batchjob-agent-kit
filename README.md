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
- `batchjob-cli model list --step-type image-generate`
- `batchjob-cli model get <model-id>`
- `batchjob-cli template list`
- `batchjob-cli template schema <template-id>`
- `batchjob-cli template download <template-id>`
- `batchjob-cli template validate-file <template-id> <xlsx-path>`
- `batchjob-cli template submit-file <template-id> <xlsx-path>`
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

## Install With Homebrew

```bash
brew install cocovs/tap/batchjob-cli
```

Or:

```bash
brew tap cocovs/tap
brew install batchjob-cli
```

Homebrew installs the CLI only. If you also want the Codex or Claude skill pack, use the GitHub Release installer above.

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
./cli/batchjob-cli model list --step-type image-generate
./cli/batchjob-cli model get google/gemini-2.5-flash-image
./cli/batchjob-cli template list
./cli/batchjob-cli template schema text-image-v1
./cli/batchjob-cli template download text-image-v1 --output-file ./text-image-v1.xlsx
./cli/batchjob-cli template validate-file text-image-v1 ./filled-text-image-v1.xlsx
./cli/batchjob-cli template submit-file text-image-v1 ./filled-text-image-v1.xlsx
./cli/batchjob-cli run submit text-image-v1 -f examples/text-image-v1.input.jsonl
./cli/batchjob-cli run watch <run-id>
./cli/batchjob-cli artifact list <run-id>
./cli/batchjob-cli artifact download <run-id> --output-dir ./downloads
```

If `template list` returns `no templates`, the target environment likely has not imported official template seed data yet.

## Model Discovery

Use model discovery when you need to understand which executable models are currently
available for one step type:

```bash
./cli/batchjob-cli model list --step-type text-generate
./cli/batchjob-cli model list --step-type image-generate --provider vertex
./cli/batchjob-cli model get google/gemini-2.5-flash-image
```

`model list` is step-type scoped on purpose. Common values are:

- `text-generate`
- `image-generate`
- `video-generate`

## Input File Format

`run submit` accepts:

- JSONL with one flat object per line
- JSON array of flat objects

The field names must match the template schema. Starter files live under `examples/`.

## Excel Template Workflow

For official templates, users can also work through Excel instead of JSON/JSONL:

```bash
./cli/batchjob-cli template download text-image-v1 --output-file ./text-image-v1.xlsx
./cli/batchjob-cli template validate-file text-image-v1 ./filled-text-image-v1.xlsx
./cli/batchjob-cli template submit-file text-image-v1 ./filled-text-image-v1.xlsx
./cli/batchjob-cli run watch <run-id>
```

`template submit-file` uploads the filled workbook and directly creates a run.
