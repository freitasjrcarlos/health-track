
# Health Track Service

This is a simple **health check** service implemented in Go. It makes HTTP requests to URLs configured in an environment variable and checks the response status of each. The service runs continuously, checking the URLs at regular intervals (configured in the code).

URLs that return a 200 (OK) status are logged as successful, and URLs that return other statuses or connection errors are logged as failures.

## Features

- Checks the health of multiple services via HTTP.
- Logs the results to the terminal, using colors to indicate success or failure:
  - **Green** for success (status 200).
  - **Red** for errors or statuses other than 200.
- Continuous loop to check the URLs every 30 seconds (configurable).

## Prerequisites

Before running the project, make sure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.19 or higher)
- [Docker](https://www.docker.com/get-started) (optional, if you want to run it in a container)

## How to Run the Service

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/health-track.git
cd health-track
```

### 2. Configure environment variables

The service uses an environment variable called `CHECK_URLS` to determine which URLs to check. Create a `.env` file in the root of the project with the following content:

```env
CHECK_URLS=http://example.com/health,http://example2.com/health
```

**Important**: Separate the URLs with a comma and do not include spaces between them.

### 3. Run the project locally

#### Install dependencies

Run the following command to install the Go dependencies defined in the `go.mod` file:

```bash
go mod tidy
```

#### Start the service

Once the dependencies are installed, you can start the service with the following command:

```bash
go run main.go
```

This will start the health check service, which will check the URLs defined in the `.env` file every 30 seconds and print the results to the terminal.

### 4. Customize the check interval

In the `main.go` file, you can adjust the time interval between checks by modifying the line that contains `time.Sleep(30 * time.Second)`.

Example to run the health check every 60 seconds:

```go
time.Sleep(60 * time.Second)
```

## Project Structure

```
├── .env                # Environment file with URLs to check
├── main.go             # Main service code
├── go.mod              # Go module file
├── go.sum              # Go module dependency checksums
└── README.md           # This file
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
