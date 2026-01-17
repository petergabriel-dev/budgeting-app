---
description: For generating bug fix prompts.
---

# WORKFLOW: Bug Fix Planning

**Trigger:** `/generate-bug-fix`

## Step 1: Read Report

**Prompt:** "Please indicate which report in `bug-reports/` I should use (or tag it in your prompt).
I will read that document to understand the root cause and the proposed solution."

## Step 2: Generate Tasks (With Skill Tags)

**Prompt:** "Based on the 'Proposed Solution' in the report, I will now generate atomic, step-by-step coding tasks.

I will:
1.  Create a header in `docs/prompts.md`: `## Bug Fix: [Bug Name]`
2.  Create tasks to implement the fix (starting with a reproduction test case).
3.  Append these tasks to `docs/prompts.md`.

**CRITICAL INSTRUCTION: SKILL TAGGING**
When generating tasks, you MUST analyze the technology involved and append the specific Skill Trigger to the task description:

1.  **Backend / Go Tasks:**
    -   If the fix involves `gin`, `handler`, `service`, or `middleware`.
    -   **Append:** `(Use Skill: go-feature)`

2.  **Frontend / React Tasks:**
    -   If the fix involves `components`, `pages`, `hooks`, or `TanStack Query`.
    -   **Append:** `(Use Skill: react-component)`

3.  **Database / SQL Tasks:**
    -   If the fix involves `goose` migrations, `sqlc` queries, or schema changes.
    -   **Append:** `(Use Skill: db-migration)`

**Example Output:**
- 'Task 1: Fix validation logic in user handler. (Use Skill: go-feature)'"

## Step 3: Confirmation

**Prompt:** "The fix strategy has been converted into actionable tasks in `prompts.md`. You may now execute them using the standard loop."