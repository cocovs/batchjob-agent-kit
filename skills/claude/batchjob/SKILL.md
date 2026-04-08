---
name: batchjob
description: Use batchjob-cli to discover templates, submit batch runs, watch execution, and download artifacts.
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
3. Discover templates:
   `batchjob-cli template list`
4. Inspect schema:
   `batchjob-cli template schema <template-id>`
5. Submit a run:
   `batchjob-cli run submit <template-id> -f rows.jsonl`
6. Watch the run:
   `batchjob-cli run watch <run-id>`
7. Download outputs when ready:
   `batchjob-cli artifact download <run-id>`

## Current MVP Scope

The first public CLI release covers:

- environment verification
- template discovery
- official template row submission
- run watching
- artifact listing and download
