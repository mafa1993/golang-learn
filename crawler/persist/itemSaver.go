package persist

import "log"

func ItemSave() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver:got item #%d: %v", itemCount, item)
		}
	}()
	return out
}
