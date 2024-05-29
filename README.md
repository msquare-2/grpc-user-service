# gRPC User Service

This is a Golang gRPC service that provides functionality to search users.

## Project Structure

```Bash
grpc-user-service/
├── proto/
│ └── user.proto
├── server/
│ ├── server.go
│ ├── user_service.go
│ └── user_service_test.go
├── main.go
├── Dockerfile
└── README.md
```


- `proto/`: Contains the Protocol Buffers file for defining the gRPC service interface.
- `server/`: Contains the implementation of the gRPC service and its tests.
- `main.go`: Entry point of the application.
- `Dockerfile`: Docker configuration file for containerization.
- `README.md`: Instructions and documentation for the project.

## Usage

### Prerequisites

- Go 1.16 or higher installed.
- Docker (optional, for containerization).

### Building and Running the Application

1. Clone the repository:
```sh
    git clone 
```

2. Navigate to the project directory:
```sh
    cd grpc-user-service
```

3. Build the application:
```sh
    go build -o grpc-user-service ./main.go
```

4. Run the application:
```sh
    ./grpc-user-service
```


### Accessing gRPC Service Endpoints

The gRPC service will be running on port `50051` by default.

### Docker

To build a Docker image for the application:
```sh
    sudo docker build -t grpc-user-service .
```

To run the Docker container:
```sh
    sudo docker run -p 50051:50051 grpc-user-service
```


## API Reference

The gRPC service provides the following methods:

- `GetUserByID`: Fetches user details based on user ID.
- `GetUsersByIDs`: Retrieves a list of user details based on a list of user IDs.
- `SearchUsers`: Searches for user details based on specific criteria.

## Testing

Unit tests are implemented for the service methods. Run the tests using:
```sh
    go test ./server -v
```


## Contributors

- [Mustafa](https://github.com/msquare-2)

## License

This project is licensed under the [MIT License](LICENSE).
