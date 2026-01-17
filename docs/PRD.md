# Product Requirement Document: "Budget Dashboard"

**Date:** January 08, 2026
**Platform:** Web-Only (Desktop/Tablet Dashboard)
**Currency:** Multi-currency (User defines currency per account)
**Core Philosophy:** Manual Entry, Privacy-First, Strict User Isolation.

---

## 1. Project Overview

A web-based personal finance dashboard designed to track Net Worth, Cash Flow, and Budget Adherence. It focuses on intentional spending by manually logging transactions and classifying them as **Needs**, **Wants**, or **Savings**.

### Key Differentiators

* **Privacy & Isolation:** Each user has a private, secure dashboard. Data is never shared between profiles.
* **Safety:** "Safe-to-Spend" calculator subtracts upcoming bills from current cash.
* **Simplicity:** Investments are tracked by "Cash In" (Cost Basis), not market fluctuations.
* **Discipline:** Paying off debt is strictly categorized as a "Need."

---

## 2. Authentication & Security (New)

To ensure privacy and data integrity, the app requires a secure authentication layer.

* **Auth System:**
* **Sign Up:** Email & Password required.
* **Login:** Secure session management (e.g., JWT or Secure Cookies).
* **Password Security:** Passwords must be hashed (e.g., bcrypt) before storage. Never store plain text passwords.


* **Data Isolation Rule (Row-Level Security):**
* Every database query must be scoped to the authenticated user.
* *Technical Rule:* All queries for Accounts, Transactions, and Bills must include `WHERE user_id = current_user.id`.
* A user **cannot** see, edit, or access another user's wallets or transactions.



---

## 3. User Experience Flow

### A. Authentication (The Gate)

1. **Landing Page:** Simple login/signup form.
2. **Verification:** (Optional for MVP) Email verification to prevent bot accounts.
3. **Session Start:** Upon successful login, the system retrieves *only* that user's specific data.

### B. Onboarding (First Run)

1. **Wallet Setup:** User manually creates accounts:
* *Input:* Name, Type, **Currency** (e.g., USD, PHP).
* *Liquid:* "BPI Payroll", "GCash", "Physical Wallet".
* *Assets:* "COL Financial", "Pag-IBIG MP2".


3. **Opening Balances:** User inputs the current balance for each account.


### C. The "Happy Path" (Daily Use)

1. **Check Status:** User logs in to see their specific "Safe-to-Spend" amount.
2. **Log Activity:** User clicks "Add Transaction" or checks off a pending bill.
3. **Immediate Feedback:** Dashboards update instantly.

---

## 4. Functional Modules

### Module 1: The Dashboard (Home)

The interface is divided into four main quadrants.

| Component | Functionality |
| --- | --- |
| **Global Header** | Displays **Total Net Worth** (Liquid + Assets) and **Current Month**. |
| **Safe-to-Spend** | Calculation: `(Total Liquid Cash) - (Sum of Pending Bills for Month)`. |
| **Commitments** | A checklist of unpaid bills sorted by Due Date. |
| **Account List** | Grouped by type: **Liquid** (Spendable) vs. **Investments** (Saved). |

### Module 2: Transaction Engine (The Core)

Every transaction must have the following logic applied:

* **Expense:**
* *Input:* Amount, Source (e.g., GCash), Category (e.g., Food).
* *Logic:* Deducts from Source.


* **Income:**
* *Input:* Amount, Destination (e.g., BPI), Source (e.g., Salary).
* *Logic:* Increases Source balance.


* **Transfer:**
* *Trigger:* User selects "Transfer" action.
* *Input:* Source Accounts (e.g., Bank) -> Destination Account (e.g., Cash).
* *Logic:* **Zero Net Worth Change.** Moves money from Source to Destination.



### Module 3: Bill & Debt Tracker

* **Commitment Entry:** User adds a bill with `Name`, `Amount`, `Due Date`.
* **Status:** Default is `Pending`.
* **Payment Action:** User clicks "Mark as Paid."
* System converts bill to an **Expense Transaction**.
* **Category:** Debt Repayment.



### Module 4: Investment Manager (Manual)

* **Valuation Model:** Cost Basis (Cash Flow Based).
* **Updates:** The balance of an investment account only changes when the user *transfers money into it*.
* **No Live Data:** The app does not fetch stock prices. It reflects "How much I have saved," not "Current Market Value."

---

## 5. Data Architecture (Schema Logic)

To support the features above, **every table must link back to the User**.

### A. Tables

1. **Users:**
* `ID` (Primary Key).
* `Email`, `PasswordHash`, `CreatedAt`.


2. **Accounts:**
* `ID`, **`UserID`** (Foreign Key).
* `Name`, `Balance`, `Currency` (ISO Code).
* `Type`: **Liquid** (Cash/Bank/E-wallet) OR **Asset** (Investment).


3. **Categories:**
* `ID`, **`UserID`** (Foreign Key).
* `Name` (e.g., "Groceries", "Debt").
* `Tag`: **Need**, **Want**, or **Savings/Investments** (For reporting).
* `IsInvestment`: Boolean (True triggers Transfer logic).
* *(Note: You can include a default set of system categories, but users should be able to make their own).*


4. **Transactions:**
* `ID`, **`UserID`** (Foreign Key).
* `Amount`, `Date`, `Note`.
* `Type`: Expense / Income / Transfer.
* `SourceAccountID`, `DestinationAccountID` (nullable).
* `CategoryID`.


5. **Bills:**
* `ID`, **`UserID`** (Foreign Key).
* `Name`, `Amount`, `DueDate`.
* `IsRecurring` (Boolean).
* `Status`: Pending / Paid.



---

## 6. Technical Constraints & Logic Rules

1. **Auth Middleware:** All API routes (except Login/Register) must be protected by authentication middleware.
2. **Currency Formatting:** Display amounts with their respective account's currency symbol/code.
3. **Negative Balances:** Accounts can go negative (e.g., Overdraft), but the UI should warn the user.
4. **Date Handling:** All "Month Views" reset on the 1st of the month.
5. **Edit/Delete:** Users can edit past transactions. If a transaction is deleted, the account balances must revert automatically.