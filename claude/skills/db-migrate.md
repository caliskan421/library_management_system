# Skill: Database Migration

## Description
Run or verify database migrations for LibraNet.

## How Migrations Work
LibraNet uses Bun's `CreateTable().IfNotExists()` for automatic schema creation.
Tables are created on server startup in `database.Migrate()`.

## Tables
| Table         | Model       | Key Columns                              |
|---------------|-------------|------------------------------------------|
| users         | model.User  | id (uuid pk), name, email (unique), password, role |
| books         | model.Book  | id (uuid pk), title, author, isbn (unique), total_copies, available_copies |
| reservations  | model.Reservation | id (uuid pk), user_id (fk), book_id (fk), status, due_date |

## Manual Steps
```bash
# Create database
createdb libranet

# Run server (auto-migrates)
cd backend && make run

# Or just build and run migration manually by starting the server once
```

## Token Optimization
- Don't re-read migration code if only adding a new column - just check the model file
- Table creation is idempotent (IfNotExists), safe to run multiple times
