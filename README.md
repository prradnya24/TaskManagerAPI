TaskManagerAPI
A lightweight and efficient Task Management REST API built with Go (Gin framework). This backend service provides endpoints to create, read, update, and delete tasks. It’s designed to be simple, scalable, and easy to extend with a future web frontend.

🚀 Features
Create new tasks with title and description

Retrieve all tasks or a single task by ID

Update existing tasks

Delete tasks

Built with Go for speed and simplicity

🛠️ Tech Stack
Go (Gin framework) – Web framework

Postman – API testing

Database – (MySQL)

📦 Installation
Clone the repository:

bash
git clone https://github.com/prradnya24/TaskManagerAPI.git
cd TaskManagerAPI
Install dependencies:

bash
go mod tidy
Run the server:

bash
go run main.go
By default, the API runs on http://localhost:8080.

🔑 API Endpoints
Method	Endpoint	Description
POST	/tasks	Create a new task
GET	/tasks	Get all tasks
GET	/tasks/:id	Get task by ID
PUT	/tasks/:id	Update task by ID
DELETE	/tasks/:id	Delete task by ID


Example Request
bash
POST /tasks
Content-Type: application/json

{
  "title": "Finish README",
  "description": "Write documentation for TaskManagerAPI"
}
🧪 Testing
You can test the API using:

Postman (recommended)

cURL

Or integrate with your frontend later

🌐 Future Plans
Add a web frontend (React, Vue, or Angular)

Implement authentication (JWT or OAuth)

Deploy to cloud platforms (Heroku, AWS, or Azure)

🤝 Contributing
Contributions are welcome!

Fork the repo
Create a new branch (feature/your-feature)

Commit changes
Push and open a Pull Request

📄 License
This project is licensed under the MIT License.
