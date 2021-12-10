## Employee API Server 
Implementation of **Clean Architecture** principles in REST API Server built with Go.

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
  git clone git@github.com:hmsayem/employee-server.git
```
Go to the project directory

```bash
  cd employee-server
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
docker build -t employee-server .
```
Run container.

```bash
 docker run  --mount type=bind,source=/path/to/project-private-key.json,target=/run/secrets/employee-server-key.json,readonly -p 8000:8000 employee-server
```


