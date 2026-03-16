// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package fsutils_test

import (
	fsutils "github.com/hujiali30001/freecdn-node/internal/utils/fs"
	"testing"
)

func TestCheckDiskWritingSpeed(t *testing.T) {
	t.Log(fsutils.CheckDiskWritingSpeed())
}

func TestCheckDiskIsFast(t *testing.T) {
	t.Log(fsutils.CheckDiskIsFast())
}