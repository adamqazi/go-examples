# Concurrency Rate Limiting

## Tasks

- [x] Accept a text file containing a list of URLs separated by newlines.
- [x] Get the contents of the URL and calculate the MD5 for each URL concurrently.
- [x] Log the MD5s maintaining the order in which they were read from the file.
- [x] Allow user to limit the number of goroutines running concurrently (default limit is set to 100).

## Improvements

- Add support for different hash functions; the user should provide the hash function to be used.
- Use a custom `http.Client` instead of the default, `http.DefaultClient`.
- Add unit tests.
- Containerize the application so that users don't need to worry about setting up the environment.

## Setup

In order to run the application you need to have [Go](https://go.dev/doc/install) installed.

The application uses Go Modules so when running the `go build` command all the dependencies listed in the `go.mod` file will automatically be installed.

## Usage

1. First and foremost you need to make sure that the setup has been completed.
2. Once the setup has been completed, you can build the binary using the `go build` command.
3. You can run the executable by typing `./concurrency-rate-limiting`, although this will prompt you to provide the filename.
4. A file already exists in the directory and can be used by running the same executable and providing an argument `./concurrency-rate-limiting -filename urls.txt`.
5. To view all the arguments that the application accepts use the `-help/-h` flag.
6. You can also limit set a threshold on the number of goroutines running concurrently by providing the `-threshold` argument.
7. Once done, you can use the `go clean` command to remove object files and cached files.

### Note: The following command-line flag syntax is permitted

- `-flag`
- `--flag` // double dashes are also permitted
- `-flag=x`
- `-flag x`  // non-boolean flags only
