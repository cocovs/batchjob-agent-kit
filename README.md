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

Authentication is environment-variable based:

```bash
export BATCHJOB_SERVER="https://batchjob-test.shengsuanyun.com/batch"
export BATCHJOB_TOKEN="your-token"
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
```

If `template list` returns `no templates`, the target environment likely has not imported official template seed data yet.
