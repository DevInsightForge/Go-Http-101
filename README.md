# Go HTTP 101

## Overview

This project is a web API developed in Go (Golang) v1.22 for handling HTTP requests. It provides endpoints for managing tasks and utilizes Docker for containerization, making it easy to deploy across different environments. The project follows a clean and organized structure to enhance maintainability and scalability.

## Features

- **HTTP Server**: Implements a robust HTTP server for handling incoming requests.
- **Routing**: Utilizes a flexible router setup for managing API endpoints efficiently.
- **Middleware**: Incorporates middleware for handling cross-origin resource sharing (CORS) and logging.
- **Database Integration**: Supports MongoDB as the backend database for storing task-related information.
- **Error Handling**: Implements utility functions for generating consistent and informative error responses.
- **Platform Independence**: Provides binaries for multiple platforms (Windows, Linux, macOS) to ensure compatibility.

## Project Structure

 ```
 .
├── Dockerfile
├── Makefile
├── bin
│   ├── darwin_amd64
│   │   └── gohttp101
│   ├── linux_amd64
│   │   └── gohttp101
│   └── windows_amd64
│       └── gohttp101.exe
├── cmd
│   ├── api
│   │   ├── http-server
│   │   │   └── http.server.go
│   │   ├── router-setup
│   │   │   ├── base.router.go
│   │   │   └── task.router.go
│   │   └── server.go
│   └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── internal
    └── application
        ├── config
        │   └── env.go
        ├── database
        │   └── mongodb.go
        ├── handler
        │   └── task.handler.go
        ├── middleware
        │   ├── cors.middleware.go
        │   └── logger.middleware.go
        ├── model
        │   ├── base
        │   │   └── base.model.go
        │   └── task.model.go
        ├── repository
        │   ├── base
        │   │   └── base.repository.go
        │   └── task.repository.go
        ├── service
        └── utility
            └── response.utility.go
 ```
 
## Setup and Installation

1. **Clone the repository:**  
   Clone `https://github.com/DevInsightForge/Go-Http-101`
2. **Install Dependencies:**  
   Run `go mod download`.
3. **Build the Application:**  
   Run `make`.
4. **Run the Application:**  
   Execute the appropriate binary for your platform located in the `bin` directory.

## Configuration

- **Environment Variables:**  
  Update the `.env.example` file in the `root` directory to set environment-specific configurations such as database connection details, server port, etc.

## Usage

- **Endpoints:**  
  Define and implement additional endpoints as per the requirements in the `router-setup` directory.
- **Middleware:**  
  Customize middleware functions in the `middleware` directory to suit specific needs.
- **Error Handling:**  
  Extend error handling logic in the `utility` directory for handling various error scenarios gracefully.

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/feature-name`)
3. Make changes and commit (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/feature-name`)
5. Create a pull request

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

- [Golang Documentation](https://golang.org/doc/)
- [Docker Documentation](https://docs.docker.com/)
- [MongoDB Documentation](https://docs.mongodb.com/)
- [Open Source Community](https://github.com/)

## Contact

For any inquiries or feedback, please contact [Project Maintainer](mailto:imzihad@gmail.com).
