# Feature: Core Infrastructure

## Task 1: Backend Project Scaffolding
**Goal:** Initialize the Go backend with Gin router and database tooling.
- [x] Create `src/backend/` directory structure per `specs/ARCHITECTURE.md`. (Use Skill: create_go_feature)
- [x] Initialize Go module (`go mod init`).
- [x] Configure `sqlc` with `sqlc.yaml`.
- [x] Configure `goose` for migrations in `internal/database/migrations/`.
- [x] Create `docker-compose.yml` for PostgreSQL.
- [x] Create `cmd/server/main.go` entry point with Gin router.

## Task 2: Frontend Project Scaffolding
**Goal:** Initialize the React frontend with Vite and React Router.
- [x] Initialize Vite + React + TypeScript (`npx create-vite`). (Use Skill: create_react_feature)
- [x] Configure Vite proxy for API requests (`vite.config.ts`).
- [x] Set up React Router in `src/app/routes.tsx`.
- [x] Create `src/lib/api.ts` with Axios instance.
- [x] Install TanStack Query and configure `QueryClientProvider`.

---

# Feature: Authentication System

## Task 3: Database Schema (Users)
**Goal:** Create the `users` table.
- [ ] Create `goose` migration: `users` table (id, email, password_hash, created_at). (Use Skill: db-migration)
- [ ] Run `goose up`.
- [ ] Create `sqlc` queries: `CreateUser`, `GetUserByEmail`, `GetUserByID`.
- [ ] Run `sqlc generate`.

## Task 4: Backend - Auth Feature
**Goal:** Implement user registration, login, and session management.
**Dependencies:** Task 3
- [ ] Create `src/backend/internal/features/auth/interface.go`. (Use Skill: create_go_feature)
- [ ] Create `src/backend/internal/features/auth/service.go`:
  - `Register(email, password)`: Hash password, create user.
  - `Login(email, password)`: Verify password, create session.
  - `Logout(sessionToken)`: Invalidate session.
  - `ValidateSession(token)`: Return user.
- [ ] Create `src/backend/internal/features/auth/handler.go`:
  - `POST /api/v1/auth/register`
  - `POST /api/v1/auth/login`
  - `POST /api/v1/auth/logout`
  - `GET /api/v1/auth/me`
- [ ] Implement Session middleware (HttpOnly cookies).
- [ ] Implement CSRF protection (Double Submit Cookie).
- [ ] Register routes in `main.go`.

## Task 5: Frontend - Auth Feature
**Goal:** Implement login/register forms and auth context.
**Dependencies:** Task 4
- [ ] Create `src/frontend/src/features/auth/types/index.ts`. (Use Skill: create_react_feature)
- [ ] Create `src/frontend/src/features/auth/hooks/useAuth.ts`:
  - `useLogin()`, `useRegister()`, `useLogout()`, `useCurrentUser()`.
- [ ] Create `src/frontend/src/features/auth/components/LoginForm.tsx`.
- [ ] Create `src/frontend/src/features/auth/components/RegisterForm.tsx`.
- [ ] Create `src/frontend/src/contexts/AuthContext.tsx`.
- [ ] Create `ProtectedRoute` component.
- [ ] Update `App.tsx` with routes: `/login`, `/register`, `/dashboard`.

## Task 6: Auth Verification
**Goal:** Verify authentication flow.
- [ ] **Automated:** Create `src/backend/tests/integration/auth_test.go`:
  - Register user -> Verify 201.
  - Login -> Verify HttpOnly cookie set.
  - Access `/auth/me` with cookie -> Verify user returned.
  - Access `/auth/me` without cookie -> Verify 401.
- [ ] **Manual:**
  - Register new user.
  - Login -> Verify redirect to dashboard.
  - Refresh page -> Verify session persists.
  - Logout -> Verify redirect to login.

---

# Feature: Dashboard Layout

## Task 7: Frontend - Dashboard Shell
**Goal:** Create the main dashboard layout with navigation.
- [ ] Create `src/frontend/src/components/DashboardLayout.tsx`. (Use Skill: create_react_feature)
- [ ] Implement responsive sidebar/navbar.
- [ ] Add global loading indicator.
- [ ] Integrate with React Router (SPA navigation).

---

# Feature: Account Management

