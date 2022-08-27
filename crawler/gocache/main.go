package gocache

import "sync"

type MapCache struct {
	data map[string]interface{}
	lock sync.RWMutex
}

func main() {

}