package evtx

import "encoding/binary"

const EventChunkByte = 0x80

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
	LastRecordNum := convertBuffer(data[16:32])
	_ = binary.Read(LastRecordNum, binary.LittleEndian, &_chunk.LastRecordNum)

	return _chunk
}
