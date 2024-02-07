# gRPC, Protobuf and MySQL Example

This project is an example of how gRPC, Protobuf and MySQL connections can work together. It uses the Go programming language and the Goqu library for database operations.

## Components

- **gRPC**: gRPC is a high-performance, open-source universal RPC framework. In this project, gRPC is used to define the service and the message request format in a .proto file, and to handle the server and client logic.

- **Protobuf**: Protobuf (Protocol Buffers) is a method of serializing structured data. It's useful when you're writing data which you want to send across a network or store in a file. In this project, Protobuf is used to define the service and the message request format in a .proto file.

- **MySQL**: MySQL is a relational database management system. In this project, MySQL is used to store account data.

- **Goqu**: Goqu is a SQL builder and query library for Go. In this project, Goqu is used to interact with the MySQL database.

## How to Run Locally

1. Clone the repository to your local machine.

2. Install the necessary dependencies.

3. Set the environment variable `ENV` to either `dev` or `prod`. This will determine which configuration file (dev.yaml or prod.yaml) the application uses. The configuration file contains the SQL address for the MySQL database.

4. Run the server code by executing the following command in your terminal: `go run server/server.go`. This will start a gRPC server that listens for `CreateAccount` and `GetAccount` requests.

5. Install grpcurl if you haven't already. It's a command-line tool that lets you interact with gRPC servers. You can install it with the following command: `brew install grpcurl` (for macOS) or `choco install grpcurl` (for Windows).

6. Run the client code using grpcurl. For example, to send a `CreateAccount` request, you can use the following command: `grpcurl -d '{"name": "John Doe", "email": "john.doe@example.com"}' -plaintext localhost:50051 accountmanager.AccountManager/CreateAccount`. This will send a `CreateAccount` request to the server.

7. Similarly, to send a `GetAccount` request, use the following command: `grpcurl -d '{"id": "your-account-id"}' -plaintext localhost:50051 accountmanager.AccountManager/GetAccount`. Replace "your-account-id" with the actual ID of the account you want to retrieve.

8. Check the MySQL database. You should see the new account created by the `CreateAccount` request.

## Note

This is a basic example and does not include any authentication or authorization checks. It's not suitable for production use without additional security measures.
