// REST: Server-defined contract (The client gets what the server gives)
// GET /users/1
{
  "id": "1",
  "name": "John Doe",
  "email": "john@example.com",
  "posts": [...] // Fixed structure
}

// GraphQL: Client-defined contract (The client requests what it needs)
query {
  user(id: "1") {
    name
    posts {
      title
    }
  }
}
