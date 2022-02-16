### Clean Architecture Principles in REST API
An Implementation of Clean Architecture principles in REST API server with Go.

**Fundamentals of Clean Architecture:**
- Independent of frameworks.
- Testable 
- Independent of UI.
- Independent of database.
- Independent of any external agency.

### API Reference

#### Get all employees

```http
  GET /employees
```


#### Add an employee

```http
  POST /employees
```

### Run

Clone the project

```bash
  git clone https://github.com/hmsayem/clean-architecture-implementation.git
```
Go to the project directory

```bash
  cd clean-architecture-implementation
```

Copy all third-party dependencies to vendor folder.

```bash
  go mod vendor
```

Export environment variables.

```bash
GOOGLE_APPLICATION_CREDENTIALS=/path/to/project-private-key.json

SERVER_PORT=:8000
```

Start the server.

```bash
  go run .
```

### Docker

Build image.

```bash
docker build -t rest-server .
```
Run container.

```bash
 docker run  --mount type=bind,source=/path/to/project-private-key.json,target=/run/secrets/project-private-key.json,readonly -p 8000:8000 rest-server
```

## Example
```bash
❯ curl -XGET  "http://localhost:8000/employees" | jq
[
  {
    "id": 1,
    "name": "Masudur Rahman",
    "title": "Senior Software Engineer",
    "team": "ByteBuilders",
    "email": "masud@appscode.com"
  },
  {
    "id": 2,
    "name": "Kamol Hasan",
    "title": "Senior Software Engineer",
    "team": "KubeDB",
    "email": "kamol@appscode.com"
  },
  {
    "id": 3,
    "name": "Alif Biswas",
    "title": "Software Engineer",
    "team": "KubeDB",
    "email": "alif@appscode.com"
  },
  {
    "id": 4,
    "name": "Piyush Kanti Das",
    "title": "Software Engineer",
    "team": "Stash",
    "email": "piyush@appscode.com"
  }
]

```

