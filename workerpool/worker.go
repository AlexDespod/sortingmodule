package workerpool

import (
	"context"
	"fmt"
	"sync"

	"github.com/AlexDespod/sortingmodule/structs"
	"github.com/AlexDespod/sortingmodule/utils"
)

// Worker контролирует всю работу
type Worker struct {
	ID       int
	dataChan chan structs.DataChanItem
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewWorker возвращает новый экземпляр worker-а
func NewWorker(ctx context.Context, cancel context.CancelFunc, channel chan structs.DataChanItem, ID int) *Worker {
	return &Worker{
		ID:       ID,
		dataChan: channel,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// запуск worker
func (wr *Worker) Start(wg *sync.WaitGroup, c Config) {
	fmt.Printf("Starting worker %d\n", wr.ID)

	wg.Add(1)

	go func() {
		defer wg.Done()
		select {
		case <-wr.ctx.Done():
			return
		default:
		}
		err := utils.ProcessDataAsync(wr.ctx, wr.cancel, wr.dataChan, c.PerChunk, wr.ID)
		if err != nil {
			wr.cancel()
		}
	}()
}

func Process(id int, data structs.DataChanItem) error {

	return nil
}
