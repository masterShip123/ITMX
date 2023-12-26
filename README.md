# ITMX
#Setup and Run the Golang REST Application:
1. Install Required Packages:
Make sure you have the required packages installed. In the terminal, run:

- go get -u github.com/gin-gonic/gin
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/sqlite
2. Download the Code:
Download the Golang code for the REST application and the data initialization script.

3. Create SQLite Database File:
Before running the application, create an SQLite database file named test.db in the same directory as the application.

4. Run the Application:
In the terminal, navigate to the directory containing the Golang code for the REST application. Run the following command:
go run /local/main.go
This will start the Golang REST application on port 8080.

#Setup and Run the Data Initialization Script:
1. Download the Code:
Download the Golang code for the data initialization script.

2. Run the Script:
In the terminal, navigate to the directory containing the Golang code for the data initialization script. Run the following command:

go run /migrationDB/main.go
This will insert sample customer data into the SQLite database.

#Interact with the REST API:
Now that both the application and the data initialization script are running, you can interact with the REST API using tools like curl, Postman, or any HTTP client. Here are some example requests:

- Create a Customer:

curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice","age":28}' http://localhost:8080/customers

- Get a Customer by ID:

curl http://localhost:8080/customers/{id}

- Update a Customer by ID:

curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Name","age":30}' http://localhost:8080/customers/{id}

- Delete a Customer by ID:

curl -X DELETE http://localhost:8080/customers/{id}
Make sure to replace {id} with an actual customer ID in the above commands.

By following these steps, you should be able to set up and run the Golang REST application and initialize the database with sample data using the provided script.
