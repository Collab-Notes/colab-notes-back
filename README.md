# colab-notes-back

## Go Installation

To run this project, you need to have Go installed on your system:

1. Download Go from the [official website](https://golang.org/dl/)
2. Follow the installation instructions according to your operating system:
   - **Windows**: Run the downloaded installer and follow the instructions
   - **macOS**: Use the installation package or `brew install go`
   - **Linux**: Extract the downloaded file to `/usr/local` or use your distribution's package manager

Verify the installation by running:

```sh
go version
```

## Swagger

To generate the documentation, run:

```sh
swag init
```

If you don't have the library downloaded, do it with:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

## Running the Project

Before running the following, you **MUST** create the documentation files using the above Swagger command.

### Standard Method

To run the project using the standard Go command:

1. Clone the repository
2. Navigate to the project directory
3. Run the application:

```sh
go run main.go
```

### Using Air for Hot-Reloading

For development, you can use Air to automatically reload the application when code changes are detected:

1. Install Air:

    ```sh
    go install github.com/cosmtrek/air@latest
    ```

2. Run the application with Air:

    ```sh
    air
    ```

Air will watch for file changes and automatically restart the server.
