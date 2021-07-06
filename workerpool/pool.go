package workerpool

import (
	"bufio"
	"context"
	"io"
	"sync"

	"github.com/AlexDespod/sortingmodule/structs"
	"github.com/AlexDespod/sortingmodule/utils"
)

// Pool воркера
type Pool struct {
	Config
	dataStream chan structs.DataChanItem
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
}

// NewPool инициализирует новый пул с заданными задачами и

func NewPool(c Config) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &Pool{
		Config:     c,
		dataStream: make(chan structs.DataChanItem, 1000),
		ctx:        ctx,
		cancel:     cancel,
	}
}

type Config struct {
	Filename    string
	PerChunk    int
	Concurrency int
	chunksDir   string
}

// Run запускает всю работу в Pool и блокирует ее до тех пор,
// пока она не будет закончена.
func (p *Pool) Run() error {
	for i := 1; i <= p.Concurrency; i++ {
		worker := NewWorker(p.ctx, p.cancel, p.dataStream, i, p.chunksDir)
		worker.Start(&p.wg, p.Config)
	}

	inFile, err := utils.GetFile(p.Filename)

	if err != nil {
		inFile.Close()
		close(p.dataStream)
		return err
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)

	for {

		select {

		case <-p.ctx.Done():

			close(p.dataStream)
			p.wg.Wait()

			// handling errors here )

			return p.ctx.Err()
		default:

		}

		item, err := utils.ReadOneLine(reader)

		if err != nil && err != io.EOF {
			close(p.dataStream)
			p.cancel()
			return err
		}

		if err == io.EOF {
			p.dataStream <- structs.DataChanItem{SortItem: item, Err: err}
			break
		}
		p.dataStream <- structs.DataChanItem{SortItem: item, Err: nil}
	}
	close(p.dataStream)

	p.wg.Wait()

	return nil
}
