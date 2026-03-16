// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package nodes

import (
	"github.com/hujiali30001/freecdn-common/pkg/nodeutils"
	"github.com/hujiali30001/freecdn-common/pkg/serverconfigs"
	"github.com/hujiali30001/freecdn-node/internal/remotelogs"
)

// 健康检查
func (this *HTTPRequest) doHealthCheck(key string, isHealthCheck *bool) (stop bool) {
	this.tags = append(this.tags, "healthCheck")

	this.RawReq.Header.Del(serverconfigs.HealthCheckHeaderName)

	data, err := nodeutils.Base64DecodeMap(key)
	if err != nil {
		remotelogs.Error("HTTP_REQUEST_HEALTH_CHECK", "decode key failed: "+err.Error())
		return
	}
	*isHealthCheck = true

	this.web.StatRef = nil

	if !data.GetBool("accessLogIsOn") {
		this.disableLog = true
	}

	if data.GetBool("onlyBasicRequest") {
		return true
	}

	return
}
