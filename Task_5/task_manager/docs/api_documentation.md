Task Management API with MongoDB
Base URL
http://localhost:8080


Authorization: Bearer YOUR_TOKEN_HERE

Endpoints

1. User Registration

Endpoint: `/register`
Method: `POST`
Request Body:


{
    "username": "string",
    "password": "string",
    "role": "optional_string"
}

Response:
201 Created: User registered successfully.
400 Bad Request: Invalid payload or username already taken.


2. User Login
Endpoint: /login
Method: POST
Request Body:
{
    "username": "string",
    "password": "string"
}
Response:
200 OK: Returns the JWT token.
401 Unauthorized: Invalid username or password.


3. Get Tasks
Endpoint: /tasks
Method: GET
Authorization: Required (User or Admin role)
Response:
200 OK: Returns a list of tasks.
403 Forbidden: Insufficient permissions.

4. Create Task
Endpoint: /tasks
Method: POST
Request Body:
{
    "title": "string",
    "description": "string",
    "status": "string"
}
Authorization: Required (User or Admin role)
Response:
201 Created: Task created successfully.
400 Bad Request: Invalid task data.
403 Forbidden: Insufficient permissions.


5. Update Task
Endpoint: /tasks/{id}
Method: PUT
Request Body:
{
    "title": "string",
    "description": "string",
    "status": "string"
}

Authorization: Required (Admin role)
Response:
200 OK: Task updated successfully.
404 Not Found: Task not found.
403 Forbidden: Insufficient permissions.


6. Delete Task
Endpoint: /tasks/{id}
Method: DELETE
Authorization: Required (Admin role)
Response:
200 OK: Task deleted successfully.
404 Not Found: Task not found.
403 Forbidden: Insufficient permissions.