---
description: For generating the project architecture
---

# WORKFLOW: Generate Architecture (Smart Context)

**Trigger:** `/generate-architecture`

## Step 1: Context Analysis

**Prompt:** "I will check the status of `specs/FEATURES.md`.
**Condition:** Is the file missing, empty, or does it only contain a header?
-   **IF YES:** We are in the **Planning Phase (Greenfield)**.
-   **IF NO:** We are in the **Maintenance Phase (Brownfield)**."

## Step 2: Execution (Branching Logic)

**Prompt:**
"**SCENARIO A: Planning Phase (New Project)**
1.  Acknowledge this is a new build.
2.  Ask: 'What are the high-level goals? What is the preferred Tech Stack (Language, Framework, DB)?'
3.  **Wait for User Input.**
4.  Generate `specs/ARCHITECTURE.md` based on the answers.
5.  **Crucial:** Automatically create/overwrite `specs/FEATURES.md` with a 'Core Infrastructure' entry so the project is initialized.

**SCENARIO B: Maintenance Phase (Existing Project)**
1.  Read `specs/FEATURES.md` and `specs/ARCHITECTURE.md`.
2.  Perform a **Dependency Audit**:
    *   Flag unused dependencies in Architecture.
    *   Add missing dependencies required by Features.
3.  Update `specs/ARCHITECTURE.md` to match the reality of the features."

## Step 3: Save and Confirm

**Prompt:** "I have completed the architecture generation.
-   If new: Architecture created and Features initialized.
-   If existing: Architecture synchronized with Features."