## Task 8: Database Schema (Accounts)
**Goal:** Create the `accounts` table.
- [ ] Create `goose` migration: `accounts` table (id, user_id, name, type, currency, balance). (Use Skill: db-migration)
- [ ] Run `goose up`.
- [ ] Create `sqlc` queries: `CreateAccount`, `ListAccountsByUser`, `GetAccount`, `UpdateAccount`, `DeleteAccount`.
- [ ] Run `sqlc generate`.

## Task 9: Backend - Accounts Feature
**Goal:** Implement account CRUD with Row-Level Security.
**Dependencies:** Task 8, Task 4 (Auth)
- [ ] Create `src/backend/internal/features/accounts/interface.go`. (Use Skill: create_go_feature)
- [ ] Create `service.go`:
  - `Create(userID, name, type, currency, balance)`
  - `List(userID)` - **MUST filter by user_id**
  - `Get(userID, accountID)` - Verify ownership
  - `Update(userID, accountID, ...)`
  - `Delete(userID, accountID)`
- [ ] Create `handler.go`:
  - `POST /api/v1/accounts`
  - `GET /api/v1/accounts`
  - `GET /api/v1/accounts/:id`
  - `PUT /api/v1/accounts/:id`
  - `DELETE /api/v1/accounts/:id`
- [ ] Register routes (protected by session middleware).

## Task 10: Frontend - Accounts Feature
**Goal:** UI for viewing and managing accounts.
**Dependencies:** Task 9
- [ ] Create `src/frontend/src/features/accounts/types/index.ts`. (Use Skill: create_react_feature)
- [ ] Create `hooks/useAccounts.ts`: `useAccounts()`, `useCreateAccount()`.
- [ ] Create `components/AccountList.tsx`.
- [ ] Create `components/AccountForm.tsx` (with Currency dropdown).
- [ ] Add "Accounts" section to Dashboard.

## Task 11: Account Management Verification
**Goal:** Verify account CRUD and data isolation.
- [ ] **Automated:** Create `src/backend/tests/integration/accounts_test.go`:
  - User A creates account -> Verify 201.
  - User A lists accounts -> See own account only.
  - User B tries to access User A's account -> Verify 404 or 403.
- [ ] **Manual:**
  - Create accounts (Liquid: "Cash", Asset: "Stocks").
  - Verify list updates.
  - Edit account -> Verify change.
  - Delete account -> Verify removed.

---

# Feature: Categories

## Task 12: Database Schema (Categories)
**Goal:** Create the `categories` table.
- [ ] Create `goose` migration: `categories` table (id, user_id, name, tag, is_investment). (Use Skill: db-migration)
- [ ] Run `goose up`.
- [ ] Create `sqlc` queries: `CreateCategory`, `ListCategoriesByUser`, `UpdateCategory`, `DeleteCategory`.
- [ ] Run `sqlc generate`.

## Task 13: Backend - Categories Feature
**Goal:** Implement category CRUD with Need/Want/Savings tags.
**Dependencies:** Task 12
- [ ] Create `src/backend/internal/features/categories/`. (Use Skill: create_go_feature)
- [ ] Implement service and handler (pattern same as Accounts).
- [ ] Register routes.

## Task 14: Frontend - Categories Feature
**Goal:** UI for managing categories.
**Dependencies:** Task 13
- [ ] Create `src/frontend/src/features/categories/`. (Use Skill: create_react_feature)
- [ ] Implement hooks and components (pattern same as Accounts).
- [ ] Add Tag selector (Need/Want/Savings) to form.

---

# Feature: Transaction Engine

## Task 15: Database Schema (Transactions)
**Goal:** Create the `transactions` table.
- [ ] Create `goose` migration: `transactions` table (id, user_id, amount, date, note, type, source_account_id, destination_account_id, category_id). (Use Skill: db-migration)
- [ ] Run `goose up`.
- [ ] Create `sqlc` queries: `CreateTransaction`, `ListTransactionsByUser`, `DeleteTransaction`.
- [ ] Run `sqlc generate`.

## Task 16: Backend - Transactions Feature
**Goal:** Implement Expense, Income, and Transfer logic.
**Dependencies:** Task 15, Task 9 (Accounts)
- [ ] Create `src/backend/internal/features/transactions/`. (Use Skill: create_go_feature)
- [ ] Implement service:
  - **Expense:** Deduct from `source_account.balance`.
  - **Income:** Add to `destination_account.balance`.
  - **Transfer:** Deduct from source, add to destination (zero net worth change).
- [ ] Create handler: `POST /api/v1/transactions`, `GET /api/v1/transactions`, `DELETE /api/v1/transactions/:id`.
- [ ] **Critical:** Wrap balance updates + transaction creation in a database transaction.

