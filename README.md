# Task Management API

Simple CRUD API project using Golang.

## Features

- Create Task
- Get All Tasks
- Get Task By ID
- Update Task
- Delete Task

## Project Structure

entity/
repository/
service/
handler/
router/
middleware/
utils/

## Run Project

### Clone project

git clone <your-repository-url>

### Go to project folder

cd task_management

### Run project

go run .

Server will run on:

http://localhost:8080

## API Endpoints

### Create Task

POST /tasks

### Get All Tasks

GET /tasks

### Get Task By ID

GET /tasks/{id}

### Update Task

PUT /tasks/{id}

### Delete Task

DELETE /tasks/{id}

## Run Tests

go test ./...