# Todo golang project
This project is just a study case for apply basic thing in golang.  
Basically a api for CRUD todo list.

## Features / Routes

- `GET /todo/` return a list of todo. 
- `GET /todo/{id}` return a describe of indentified todo.
- `POST /todo/` add a new todo.
- `PUT /todo/{id}` edit indentified todo.
- `DELETE /todo/{id}` delete indentified todo.

## Todo data
- **id** - identifier of todo.
- **name** - name of todo.
- **description** - description of todo.
- **created** - created date of todo.
- **updated** - updated date of todo.