package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/caiquetorres/compression-tool/encoder"
)

type commandOptions struct {
	Type       string
	InputFile  string
	OutputFile string
}

func main() {
	errLogger := log.New(os.Stderr, "ccct: ", 0)
	if len(os.Args) < 2 {
		errLogger.Println("missing command, expected 'encode' or 'decode'")
		return
	}
	command := os.Args[1]
	options, err := parseCommands(command, os.Args[2:])
	if err != nil {
		errLogger.Println(err)
		return
	}
	switch options.Type {
	case "encode":
		err := encode(options.InputFile, options.OutputFile)
		if err != nil {
			log.Fatal(err)
		}
	case "decode":
		decode(options.InputFile, options.OutputFile)
	default:
		errLogger.Printf("unknown command: %s\n", options.Type)
	}
}

func parseCommands(command string, args []string) (*commandOptions, error) {
	encodeCmd := flag.NewFlagSet("encode", flag.ContinueOnError)
	decodeCmd := flag.NewFlagSet("decode", flag.ContinueOnError)
	var inputFile, ouputFile string
	switch command {
	case "encode":
		encodeCmd.StringVar(&inputFile, "f", "", "file to encode")
		encodeCmd.StringVar(&ouputFile, "o", "", "output file")
		if err := encodeCmd.Parse(args); err != nil {
			return nil, fmt.Errorf("failed to parse 'encode' command: %w", err)
		}
	case "decode":
		decodeCmd.StringVar(&inputFile, "f", "", "file to decode")
		decodeCmd.StringVar(&ouputFile, "o", "", "output file")
		if err := decodeCmd.Parse(args); err != nil {
			return nil, fmt.Errorf("failed to parse 'decode' command: %w", err)
		}
	default:
		return nil, fmt.Errorf("unknown command: %s", command)
	}
	if inputFile == "" {
		return nil, fmt.Errorf("-f flag (input file) is required for '%s' command", command)
	}
	if ouputFile == "" {
		return nil, fmt.Errorf("-o flag (output file) is required for '%s' command", command)
	}
	return &commandOptions{
		Type:       command,
		InputFile:  inputFile,
		OutputFile: ouputFile,
	}, nil
}

func encode(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}
	outputFile, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	return encoder.Encode(inputFile, outputFile)
}

func decode(srcFilePath, outFilePath string) {
	fmt.Println("decoding...")
}
