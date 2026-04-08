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

## Current MVP Scope

The first public CLI release only covers environment verification and template discovery.
