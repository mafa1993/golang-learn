package persist

import (
	"crawler/engine"

	"crawler/persist"

	"gopkg.in/olivere/elastic.v5"
)

// 数据存储服务，

type ItemSaveServer struct {
	Client *elastic.Client
}

// 存储服务
func (s ItemSaveServer) Save(item engine.Item, result *string) error {

	_, err := persist.Save(s.Client, item)
	if err != nil {
		return err
	}
	*result = "ok"
	return nil

}
