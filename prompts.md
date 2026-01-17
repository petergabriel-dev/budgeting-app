# Feature: Multi-Role Authentication & Dashboards

## Task 1: Database Schema Update
**Goal:** Add support for user roles in the database.
- [x] Create a new Goose migration to add a `role` column to the `users` table.
- [x] Define the roles: 'admin', 'superadmin', 'client'. Default to 'client' or make it required.
- [x] Run the migration.
- [x] Update `sqlc` queries if necessary (e.g., `CreateUser`, `GetUser`) to include the role.

## Task 2: Backend Logic for Roles
**Goal:** Ensure the backend handles and enforces roles.
- [x] Update the `User` struct in Go to include the `Role` field.
- [x] Update the registration handler to accept a role (or default it appropriately).
- [x] Update the login handler/response to include the user's role (so the frontend knows where to redirect).
- [x] Create a middleware `RequireRole(role string)` for endpoint protection.

## Task 3: Login Screen UI
**Goal:** Create a unified login screen for all user types.
- [x] Create `src/frontend/src/pages/Login.tsx` with email/password form.
- [x] Style with modern, premium aesthetics (dark mode, gradients, micro-animations).
- [x] Integrate with auth API using TanStack Query mutation.
- [x] Handle loading and error states gracefully.

## Task 4: Frontend Dashboard Placeholders
**Goal:** Create visual destinations for each role.
- [x] Create `src/frontend/src/pages/dashboards/AdminDashboard.tsx`.
- [x] Create `src/frontend/src/pages/dashboards/SuperAdminDashboard.tsx`.
- [x] Create `src/frontend/src/pages/dashboards/ClientDashboard.tsx`.
- [x] Add simple "Welcome [Role]" text to each with basic styling.

## Task 5: Frontend Routing & Auth Logic
**Goal:** Redirect users to the correct dashboard upon login and protect routes.
- [x] Create `src/frontend/src/contexts/AuthContext.tsx` to manage auth state and user role.
- [x] Update the Login component to redirect based on the role received from the API.
    - Admin -> `/admin/dashboard`
    - Superadmin -> `/superadmin/dashboard`
    - Client -> `/dashboard`
- [x] Configure React Router with protected routes that check the user's role.
- [x] Create a `ProtectedRoute` component that redirects unauthenticated users to `/login`.

## Feature: Superadmin User Management

### Task 6: Superadmin Backend - Create Admin
- **Goal:** Allow Superadmins to create Admin users.
- **Dependencies:** Task 2 (Backend Logic for Roles)
- **Steps:**
  1. Create `src/backend/internal/shared/middleware/rbac.go` with `RequireRole(role)`.
  2. Implement `CreateAdmin` in `features/users/service.go`.
  3. Create `features/users/handler.go` with `CreateAdmin` endpoint (POST /api/users/admin).
  4. Register route in `main.go` wrapped in Superadmin middleware.

### Task 7: Superadmin Backend - List Admins
- **Goal:** Allow Superadmins to view all Admins.
- **Steps:**
  1. Implement `ListAdmins` in `features/users/service.go`.
  2. Create `ListAdmins` endpoint in handler (GET /api/users/admin).
  3. Register route.

### Task 8: Superadmin Frontend - Manage Admins
- **Goal:** UI to view and add admins.
- **Dependencies:** Task 4 (Dashboard Placeholders)
- **Steps:**
  1. Update `SuperAdminDashboard.tsx` to fetch and display list of admins.
  2. Create `AddAdminForm.tsx` component.
  3. Wire up form to `POST /api/users/admin`.

### Task 9: Superadmin Feature Verification (Automated)
- **Goal:** Verify the security and functionality of the Superadmin features via integration tests.
- **Dependencies:** Tasks 6, 7
- **Steps:**
  1. Create `src/backend/tests/integration/superadmin_test.go`.
  2. Test Case: Superadmin can create an Admin user (201 Created).
  3. Test Case: Admin/Client cannot create an Admin user (403 Forbidden).
  4. Test Case: Superadmin can list admins (200 OK).
  5. Test Case: Admin/Client cannot list admins (403 Forbidden).

