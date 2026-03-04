# Task API 🚀

A simple RESTful Task API built with Go and the Gin web framework. This API provides full CRUD (Create, Read, Update, Delete) operations for managing tasks.

![Go Version](https://img.shields.io/badge/Go-1.25-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Framework](https://img.shields.io/badge/Gin-1.11-orange)

## 📋 Features

- **Create Task** - Add new tasks to the list
- **Get All Tasks** - Retrieve all tasks
- **Update Task** - Modify existing tasks (title and status)
- **Delete Task** - Remove tasks from the list
- **In-memory Storage** - Tasks are stored in memory (data persists while server is running)

## 🛠 Tech Stack

- **Language:** [Go](https://golang.org/) 1.25
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Architecture:** RESTful API

## 📦 Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/tawhidul36/To-Do.git
   cd go-practice
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the server:**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

| Method   | Endpoint        | Description              |
|----------|-----------------|--------------------------|
| `POST`   | `/tasks`        | Create a new task        |
| `GET`    | `/tasks`        | Get all tasks            |
| `PUT`    | `/tasks/:id`    | Update a task by ID      |
| `DELETE` | `/tasks/:id`    | Delete a task by ID      |

### Request & Response Examples

#### Create a Task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "done": false}'
```

**Response:**
```json
{
  "id": 1,
  "title": "Learn Go",
  "done": false
}
```

#### Get All Tasks
```bash
curl -X GET http://localhost:8080/tasks
```

**Response:**
```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "done": false
  }
]
```

#### Update a Task
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "done": true}'
```

**Response:**
```json
{
  "id": 1,
  "title": "Learn Go",
  "done": true
}
```

#### Delete a Task
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

**Response:**
```json
{
  "message": "Deleted"
}
```

## 📝 Task Structure

| Field  | Type    | Description              |
|--------|---------|--------------------------|
| `id`   | integer | Unique task identifier   |
| `title`| string  | Task title               |
| `done` | boolean | Task completion status   |

## 🤝 Contributing

Feel free to fork this repository and submit pull requests!

## 📄 License

This project is licensed under the MIT License.
