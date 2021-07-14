# hcp

## Name

    hcp -- HTTP copy file

## Synopsis

``` bash
hcp [:port] file [...]
```

## Description

`hcp` stands for "HTTP cp". It is a simple utility written in Go to allow users to copy their local file via a temporary HTTP Server.

Parameters:

- `port`: the port on which the HTTP server will listen. This parameter is optional, if not specified, the default port is 8000.
- `file`: the path to a file to be copied (no matter if relative or absolute).

## Examples

To copy files on the default port (8000):
``` bash
hcp foo.txt bar.csv
```

To copy files on a custom port (for example, 777):
``` bash
hcp :777 foo.txt bar.csv
```
## Dependencies

This program depends on the Go packet [Mux](https://github.com/gorilla/mux). To install it, see the official guide.

## Compile

To comile the program is sufficient to call `go build` on `hcp.go`:

``` bash
go build hcp.go
```

Then you can relocate the executable file anywhere you want.

You can also choose to not compile this program and to run it as a script:

``` bash
go run hcp.go [:port] file [...]
```

