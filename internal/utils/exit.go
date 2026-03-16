// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package utils

import (
	"github.com/hujiali30001/freecdn-node/internal/events"
	"os"
)

func Exit() {
	events.Notify(events.EventTerminated)
	os.Exit(0)
}