### Task 10: Superadmin Feature Verification (Manual)
- **Goal:** Manually verify the user journey.
- **Dependencies:** Task 8
- **Steps:**
  1. Log in as `manager@test.local` (Superadmin).
  2. Navigate to Dashboard and verify "System Administrators" list is visible.
  3. Click "Add New Admin", fill out form, submit.
  4. Verify success message and list update.
  5. Logout and log in as the newly created Admin.
  6. Verify redirection to Admin Dashboard.

## Feature: Enhanced Superadmin Admin Management

### Task 11: Database Schema Update (Last Login)
- **Goal:** Track user authentication timestamps.
- **Steps:**
  1. [x] Create a `goose` migration to add `last_login_at` (timestamptz, nullable) to `users` table.
  2. [x] Run migration.
  3. [x] Update `sqlc` queries (`GetUser`, `ListAmins`) to return `last_login_at`.
  4. [x] Update `User` struct in Go.
  5. [x] Update Login Handler to set `last_login_at` = NOW() on successful authentication.

### Task 12: Backend Admin Management (Update/Delete)
- **Goal:** Enable updating and deleting admin users.
- **Steps:**
  1. [x] Create `UpdateUserQuery` in `sqlc` (fields: name, email, password_hash if provided, role).
  2. [x] Create `DeleteUserQuery`.
  3. [x] Implement `UpdateAdmin` in `features/users/handler.go` (PUT /api/users/admin/:id). Handle optional password update (hash it!).
  4. [x] Implement `DeleteAdmin` in `features/users/handler.go` (DELETE /api/users/admin/:id).
  5. [x] Register routes protected by Superadmin middleware.

### Task 13: Frontend List Update (Last Login)
- **Goal:** Display the new data in the admin list.
- **Steps:**
  1. [x] Update `SuperAdminDashboard.tsx` logic to display the "Last Login" column.
  2. [x] Format the date beautifully (e.g., "Start of shift", "2 hours ago" or specific date).

### Task 14: Frontend Edit & Delete Actions
- **Goal:** UI controls for managing admins.
- **Steps:**
  1. [x] Add "Edit" and "Delete" actions to the Admin table rows.
  2. [x] Create `EditAdminModal` component (Name, Email, New Password field).
  3. [x] Wire Edit form to `PUT /api/users/admin/:id`.
  4. [x] Create Delete confirmation dialog.
  5. [x] Wire Delete action to `DELETE /api/users/admin/:id`.
  6. [x] Ensure list refreshes after actions.

### Task 15: Verification (Enhanced Feature)
- **Goal:** Verify the new capabilities.
- **Testing:**
  1. [x] **Automated:** Integration test for `UpdateAdmin` (change name/password) and `DeleteAdmin`.
  2. [x] **Manual:** Login as Superadmin -> Edit an Admin's name -> Verify change reflected. Change password -> Logout -> Login as Admin with new password -> Verify success. Delete Admin -> Verify row removed.


## Feature: Client Management (Admin)

### Task 16: Backend - Create & List Clients
- **Goal:** Allow Admins to create and list Client users.
- **Dependencies:** Task 2 (Roles), Task 11 (Last Login)
- **Steps:**
  1. [x] Create `ListClients` query in `sqlc` (select all users where role = 'client').
  2. [x] Implement `CreateClient` in `features/users/service.go` (ensure role is hardcoded to 'client').
  3. [x] Implement `ListClients` in service.
  4. [x] Create handlers: `POST /api/users/client` and `GET /api/users/client`.
  5. [x] Register routes protected by Admin middleware (Admins and Superadmins can access).

### Task 17: Backend - Edit & Delete Clients
- **Goal:** Allow Admins to manage Client accounts.
- **Steps:**
  1. [x] Reuse `UpdateUserQuery` and `DeleteUserQuery`.
  2. [x] Implement `UpdateClient` and `DeleteClient` in handlers (ensure target user is actually a client to prevent privilege escalation or accidental admin deletion).
  3. [x] Register routes: `PUT /api/users/client/:id` and `DELETE /api/users/client/:id`.

