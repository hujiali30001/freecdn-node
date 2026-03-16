// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodes_test

import (
	"github.com/hujiali30001/freecdn-common/pkg/nodeconfigs"
	"github.com/hujiali30001/freecdn-node/internal/caches"
	"github.com/hujiali30001/freecdn-node/internal/nodes"
	"testing"
)

func TestHTTPCacheTaskManager_Loop(t *testing.T) {
	// initialize cache policies
	config, err := nodeconfigs.SharedNodeConfig()
	if err != nil {
		t.Fatal(err)
	}
	caches.SharedManager.UpdatePolicies(config.HTTPCachePolicies)

	var manager = nodes.NewHTTPCacheTaskManager()
	err = manager.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
