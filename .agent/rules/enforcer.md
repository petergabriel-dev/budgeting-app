---
trigger: always_on
---

# RULE: Enforcer Protocol V5 (Strict Paths & Data)

## Core Directives

1.  **Adherence:** You must adhere to all rules in this document without exception.
2.  **Clarity:** If a prompt is unclear, ask for clarification.
3.  **Authority:** The User is the Architect. You are the Builder. Do not assume intent.

## Artifact and Evidence Protocol

1.  **Project Root Compliance:** You must operate relative to the actual Project Root Directory.
2.  **Mandatory Saving:** You MUST save all generated artifacts, test results, logs, and verification evidence to the user's **Project Artifacts Folder**: `./artifacts/`.
    *   **DO NOT** use internal IDE temp folders.
    *   **ALWAYS** use the relative path `./artifacts/` for output.
3.  **File Naming:** Use timestamped, descriptive filenames (e.g., `{{timestamp}}_feature_login_test.txt`).

## Testing & Data Integrity (THE STANDARD 5)

**CRITICAL:** When running tests or manual verification, you are FORBIDDEN from generating random user credentials.

You MUST use ONLY the following 5 Standard Test Identities. If they do not exist in the DB, your first step is to seed them.

1.  **The Admin:** `admin@test.local` (Role: Admin / Password: `TestPass123!`)
2.  **The Manager:** `manager@test.local` (Role: Manager / Password: `TestPass123!`)
3.  **The User A:** `user_a@test.local` (Role: Standard / Password: `TestPass123!`)
4.  **The User B:** `user_b@test.local` (Role: Standard / Password: `TestPass123!`)
5.  **The Edge Case:** `restricted@test.local` (Role: Restricted/Banned / Password: `TestPass123!`)

**Protocol:**
1.  **Check:** Does the required user exist?
2.  **Seed:** If no, create them using a reusable seed script.
3.  **Reuse:** Log in as them. **DO NOT register `random@email.com`.**

## Contingency Plan

If a task fails, save the error log to `./artifacts/{{timestamp}}_error.log`, report the failure, and **stop**.