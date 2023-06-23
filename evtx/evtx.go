package evtx

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

const EVENT_HEADER_BYTE = 0x80

type EventHeader struct {
	Magic             string
	oldestChunk       int64
	currentChunkNum   int64
	nextRecordNum     int64
	headerPart1Length uint32
	minorVersion      uint16
	majorVersion      uint16
	headerSize        uint16
	chunkCount        uint16
	unKnown           string
	flags             uint32
	checkSum          uint32
}

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

	fmt.Println(header_byte[24:32])

	_header := createEventHeader(header_byte)

	fmt.Println(_header.Magic)
	fmt.Println(_header.oldestChunk)
	fmt.Println(_header.currentChunkNum)
	fmt.Println(_header.nextRecordNum)
}

func createEventHeader(header_byte []byte) EventHeader {
	_header := EventHeader{}
	_header.Magic = string(header_byte[0:8])

	// old est chunk
	oldestchunk := convert2int64(header_byte[8:16])
	binary.Read(oldestchunk, binary.LittleEndian, &_header.oldestChunk)

	// current chunk number
	currentChunkNum := convert2int64(header_byte[16:24])
	binary.Read(currentChunkNum, binary.LittleEndian, &_header.currentChunkNum)

	// next Record num
	nextRecordNum := convert2int64(header_byte[24:32])
	binary.Read(nextRecordNum, binary.LittleEndian, &_header.nextRecordNum)

	return _header
}

func convert2int64(data []byte) *bytes.Buffer {
	bufSlice := []byte(data)
	buf := bytes.NewBuffer(bufSlice)
	return buf
}
