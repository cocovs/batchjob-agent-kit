---
name: batchjob
description: Use batchjob-cli to discover executable models and templates, download or submit official template Excel files, watch execution, and download artifacts.
---

# batchjob

Use this skill when the user wants to use hosted BatchJob capabilities through `batchjob-cli`.

## When To Use

- The user needs batch-oriented generation workflows.
- The user is a developer or uses an agent that can call a CLI.
- The task is better modeled as a template-driven batch run than as an in-chat one-off answer.

## First Steps

1. Ensure:
   `BATCHJOB_SERVER`
   `BATCHJOB_TOKEN`
2. Run:
   `batchjob-cli doctor`
3. Discover executable models when needed:
   `batchjob-cli model list --step-type image-generate`
   `batchjob-cli model get <model-id>`
4. Discover templates:
   `batchjob-cli template list`
5. Inspect schema:
   `batchjob-cli template schema <template-id>`
6. If you want the official Excel workflow:
   `batchjob-cli template download <template-id>`
   `batchjob-cli template validate-file <template-id> <xlsx-path>`
   `batchjob-cli template submit-file <template-id> <xlsx-path>`
7. Or submit a run from JSON/JSONL rows:
   `batchjob-cli run submit <template-id> -f rows.jsonl`
8. Watch the run:
   `batchjob-cli run watch <run-id>`
9. Download outputs when ready:
   `batchjob-cli artifact download <run-id>`

## Current MVP Scope

The first public CLI release covers:

- environment verification
- executable model discovery
- template discovery
- official template Excel download / validation / submission
- official template row submission
- run watching
- artifact listing and download
