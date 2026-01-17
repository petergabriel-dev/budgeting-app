---
description: For generating the db design.
---

# WORKFLOW: Generate DB Design

**Trigger:** `/generate-db-design`

## Step 1: Read Architecture

**Prompt:** "I will now read the `specs/ARCHITECTURE.md` to understand the system context for the database design."

## Step 2: Gather Data Model Requirements

**Prompt:** "Based on the architecture, please describe the data entities, their fields, and their relationships (e.g., users, posts, comments)."

## Step 3: Generate Document

**Prompt:** "Thank you. I will now generate the `specs/DB_DESIGN.md` document with the full data model and schema."

## Step 4: Save and Confirm

**Prompt:** "The `specs/DB_DESIGN.md` document has been created. Please review it before we generate the task queue."