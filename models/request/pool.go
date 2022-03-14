package request

import (
	"afkl/fumofuzzer/models/response"
	"log"
	"sync"

	"github.com/panjf2000/ants/v2"
)

var TaskWaitGroup sync.WaitGroup

type RequetsPool struct {
	Pool *ants.PoolWithFunc
}

func (pool *RequetsPool) Wait() {
	TaskWaitGroup.Wait()
}

func (pool *RequetsPool) Close() {
	pool.Pool.Release()
}

func (pool *RequetsPool) Submit(fuzz *FuzzRequestTemplate, coll response.FuzzResponseCollector) {
	TaskWaitGroup.Add(1)
	defer TaskWaitGroup.Done()

	reqChan := make(chan *FuzzRequest)
	fuzz.Iterator.Exec(fuzz.Payloads)

	go func() {
		for req := range reqChan {
			if err := pool.Pool.Invoke(req); err != nil {
				log.Println(err.Error())
			}
		}
	}()

	for {
		data, alive := <-fuzz.Iterator.Channel()
		if alive {
			reqChan <- fuzz.GenerateFuzzRequest(data, coll)
			TaskWaitGroup.Add(1)
		} else {
			break
		}
	}

}

func FuzzRequestWorker(fuzzReqInterface interface{}) {
	defer TaskWaitGroup.Done()
	fuzzReq, ok := fuzzReqInterface.(*FuzzRequest)
	if !ok {
		log.Fatalln("Failed to convert to type FuzzRequest")
	}

	resp, err := fuzzReq.Request.Send()
	if err != nil {
		// fuzzReq.Channel() <- resp
		log.Printf("%v => [ERR]%s\n", fuzzReq.Data, resp.Request.URL)
	} else {
		fuzzReq.collector.Channel() <- resp
		// log.Printf("%v => [%d]%s\n", fuzzReq.Data, resp.StatusCode(), resp.Request.URL)
	}
}

func InitRequestPool(threadNum int) *RequetsPool {
	pool, err := ants.NewPoolWithFunc(threadNum, FuzzRequestWorker)
	if err != nil {
		log.Fatalf("Init Request Pool Error: %s\n", err.Error())
	}

	return &RequetsPool{
		Pool: pool,
	}
}
