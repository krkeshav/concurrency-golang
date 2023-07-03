package producer_consumer

import "sync"

type Restaurent struct {
	OrdersToDeliver chan Orders
	CloseRestaurent chan bool
}

type Orders struct {
	OrderId   uint64
	OrderName string
}

func StartProducerConsumer() {
	restaurent := Restaurent{
		OrdersToDeliver: make(chan Orders),
		CloseRestaurent: make(chan bool),
	}
	mwg := &sync.WaitGroup{}
	mwg.Add(2)
	go Producer(&restaurent, mwg)
	go Consumer(&restaurent, mwg)
	mwg.Wait()
}
