package healthroutes

const URL = "status"

type StatusRoutes struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
	Memory   Memory   `json:"memory"`
}

type Server struct {
	ConnectionsActive   int `json:"connections_active"`
	ConnectionsReading  int `json:"connections_reading"`
	ConnectionsAccepted int `json:"connections_accepted"`
	ConnectionsWriting  int `json:"connections_writing"`
	TotalRequests       int `json:"total_requests"`
	ConnectionsWaiting  int `json:"connections_waiting"`
	ConnectionsHandled  int `json:"connections_handled"`
}

type Database struct {
	Reachable bool `json:"reachable"`
}

type MemoryInfo struct {
	AllocatedSlabs string `json:"allocated_slabs"`
	Capacity       string `json:"capacity"`
}

type WorkersLuaVms struct {
	Pid             int    `json:"pid"`
	HTTPAllocatedGc string `json:"http_allocated_gc"`
}

type Memory struct {
	LuaSharedDict map[string]MemoryInfo `json:"lua_shared_dicts"`
	WorkersLuaVms []WorkersLuaVms       `json:"workers_lua_vms"`
}
