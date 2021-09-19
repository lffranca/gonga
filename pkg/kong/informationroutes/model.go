package informationroutes

const URL = "/"

type InformationRoutes struct {
	Plugins       Plugins       `json:"plugins"`
	LuaVersion    string        `json:"lua_version"`
	Configuration Configuration `json:"configuration"`
	PIDs          PIDs          `json:"pids"`
	Tagline       string        `json:"tagline"`
	Hostname      string        `json:"hostname"`
	Version       string        `json:"version"`
	Timers        Timers        `json:"timers"`
	NodeID        string        `json:"node_id"`
}

type Plugins struct {
	EnabledInCluster  []string        `json:"enabled_in_cluster"`
	AvailableOnServer map[string]bool `json:"available_on_server"`
}

type AdminListeners struct {
	SSL           bool   `json:"ssl"`
	Deferred      bool   `json:"deferred"`
	HTTP2         bool   `json:"http2"`
	Port          int    `json:"port"`
	Listener      string `json:"listener"`
	ProxyProtocol bool   `json:"proxy_protocol"`
	IP            string `json:"ip"`
	Bind          bool   `json:"bind"`
	BacklogD      bool   `json:"backlog=%d+"`
	ReusePort     bool   `json:"reuseport"`
}

type ProxyListeners struct {
	Ssl           bool   `json:"ssl"`
	Deferred      bool   `json:"deferred"`
	HTTP2         bool   `json:"http2"`
	Port          int    `json:"port"`
	Listener      string `json:"listener"`
	ProxyProtocol bool   `json:"proxy_protocol"`
	IP            string `json:"ip"`
	Bind          bool   `json:"bind"`
	BacklogD      bool   `json:"backlog=%d+"`
	ReusePort     bool   `json:"reuseport"`
}

type NginxAdminDirectives struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EnabledHeaders struct {
	XKongProxyLatency    bool `json:"X-Kong-Proxy-Latency"`
	Via                  bool `json:"Via"`
	XKongResponseLatency bool `json:"X-Kong-Response-Latency"`
	XKongAdminLatency    bool `json:"X-Kong-Admin-Latency"`
	XKongUpstreamLatency bool `json:"X-Kong-Upstream-Latency"`
	XKongUpstreamStatus  bool `json:"X-Kong-Upstream-Status"`
	ServerTokens         bool `json:"server_tokens"`
	LatencyTokens        bool `json:"latency_tokens"`
	Server               bool `json:"Server"`
}
type StatusListeners struct {
	Ssl      bool   `json:"ssl"`
	IP       string `json:"ip"`
	Listener string `json:"listener"`
	Port     int    `json:"port"`
}

