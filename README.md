# Project Simple API

## Overview

This document provides instructions on how to set up your database configuration for the [Your Project Name] project using a `.env` file.

## Prerequisites

Before you start, make sure you have the following:

- [ ] Docker installed on your machine.
- [ ] A running instance of your database (e.g., MySQL, PostgreSQL, SQLite).
- [ ] Database credentials (username, password, host, port).
- [ ] A copy of the `.env.example` file.

## Instructions

1. **Copy the `.env.example` file:**

   If you don't have an `.env` file, start by copying the provided `.env.example` file.
   
   ```
   cp .env.example .env
   ```

2. Open the .env file:

    Use your preferred text editor to open the .env file.

    ```
    nano .env
    ```
3. Update the Database Configuration:

    Locate the database configuration section in the .env file. Update it to match your provided database information:
    ```
    DB_HOST=your_database_host
    DB_PORT=your_database_port
    DB_DATABASE=your_database_name
    DB_USERNAME=your_database_username
    DB_PASSWORD=your_database_password
    DB_DEBUG=true
    DB_MIGRATION=false
    ```
    Replace the placeholder values (your_database_name, your_database_user, your_database_password) with your actual database information. Additionally, set DB_DEBUG to true and DB_MIGRATION to false.

4. Logger Configuration:

    Add the following lines for logger configuration:

    ```
    LOGGER_LOGS_WRITE=true
    LOGGER_FOLDER_PATH=./logs
    ```
    Adjust the LOGGER_FOLDER_PATH based on your preferred folder structure.

5. Save and Close the File:

    Save the changes and close the .env file.

6. Verify the Configuration:

    Make sure your application can connect to the database using the updated configuration. You can do this by running a database-related task or checking your application logs.

7. Run Unit Testing:

    Execute the following command to run unit tests and generate a coverage report:

    ```
    make test-coverage
    ```
8. Build and Run in Docker:

    Use the following command to build and run your application in Docker:

    ```
    make build
    ```
    This assumes you have a Makefile with a target named build that contains the necessary Docker commands.

9. Export Postman Collection:

    Use Postman to export the provided collection file (Simple Api.postman_collection.json) to your local machine.

10. Start or Restart Your Application:

    If your application was already running, you may need to restart it to apply the new database configuration.


# API Documentation

This documentation outlines the usage of the Simple API, providing details on various endpoints and their functionalities.

## Endpoints

### 1. Add Product

- **Method:** POST
- **Endpoint:** `localhost:7690/products`
- **Request Body:**
    ```json
    {
        "title": "Mie Sedap rasa soto",
        "description": "Taburan ayam gurih nikmat di setiap kemasan",
        "rating": 9.0,
        "image": "http://google.com/image.jpg"
    }
    ```
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": null
    }
    ```

### 2. Get List Product

- **Method:** GET
- **Endpoint:** `localhost:7690/products?page=1&limit=10`
- **Query Parameters:**
    - `title` (disabled)
    - `rating` (disabled)
    - `page`: 1
    - `limit`: 10
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": [
            {
                "id": "22c8e385-6d60-4ddb-87b2-3fb543d43177",
                "title": "Mie indomi Rasa ayam Soto",
                "description": "Taburan ayam gurih nikmat di setiap kemasan",
                "rating": 8.1,
                "image": "http://google.com/image.jpg"
            },
            {
                "id": "a1b91cb9-c4a5-408f-ad28-5f32e197d954",
                "title": "Mie indomi Rasa ayam Bawang",
                "description": "Taburan ayam gurih nikmat di setiap kemasan",
                "rating": 8.1,
                "image": "http://google.com/image.jpg"
            }
        ],
        "pagination": {
            "page": 1,
            "limit": 10,
            "totalData": 2,
            "totalPage": 1
        }
    }
    ```

### 3. Get Product Detail

- **Method:** GET
- **Endpoint:** `localhost:7690/products/22c8e385-6d60-4ddb-87b2-3fb543d43177`
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": {
            "id": "22c8e385-6d60-4ddb-87b2-3fb543d43177",
            "title": "Mie indomi Rasa ayam Soto",
            "description": "Taburan ayam gurih nikmat di setiap kemasan",
            "rating": 8.1,
            "image": "http://google.com/image.jpg",
            "createdAt": "2023-12-20T00:00:49.591+07:00",
            "updatedAt": "2023-12-20T00:00:49.591+07:00",
            "deletedAt": null
        }
    }
    ```

### 4. Update Product

- **Method:** PUT
- **Endpoint:** `localhost:7690/products/b34e8eac-ac43-4163-b9ad-49f15644b4fa`
- **Request Body:**
    ```json
    {
        "title": "Mie Sedap rasa soto lamongan",
        "description": "Taburan ayam gurih nikmat di setiap kemasan",
        "rating": 9.1,
        "image": "http://google.com/image.jpg"
    }
    ```
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": null
    }
    ```

### 5. Delete Product

- **Method:** DELETE
- **Endpoint:** `localhost:7690/products/22c8e385-6d60-4ddb-87b2-3fb543d43177`
- **Response:**
    ```json
    {
        "code": 200,
        "message": "Success",
        "data": null
    }
    ```


