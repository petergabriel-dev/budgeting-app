# Features

This document tracks all implemented and planned features, their technical requirements, and verification strategies.

---

## Core Infrastructure

- [ ] **Project Scaffolding**
  - [ ] Initialize Go module in `src/backend`
  - [ ] Initialize React/Vite project in `src/frontend`
  - [ ] Configure Vite proxy for API requests
  - [ ] Set up `sqlc` for database queries
  - [ ] Set up `sqlc` for database queries
  - [ ] Set up React Router for client-side routing

- [ ] **Authentication System**
  - [ ] User registration
  - [ ] User login (HttpOnly Session Cookies)
  - [ ] Logout functionality
  - [ ] Session middleware
  - [ ] CSRF protection (Double Submit Cookie)

- [ ] **Dashboard Layout**
  - [ ] Responsive sidebar/navbar
  - [ ] SPA navigation (React Router)
  - [ ] Global loading indicator

---

### 1. Account Management (Data Isolation)
- [ ] List user's accounts (Liquid & Asset)
- [ ] Create new account (Name, Type, Currency, Opening Balance)
- [ ] Edit account details
- [ ] Row-Level Security: Ensure users only see their own accounts

### 2. Transaction Engine
- [ ] **Expense**: Deduct from Source Account
- [ ] **Income**: Add to Destination Account
- [ ] **Transfer**: Move from any Source to Destination (Zero Net Worth Change)
- [ ] Transaction History View (Filtered by Month)

### 3. Bill & Debt Tracker
- [ ] Create recurring/one-time bills
- [ ] Dashboard "Commitments" List (Sorted by due date)
- [ ] "Mark as Paid" Action (Converts Bill -> Expense Transaction)
- [ ] Safe-to-Spend Calculation (Liquid Cash - Pending Bills)

### 4. Categories
- [ ] Manage custom categories (with Need/Want/Savings tags)
- [ ] "Is Investment" toggle for legacy tracking

## Planned Features

> Features will be added here as they are planned and implemented.