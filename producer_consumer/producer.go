package producer_consumer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// This file contains the code for producer which in this case is a restaurent which
// keeps producing multiple orders and the consumer i.e delivery men will pick it up

var dishes []string = []string{
	"Butter Chicken",
	"Pizza",
	"Burger",
	"Biryani",
	"Pasta",
}

type OrderIdGenerator struct {
	initialOrderId int
	mx             sync.Mutex
}

func (o *OrderIdGenerator) GetOrderId() int {
	defer o.mx.Unlock()
	o.mx.Lock()
	o.initialOrderId++
	return o.initialOrderId
}

func Producer(restaurent *Restaurent, mwg *sync.WaitGroup) {
	defer mwg.Done()
	// We definer an initial orderId for order Generator struct
	orderIdGenerator := OrderIdGenerator{
		initialOrderId: 1000,
	}

	// We have 5 chefs who can make 10 dishes everyday and after that they are done for the day
	// These 5 chefs can make dishes concurrently.
	wg := &sync.WaitGroup{}

	for chef := 1; chef < 6; chef++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				dishToCook := dishes[rand.Intn(len(dishes))]
				orderId := orderIdGenerator.GetOrderId()
				fmt.Printf("Chef %v has started cooking order %v which is %v\n", c, orderId, dishToCook)
				time.Sleep(time.Second * 2)
				fmt.Printf("Chev %v has cooked order %v and dish is ready to deliver\n", c, orderId)
				order := Orders{
					OrderId:   uint64(orderId),
					OrderName: dishToCook,
				}
				restaurent.OrdersToDeliver <- order
			}

		}(chef)
	}

	wg.Wait()
	fmt.Println("All cooks are done for the day!!!")
	restaurent.CloseRestaurent <- true
}
