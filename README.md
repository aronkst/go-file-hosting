# Go File Hosting

This open source project is a straightforward microservice designed for file download or upload operations, accessible via a web server. Built using Go, it facilitates deployment through Docker for seamless execution.

The microservice functions through two primary routes accessed via POST requests. Users can either specify a file URL for downloading or directly upload a file to the server. Once processed, the files become accessible within a built-in mini file hosting environment within the application, accessible via unique URLs.

For more comprehensive understanding and setup, the project includes a Makefile with essential commands for execution. Additionally, it provides examples demonstrating usage scenarios, including practical `curl` commands for quick integration and testing.

# How to use

Run the application:

```
make run
```

Start the application:

```
make start
```

Stop the application:

```
make stop
```

Example of how to upload a file:

(replace FILE with the file to sent/upload.)

```
make test-file file=FILE
```

Example of how to download a file:

(replace URL with url file to download.)

```
make test-url url=URL
```
