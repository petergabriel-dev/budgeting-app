---
description: For generating bug reports.
---

# WORKFLOW: Bug Investigation

**Trigger:** `/generate-bug-report`

## Step 1: Investigation & Analysis

**Prompt:** "I see you have described a bug. I will NOT fix it yet.
Instead, I will now investigate the codebase to find the root cause.
I will:
1.  Read the relevant source files associated with your description.
2.  Trace the logic flow to identify where it breaks.
3.  Determine *why* it is happening (logic error, missing dependency, state mismatch, etc.)."

## Step 2: Generate Report

**Prompt:** "I have finished the investigation. I will now save a detailed report to the `bug-reports/` directory.
Filename format: `bug-reports/{{timestamp}}_[short_bug_name].md`

The content will include:
-   **Description:** What is happening.
-   **Root Cause Analysis:** The technical reason for the failure.
-   **Affected Files:** List of files involved.
-   **Proposed Solution:** A high-level strategy for the fix (not the code itself).

I will save this file now."

## Step 3: Confirmation

**Prompt:** "Investigation complete. Report saved to `bug-reports/`.
When you are ready to plan the fix, run `/generate-bug-fix` and reference this file."