package main

import "syncronization/syncs"

func main() {
	syncs.SyncSendReceive(5)
	syncs.StartProducerConsumer(10)
	syncs.StartProducerConsumer(5)
}
