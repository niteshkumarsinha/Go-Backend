# Chapter 8: PostgreSQL Integration

**Directory:** `src/8-postgres`

This chapter introduces persistence by connecting the Go application to a PostgreSQL database using the `database/sql` package and the `lib/pq` driver. It also covers containerizing the application and database with Docker.

### Key Concepts

- **`database/sql`**: Go's standard interface for communicating with SQL databases.
- **`lib/pq`**: A pure Go Postgres driver for the `database/sql` package. Imported anonymously (`_ "github.com/lib/pq"`) to register itself.
- **`sql.Open`**: Opens a database connection (lazily).
- **`db.Query`**: executing queries that return rows (e.g., `SELECT`).
- **`db.QueryRow`**: Executing a query expected to return at most one row.
- **`db.Exec`**: Executing queries that don't return rows (e.g., `INSERT`, `UPDATE`, `DELETE`).
- **Refining API Handlers**: Implementing full CRUD (Create, Read, Update, Delete) handlers backed by a real database.
- **Docker Compose**: Orchestrating the Go application and PostgreSQL database services together. Setting up persistent volumes and environment variables.

### Example

```go
// Connecting to the database
connStr := "host=postgres user=postgres password=password dbname=testdb sslmode=disable"
db, err := sql.Open("postgres", connStr)

// Querying data
rows, err := db.Query("SELECT id, name FROM users")
for rows.Next() {
    var u User
    rows.Scan(&u.ID, &u.Name)
    // ...
}

// Inserting data
db.Exec("INSERT INTO users (name) VALUES ($1)", "John Doe")
```

### API Testing

You can interact with the API using `curl`.

**1. Create a User**
```bash
curl -X POST http://localhost:8080/users/create \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice"}'
```

**2. List All Users**
```bash
curl http://localhost:8080/users
```

**3. Update a User**
```bash
# Replace 1 with the actual ID returned from creation
curl -X PUT http://localhost:8080/users/update \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "name": "Alice Updated"}'
```

**4. Delete a User**
```bash
# Replace 1 with the actual ID
curl -X DELETE http://localhost:8080/users/delete \
     -H "Content-Type: application/json" \
     -d '{"id": 1}'
```

