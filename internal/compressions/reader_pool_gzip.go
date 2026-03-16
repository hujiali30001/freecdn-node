// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package compressions

import (
	teaconst "github.com/hujiali30001/freecdn-node/internal/const"
	"io"
)

var sharedGzipReaderPool *ReaderPool

func init() {
	if !teaconst.IsMain {
		return
	}

	sharedGzipReaderPool = NewReaderPool(CalculatePoolSize(), func(reader io.Reader) (Reader, error) {
		return newGzipReader(reader)
	})
}
