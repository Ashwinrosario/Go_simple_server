# Simple Go Server with MongoDB

This is a simple Go server that interacts with a MongoDB database. The server exposes several API endpoints to manage user data (CRUD operations).

## Prerequisites

To run this project, you need the following tools installed on your machine:

- Go (1.16 or later)
- MongoDB
- `godotenv` package for loading environment variables

## Setup and Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/ashwinrosario/Go_simple_server.git
    cd simple-server
    ```

2. **Install dependencies:**

    This project uses the `github.com/joho/godotenv` and `go.mongodb.org/mongo-driver/mongo` packages. These dependencies should be installed automatically when you run the project, but you can also install them manually if needed.

    ```bash
    go get github.com/joho/godotenv
    go get go.mongodb.org/mongo-driver/mongo
    ```

3. **Create the `.env` file:**

    Create a `.env` file in the root directory and define the following variables:

    ```plaintext
    MONGO_URI=mongodb://localhost:27017
    DB_NAME=your_database_name
    COLLECTION_NAME=users
    SERVER_PORT=:8080
    ```

    Replace the `MONGO_URI` with your MongoDB connection string, `DB_NAME` with your desired database name, `COLLECTION_NAME` with your desired collection name, and `SERVER_PORT` with the desired port for the server.

4. **Run the server:**

    Run the server using the following command:

    ```bash
    go run main.go
    ```

    The server will start running on the port specified in your `.env` file (default: `8080`).

    You should see the following output:

    ```
    Successfully connected to MongoDB!
    Server is running on http://localhost:8080
    ```

## API Endpoints

### 1. **GET /users**

   Retrieve all users in the database.

   **Request:**
   ```http
   GET http://localhost:8080/users
