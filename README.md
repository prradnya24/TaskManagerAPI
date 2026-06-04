# 📝 Task Manager API

A lightweight and efficient **Task Management REST API** built with **Go** and the **Gin** framework. This backend service provides a full CRUD interface to manage tasks, backed by a **MySQL** database using **GORM**.

---

## 🚀 Features

- ✅ Create new tasks with title, description, and status
- ✅ Retrieve all tasks
- ✅ Retrieve a single task by ID
- ✅ Update existing tasks
- ✅ Delete tasks
- ✅ Environment-based configuration (no hardcoded credentials)
- ✅ Clean layered architecture (models / controllers / routes / config)

---

## 🛠️ Tech Stack

| Tool | Purpose |
|---|---|
| [Go](https://golang.org/) | Programming language |
| [Gin](https://github.com/gin-gonic/gin) | HTTP web framework |
| [GORM](https://gorm.io/) | ORM for database interaction |
| [MySQL](https://www.mysql.com/) | Relational database |
| [godotenv](https://github.com/joho/godotenv) | Environment variable loader |
| [Postman](https://www.postman.com/) | API testing |

---

## 📁 Project Structure

```
task-manager/
├── .env                  # DB credentials (not committed to Git)
├── .gitignore
├── main.go               # Entry point
├── config/
│   └── db.go             # Database connection
├── models/
│   └── task.go           # Task struct / DB schema
├── controllers/
│   └── task.go           # CRUD logic
└── routes/
    └── routes.go         # Route definitions
```

---

## 📌 API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/tasks` | Create a new task |
| `GET` | `/tasks` | Get all tasks |
| `GET` | `/tasks/:id` | Get a single task by ID |
| `PUT` | `/tasks/:id` | Update a task by ID |
| `DELETE` | `/tasks/:id` | Delete a task by ID |

---

## 📂 Task Model

```json
{
  "id": 1,
  "title": "Build REST API",
  "description": "Complete the task manager project",
  "status": "pending"
}
```

---

## ⚙️ Getting Started

### Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [MySQL](https://www.mysql.com/)
- [Git](https://git-scm.com/)

---

### 1. Clone the repository

```bash
git clone https://github.com/your-username/task-manager-api.git
cd task-manager-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set up the database

Create a MySQL database:

```sql
CREATE DATABASE task_manager;
```

### 4. Configure environment variables

Create a `.env` file in the root folder:

```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=task_manager
```

### 5. Run the server

```bash
go run main.go
```

Server will start at: `http://localhost:8080`

> The `tasks` table will be created automatically via GORM AutoMigrate.

---

## 🧪 Testing the API (Postman)

**Create a Task**
```
POST http://localhost:8080/tasks
Content-Type: application/json

{
  "title": "Buy groceries",
  "description": "Milk, eggs, bread",
  "status": "pending"
}
```

**Get All Tasks**
```
GET http://localhost:8080/tasks
```

**Get Task by ID**
```
GET http://localhost:8080/tasks/1
```

**Update a Task**
```
PUT http://localhost:8080/tasks/1
Content-Type: application/json

{
  "status": "done"
}
```

**Delete a Task**
```
DELETE http://localhost:8080/tasks/1
```

---

## 🔒 Security Notes

- Database credentials are stored in `.env` and never hardcoded
- `.env` is listed in `.gitignore` and never pushed to GitHub

---

## 🔮 Future Improvements

- [ ] JWT Authentication (login & signup)
- [ ] Docker support
- [ ] Pagination for GET /tasks
- [ ] Input validation middleware
- [ ] Frontend integration

---

## 👤 Author

**Your Name**
- GitHub: [@prradnya24](https://github.com/prradnya24)

---

