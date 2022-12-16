### go.mod file

The go.mod file is like a python 'requirements.tct' file,
it contains a list of the go modules need to run the go code
in this directory.

Created the `go.mod` file:
```
$ go mod init github.com/joshuamhtsang/network_prog_with_go
go: creating new go.mod: module github.com/joshuamhtsang/network_prog_with_go
```


### Running the Go files

To run the test functions:

```
$ go test -v listen_test.go
$ go test -v dial_test.go
```

Note that the '-v' option cause the t.Logf() functions
to actually output to terminal, otherwise these outputs
are not printed to terminal.