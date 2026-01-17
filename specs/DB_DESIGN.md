# Database Design

## 1. Overview
The database uses PostgreSQL. All tables are designed to strictly enforce Row-Level Security principles by including `user_id` as a foreign key where applicable.

---

## 2. Entity Relationship Diagram

```mermaid
erDiagram
    Users ||--o{ Accounts : has
    Users ||--o{ Transactions : has
    Users ||--o{ Bills : has
    Users ||--o{ Categories : has
    
    Accounts ||--o{ Transactions : source
    Accounts ||--o{ Transactions : destination
    Categories ||--o{ Transactions : classifies

    Users {
        uuid id PK
        string email
        string password_hash
        timestamp created_at
    }

    Accounts {
        uuid id PK
        uuid user_id FK
        string name
        string type "Liquid, Asset"
        string currency "ISO Code"
        decimal balance
    }

    Categories {
        uuid id PK
        uuid user_id FK
        string name
        string tag "Need, Want, Savings"
        boolean is_investment
    }

    Transactions {
        uuid id PK
        uuid user_id FK
        decimal amount
        timestamp date
        string note
        string type "Expense, Income, Transfer"
        uuid source_account_id FK "Nullable"
        uuid destination_account_id FK "Nullable"
        uuid category_id FK
    }

    Bills {
        uuid id PK
        uuid user_id FK
        string name
        decimal amount
        date due_date
        boolean is_recurring
        string status "Pending, Paid"
    }
```

---

## 3. Schema Definitions

### 3.1 Users
| Column | Type | Constraints | Description |
|---|---|---|---|
| `id` | UUID | PK | Unique identifier |
| `email` | VARCHAR | Unique, Not Null | User email |
| `password_hash` | VARCHAR | Not Null | Bcrypt hash |
| `created_at` | TIMESTAMPTZ | Default Now() | Account creation time |

### 3.2 Accounts
| Column | Type | Constraints | Description |
|---|---|---|---|
| `id` | UUID | PK | Unique identifier |
| `user_id` | UUID | FK -> Users.id | Owner |
| `name` | VARCHAR | Not Null | E.g., "BPI", "Cash" |
| `type` | VARCHAR | Not Null | `liquid` or `asset` |
| `currency` | VARCHAR | Not Null | ISO 4217 Code (e.g. USD, PHP) |
| `balance` | DECIMAL | Not Null | Current balance |

### 3.3 Categories
| Column | Type | Constraints | Description |
|---|---|---|---|
| `id` | UUID | PK | Unique identifier |
| `user_id` | UUID | FK -> Users.id | Owner |
| `name` | VARCHAR | Not Null | E.g., "Food", "Rent" |
| `tag` | VARCHAR | Not Null | `need`, `want`, `savings` |
| `is_investment` | BOOLEAN | Default False | If true, triggers transfer logic |

### 3.4 Transactions
| Column | Type | Constraints | Description |
|---|---|---|---|
| `id` | UUID | PK | Unique identifier |
| `user_id` | UUID | FK -> Users.id | Owner |
| `amount` | DECIMAL | Not Null | Transaction value |
| `date` | TIMESTAMPTZ | Not Null | Transaction time |
| `note` | TEXT | | Optional description |
| `type` | VARCHAR | Not Null | `expense`, `income`, `transfer` |
| `source_account_id` | UUID | FK -> Accounts.id | Nullable (for income) |
| `destination_account_id` | UUID | FK -> Accounts.id | Nullable (for expense) |
| `category_id` | UUID | FK -> Categories.id | Classification |

### 3.5 Bills
| Column | Type | Constraints | Description |
|---|---|---|---|
| `id` | UUID | PK | Unique identifier |
| `user_id` | UUID | FK -> Users.id | Owner |
| `name` | VARCHAR | Not Null | Bill name |
| `amount` | DECIMAL | Not Null | Amount due |
| `due_date` | DATE | Not Null | Due date recurrence |
| `is_recurring` | BOOLEAN | Default False | |
| `status` | VARCHAR | Default 'pending' | `pending`, `paid` |
