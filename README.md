# hcp

## Name

    hcp -- HTTP copy file

## Synopsis

``` bash
hcp [:port] file [...]
```

## Description

`hcp` stands for "HTTP cp". It is a simple utility written in Go that allows users to copy their local file via a temporary HTTP Server.

## Parameters:

- `port`: the port on which the HTTP server will listen to. This parameter is optional: if not given, the default port is 8000.
- `file`: the path to the files to be copied (no matter if they are relative or absolute).

## Examples

To copy `foo.txt` and `bar.txt` via the default port (8000):
``` bash
hcp foo.txt bar.csv
```

To copy them on a custom port, for example 777:
``` bash
hcp :777 foo.txt bar.csv
```
## Dependencies

This program depends on [Mux](https://github.com/gorilla/mux).

## Compile

To compile the code it is sufficient to call `go build` on `hcp.go`:

``` bash
go build hcp.go
```

The executable file can be relocated anywhere.

You can also choose to not compile the code and to run it as a script:

``` bash
go run hcp.go [:port] file [...]
```

