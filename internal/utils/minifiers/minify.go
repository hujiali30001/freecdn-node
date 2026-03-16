// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package minifiers

import (
	"github.com/hujiali30001/freecdn-common/pkg/serverconfigs"
	"net/http"
)

// MinifyResponse minify response body
func MinifyResponse(config *serverconfigs.HTTPPageOptimizationConfig, url string, resp *http.Response) error {
	// stub
	return nil
}
