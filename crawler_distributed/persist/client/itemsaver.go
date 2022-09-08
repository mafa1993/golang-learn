package client

// engine发送item到这个client  client发送rpc调用
// 将原本的itemsaver改为这个

import (
	"crawler/engine"
	"log"
	"rpcc"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func ItemSave(host string) (chan engine.Item, error) {
	// 创建rpc客户端
	client, err := rpcc.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("item saver:got item #%d: %v", itemCount, item)
			//todo rpc save
			var result string
			err = client.Call("ItemSaveServer.Save", item, &result)

			if err != nil {
				log.Println(err)
			}
		}
	}()
	return out, nil
}
