package producer_consumer

import (
	"fmt"
	"sync"
	"time"
)

// The Consumer function corresponds to delivery men who will pick up orders and deliver it
func Consumer(restaurent *Restaurent, mwg *sync.WaitGroup) {
	defer mwg.Done()

	for deliveryMan := 1; deliveryMan < 6; deliveryMan++ {
		go func(deliveryManNum int) {
			for {
				orderToDeliver := <-restaurent.OrdersToDeliver
				fmt.Printf("Order %v recieved and delivery man number %v has started delivery\n", orderToDeliver.OrderId, deliveryManNum)
				time.Sleep(time.Second)
				fmt.Printf("Order %v was delivered by delivery man number %v\n", orderToDeliver.OrderId, deliveryManNum)
			}
		}(deliveryMan)
	}

	<-restaurent.CloseRestaurent
	fmt.Println("It seems the restaurent has closed")
}
