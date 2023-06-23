package main

import "event-parser/evtx"

const EVENT_LOG = "System.evtx"

func main() {
	evtx.Open(EVENT_LOG)
}
