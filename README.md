# Airplane-Dealers-Divar


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