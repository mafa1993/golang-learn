package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"log"
	"rpcc"

	"gopkg.in/olivere/elastic.v5"
)

// func ItemSave() (chan engine.Item, error) {
// 	out := make(chan engine.Item)

// 	// 创建tcp客户端，不进行集群状态检查
// 	client, err := elastic.NewClient(elastic.SetSniff(false))
// 	if err != nil {
// 		return nil, err
// 	}

// 	go func() {
// 		itemCount := 0
// 		for {
// 			item := <-out
// 			itemCount++
// 			log.Printf("item saver:got item #%d: %v", itemCount, item)
// 			Save(client, item)
// 		}
// 	}()
// 	return out, nil
// }

// 信息存入到es
/**
 * @param item interface{}  json 数据，存入到es
 * @return id，error  存入到es的id和 此函数运行过程中的报错
 */
func Save(client *elastic.Client, item engine.Item) (id string, err error) {

	// 容错处理
	if item.Id == "" {
		return "", errors.New("id不能为空")
	}

	// 自己指定id录入
	resp, err := client.Index().Index("zhenai").Type("doc").Id(item.Id).BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil

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
