package evtx

import (
	"fmt"
	"os"
)

func Open(filePath string) {
	_file, err := os.Open(filePath)

	if err != nil {
		fmt.Print("Error: failed open file")
	}
	defer _file.Close()

	header_byte := make([]byte, EVENT_HEADER_BYTE)
	_, err = _file.Read(header_byte)

	if err != nil {
		fmt.Print("Error: fail read file")
	}
	defer _file.Close()

	_header := parseEventHeader(header_byte)
	fmt.Printf("%+v\n", _header)
}
