# simdisk

A simple file sharing tool that uses a local disk file system to share files between clients.

## Before you run

Rename `config.example.toml` to `config.toml` and edit it to your needs.

## Running
To run the server, use the following command:

```bash
go run main.go
```

## Building and Releasing

To build the server, use the following command:

```bash
go build -o build/simdisk
```

To build the server for a specific OS and architecture, use the environment variables `GOOS` and `GOARCH`. 

For example, to build for Linux on amd64 architecture, you can use:

```bash
env GOOS=linux GOARCH=amd64 go build -o build/simdisk-linux-amd64
```

This produces a binary in the `build` directory named `simdisk-linux-amd64`.
It is a static binary, so it can be run on any amd64 Linux system without needing to install Go.
