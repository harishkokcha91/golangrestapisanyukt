Here is an example of a `setup.md` file based on your instructions:

---

# Project Setup and Run Guide

## Prerequisites

- Docker installed and running on your local machine
- Go installed on your local machine
- Visual Studio Code (VSCode) or any other code editor
- A REST client like VSCode RestClient extension or Postman

## Steps to Run the Project

### 1. Start Docker Daemon

First, ensure the Docker daemon is running. You can start it with the following command:

```sh
sudo dockerd
```

### 2. Start PostgreSQL with Docker Compose

Use Docker Compose to run the PostgreSQL image:

```sh
docker-compose up -d
```

### 3. Create the Database

If the database has not been created, run the `CreateDatabase.go` file:

```sh
go run CreateDatabase.go
```

### 4. Migrate the Database

If the tables are not present in the database, you need to run the migration script located in the `migrate` folder:

```sh
go run migrate/migrate.go
```

### 5. Run the Main Application

Start your main application server by running the main file:

```sh
go run main.go
```

### 6. Test the APIs

You can test the APIs using the REST client of your choice. 

#### Using VSCode RestClient Extension:

- Install the RestClient extension from the VSCode marketplace.
- Use the `.rest` files located in the `restclient` folder to test the APIs.

#### Using Postman:

- Import the requests from the `restclient` folder into Postman and run them.

### Additional Notes

- Ensure you have the necessary environment variables set up in your `.env` or `app.env` file as required by the application.
- You can refer to Docker and PostgreSQL documentation for troubleshooting if you encounter any issues.

---

Feel free to adjust the steps or add any additional information as needed.