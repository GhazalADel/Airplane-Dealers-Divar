# Airplane-Dealers-Divar
Airplane Dealers Divar is a cutting-edge advertising platform designed to facilitate the buying and selling of airplanes by airlines. This platform serves as a centralized hub where airlines can showcase their available aircraft for potential buyers, providing a seamless and efficient process for aircraft transactions. The project is developed using the Go programming language, leveraging its robustness and concurrency features to deliver a high-performance and scalable solution.


## Features

- User registration and login
- Payment gateway creation and verification
- Ads
  - Ads Bookmarks
  - Repair Request
  - Expert Check Request
  - Ads Filtering
 - Diffrent User Roles (Matin <<Super User>>, Admin, Expert, Airline)
 - Admin Panel for managing users, configurations, and Ads

## Setup

1. Clone the repository to your local machine:

```bash
git clone https://github.com/zereshk-quera/Airplane-Dealers-Divar.git
```

2. Navigate into the project directory:

```bash
cd Airplane-Dealers-Divar
```

3. Install the required dependencies:

```bash
go mod download
```

4. Run the project:

```bash
go run .
```

The server will start running at `localhost:8080`.

5. Access the Swagger API documentation:

Swagger URL: [http://localhost:8080/swagger/](http://localhost:8080/swagger/)

The Swagger URL provides access to the Swagger API documentation for the Airplane-Dealers-Divar Project.

## Swagger
start using swagger [echo-swagger man page](https://github.com/swaggo/echo-swagger)

The Swagger UI provides a user-friendly interface to explore the API endpoints, view request/response details, and even test the API by sending requests directly from the Swagger UI.

Open a web browser and navigate to the Swagger UI URL. The default URL is typically http://localhost:8080/swagger/index.html.

## Run Tests

To run tests for your Go project using the `go test` command, follow these steps:

1. Open your project's terminal or command prompt.

2. Navigate to the root directory of your Go project.

3. Run the following command to execute all tests in your project and its subdirectories:

```shell
   go test ./...
```

   - The `go test` command is used to run tests in Go.
   - The `./...` argument instructs Go to run tests in the current directory and all its subdirectories.

4. Go will execute the tests and display the results in the terminal or command prompt. You'll see output indicating which tests passed or failed, along with any error messages or stack traces.

Running tests with `go test ./...` is a convenient way to execute all tests within your project. It recursively traverses the project directory structure, identifying and running tests in each package.

Make sure you have the necessary test files and test functions defined in your project. The test files should have a `_test.go` suffix, and the test functions should start with the word `Test`. For example, a test function could be named `TestMyFunction`.

By running tests regularly, you can ensure the correctness and reliability of your code, identify and fix any issues, and maintain the quality of your Go project.
