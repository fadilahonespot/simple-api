# Project Name Database Configuration

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

   ```bash
   cp .env.example .env

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

