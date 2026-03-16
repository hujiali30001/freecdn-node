package nodes

import (
	"github.com/hujiali30001/freecdn-node/internal/utils/testutils"
	"testing"
)

func TestAPIStream_Start(t *testing.T) {
	if !testutils.IsSingleTesting() {
		return
	}

	apiStream := NewAPIStream()
	apiStream.Start()
}