## Task 17: Frontend - Transactions Feature
**Goal:** UI for logging and viewing transactions.
**Dependencies:** Task 16
- [ ] Create `src/frontend/src/features/transactions/`. (Use Skill: create_react_feature)
- [ ] Create `LogTransactionForm.tsx` with:
  - Type selector (Expense/Income/Transfer).
  - Account dropdowns (Source/Destination based on type).
  - Category dropdown.
  - Amount & Date.
- [ ] Create `TransactionHistory.tsx` (filtered by month).
- [ ] Add "Transactions" tab to Dashboard.

## Task 18: Transaction Engine Verification
**Goal:** Verify transaction logic and balance updates.
- [ ] **Automated:**
  - Create Expense -> Verify source account balance decreases.
  - Create Income -> Verify destination account balance increases.
  - Create Transfer -> Verify source decreases, destination increases, net worth unchanged.
  - Delete transaction -> Verify balances revert.
- [ ] **Manual:** Full flow test via UI.

---

# Feature: Bill & Debt Tracker

## Task 19: Database Schema (Bills)
**Goal:** Create the `bills` table.
- [ ] Create `goose` migration: `bills` table (id, user_id, name, amount, due_date, is_recurring, status). (Use Skill: db-migration)
- [ ] Run `goose up`.
- [ ] Create `sqlc` queries.
- [ ] Run `sqlc generate`.

## Task 20: Backend - Bills Feature
**Goal:** Implement bill CRUD and "Mark as Paid" action.
**Dependencies:** Task 19, Task 16 (Transactions)
- [ ] Create `src/backend/internal/features/bills/`. (Use Skill: create_go_feature)
- [ ] Implement service:
  - `MarkAsPaid(billID)`: Create Expense transaction, set bill status to `paid`.
- [ ] Create handler:
  - `POST /api/v1/bills`, `GET /api/v1/bills`, `PUT /api/v1/bills/:id`, `DELETE /api/v1/bills/:id`.
  - `POST /api/v1/bills/:id/pay` (Mark as Paid).

## Task 21: Frontend - Bills Feature
**Goal:** UI for viewing, creating, and paying bills.
**Dependencies:** Task 20
- [ ] Create `src/frontend/src/features/bills/`. (Use Skill: create_react_feature)
- [ ] Create `BillList.tsx` (sorted by due date, pending first).
- [ ] Create `BillForm.tsx`.
- [ ] Add "Pay" button that calls `/pay` endpoint.
- [ ] Add "Commitments" section to Dashboard.

## Task 22: Safe-to-Spend Calculation
**Goal:** Display Safe-to-Spend on Dashboard.
**Dependencies:** Task 21, Task 9 (Accounts)
- [ ] Backend: Create `GET /api/v1/dashboard/summary` endpoint returning:
  - Total Liquid Balance.
  - Sum of Pending Bills for current month.
  - Safe-to-Spend = Liquid - Pending Bills.
- [ ] Frontend: Create `SafeToSpendWidget.tsx` on Dashboard.

## Task 23: Bill Tracker Verification
**Goal:** Verify bill payment and Safe-to-Spend.
- [ ] **Automated:**
  - Create bill -> Verify Safe-to-Spend decreases.
  - Pay bill -> Verify Expense transaction created, account balance decreases.
- [ ] **Manual:** Full flow test via UI.

---

# Feature: Investment Manager

## Task 24: Investment Tracking Logic
**Goal:** Ensure Investment accounts update only via Transfers.
**Dependencies:** Task 16 (Transactions)
- [ ] Verify: Transfers to `Asset` type accounts correctly update balance.
- [ ] Verify: No direct "Income" can be added to Asset accounts (optional rule, or UI guidance).
- [ ] Add Net Worth widget to Dashboard (Liquid + Assets).

---

## Verification: End-to-End

### Task 25: Full User Journey Test
**Goal:** Verify the complete happy path.
- [ ] **Manual:**
  1. Register new user.
  2. Create accounts: "Bank (Liquid)", "Cash (Liquid)", "Stocks (Asset)".
  3. Add income to Bank.
  4. Log expense from Cash (category: Food, tag: Want).
  5. Transfer from Bank to Stocks.
  6. Create bill due next week.
  7. Verify Safe-to-Spend.
  8. Pay bill -> Verify balance updates.
  9. Check Net Worth (should include Stocks).
