# Go File Hosting

This is a simple microservice to download or upload files and make them available via a webserver. This project was developed with Go and its execution can be done easily using Docker.

Its operation is as follows, perform a POST request in one of the two available routes with the necessary data, such as the URL of the file to be downloaded or the file to be sent/upload to the server, then this file is be available in a mini file hosting within the application, and these files can be accessed via a URL.

To know more details, there is the Makefile file with the commands necessary for its execution and example of how to use, including `curl` commands.

# How to use

Builds the Docker application, must be the first command to be executed:

```
make build
```

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

Example of how to upload a file.

(replace FILE with the file to sent/upload.)

```
make test-file file=FILE
```

Example of how to download a file.

(replace URL with url file to download.)

```
make test-url url=URL
```
