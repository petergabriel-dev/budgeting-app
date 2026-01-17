---
description: For adding new features to the project
---

# WORKFLOW: Generate Feature (Registry Aware)

**Trigger:** `/generate-feature`

## Step 1: Registry & Dependency Check

**Prompt:** "I am ready to scope a new feature.
First, read `specs/FEATURES.md` and `specs/ARCHITECTURE.md`.
Check for two things:
1.  **Duplication:** Does this feature (or a similar one) already exist in `FEATURES.md`?
2.  **Feasibility:** Does `ARCHITECTURE.md` support the necessary technologies? (e.g., if the user wants Realtime Chat, do we have WebSockets defined?)

Report your findings. If dependencies are missing, suggest adding them to Architecture first."

## Step 2: Define & Register

**Prompt:** "Based on the check, please generate the entry for this new feature.
Then, append it to `specs/FEATURES.md`.
Use this format:

## [Feature Name]
- **Status:** Planned
- **User Story:** [Story]
- **Tech Stack:** [Libraries/Components used]
- **Implementation:** [Key files to be created/modified]
- **Testing Strategy:** [How we verify this]

Ensure the file is saved."

## Step 3: Generate Tasks

**Prompt:** "Now, generate the atomic tasks required to build this feature based on the entry we just created.
Append these tasks to `prompts.md` under the header `## Feature: [Feature Name]`."

## Step 4: Confirm

**Prompt:** "Feature registered in `specs/FEATURES.md` and tasks queued in `docs/prompts.md`. Ready for execution."