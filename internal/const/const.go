package teaconst

const (
	Version = "0.7.0"

	// GoEdgeVersion 是本项目基于的 GoEdge 上游版本
	GoEdgeVersion = "0.7.0"

	ProductName = "Edge Node"
	ProcessName = "edge-node"

	Role = "node"

	EncryptMethod = "aes-256-cfb"

	// SystemdServiceName systemd
	SystemdServiceName = "edge-node"

	AccessLogSockName    = "edge-node.accesslog"
	CacheGarbageSockName = "edge-node.cache.garbage"

	EnableKVCacheStore = true // determine store cache keys in KVStore or sqlite
)
