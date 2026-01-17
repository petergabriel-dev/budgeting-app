---
description: For logging prompts and tasks after every execution.
---

# WORKFLOW: Log Work (Session Summary)

**Trigger:** `/log-work`

## Step 1: Context Analysis & Logging

**Prompt:** "Please review the chat history for the current task/feature we have been working on.
Identify the sequence of prompts and actions leading up to this point (not just the last message).

Consolidate this session into a single log entry and append it to `docs/execution_log.md` using this strict format:

---
**Timestamp:** {{timestamp}}
**Task Context:** [Synthesize the user's original goal and subsequent adjustments]
**Implementation:** [Summary of files created, functions modified, and logic added]
**Evidence:** [List the artifacts or specific files saved during this session]
**Status:** COMPLETED

Do not ask for input. Derive the summary entirely from our chat history and save it immediately."

## Step 2: Verification

**Prompt:** "I have synthesized the recent session history and updated the `docs/execution_log.md`. Please verify the entry."