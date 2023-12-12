# Cli Password Generator Go

## Description

This project is a password generator implemented in [Go](https://golang.org/). It utilizes the [Cobra CLI](https://github.com/spf13/cobra) library for command-line interface functionality.

## Usage

To use this project, follow these steps:

1. Run `go build` to build the project.
2. Run `go install` to install the project.
3. Execute the following command to generate a password:

```bash
passwordGen generate -l 20 -d -s
```

After running the command, you will see the generated password printed to the console, and it will also be copied to the clipboard. The output will look similar to the following:

```bash
Generating password...
y@7Fp&4oL!vN\*Q9m$5S1
Password copied to clipboard!
```