type NginxProxyDirectives struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NginxMainDirectives struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NginxEventsDirectives struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NginxHTTPDirectives struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NginxStreamDirectives struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Configuration struct {
	NginxPid                          string                  `json:"nginx_pid"`
	ProxySslEnabled                   bool                    `json:"proxy_ssl_enabled"`
	Role                              string                  `json:"role"`
	Headers                           []string                `json:"headers"`
	Prefix                            string                  `json:"prefix"`
	ClusterV2                         bool                    `json:"cluster_v2"`
	NginxAdminClientBodyBufferSize    string                  `json:"nginx_admin_client_body_buffer_size"`
	NginxConf                         string                  `json:"nginx_conf"`
	Plugins                           []string                `json:"plugins"`
	NginxKongConf                     string                  `json:"nginx_kong_conf"`
	LoadedPlugins                     map[string]bool         `json:"loaded_plugins"`
	NginxKongStreamConf               string                  `json:"nginx_kong_stream_conf"`
	KongEnv                           string                  `json:"kong_env"`
	NginxHTTPLuaRegexMatchLimit       string                  `json:"nginx_http_lua_regex_match_limit"`
	SslCertDefault                    string                  `json:"ssl_cert_default"`
	SslCertKeyDefault                 string                  `json:"ssl_cert_key_default"`
	SslCertDefaultEcdsa               string                  `json:"ssl_cert_default_ecdsa"`
	PgHost                            string                  `json:"pg_host"`
	ClientSslCertDefault              string                  `json:"client_ssl_cert_default"`
	DNSResolver                       interface{}             `json:"dns_resolver"`
	DNSHostsFile                      string                  `json:"dns_hostsfile"`
	SslCipherSuite                    string                  `json:"ssl_cipher_suite"`
	DNSNotFoundTTL                    int                     `json:"dns_not_found_ttl"`
	MemCacheSize                      string                  `json:"mem_cache_size"`
	DNSOrder                          []string                `json:"dns_order"`
	PgUser                            string                  `json:"pg_user"`
	StatusSslCertDefault              string                  `json:"status_ssl_cert_default"`
	StatusSslCertKeyDefault           string                  `json:"status_ssl_cert_key_default"`
	StatusSslCertDefaultEcdsa         string                  `json:"status_ssl_cert_default_ecdsa"`
	LuaPackageCPath                   string                  `json:"lua_package_cpath"`
	StatusSslCertKeyDefaultEcdsa      string                  `json:"status_ssl_cert_key_default_ecdsa"`
	HostPorts                         map[string]int          `json:"host_ports"`
	AnonymousReports                  bool                    `json:"anonymous_reports"`
	SslCert                           []string                `json:"ssl_cert"`
	CassandraKeyspace                 string                  `json:"cassandra_keyspace"`
	CassandraTimeout                  int                     `json:"cassandra_timeout"`
	PgTimeout                         int                     `json:"pg_timeout"`
	UpstreamKeepaliveIdleTimeout      int                     `json:"upstream_keepalive_idle_timeout"`
	UpstreamKeepaliveMaxRequests      int                     `json:"upstream_keepalive_max_requests"`
	Kic                               bool                    `json:"kic"`
	AdminListeners                    []AdminListeners        `json:"admin_listeners"`
	ProxyListeners                    []ProxyListeners        `json:"proxy_listeners"`
	StreamListeners                   interface{}             `json:"stream_listeners"`
	PluginServerNames                 interface{}             `json:"pluginserver_names"`
	UpstreamKeepalivePoolSize         int                     `json:"upstream_keepalive_pool_size"`
	LuaPackagePath                    string                  `json:"lua_package_path"`
	GoPluginsDir                      string                  `json:"go_plugins_dir"`
	PortMaps                          []string                `json:"port_maps"`
	GoPluginServerEXE                 string                  `json:"go_pluginserver_exe"`
	AdminListen                       []string                `json:"admin_listen"`
	StatusListen                      []string                `json:"status_listen"`
	StreamListen                      []string                `json:"stream_listen"`
	ClusterListen                     []string                `json:"cluster_listen"`
	AdminSslCert                      []string                `json:"admin_ssl_cert"`
	AdminSslCertKey                   []string                `json:"admin_ssl_cert_key"`
	StatusSslCert                     interface{}             `json:"status_ssl_cert"`
	StatusSslCertKey                  interface{}             `json:"status_ssl_cert_key"`
	DbResurrectTTL                    int                     `json:"db_resurrect_ttl"`
	NginxUser                         string                  `json:"nginx_user"`
	NginxMainUser                     string                  `json:"nginx_main_user"`
	NginxDaemon                       string                  `json:"nginx_daemon"`
	NginxMainDaemon                   string                  `json:"nginx_main_daemon"`
	NginxWorkerProcesses              string                  `json:"nginx_worker_processes"`
	NginxMainWorkerProcesses          string                  `json:"nginx_main_worker_processes"`
	TrustedIps                        interface{}             `json:"trusted_ips"`
	RealIPHeader                      string                  `json:"real_ip_header"`
	NginxProxyRealIPHeader            string                  `json:"nginx_proxy_real_ip_header"`
	RealIPRecursive                   string                  `json:"real_ip_recursive"`
	NginxProxyRealIPRecursive         string                  `json:"nginx_proxy_real_ip_recursive"`
	ClientMaxBodySize                 string                  `json:"client_max_body_size"`
	NginxHTTPClientMaxBodySize        string                  `json:"nginx_http_client_max_body_size"`
	NginxStatusDirectives             interface{}             `json:"nginx_status_directives"`
	NginxHTTPClientBodyBufferSize     string                  `json:"nginx_http_client_body_buffer_size"`
	NginxAdminDirectives              []NginxAdminDirectives  `json:"nginx_admin_directives"`
	PgPassword                        string                  `json:"pg_password"`
	PgSsl                             bool                    `json:"pg_ssl"`
	PgSslVerify                       bool                    `json:"pg_ssl_verify"`
	NginxSUpstreamDirectives          interface{}             `json:"nginx_supstream_directives"`
	PgSemaphoreTimeout                int                     `json:"pg_semaphore_timeout"`
	NginxSProxyDirectives             interface{}             `json:"nginx_sproxy_directives"`
	StreamProxySslEnabled             bool                    `json:"stream_proxy_ssl_enabled"`
	NginxHTTPUpstreamDirectives       interface{}             `json:"nginx_http_upstream_directives"`
	NginxHTTPStatusDirectives         interface{}             `json:"nginx_http_status_directives"`
	ClientSsl                         bool                    `json:"client_ssl"`
	CassandraPort                     int                     `json:"cassandra_port"`
	PgDatabase                        string                  `json:"pg_database"`
	CassandraSsl                      bool                    `json:"cassandra_ssl"`
	CassandraSslVerify                bool                    `json:"cassandra_ssl_verify"`
	CassandraWriteConsistency         string                  `json:"cassandra_write_consistency"`
	DNSErrorTTL                       int                     `json:"dns_error_ttl"`
	DNSStaleTTL                       int                     `json:"dns_stale_ttl"`
	Database                          string                  `json:"database"`
	NginxAdminClientMaxBodySize       string                  `json:"nginx_admin_client_max_body_size"`
	NginxStreamSslSessionTickets      string                  `json:"nginx_stream_ssl_session_tickets"`
	EnabledHeaders                    EnabledHeaders          `json:"enabled_headers"`
	ClientSslCertKeyDefault           string                  `json:"client_ssl_cert_key_default"`
	NginxStreamSslPreferServerCiphers string                  `json:"nginx_stream_ssl_prefer_server_ciphers"`
	CassandraReadConsistency          string                  `json:"cassandra_read_consistency"`
	AdminSslCertKeyDefault            string                  `json:"admin_ssl_cert_key_default"`
	DNSNoSync                         bool                    `json:"dns_no_sync"`
	CassandraLbPolicy                 string                  `json:"cassandra_lb_policy"`
	StatusListeners                   []StatusListeners       `json:"status_listeners"`
	ErrorDefaultType                  string                  `json:"error_default_type"`
	NginxErrLogs                      string                  `json:"nginx_err_logs"`
	ProxyListen                       []string                `json:"proxy_listen"`
	CassandraContactPoints            []string                `json:"cassandra_contact_points"`
	ClusterListeners                  interface{}             `json:"cluster_listeners"`
	CassandraReplStrategy             string                  `json:"cassandra_repl_strategy"`
	SslCertKeyDefaultEcdsa            string                  `json:"ssl_cert_key_default_ecdsa"`
	DbCacheTTL                        int                     `json:"db_cache_ttl"`
	CassandraReplFactor               int                     `json:"cassandra_repl_factor"`
	CassandraDataCenters              []string                `json:"cassandra_data_centers"`
	CassandraSchemaConsensusTimeout   int                     `json:"cassandra_schema_consensus_timeout"`
	NginxHTTPSslSessionTimeout        string                  `json:"nginx_http_ssl_session_timeout"`
	PgMaxConcurrentQueries            int                     `json:"pg_max_concurrent_queries"`
	PgPort                            int                     `json:"pg_port"`
	SslProtocols                      string                  `json:"ssl_protocols"`
	ClientBodyBufferSize              string                  `json:"client_body_buffer_size"`
	NginxHTTPSslProtocols             string                  `json:"nginx_http_ssl_protocols"`
	NginxStreamSslProtocols           string                  `json:"nginx_stream_ssl_protocols"`
	SslPreferServerCiphers            string                  `json:"ssl_prefer_server_ciphers"`
	NginxHTTPSslPreferServerCiphers   string                  `json:"nginx_http_ssl_prefer_server_ciphers"`
	DbCacheWarmupEntities             []string                `json:"db_cache_warmup_entities"`
	SslDhparam                        string                  `json:"ssl_dhparam"`
	NginxHTTPSslDhparam               string                  `json:"nginx_http_ssl_dhparam"`
	NginxStreamSslDhparam             string                  `json:"nginx_stream_ssl_dhparam"`
	SslSessionTickets                 string                  `json:"ssl_session_tickets"`
	NginxHTTPSslSessionTickets        string                  `json:"nginx_http_ssl_session_tickets"`
	DbUpdateFrequency                 int                     `json:"db_update_frequency"`
	SslSessionTimeout                 string                  `json:"ssl_session_timeout"`
	DbUpdatePropagation               int                     `json:"db_update_propagation"`
	NginxStreamSslSessionTimeout      string                  `json:"nginx_stream_ssl_session_timeout"`
	ProxyAccessLog                    string                  `json:"proxy_access_log"`
	ProxyErrorLog                     string                  `json:"proxy_error_log"`
	ProxyStreamAccessLog              string                  `json:"proxy_stream_access_log"`
	ProxyStreamErrorLog               string                  `json:"proxy_stream_error_log"`
	AdminAccessLog                    string                  `json:"admin_access_log"`
	AdminErrorLog                     string                  `json:"admin_error_log"`
	StatusAccessLog                   string                  `json:"status_access_log"`
	StatusErrorLog                    string                  `json:"status_error_log"`
	LogLevel                          string                  `json:"log_level"`
	NginxProxyDirectives              []NginxProxyDirectives  `json:"nginx_proxy_directives"`
	NginxUpstreamDirectives           interface{}             `json:"nginx_upstream_directives"`
	NginxOptimizations                bool                    `json:"nginx_optimizations"`
	LuaSslTrustedCertificate          interface{}             `json:"lua_ssl_trusted_certificate"`
	LuaSslVerifyDepth                 int                     `json:"lua_ssl_verify_depth"`
	LuaSslProtocols                   string                  `json:"lua_ssl_protocols"`
	NginxHTTPLuaSslProtocols          string                  `json:"nginx_http_lua_ssl_protocols"`
	NginxStreamLuaSslProtocols        string                  `json:"nginx_stream_lua_ssl_protocols"`
	LuaSocketPoolSize                 int                     `json:"lua_socket_pool_size"`
	WorkerConsistency                 string                  `json:"worker_consistency"`
	ClusterControlPlane               string                  `json:"cluster_control_plane"`
	CassandraUsername                 string                  `json:"cassandra_username"`
	SslCiphers                        string                  `json:"ssl_ciphers"`
	ClusterMTLS                       string                  `json:"cluster_mtls"`
	NginxMainDirectives               []NginxMainDirectives   `json:"nginx_main_directives"`
	NginxEventsDirectives             []NginxEventsDirectives `json:"nginx_events_directives"`
	NginxHTTPDirectives               []NginxHTTPDirectives   `json:"nginx_http_directives"`
	ClusterDataPlanePurgeDelay        int                     `json:"cluster_data_plane_purge_delay"`
	ClusterOcsp                       string                  `json:"cluster_ocsp"`
	UntrustedLua                      string                  `json:"untrusted_lua"`
	UntrustedLuaSandboxRequires       interface{}             `json:"untrusted_lua_sandbox_requires"`
	UntrustedLuaSandboxEnvironment    interface{}             `json:"untrusted_lua_sandbox_environment"`
	PgRoSslVerify                     bool                    `json:"pg_ro_ssl_verify"`
	AdminSslCertKeyDefaultEcdsa       string                  `json:"admin_ssl_cert_key_default_ecdsa"`
	SslCertKey                        []string                `json:"ssl_cert_key"`
	SslCertCsrDefault                 string                  `json:"ssl_cert_csr_default"`
	AdminAccLogs                      string                  `json:"admin_acc_logs"`
	NginxAccLogs                      string                  `json:"nginx_acc_logs"`
	NginxMainWorkerRLimitNoFile       string                  `json:"nginx_main_worker_rlimit_nofile"`
	AdminSslCertDefaultEcdsa          string                  `json:"admin_ssl_cert_default_ecdsa"`
	NginxEventsWorkerConnections      string                  `json:"nginx_events_worker_connections"`
	AdminSslCertDefault               string                  `json:"admin_ssl_cert_default"`
	NginxEventsMultiAccept            string                  `json:"nginx_events_multi_accept"`
	WorkerStateUpdateFrequency        int                     `json:"worker_state_update_frequency"`
	StatusSslEnabled                  bool                    `json:"status_ssl_enabled"`
	AdminSslEnabled                   bool                    `json:"admin_ssl_enabled"`
	CassandraRefreshFrequency         int                     `json:"cassandra_refresh_frequency"`
	NginxStreamDirectives             []NginxStreamDirectives `json:"nginx_stream_directives"`
	PgRoSsl                           bool                    `json:"pg_ro_ssl"`
}

type PIDs struct {
	Master  int   `json:"master"`
	Workers []int `json:"workers"`
}

type Timers struct {
	Running int `json:"running"`
	Pending int `json:"pending"`
}
