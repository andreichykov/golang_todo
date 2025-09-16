****Project: TODO API (Go)****


🎯 **Goal**

Build a simple REST API in Go for managing TODO list items.

• Each TODO item has: id, title, description, status (e.g., pending/done), createdAt, updatedAt.

• Core functionality: create, get, update, delete TODO items.

• Future extensions: support multiple users & a client app.


✅ **Requirements**

1. **Basic API (MVP)**

• API should be built using Go (preferably net/http, gorilla/mux, or chi).

• Provide REST endpoints for TODO items:

- POST /todos → Create a new TODO.

- GET /todos → List all TODOs.

- GET /todos/{id} → Get a single TODO by ID.

- PUT /todos/{id} → Update a TODO.

- DELETE /todos/{id} → Delete a TODO.

• Responses must be in JSON.

2. **Data Storage**

• For MVP, use in-memory storage (map or slice).

• Later, replace with database (SQLite or PostgreSQL).

3. **Code Structure**

• Split code into packages:

- models → TODO struct & validation.

- handlers → HTTP handlers.

- store → storage implementation (in-memory first).

- main.go → server entrypoint.

4. **Future Features (not in MVP)**

• Multiple users (each TODO belongs to a user).

• Authentication (JWT or sessions).

• Client app (web or mobile) that consumes the API.
