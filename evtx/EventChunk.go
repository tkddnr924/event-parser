package evtx

import "encoding/binary"

const EventChunkByte = 0x80
const EventChunkSize = 0x200

type EventChunk struct {
	Magic               string
	FirstRecordNum      int64
	LastRecordNum       int64
	FirstRecordId       int64
	LastRecordId        int64
	HeaderSize          uint32
	LastRecordOffset    uint32
	FreeSpaceOffset     uint32
	EventRecordChecksum uint32
	UnKnown             string
	HeaderCRC           uint32
}

func ParseEventChunk(data []byte) EventChunk {
	_chunk := EventChunk{}

	// Magic
	_chunk.Magic = string(data[0:8])

	// first record num
	FirstRecordNum := convertBuffer(data[8:16])
	_ = binary.Read(FirstRecordNum, binary.LittleEndian, &_chunk.FirstRecordNum)

	// last record num
	LastRecordNum := convertBuffer(data[16:24])
	_ = binary.Read(LastRecordNum, binary.LittleEndian, &_chunk.LastRecordNum)

	// first event record identifier
	FirstRecordId := convertBuffer(data[24:32])
	_ = binary.Read(FirstRecordId, binary.LittleEndian, &_chunk.FirstRecordId)

	// Last event record identifier
	LastRecordId := convertBuffer(data[32:40])
	_ = binary.Read(LastRecordId, binary.LittleEndian, &_chunk.LastRecordId)

	// Header Size
	HeaderSize := convertBuffer(data[40:44])
	_ = binary.Read(HeaderSize, binary.LittleEndian, &_chunk.HeaderSize)

	// LastRecordOffset
	LastRecordOffset := convertBuffer(data[44:48])
	_ = binary.Read(LastRecordOffset, binary.LittleEndian, &_chunk.LastRecordOffset)

	// FreeSpaceOffset
	FreeSpaceOffset := convertBuffer(data[48:52])
	_ = binary.Read(FreeSpaceOffset, binary.LittleEndian, &_chunk.FreeSpaceOffset)

	// EventRecordChecksum
	EventRecordChecksum := convertBuffer(data[52:56])
	_ = binary.Read(EventRecordChecksum, binary.LittleEndian, &_chunk.EventRecordChecksum)

	// UnKnown
	_chunk.UnKnown = string(data[56:124])

	// HeaderCRC
	HeaderCRC := convertBuffer(data[124:128])
	_ = binary.Read(HeaderCRC, binary.LittleEndian, &_chunk.HeaderCRC)

	return _chunk
}
