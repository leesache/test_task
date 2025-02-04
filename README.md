# DeNet2.0

DeNet2.0 is a Go-based RESTful API designed to manage user registration, authentication, referrals, and task completions with PostgreSQL as the database.

## Features
- **User Registration & Login**: Secure JWT-based authentication.
- **Protected routes to each endpoint**: Protected with JWT.
- **Referral System**: Users can refer others, they both will earn rewards
- **Task Completion**: Users complete tasks to update their balance.
- **Graceful Shutdown**: Ensures clean termination of the server.
- **Error Handling**: Custom error hierarchy for meaningful responses.

## Setup

### Prerequisites
- Docker and Docker Compose installed.

### Run Locally
1. Clone the repository:
   git clone https://github.com/yourusername/DeNet2.0.git
   cd DeNet2.0
2. Start the application:
   docker-compose up --build

### API Endpoints
- POST /users/register: Register a new user.
  Request Body:
    {
      "user_name": "testuser",
      "user_password": "password123"
    }
- POST /users/:id/referrer: Apply a referral code.
  Request Body:
    {
      "referrer_id": 2
    }
- POST /users/{id}/task/complete: Complete a task and update balance.
- GET /users/leardearboard: Get leadearboard sorted by Balance DESC
- GET /users/{id}/status

### Tech Stack
- Backend: Go (Gin framework)
- Database: PostgreSQL
- Authentication: JWT
- Error Handling: Custom error hierarchy
- Deployment: Docker

## Error Responses
The API provides structured error responses with appropriate HTTP status codes:
- 400 Bad Request: Invalid input or missing parameters.
- 401 Unauthorized: Invalid credentials or unauthorized access.
- 404 Not Found: Resource not found.
- 409 Conflict: Invalid operation (e.g., self-referral or task already completed).
- 500 Internal Server Error: Unexpected server errors.
= etc.

## Graceful Shutdown
The server supports graceful shutdown to ensure:
- Active requests are completed before termination.
- Database connections and other resources are properly released.

To trigger a graceful shutdown, send a termination signal (SIGINT or SIGTERM) to the running container:
docker kill -s SIGTERM denet-web

## Contributing
Feel free to contribute by opening issues or submitting pull requests. All contributions are welcome!

## License
This project is licensed under the MIT License. See the LICENSE file for details.