package request

import (
	"log"
	"sync"

	"github.com/panjf2000/ants/v2"
)

type RequetsPool struct {
	Pool          *ants.PoolWithFunc
	TaskWaitGroup sync.WaitGroup
}

func fetch(value interface{}) {

}

func Submit() {

}

func InitRequestPool(threadNum int) *RequetsPool {
	pool, err := ants.NewPoolWithFunc(threadNum, fetch)
	if err != nil {
		log.Fatalf("Init Request Pool Error: %s\n", err.Error())
	}

	return &RequetsPool{
		Pool: pool,
	}
}
