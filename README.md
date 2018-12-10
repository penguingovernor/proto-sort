# Proto-Sort

Proto-Sort is a simple cli tool that was a learning exercise for protocol buffers.

## Installation

`go get -u github.com/penguingovernor/proto-sort`

## Usage

proto-sort gets input via cli args and stores them in a protobuf file.

Then the user can chose to view the entered items or sort them.

i.e.

```shell
$ proto-sort add 1 842 43 756 93 3
$ proto-sort add 9 8
$ proto-sort list # Prints 1 842 43 756 93 3 9 8
$ proto-sort sort # Prints 1 3 8 9 43 93 756 842
```

For more help regarding flags and commands run `proto-sort help`
