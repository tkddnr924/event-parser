package main

import "event-parser/evtx"

const EventLog = "System.evtx"

func main() {
	evtx.Open(EventLog)
}
