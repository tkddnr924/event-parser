package evtx

import (
	"encoding/binary"
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

func parseEventHeader(header_byte []byte) EventHeader {
	_header := EventHeader{}

	// Magic (8 bytes)
	_header.Magic = string(header_byte[0:8])

	// old est chunk (8 bytes)
	oldestchunk := convertBuffer(header_byte[8:16])
	binary.Read(oldestchunk, binary.LittleEndian, &_header.oldestChunk)

	// current chunk number (8 bytes)
	currentChunkNum := convertBuffer(header_byte[16:24])
	binary.Read(currentChunkNum, binary.LittleEndian, &_header.currentChunkNum)

	// next Record num (8 bytes)
	nextRecordNum := convertBuffer(header_byte[24:32])
	binary.Read(nextRecordNum, binary.LittleEndian, &_header.nextRecordNum)

	// header part 1 length : uint32 (4 bytes)
	headerPart1Length := convertBuffer(header_byte[32:36])
	binary.Read(headerPart1Length, binary.LittleEndian, &_header.headerPart1Length)

	// minorVersion (2 bytes)
	minorVersion := convertBuffer(header_byte[36:38])
	binary.Read(minorVersion, binary.LittleEndian, &_header.minorVersion)

	// majorVersion (2 bytes)
	majorVersion := convertBuffer(header_byte[38:40])
	binary.Read(majorVersion, binary.LittleEndian, &_header.majorVersion)

	// headerSize (2 bytes)
	headerSize := convertBuffer(header_byte[40:42])
	binary.Read(headerSize, binary.LittleEndian, &_header.headerSize)

	// chunk count (2 bytes)
	chunkCount := convertBuffer(header_byte[42:44])
	binary.Read(chunkCount, binary.LittleEndian, &_header.chunkCount)

	// unknown char 76 bytes
	_header.unKnown = string(header_byte[44:120])

	// flags (4 bytes)
	flags := convertBuffer(header_byte[120:124])
	binary.Read(flags, binary.LittleEndian, &_header.flags)

	// checkSum (4 bytes)
	checkSum := convertBuffer(header_byte[124:128])
	binary.Read(checkSum, binary.LittleEndian, &_header.checkSum)

	return _header
}
