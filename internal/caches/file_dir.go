// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package caches

import "github.com/hujiali30001/freecdn-common/pkg/serverconfigs/shared"

type FileDir struct {
	Path     string
	Capacity *shared.SizeCapacity
	IsFull   bool
}
