package evtx

import (
	"fmt"
	"os"
)

func Open(filePath string) {

	// Open event log file
	_file, err := os.Open(filePath)

	if err != nil {
		fmt.Print("Error: failed open file")
	}
	defer func(_file *os.File) {
		err := _file.Close()
		if err != nil {

		}
	}(_file)

	// Parse event header
	headerByte := make([]byte, EventHeaderByte)
	_, err = _file.Read(headerByte)

	if err != nil {
		fmt.Print("Error: fail read file")
	}
	defer func(_file *os.File) {
		err := _file.Close()
		if err != nil {

		}
	}(_file)

	_header := parseEventHeader(headerByte)

	// Parse event chunk
	_, err = _file.Seek(int64(_header.headerSize), 0)
	chunkByte := make([]byte, EventChunkByte)

	if err != nil {
		fmt.Print("Error: fail seek file")
	}
	defer func(_file *os.File) {
		err := _file.Close()
		if err != nil {

		}
	}(_file)

	_chunk := ParseEventChunk(chunkByte)

	fmt.Printf("%+v\n", _header)
	fmt.Printf("%+v\n", _chunk)
}
