package persist

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSave() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("item saver:got item #%d: %v", itemCount, item)
			Save(item)
		}
	}()
	return out
}

// 信息存入到es
func Save(item interface{}) {
	// 创建tcp客户端，不进行集群状态检查
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Index().Index("zhenai").Type("doc").BodyJson(item).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v",resp)
}
