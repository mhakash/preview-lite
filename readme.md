# Preview-Lite File Server

This project provides a lightweight file server to serve static content for your frontend application.

## Usage

1. **Run the binary in the `index.html` directory:**

    ```sh
    ./preview-lite
    ```

2. **Or specify the location of the files:**

    ```sh
    ./preview-lite "location-of-files"
    ```

## Main Features

- Serves static content for frontend applications.
- Easy to use by running a single binary.

## Getting Started

1. **Build the project:**

    ```sh
    go build -o preview-lite main.go
    ```

2. **Run the binary as described in the usage section.**

## Development

To build and run the project during development, you can use the provided `makefile`:

1. **Build the project:**

    ```sh
    make build
    ```

2. **Run the project:**

    ```sh
    make run
    ```

## License

This project is licensed under the MIT License.
