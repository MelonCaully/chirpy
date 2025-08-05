# 🐦 Chirpy 

A lightweight microblogging backend built with Go, designed to be fast, secure, and testable. Chirpy supports user authentication, refresh tokens, and chirp creation — making it ideal for learning backend architecture or powering minimalist web apps.

The project is structured for clarity and modularity, with separation between routing, authentication, and database layers.

## 🎯 Motivation

Writing backend services should be a clean, hands-on experience. Chirpy is designed to help you:

- 🛠️ Practice building REST APIs with Go and PostgreSQL

- 🔐 Implement secure user login with hashed passwords and refresh tokens

- 🐦 Create, read, update, and delete “chirps” — short messages posted by users

- 📊 Explore service metrics, readiness endpoints, and middleware patterns

Whether you’re learning Go, prepping for interviews, or building a foundation for bigger systems, Chirpy is built to help you go from idea to implementation quickly.

## ⚙️ Installation & Setup

### Prerequisites

Install the following before running the project:

- [Go](https://go.dev/dl/) — for building and running the app

- [PostgreSQL](https://www.postgresql.org/download/) — for database support

- [Goose](https://github.com/pressly/goose) — for database migrations

- [sqlc](https://github.com/kyleconroy/sqlc) — for generating type-safe Go from SQL

## 🛠️ Configuration

### 1. Clone the Project

```bash
git clone https://github.com/your-username/chirpy
cd chirpy
```

### 2. Set Up Enviroment

Create a .env file or configure envirement variables as needed for:

- DB_URL
- PLATFORM
- JWT_SECRET
- POLKA_KEY

### 3. Run Migrations

```bash
goose -dir sql/schema postgres "$DATABASE_URL" up
```

### 4. Launch the Server

```bash
go run .
```
The app should now be running on http://localhost:8080

# 🚀 Features

✅ User registration and login with hashed passwords

🔄 Refresh token support and revocation

🐦 Chirp creation, retrieval, filtering, and deletion

🔐 JWT-based authentication middleware

🔴 Readiness and metrics endpoints

📜 Database migrations with Goose

💬 Fully typed SQL queries with sqlc

# 🧪 Testing

Testing is currently done manually or via CLI test suites other than auth.go which has unit tests.

Planned future support:

✅ Unit tests for database logic

✅ Integration tests for endpoints

# 🤝 Contributing

If you'd like to help improve this project:

```bash
git clone https://github.com/MelonCaully/chirpy
cd chirpy
```

Then fork and submit a pull request to the main branch.

Contributions to database optimization, authentication security, or new features (e.g., follow system, timeline feed) are always welcome!