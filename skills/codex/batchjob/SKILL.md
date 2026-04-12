---
name: batchjob
description: Use batchjob-cli to discover executable models and templates, submit batch runs, watch execution, and download artifacts.
---

# batchjob

Use this skill when the user wants to run hosted batch generation tasks through `batchjob-cli`.

## When To Use

- The user wants batch text-to-image or text-to-image-to-video generation.
- The user is comfortable using a developer tool or agent-assisted CLI workflow.
- The task can be expressed as repeated structured rows instead of a one-off chat response.

## When Not To Use

- The user only needs a single immediate generation in chat.
- The task is exploratory and not yet structured enough for batch input.
- `BATCHJOB_SERVER` or `BATCHJOB_TOKEN` is not configured and the user does not want setup help.

## Command Pattern

1. Check environment:
   `batchjob-cli doctor`
2. Discover executable models when step-level model choice matters:
   `batchjob-cli model list --step-type image-generate`
   `batchjob-cli model get <model-id>`
3. Discover available templates:
   `batchjob-cli template list`
4. Inspect one template:
   `batchjob-cli template schema <template-id>`
5. Prepare a JSONL or JSON array file that matches the schema.
6. Submit a run:
   `batchjob-cli run submit <template-id> -f rows.jsonl`
7. Watch the run:
   `batchjob-cli run watch <run-id>`
8. List or download artifacts:
   `batchjob-cli artifact list <run-id>`
   `batchjob-cli artifact download <run-id>`

## Current MVP Scope

The public CLI MVP currently supports:

- `batchjob-cli doctor`
- `batchjob-cli model list --step-type <step-type>`
- `batchjob-cli model get <model-id>`
- `batchjob-cli template list`
- `batchjob-cli template schema <template-id>`
- `batchjob-cli run submit <template-id> -f rows.jsonl`
- `batchjob-cli run watch <run-id>`
- `batchjob-cli artifact list <run-id>`
- `batchjob-cli artifact download <run-id>`
