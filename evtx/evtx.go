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
		_ = _file.Close()
	}(_file)

	// Parse event header
	headerByte := make([]byte, EventHeaderByte)
	_, err = _file.Read(headerByte)

	if err != nil {
		fmt.Print("Error: fail read file")
	}
	defer func(_file *os.File) {
		_ = _file.Close()
	}(_file)

	_header := parseEventHeader(headerByte)

	// Parse event chunk
	// chunk offset == 0x00 + headerSize
	_, err = _file.Seek(int64(_header.headerSize), 0)
	chunkByte := make([]byte, EventChunkByte)

	if err != nil {
		fmt.Print("Error: fail seek file")
	}
	defer func(_file *os.File) {
		_ = _file.Close()
	}(_file)

	_, err = _file.Read(chunkByte)
	if err != nil {
		fmt.Print("Error: fail read event chunk")
	}
	defer func(_file *os.File) {
		_ = _file.Close()
	}(_file)

	_chunk := ParseEventChunk(chunkByte)

	fmt.Printf("%+v\n", _header)
	fmt.Printf("%+v\n", _chunk)
}