### Task 18: Frontend - Client Management UI
- **Goal:** Interface for Admins to view/manage clients.
- **Dependencies:** Task 4 (Dashboard Placeholders)
- **Steps:**
  1. [x] Create `src/frontend/src/api/clients.ts` (copy pattern from `users.ts`).
  2. [x] Update `AdminDashboard.tsx`:
     - [x] Add "Clients Management" section (similar to Superadmin's Admin list).
     - [x] Display Client table (Name, Email, Last Login, Status).
  3. [x] Create `AddClientForm` component.
  4. [x] Wire up "Add Client" button.

### Task 19: Frontend - Edit & Delete Client UI
- **Goal:** Actions for Client rows.
- **Steps:**
  1. [x] Add "Edit" and "Delete" buttons to the client table.
  2. [x] Reuse `EditUserModal` (refactor `EditAdminModal` to be generic if needed, or create copy `EditClientModal`).
  3. [x] Connect to client API endpoints.
  4. [x] Implement optimistic updates or refetch on success.

### Task 20: Client Management Verification
- **Goal:** Verify Admin capabilities.
- **Testing:**
  1. [x] **Automated:** Integration tests for Client CRUD. Ensure Admins can create Clients. Ensure Clients cannot create Clients.
  2. [x] **Manual:**
     - [x] Log in as Admin (`admin@test.local`).
     - [x] Create a new Client.
     - [x] Verify Client appears in list.
     - [x] Edit Client details.
     - [x] Delete Client.

## Feature: Package Management (Admin)

### Task 21: Database Schema Update (Packages & Subscriptions)
- **Goal:** Create tables for packages and subscriptions.
- **Steps:**
  - [x] Create a `goose` migration to add `packages` and `subscriptions` tables (reference `specs/DB_DESIGN.md`).
  - [x] Create necessary enums (e.g., `subscription_status`).
  - [x] Run migration.
  - [x] Generate `sqlc` code for new tables.

### Task 22: Backend - Package Management (CRUD)
- **Goal:** Allow Admins to manage packages.
- **Dependencies:** Task 21
- **Steps:**
  - [x] Create `sqlc` queries for `CreatePackage`, `ListPackages`, `GetPackage`, `UpdatePackage`, `DeletePackage`.
  - [x] Implement `PackageService` methods in `internal/features/packages/service.go`.
  - [x] Create handlers in `internal/features/packages/handler.go`.
  - [x] Register routes: `POST/GET /api/packages`, `PUT/DELETE /api/packages/:id`.
  - [x] Protect routes with `RequireRole('admin')` (allows Admin and Superadmin).

### Task 23: Backend - Subscription Management
- **Goal:** Allow Admins to assign packages to clients.
- **Dependencies:** Task 21, Task 22
- **Steps:**
  - [x] Create `sqlc` queries for `CreateSubscription`, `ListSubscriptions`, `UpdateSubscription` (status), `DeleteSubscription`.
  - [x] Implement `SubscriptionService` methods.
  - [x] Create handlers: `POST /api/subscriptions` (assign), `GET /api/subscriptions` (filter by client/status).
  - [x] Register routes protected by Admin middleware.

### Task 24: Frontend - Package Management UI
- **Goal:** Interface for Admins to manage service packages.
- **Steps:**
  - [x] Create `src/frontend/src/api/packages.ts`.
  - [x] Update `AdminDashboard.tsx`:
     - [x] Add "Packages" tab/section.
     - [x] Display table/grid of existing packages.
  - [x] Create `AddPackageForm` modal/page.
  - [x] Create `EditPackageModal`.
  - [x] Implement Delete confirmation.

### Task 25: Frontend - Client Subscription UI
- **Goal:** Interface for Admins to assign packages to clients.
- **Dependencies:** Task 24, Task 18 (Client Management)
- **Steps:**
  - [x] Create `src/frontend/src/api/subscriptions.ts`.
  - [x] Update `AdminDashboard.tsx` (Clients section) or Client Details View:
     - [x] Add "Assign Package" action.
     - [x] Show current subscription status for each client.
  - [x] Create `AssignPackageModal` (Select Package, Start Date).

### Task 26: Package Management Verification
- **Goal:** Verify Admin capabilities for packages and subscriptions.
- **Testing:**
  - [x] **Automated:** Integration tests for Package CRUD and Subscription flows.
  - [x] **Manual:**
     - [x] Log in as Admin.
     - [x] Create a "Gold Tier" package.
     - [x] Assign "Gold Tier" to a Client.
     - [x] Verify Client sees the subscription (if Client Dashboard is ready, otherwise check DB/Admin view).

### [x] Task 27: Frontend - Multi-Package Support UI
- **Goal:** Ensure Admin Dashboard handles multiple active packages per client.
- **Dependencies:** Task 25
- **Steps:**
  1. [x] Refactor `AdminDashboard.tsx` "Package" column to handle an array of subscriptions.
  2. [x] Update `AssignPackageModal` to allow stacking assignments (checking for duplicates if necessary).
  3. [x] Verify UI displays multiple packages correctly (e.g., "Package A, Package B").

## Feature: Client Dashboard & Requests

### [x] Task 28: Client Dashboard - View Subscription
- **Goal:** Allow Clients to see their active packages.
- **Dependencies:** Task 21 (Schema), Task 23 (Back Subscriptions)
- **Steps:**
  - [x] Create `src/frontend/src/pages/dashboards/ClientDashboard.tsx` (enhance placeholder).
  - [x] Implement `useQuery` to fetch `my_subscription` (uses `GET /api/v1/client/subscriptions` with user isolation).
  - [x] Display Active Package details (Name, Deliverables, Start Date).

### [x] Task 29: Database Schema Update (Requests & Revisions)
- **Goal:** Tables for tracking client requests.
- **Steps:**
  - [x] Migration for `requests` table (user_id, title, content, status, priority, due_date).
  - [x] `sqlc` generation.

### [x] Task 30: Backend - Requests Management
- **Goal:** API for Clients to submit/view requests and Admins to manage them.
- **Steps:**
  - [x] `RequestService` with CRUD.
  - [x] `RequestHandler`:
    - [x] Client: `POST /api/client/requests`, `GET /api/client/requests`
    - [x] Admin: `GET /api/requests` (all), `PUT /api/requests/:id` (status update).

### [x] Task 31: Frontend - Request Submission
- **Goal:** UI for clients to submit requests.
- **Steps:**
  - [x] Add "Requests" tab to Client Dashboard.
  - [x] Create `CreateRequestForm` component.
  - [x] Display list of submitted requests with status tags.

## Feature: Deliverables Management

### [x] Task 32: Database Schema Update (Deliverables/Files)
- **Goal:** Track file uploads associated with requests.
- **Steps:**
  1. [x] Migration for `deliverables` table (request_id, file_url, file_type, uploaded_by_admin_id).

### [x] Task 33: Backend - File Uploads (Local/Bucket)
- **Goal:** Handle file storage.
- **Steps:**
  1. [x] Create `FileService` (for now, local filesystem storage in `uploads/` or a mock S3 interface).
  2. [x] Endpoints: `POST /api/upload` (returns URL).

### [x] Task 34: Frontend - Deliverables UI
- **Goal:** Admins upload, Clients download.
- **Steps:**
  1. [x] Admin: In "Request Details" view, add "Upload Deliverable" button.
  2. [x] Client: In "Request Details" view, show list of deliverables with "Download" links.

## Feature: Analytics (Admin)

### Task 35: Backend - Analytics Endpoints
- **Goal:** Access high-level stats.
- **Steps:**
  1. [x] Create custom SQL queries for: Count Active Clients, Count Pending Requests, Total Revenue (if applicable).
  2. [x] Create `GET /api/analytics/dashboard`.

### Task 36: Frontend - Admin Analytics Widgets
- **Goal:** Visual overview on Admin Dashboard.
- **Steps:**
  1. [x] Create `StatCard` component.
  2. [x] Fetch analytics data and display at top of Admin Dashboard.

### [x] Task 37: Database Schema Update (Pending Status)
- **Goal:** Support 'pending' subscription status.
- **Steps:**
  - [x] Create migration: `ALTER TYPE subscription_status ADD VALUE 'pending';`.
  - [x] Run migration.

### [x] Task 38: Backend - Package Requests API
- **Goal:** Endpoints for requesting and approving packages.
- **Steps:**
  - [x] `POST /api/subscriptions/request`: Create pending subscription (Client only).
  - [x] `PUT /api/subscriptions/:id/approve`: Set status to active (Admin only).
  - [x] Ensure `ListPackages` is accessible to Clients.

### [x] Task 39: Frontend - Client Package Request UI
- **Goal:** Clients can browse and request.
- **Steps:**
  - [x] Update `ClientDashboard` to list packages.
  - [x] Add "Request" button.
  - [x] Display "Pending" badge for unapproved subscriptions.

### [x] Task 40: Frontend - Admin Approval UI
- **Goal:** Admins can view and approve requests.
- **Steps:**
  - [x] Update `AdminDashboard` Subscription/Client list to highlight Pending items.
  - [x] Add "Approve" action button.

### [x] Task 41: UI Refactor - Admin Requests Tab
- **Goal:** Move "Pending Package Requests" to a dedicated tab in Admin Dashboard.
- **Steps:**
  - [x] Update `AdminDashboard.tsx` state to include 'requests' tab.
  - [x] Implement "Package Requests" tab content (list of pending subscriptions).
  - [x] Move approval logic to this new view.
  - [x] Add badge count to tab label (e.g., "Requests (3)").

### [x] Task 42: Bug Fix - Prevent Duplicate Subscriptions
- **Goal:** Ensure clients cannot request or be assigned the same package multiple times if already active/pending.
- **Steps:**
  - [x] Create DB migration: unique partial index on `subscriptions(user_id, package_id)` where status IN ('active', 'pending').
  - [x] Update `SubscriptionService.RequestSubscription` to check for existing active/pending.
  - [x] Update `SubscriptionService.CreateSubscription` (Admin assign) to check for existing.
  - [x] Update `AvailablePackages.tsx`: check if package is already in `mySubscriptions` (active or pending) and disable "Request" button/text.
  - [x] Verify via test: `TestClientCannotRequestDuplicate`.

### [x] Task 43: Admin - Remove Package from Client
- **Goal:** Enable Admins to remove or cancel a client's active package subscription.
- **Steps:**
  - [x] Backend: Verify `DELETE /api/subscriptions/:id` or `PUT status='cancelled'` logic (prefer Cancel over Delete for history? or Delete for "mistake"?). Let's implement Soft Delete (Cancel) primarily, or hard Delete if just assigned.
     - *Decision*: Use `DELETE` endpoint for removal.
  - [x] Frontend (`AdminDashboard.tsx`): Update the Client List -> Packages column.
     - [x] Add an "X" or "Remove" button next to each package badge.
     - [x] Add confirmation modal ("Are you sure you want to remove package X from user Y?").
  - [x] Verify removal updates the list immediately.

## Bug Fix: Startup Race Condition
**Goal:** Prevent frontend from failing with proxy errors on startup.
- [x] Modify `start.sh` to wait for backend port 8080.
- [x] Verify fix by running `./start.sh` and checking logs.

## Bug Fix: Package Deletion Failure (FK Constraint & UI Hang)
**Goal:** Handle package deletion failures gracefully.
- [x] Backend: Update `DeletePackage` handler to return 409 Conflict if subscriptions exist.
- [x] Frontend: Add `onError` to `deletePackageMutation` to show error and close modal.
- [x] Verify: Attempt to delete a package with subscriptions; confirm user sees error message.

## Bug Fix: Package Deletion False Positive
**Goal:** Only count active/pending subscriptions when checking if package is in use.
- [x] Modify `CountSubscriptionsByPackage` query to filter by status.
- [x] Run `sqlc generate`.
- [x] Rebuild and verify.

## Bug Fix: Package Cascade Delete
**Goal:** Add ON DELETE CASCADE to subscriptions.package_id FK.
- [x] Create migration to alter FK constraint.
- [x] Remove service-level CountSubscriptionsByPackage check (now handled by DB).
- [x] Run migration and verify.

## Bug Fix: Backend Route Wildcard Conflict
**Goal:** Fix server panic caused by conflicting wildcard parameter names in Gin routes.
**Reference:** `bug-reports/20260113_backend_route_panic.md`
- [x] Update `RegisterRoutes` in `deliverables/handler.go`: Change route from `/:requestId/deliverables` to `/:id/deliverables`.
- [x] Update `RegisterClientRoutes` in `deliverables/handler.go`: Change route from `/:requestId/deliverables` to `/:id/deliverables`.
- [x] Update `ListByRequest` function in `deliverables/handler.go`: Change `c.Param("requestId")` to `c.Param("id")`.
- [x] Restart backend server and verify no panic occurs.
- [x] Test: Verify Admin can list deliverables for a request via API.
- [x] Test: Verify Client can list deliverables for a request via API.

## Feature: Deliverable Review Workflow
**Goal:** Implement a review workflow where Admins submit deliverables for client review, and Clients can approve or request revisions.

### Task 44: Database Schema Update (Deliverable Status)
- [x] Create migration `20260113114000_add_deliverable_status.sql`:
  - Add `deliverable_status` enum: `pending_review`, `approved`, `revision_requested`.
  - Add `status` column to `deliverables` table (default: `pending_review`).
  - Add `revision_notes` column (TEXT, nullable) for client feedback.
- [x] Run migration.
- [x] Run `sqlc generate`.

### Task 45: Backend - Status Update API
- [x] Add `UpdateDeliverableStatus` query in `deliverables.sql`.
- [x] Add `UpdateDeliverableStatus(ctx, id, status, notes)` in `service.go`.
- [x] Add `UpdateStatus` handler in `handler.go`:
  - `PUT /deliverables/:id/status` for Admin (resubmit).
  - `PUT /client/deliverables/:id/status` for Client (approve/revision).
- [x] Register routes in `main.go`.

### Task 46: Frontend - Admin Deliverable Status UI
- [x] Update `api/deliverables.ts`: Add `updateStatus(id, status)` to admin API.
- [x] Update `RequestDetailsModal.tsx` (Admin):
  - [x] Display status badge on each deliverable.
  - [x] Show "Resubmit" button for `revision_requested` items.
  - [x] Display client's `revision_notes` if present.

### Task 47: Frontend - Client Review UI
- [x] Update `api/deliverables.ts`: Add `updateStatus(id, status, notes?)` to client API.
- [x] Update `RequestDetailsModal.tsx` (Client):
  - [x] Display status badge on each deliverable.
  - [x] For `pending_review`: Show "Approve" and "Request Revision" buttons.
- [x] Create `RevisionNotesModal.tsx`: Modal for client to enter revision notes.

### Task 48: Deliverable Review Workflow Verification
- [x] **Automated:** Create `src/backend/tests/integration/deliverables_test.go`:
  - Admin can update status to `pending_review`.
  - Client can update status to `approved` or `revision_requested`.
  - Role enforcement (403 for unauthorized).
- [x] **Manual:**
  - [x] Admin uploads → Status shows "Pending Review".
  - [x] Client approves → Status shows "Approved".
  - [x] Client requests revision with notes → Status shows "Revision Requested".
  - [x] Admin resubmits → Status returns to "Pending Review".

### Task 49: Bug Fix - Request Completion Logic
- [x] **Analysis:** `UpdateStatus` only updated deliverable, didn't check if all deliverables were approved.
- [x] **Backend:**
  - Update `UpdateStatus` to check if all deliverables are approved.
  - If yes, auto-update request status to `completed`.
- [x] **Verification:**
  - Verified with new cycle (Admin upload -> Client approve -> Request auto-completes).

### Task 50: Data Migration - Auto-complete Approved Requests
- [x] **Context:** Fix existing requests that were approved before auto-complete logic.
- [x] **Migration:** Create `20260113124700_complete_approved_requests.sql` to update requests where all deliverables are approved.
- [x] **Execution:** Ran migration successfully.

### Task 51: Bug Fix - Dashboard Stats Wrong Data
- [x] **Issue:** "Pending Requests" stat showed 0 despite pending items.
- [x] **Analysis:** Stat was counting pending *subscriptions*, not *requests*. Also, requests query was disabled on Overview tab.
- [x] **Fix:**
  - Enable `activeTab` query on Overview.
  - Update stat to count `myRequests` filtered by `pending` status.


## Bug Fix: Client Status Always Active
**Goal:** Fix the hardcoded 'ACTIVE' status for clients.
**Reference:** 
- [x] Update  to conditionally render ACTIVE/INACTIVE based on subscription count.
- [x] Verify: Client with 0 packages shows INACTIVE. Client with 1+ shows ACTIVE.

## Bug Fix: Admin Stats Refresh
**Goal:** Ensure Admin Dashboard widgets (Active Clients, Pending Requests, MRR) update automatically after actions.
**Reference:** `bug-reports/20260113_144000_admin_stats_refresh.md`

### Task 52: Update Mutation Success Handlers
- **Goal:** Invalidate `['dashboardStats']` query key on relevant mutations.
- **Steps:**
  1. [x] Update `AdminDashboard.tsx`:
     - [x] `deleteMutation` (Client deletion)
     - [x] `approveMutation` (Subscription approval)
     - [x] `deleteSubscriptionMutation` (Subscription removal)
  2. [x] Update `AddClientForm.tsx`:
     - [x] `mutation` (Client creation)
  3. [x] Update `AssignPackageModal.tsx`:
     - [x] `mutation` (Package assignment)
  4. [x] Verify:
     - [x] Add a client -> "Active Clients" widget increments without refresh.
     - [x] Approve a request -> "Pending Requests" decrements, "Active Clients" (maybe) increments.

## Bug Fix: Missing Packages on Client Dashboard
**Goal:** Fix the issue where active subscriptions are not displaying on the Client Dashboard.
- [x] **Reproduction Test:**
    - [x] Create `src/backend/tests/integration/bug_client_subs_test.go`.
    - [x] Test setup: Create a client and a package, assign package to client.
    - [x] Test execution: Make a GET request to `/api/v1/subscriptions` as the client.
    - [x] Assertion: Verify the response body contains the package ID and status 'active'.
    - [x] Debug: Print the raw JSON response to confirm serialization format.
- [x] **Fix Implementation:**
    - [x] Identified Root Cause: Frontend was calling Admin-only endpoint `/api/v1/subscriptions` instead of Client endpoint `/api/v1/client/subscriptions`.
    - [x] Updated `src/frontend/src/api/subscriptions.ts` to include `listMySubscriptions` using the correct endpoint.
    - [x] Updated `src/frontend/src/pages/dashboards/ClientDashboard.tsx` to use `listMySubscriptions`.
- [x] **Verification:**
    - [x] Run the integration test and ensure it passes (after updating test to use correct endpoint).
    - [x] Manual Verification: Login as `user_a@test.local`, check Client Dashboard for "Test Package".

## Feature: Client Questionnaires (Onboarding)

### Task 53: Database Schema Update (Questionnaires)
- **Goal:** Add questionnaire tracking to users and create responses table.
- **Steps:**
  - [x] Create `20260115004900_create_questionnaires.sql` migration:
    - Add `has_completed_onboarding` (BOOLEAN) to `users`.
    - Add `has_completed_brand_questionnaire` (BOOLEAN) to `users`.
    - Create `questionnaire_responses` table (id, user_id, questionnaire_type, responses JSONB, timestamps).
  - [x] Run `goose up`.
  - [x] Create `questionnaires.sql` queries: `CreateQuestionnaireResponse`, `GetQuestionnaireResponse`, `UpdateQuestionnaireResponse`.
  - [x] Update `users.sql`: Add query to update onboarding status.
  - [x] Run `sqlc generate`.


### Task 54: Backend - Questionnaires Feature
- **Goal:** Implement questionnaire submission and retrieval API.
- **Dependencies:** Task 53
- **Steps:**
  - [x] Create `src/backend/internal/features/questionnaires/types.go`.
  - [x] Create `src/backend/internal/features/questionnaires/service.go`.
  - [x] Create `src/backend/internal/features/questionnaires/handler.go`:
    - `POST /api/v1/client/questionnaires/onboarding` - Submit onboarding questionnaire.
    - `GET /api/v1/client/questionnaires/onboarding` - Get onboarding response.
    - `GET /api/v1/client/questionnaires/onboarding/status` - Check completion status.
  - [x] Register routes in `main.go`.


### Task 55: Backend - Auth Response Update
- **Goal:** Include onboarding completion status in auth responses.
- **Dependencies:** Task 53
- **Steps:**
  - [x] Update `sessions.sql`: Add `has_completed_onboarding` to `GetSessionByToken` query.
  - [x] Run `sqlc generate`.
  - [x] Update `auth/interface.go`: Add `HasCompletedOnboarding` to `UserResponse`.
  - [x] Update `auth/service.go`: Populate `HasCompletedOnboarding` in `ValidateSession` and `userToResponse`.


### Task 56: Frontend - Type Updates
- **Goal:** Update frontend types to include onboarding status.
- **Dependencies:** Task 55
- **Steps:**
  - [x] Update `types/auth.ts`: Add `has_completed_onboarding?: boolean` to `User` interface.


### Task 57: Frontend - Questionnaires API
- **Goal:** Create API module for questionnaire operations.
- **Steps:**
  - [x] Create `src/frontend/src/api/questionnaires.ts`:
    - `submitOnboarding(data)` - Submit onboarding questionnaire.
    - `getOnboarding()` - Get current onboarding response.
    - `getOnboardingStatus()` - Check if onboarding is complete.


### Task 58: Frontend - Onboarding Guard Component
- **Goal:** Create routing guard for onboarding completion.
- **Dependencies:** Task 56
- **Steps:**
  - [x] Create `src/frontend/src/components/OnboardingGuard.tsx`.
  - [x] Component checks `user.has_completed_onboarding`.
  - [x] Redirect to `/onboarding` if not completed (client role only).


### Task 59: Frontend - Onboarding Questionnaire Page
- **Goal:** Create premium-styled "Getting to Know You" 6-step wizard form.
- **Form Structure:** 14 questions in 6 sections based on `docs/The 'Bite-Sized' Brand Intake Form.md`:
  - **Section 1 - The Origin (The "Why"):** The Spark, The Gap, The Purpose
  - **Section 2 - The Dream Client (The "Who"):** The Character, The Nightmare, The Happy Ending
  - **Section 3 - The Edge (The "What"):** The Superpower, The Non-Negotiable, The "So What?"
  - **Section 4 - The Vibe (The "Voice"):** The Celebrity, The Anti-Celebrity, The Outfit
  - **Section 5 - The Enemy (The "Positioning"):** The Villain, The Status Quo
  - **Section 6 - The Reality Check (The "Diagnostic"):** The Misconception, The Barrier
- **UX Requirements:**
  - 6 separate pages/steps (one per section) with Back/Next navigation
  - All fields required — cannot proceed to next step with empty fields
  - No progress saving (if browser closes, user starts over)
- **Dependencies:** Task 57, Task 58
- **Steps:**
  - [x] Create `src/frontend/src/pages/questionnaires/OnboardingQuestionnaire.tsx`.
  - [x] Design 6-step wizard with progress indicator and modern aesthetics.
  - [x] Implement per-step validation (all fields required before Next).
  - [x] Store responses in JSONB-compatible nested structure (in React state only).
  - [x] Submit via `questionnairesApi.submitOnboarding()` on final step.
  - [x] On success: Invalidate auth query, redirect to `/dashboard`.


### Task 60: Frontend - Routing Updates
- **Goal:** Integrate onboarding flow into app routing.
- **Dependencies:** Task 58, Task 59
- **Steps:**
  - [x] Update `App.tsx`:
    - Add `/onboarding` route (protected, clients only).
    - Update `RootRedirect` to check `has_completed_onboarding` for clients.
  - [x] Wrap `ClientDashboard` route with `OnboardingGuard`.


### Task 61: Questionnaires Feature Verification
- **Goal:** Verify the complete onboarding flow.
- **Testing:**
  - [x] **Automated:** Create `src/backend/tests/integration/questionnaires_test.go`:
    - Client can submit onboarding questionnaire.
    - User's `has_completed_onboarding` is updated.
    - Non-clients cannot access questionnaire endpoints.
  - [ ] **Manual:**
    - Create new client via Admin Dashboard.
    - Log in as new client → Verify redirect to `/onboarding`.
    - Complete questionnaire → Verify redirect to `/dashboard`.
    - Log out and back in → Verify direct dashboard access.
    - Log in as Admin/Superadmin → Verify no onboarding redirect.

---

## Bug Fix: First Questionnaire Submit Redirect Failure

**Report:** `bug-reports/20260115_questionnaire_first_submit_redirect_failure.md`
**Root Cause:** Race condition - `navigate('/dashboard')` executes before `refetchQueries()` completes, causing `OnboardingGuard` to see stale data.

### Task 62: Fix Questionnaire Submit Redirect Race Condition
- **Goal:** Ensure navigation only happens after auth data is refreshed.
- **Steps:**
  - [x] Update `src/frontend/src/pages/questionnaires/OnboardingQuestionnaire.tsx`:
    - Change `onSuccess` handler to use `async` and `await queryClient.refetchQueries()` before `navigate()`. (Use Skill: react-component)
  - [x] Manual verification:
    - Create new client, fill questionnaire, submit → Should redirect to `/dashboard` on first try.


