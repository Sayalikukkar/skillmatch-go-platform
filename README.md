# Skill Marketplace API

A RESTful API and CLI utility built using Go for a Skill Marketplace website, where **Providers** (individuals or companies) offer their skills, and **Users** (clients) request tasks to be completed. This application supports functionalities such as creating users, providers, posting skills, tasks, and managing task statuses. 

## Features

- **Provider Functionality**:
  - Create a provider (individual or company).
  - Post, update, and manage skills.
  - Make offers for tasks posted by users.
  - Update task progress and mark tasks as completed.

- **User Functionality**:
  - Create a user (individual or company).
  - Post tasks with detailed descriptions.
  - Accept or reject offers made by providers.
  - Accept or reject task completion status.

## Tech Stack

- **Go (Golang)**: The primary programming language.
- **Gin**: For building the REST API.
- **GORM**: For database interaction with MySQL.
- **MySQL**: The database used to store application data.

## Endpoints

### Provider Endpoints

- `POST /providers`: Create a new provider.
- `POST /skills`: Create a skill (only providers).
- `PUT /skills/{id}`: Update a skill.
- `POST /tasks/{id}/offer`: Make an offer for a task.
- `PUT /tasks/{id}/progress`: Update task progress with a description and timestamp.
- `PUT /tasks/{id}/complete`: Mark a task as completed.

### User Endpoints

- `POST /users`: Create a new user.
- `POST /tasks`: Create a new task (only users).
- `PUT /tasks/{id}`: Update an existing task.
- `PUT /offers/{id}/accept`: Accept a provider's offer.
- `PUT /offers/{id}/reject`: Reject a provider's offer.
- `PUT /tasks/{id}/accept`: Accept the task completion.

## Database Setup

Ensure you have MySQL installed and set up a database for the application.

### Database Schema

The application uses the following tables:

- **Users**
- **Providers**
- **Skills**
- **Tasks**
- **Offers**
- **Task Progress**
- **Task Completion**

### Database Configuration

1. Clone the repository.
2. Install Go dependencies with `go mod tidy`.
3. Update the `db` for MySQL connection details in the `db/db.go` file.

## Running the Application

### Start the API Server

1. Navigate to the project directory.
2. Run the server using the following command:

```bash
go run main.go
