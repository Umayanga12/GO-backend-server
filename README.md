# Go Backend HTTP Server

## Overview

This repository contains a Go-based backend HTTP server. The server facilitates account management and token operations, providing APIs Request Handling
## Features

- Account creation
- Token allocation
- Token burning
- Health check endpoint
- CORS support

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [Chi Router](https://github.com/go-chi/chi)
- [GoDotEnv](https://github.com/joho/godotenv)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Umayanga12/GO-backend-server.git
    cd GO-backend-server
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Set up Ganache and start the local blockchain.

4. Create a `.env` file in the root directory and set the following environment variable:

    ```plaintext
    PORT=your_desired_port
    ```

## Usage

1. Run the server:

    ```sh
    go run main.go
    ```

2. The server will start and listen on the port specified in the `.env` file. By default, it includes the following endpoint:

    - `GET /v1/healthz` - Health check endpoint to verify if the server is running.

## Code Explanation

### Main File (`main.go`)

- **Importing packages**:
  - `encoding/json`, `fmt`, `log`, `net/http`, `os` for standard Go functionalities.
  - `github.com/go-chi/chi` for routing.
  - `github.com/go-chi/cors` for handling CORS.
  - `github.com/joho/godotenv` for loading environment variables from a `.env` file.

- **Main Function**:
  - Loads environment variables from the `.env` file.
  - Retrieves the port from the environment variable.
  - Sets up the router using Chi.
  - Configures CORS to allow requests from any origin.
  - Mounts the `/v1` route with a health check endpoint.
  - Starts the server on the specified port.

- **Health Check Handler (`HandlerRediness`)**:
  - Responds with a 200 OK status to indicate server readiness.

- **Helper Functions**:
  - `responceWithJson` - Sends a JSON response.
  - `responceWithError` - Sends an error response in JSON format.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
