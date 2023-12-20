# go-cleanarch-api

## Diagram

The diagram below shows the architectural components that are implemented in this project and how it's integrated each other. 

There are two main groups of components contained by the dotted line boxes. They are **infra** and **core**.

In the **infra** group are all components not related with the application business logic. Here, they are divided in two sub-categories: **controller** and **repository**.

The sub-group **controller** contains the controller of the application, an interface for controller adapters and two implementations of this interface, using each a different web framework, [Gin](https://github.com/gin-gonic/gin) and [Fiber](https://github.com/gofiber/fiber).

The sub-group **repository** contains two implementations of repositories using the databases [MySQL](https://www.mysql.com/) and [Postgres](https://www.postgresql.org/). Here is contained all specific vendor implementations to connect to databases, open connections and make operations (select, insert, update, delete, etc).

In the **core** group are all components that are related with the application business logic. Here, they are divired in three sub-categories: **usecase** (business logic), **external** (interfaces to external components) and **domain** (application domain objects).

![go-cleanarch-api diagram](diagram.png)

## Launch MySQL and Postgres with Docker

To launch the MySQL and Postgres with the correct database, user and password to work with the instructions contained here, run the docker compose with the command below:

```sh
docker-compose up -d
```

## Run API with VSCode

You can run the application using the integrated VSCode launcher (check if the required extensions is installed) copying the json below to the `.vscode/launch.json` file. In this file, there are five different ways to run the application, being them:
* Launch with static env: run application with the variables seted in `static_env.go` file
* Launch with Fiber and MySQL: run application with Fiber web framework and MySQL database
* Launch with Fiber and Postgres: run application with Fiber web framework and Postgres database
* Launch with Gin and MySQL: run application with Gin web framework and MySQL database
* Launch with Gin and Postgres: run application with Gin web framework and Postgres database

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch with static env",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "ENV_TYPE": "STATIC"
            }
        },
        {
            "name": "Launch with Fiber and MySQL",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "ENV_TYPE": "HOST",
                "HTTP_FRAMEWORK": "FIBER",
                "DATABESE": "MYSQL",
                "MYSQL_BOOK_DATABASE": "book",
                "MYSQL_BOOK_USER": "book",
                "MYSQL_BOOK_PASSWORD": "book",
                "MYSQL_BOOK_HOST": "localhost",
                "MYSQL_BOOK_PORT": "3306",
                "POSTGRES_BOOK_DATABASE": "book",
                "POSTGRES_BOOK_USER": "book",
                "POSTGRES_BOOK_PASSWORD": "book",
                "POSTGRES_BOOK_HOST": "localhost",
                "POSTGRES_BOOK_PORT": "5432",
            }
        },
        {
            "name": "Launch with Fiber and Postgres",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "ENV_TYPE": "HOST",
                "HTTP_FRAMEWORK": "FIBER",
                "DATABESE": "POSTGRES",
                "MYSQL_BOOK_DATABASE": "book",
                "MYSQL_BOOK_USER": "book",
                "MYSQL_BOOK_PASSWORD": "book",
                "MYSQL_BOOK_HOST": "localhost",
                "MYSQL_BOOK_PORT": "3306",
                "POSTGRES_BOOK_DATABASE": "book",
                "POSTGRES_BOOK_USER": "book",
                "POSTGRES_BOOK_PASSWORD": "book",
                "POSTGRES_BOOK_HOST": "localhost",
                "POSTGRES_BOOK_PORT": "5432",
            }
        },
        {
            "name": "Launch with Gin and MySQL",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "ENV_TYPE": "HOST",
                "HTTP_FRAMEWORK": "GIN",
                "DATABESE": "MYSQL",
                "MYSQL_BOOK_DATABASE": "book",
                "MYSQL_BOOK_USER": "book",
                "MYSQL_BOOK_PASSWORD": "book",
                "MYSQL_BOOK_HOST": "localhost",
                "MYSQL_BOOK_PORT": "3306",
                "POSTGRES_BOOK_DATABASE": "book",
                "POSTGRES_BOOK_USER": "book",
                "POSTGRES_BOOK_PASSWORD": "book",
                "POSTGRES_BOOK_HOST": "localhost",
                "POSTGRES_BOOK_PORT": "5432",
            }
        },
        {
            "name": "Launch with Gin and Postgres",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "ENV_TYPE": "HOST",
                "HTTP_FRAMEWORK": "GIN",
                "DATABESE": "POSTGRES",
                "MYSQL_BOOK_DATABASE": "book",
                "MYSQL_BOOK_USER": "book",
                "MYSQL_BOOK_PASSWORD": "book",
                "MYSQL_BOOK_HOST": "localhost",
                "MYSQL_BOOK_PORT": "3306",
                "POSTGRES_BOOK_DATABASE": "book",
                "POSTGRES_BOOK_USER": "book",
                "POSTGRES_BOOK_PASSWORD": "book",
                "POSTGRES_BOOK_HOST": "localhost",
                "POSTGRES_BOOK_PORT": "5432",
            }
        },
    ]
}
```
