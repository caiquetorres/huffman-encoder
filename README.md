# Compression Tool
This repository provides a solution for [Coding Challenge #03](https://codingchallenges.fyi/challenges/challenge-huffman), which involves creating an implementation of the huffman encoder and decoder

## Build and Run

To build and run the tool, use the following commands:

```sh
# Build the project
make build

# encoding
./bin/compression_tool encode -f file_to_encode.txt -o output_file.huff

# decoding
./bin/compression_tool decode -f file_to_decode.huff -o decoded_file.txt
```

<!-- ## Test

To run tests for the tool, use the following command:

```sh
# Run tests
make test
``` -->
