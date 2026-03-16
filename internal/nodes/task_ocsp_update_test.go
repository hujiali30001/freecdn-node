// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package nodes_test

import (
	"github.com/hujiali30001/freecdn-node/internal/nodes"
	"testing"
)

func TestOCSPUpdateTask_Loop(t *testing.T) {
	var task = &nodes.OCSPUpdateTask{}
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
