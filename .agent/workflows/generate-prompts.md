---
description: 
---

# WORKFLOW: Generate Prompts (Backlog Sync)

**Trigger:** `/generate-prompts`

## Step 1: Audit Feature Registry

**Prompt:** "I will read `specs/FEATURES.md`.
I am looking for entries where the **Status** is set to **'Planned'** or **'In Progress'**.
I will also read `docs/prompts.md` to see which tasks are currently queued."

## Step 2: Identify Missing Tasks

**Prompt:** "I will compare the two documents.
For every 'Planned' feature in `FEATURES.md`:
-   **Check:** Do tasks for this feature already exist in `prompts.md`?
-   **Action:** If tasks are missing, I will generate the implementation tasks for that specific feature (using the context from Architecture and DB Design).

**Note:** I will **NOT** regenerate tasks for features marked 'Completed' or features that already have tasks queued."

## Step 3: Append & Sync (With Skill Tags)

**Prompt:** "I will append any missing tasks to `docs/prompts.md` under their respective feature headers.

**CRITICAL INSTRUCTION: SKILL TAGGING**
When generating tasks, you MUST analyze the technology involved and append the specific Skill Trigger to the task description:

1.  **Backend / Go Tasks:**
    -   If the task involves structuring a new feature module (Hander/Service/Interface).
    -   **Append:** `(Use Skill: create_go_feature)`

2.  **Frontend / React Tasks:**
    -   If the task involves structuring a new feature module (Components/Hooks).
    -   **Append:** `(Use Skill: create_react_feature)`

3.  **Database / SQL Tasks:**
    -   If the task involves `goose` migrations, `sqlc` queries, or schema changes.
    -   **Append:** `(Use Skill: db-migration)`

**Example Output:**
- 'Task 1: Scaffold Auth feature backend. (Use Skill: create_go_feature)'
- 'Task 2: Scaffold Auth feature frontend. (Use Skill: create_react_feature)'"

## Step 4: Confirmation

**Prompt:** "Sync complete. `prompts.md` has been updated with skill-tagged tasks."