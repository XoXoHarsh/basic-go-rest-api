# Go REST API Server

A basic Golang REST API server for managing students.

## Overview

This project demonstrates a simple REST API server built with Go. It supports creating a student and retrieving student data.

## Getting Started

### Prerequisites

- Go 1.23.5 or higher

### Clone the Repository

```bash
git clone https://github.com/xoxoharsh/go-student-api.git
cd go-student-api
```

### Running the Project

1. Ensure you have a valid configuration file (see [config/local.yaml](./config/local.yaml)).
2. Run the server:

   ```bash
   go run cmd/students-api/main.go -config=config/local.yaml
   ```

The server will start and listen on the address specified in the config file (default: `localhost:8082`).

## API Endpoints

- **POST /api/students**  
  Create a new student.  
  Request body must be a JSON object with `name`, `email`, and `age`.

- **GET /api/students/{id}**  
  Retrieve a student by id.

- **GET /api/students**  
  Retrieve a list of all students.
