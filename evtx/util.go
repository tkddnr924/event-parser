package evtx

import "bytes"

func convertBuffer(data []byte) *bytes.Buffer {
	bufSlice := []byte(data)
	buf := bytes.NewBuffer(bufSlice)
	return buf
}
