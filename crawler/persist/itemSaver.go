package persist

import (
	"context"
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
/**
 * @param item interface{}  json 数据，存入到es
 * @return id，error  存入到es的id和 此函数运行过程中的报错
 */
func Save(item interface{}) (id string, err error) {
	// 创建tcp客户端，不进行集群状态检查
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	resp, err := client.Index().Index("zhenai").Type("doc").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil

}
