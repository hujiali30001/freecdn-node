// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package compressions

import (
	teaconst "github.com/hujiali30001/freecdn-node/internal/const"
	"io"
)

var sharedGzipWriterPool *WriterPool

func init() {
	if !teaconst.IsMain {
		return
	}

	sharedGzipWriterPool = NewWriterPool(CalculatePoolSize(), func(writer io.Writer, level int) (Writer, error) {
		return newGzipWriter(writer)
	})
}
