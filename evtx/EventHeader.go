package evtx

import (
	"encoding/binary"
)

const EventHeaderByte = 0x80

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

func parseEventHeader(headerByte []byte) EventHeader {
	_header := EventHeader{}

	// Magic (8 bytes)
	_header.Magic = string(headerByte[0:8])

	// old est chunk (8 bytes)
	oldEstChunk := convertBuffer(headerByte[8:16])
	_ = binary.Read(oldEstChunk, binary.LittleEndian, &_header.oldestChunk)

	// current chunk number (8 bytes)
	currentChunkNum := convertBuffer(headerByte[16:24])
	_ = binary.Read(currentChunkNum, binary.LittleEndian, &_header.currentChunkNum)

	// next Record num (8 bytes)
	nextRecordNum := convertBuffer(headerByte[24:32])
	_ = binary.Read(nextRecordNum, binary.LittleEndian, &_header.nextRecordNum)

	// header part 1 length : uint32 (4 bytes)
	headerPart1Length := convertBuffer(headerByte[32:36])
	_ = binary.Read(headerPart1Length, binary.LittleEndian, &_header.headerPart1Length)

	// minorVersion (2 bytes)
	minorVersion := convertBuffer(headerByte[36:38])
	_ = binary.Read(minorVersion, binary.LittleEndian, &_header.minorVersion)

	// majorVersion (2 bytes)
	majorVersion := convertBuffer(headerByte[38:40])
	_ = binary.Read(majorVersion, binary.LittleEndian, &_header.majorVersion)

	// headerSize (2 bytes)
	headerSize := convertBuffer(headerByte[40:42])
	_ = binary.Read(headerSize, binary.LittleEndian, &_header.headerSize)

	// chunk count (2 bytes)
	chunkCount := convertBuffer(headerByte[42:44])
	_ = binary.Read(chunkCount, binary.LittleEndian, &_header.chunkCount)

	// unknown char 76 bytes
	_header.unKnown = string(headerByte[44:120])

	// flags (4 bytes)
	flags := convertBuffer(headerByte[120:124])
	_ = binary.Read(flags, binary.LittleEndian, &_header.flags)

	// checkSum (4 bytes)
	checkSum := convertBuffer(headerByte[124:128])
	_ = binary.Read(checkSum, binary.LittleEndian, &_header.checkSum)

	return _header
}
