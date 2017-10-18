// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Copyright 2017 Signal 18 SARL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <svaroqui@gmail.com>
// This source code is licensed under the GNU General Public License, version 3.
// Redistribution/Reuse of this code is permitted under the GNU v3 license, as
// an additional term, ALL code must carry the original Author(s) credit in comment form.
// See LICENSE in this directory for the integral text.

package config

type Config struct {
	BaseDir                            string `mapstructure:"monitoring-basedir"`
	WorkingDir                         string `mapstructure:"monitoring-datadir"`
	ShareDir                           string `mapstructure:"monitoring-sharedir"`
	MonitoringTicker                   int64  `mapstructure:"monitoring-ticker"`
	Socket                             string `mapstructure:"monitoring-socket"`
	TunnelHost                         string `mapstructure:"monitoring-tunnel-host"`
	TunnelCredential                   string `mapstructure:"monitoring-tunnel-credential"`
	MonitorAddress                     string `mapstructure:"monitoring-address"`
	Interactive                        bool   `mapstructure:"interactive"`
	Verbose                            bool   `mapstructure:"verbose"`
	LogFile                            string `mapstructure:"log-file"`
	LogSyslog                          bool   `mapstructure:"log-syslog"`
	LogLevel                           int    `mapstructure:"log-level"`
	User                               string `mapstructure:"db-servers-credential"`
	Hosts                              string `mapstructure:"db-servers-hosts"`
	HostsTLSCA                         string `mapstructure:"db-servers-tls-ca-cert"`
	HostsTLSKEY                        string `mapstructure:"db-servers-tls-client-key"`
	HostsTLSCLI                        string `mapstructure:"db-servers-tls-client-cert"`
	PrefMaster                         string `mapstructure:"db-servers-prefered-master"`
	IgnoreSrv                          string `mapstructure:"db-servers-ignore-hosts"`
	Timeout                            int    `mapstructure:"db-servers-connect-timeout"`
	ReadTimeout                        int    `mapstructure:"db-servers-read-timeout"`
	MariaDBBinaryPath                  string `mapstructure:"db-servers-binary-path"`
	DbServerLocality                   string `mapstructure:"db-servers-locality"`
	MasterConnectRetry                 int    `mapstructure:"replication-master-connect-retry"`
	RplUser                            string `mapstructure:"replication-credential"`
	MasterConn                         string `mapstructure:"replication-source-name"`
	ReplicationSSL                     bool   `mapstructure:"replication-use-ssl"`
	MultiMasterRing                    bool   `mapstructure:"replication-multi-master-ring"`
	MultiMasterWsrep                   bool   `mapstructure:"replication-multi-master-wsrep"`
	MultiMaster                        bool   `mapstructure:"replication-multi-master"`
	MultiTierSlave                     bool   `mapstructure:"replication-multi-tier-slave"`
	ReplicationNoRelay                 bool   `mapstructure:"replication-master-slave-never-relay"`
	SwitchWaitKill                     int64  `mapstructure:"switchover-wait-kill"`
	SwitchWaitTrx                      int64  `mapstructure:"switchover-wait-trx"`
	SwitchWaitWrite                    int    `mapstructure:"switchover-wait-write-query"`
	SwitchGtidCheck                    bool   `mapstructure:"switchover-at-equal-gtid"`
	SwitchSync                         bool   `mapstructure:"switchover-at-sync"`
	SwitchMaxDelay                     int64  `mapstructure:"switchover-max-slave-delay"`
	SwitchSlaveWaitCatch               bool   `mapstructure:"switchover-slave-wait-catch"`
	FailLimit                          int    `mapstructure:"failover-limit"`
	PreScript                          string `mapstructure:"failover-pre-script"`
	PostScript                         string `mapstructure:"failover-post-script"`
	ReadOnly                           bool   `mapstructure:"failover-readonly-state"`
	FailTime                           int64  `mapstructure:"failover-time-limit"`
	FailSync                           bool   `mapstructure:"failover-at-sync"`
	FailEventScheduler                 bool   `mapstructure:"failover-event-scheduler"`
	FailEventStatus                    bool   `mapstructure:"failover-event-status"`
	FailRestartUnsafe                  bool   `mapstructure:"failover-restart-unsafe"`
	MaxFail                            int    `mapstructure:"failover-falsepositive-ping-counter"`
	FailResetTime                      int64  `mapstructure:"failcount-reset-time"`
	FailMode                           string `mapstructure:"failover-mode"`
	FailMaxDelay                       int64  `mapstructure:"failover-max-slave-delay"`
	CheckFalsePositiveHeartbeat        bool   `mapstructure:"failover-falsepositive-heartbeat"`
	CheckFalsePositiveMaxscale         bool   `mapstructure:"failover-falsepositive-maxscale"`
	CheckFalsePositiveHeartbeatTimeout int    `mapstructure:"failover-falsepositive-heartbeat-timeout"`
	CheckFalsePositiveMaxscaleTimeout  int    `mapstructure:"failover-falsepositive-maxscale-timeout"`
	CheckFalsePositiveExternal         bool   `mapstructure:"failover-falsepositive-external"`
	CheckFalsePositiveExternalPort     int    `mapstructure:"failover-falsepositive-external-port"`
	Autorejoin                         bool   `mapstructure:"autorejoin"`
	AutorejoinFlashback                bool   `mapstructure:"autorejoin-flashback"`
	RejoinScript                       string `mapstructure:"autrejoin-script"`
	AutorejoinMysqldump                bool   `mapstructure:"autorejoin-mysqldump"`
	AutorejoinBackupBinlog             bool   `mapstructure:"autorejoin-backup-binlog"`
	AutorejoinSemisync                 bool   `mapstructure:"autorejoin-flashback-on-sync"`
	AutorejoinNoSemisync               bool   `mapstructure:"autorejoin-flashback-on-unsync"`
	AutorejoinSlavePositionalHearbeat  bool   `mapstructure:"autorejoin-slave-positional-hearbeat"`
	AutorejoinZFSFlashback             bool   `mapstructure:"autorejoin-zfs-flashback"`
	CheckType                          string `mapstructure:"check-type"`
	CheckReplFilter                    bool   `mapstructure:"check-replication-filters"`
	CheckBinFilter                     bool   `mapstructure:"check-binlog-filters"`
	RplChecks                          bool   `mapstructure:"check-replication-state"`
	ForceSlaveHeartbeat                bool   `mapstructure:"force-slave-heartbeat"`
	ForceSlaveHeartbeatTime            int    `mapstructure:"force-slave-heartbeat-time"`
	ForceSlaveHeartbeatRetry           int    `mapstructure:"force-slave-heartbeat-retry"`
	ForceSlaveGtid                     bool   `mapstructure:"force-slave-gtid-mode"`
	ForceSlaveNoGtid                   bool   `mapstructure:"force-slave-no-gtid-mode"`
	ForceSlaveSemisync                 bool   `mapstructure:"force-slave-semisync"`
	ForceSlaveReadOnly                 bool   `mapstructure:"force-slave-readonly"`
	ForceBinlogRow                     bool   `mapstructure:"force-binlog-row"`
	ForceBinlogAnnotate                bool   `mapstructure:"force-binlog-annotate"`
	ForceBinlogCompress                bool   `mapstructure:"force-binlog-compress"`
	ForceBinlogSlowqueries             bool   `mapstructure:"force-binlog-slowqueries"`
	ForceBinlogChecksum                bool   `mapstructure:"force-binlog-checksum"`
	ForceInmemoryBinlogCacheSize       bool   `mapstructure:"force-inmemory-binlog-cache-size"`
	ForceDiskRelayLogSizeLimit         bool   `mapstructure:"force-disk-relaylog-size-limit"`
	ForceDiskRelayLogSizeLimitSize     uint64 `mapstructure:"force-disk-relaylog-size-limit-size"`
	ForceSyncBinlog                    bool   `mapstructure:"force-sync-binlog"`
	ForceSyncInnoDB                    bool   `mapstructure:"force-sync-innodb"`
	ForceNoslaveBehind                 bool   `mapstructure:"force-noslave-behind"`
	Spider                             bool   `mapstructure:"spider"`
	BindAddr                           string `mapstructure:"http-bind-address"`
	HttpPort                           string `mapstructure:"http-port"`
	HttpServ                           bool   `mapstructure:"http-server"`
	HttpRoot                           string `mapstructure:"http-root"`
	HttpAuth                           bool   `mapstructure:"http-auth"`
	HttpBootstrapButton                bool   `mapstructure:"http-bootstrap-button"`
	SessionLifeTime                    int    `mapstructure:"http-session-lifetime"`
	Daemon                             bool   `mapstructure:"daemon"`
	MailFrom                           string `mapstructure:"mail-from"`
	MailTo                             string `mapstructure:"mail-to"`
	MailSMTPAddr                       string `mapstructure:"mail-smtp-addr"`
	Heartbeat                          bool   `mapstructure:"heartbeat-table"`
	MdbsProxyOn                        bool   `mapstructure:"mdbshardproxy"`
	MdbsProxyHosts                     string `mapstructure:"mdbshardproxy-servers"`
	MdbsProxyUser                      string `mapstructure:"mdbshardproxy-user"`
	MxsOn                              bool   `mapstructure:"maxscale"`
	MxsHost                            string `mapstructure:"maxscale-servers"`
	MxsPort                            string `mapstructure:"maxscale-port"`
	MxsUser                            string `mapstructure:"maxscale-user"`
	MxsPass                            string `mapstructure:"maxscale-pass"`
	MxsWritePort                       int    `mapstructure:"maxscale-write-port"`
	MxsReadPort                        int    `mapstructure:"maxscale-read-port"`
	MxsReadWritePort                   int    `mapstructure:"maxscale-read-write-port"`
	MxsMaxinfoPort                     int    `mapstructure:"maxscale-maxinfo-port"`
	MxsBinlogOn                        bool   `mapstructure:"maxscale-binlog"`
	MxsBinlogPort                      int    `mapstructure:"maxscale-binlog-port"`
	MxsDisableMonitor                  bool   `mapstructure:"maxscale-disable-monitor"`
	MxsGetInfoMethod                   string `mapstructure:"maxscale-get-info-method"`
	MxsServerMatchPort                 bool   `mapstructure:"maxscale-server-match-port"`
	HaproxyOn                          bool   `mapstructure:"haproxy"`
	HaproxyHosts                       string `mapstructure:"haproxy-servers"`
	HaproxyWritePort                   int    `mapstructure:"haproxy-write-port"`
	HaproxyReadPort                    int    `mapstructure:"haproxy-read-port"`
	HaproxyStatPort                    int    `mapstructure:"haproxy-stat-port"`
	HaproxyWriteBindIp                 string `mapstructure:"haproxy-ip-write-bind"`
	HaproxyReadBindIp                  string `mapstructure:"haproxy-ip-read-bind"`
	HaproxyBinaryPath                  string `mapstructure:"haproxy-binary-path"`
	ProxysqlOn                         bool   `mapstructure:"proxysql"`
	ProxysqlHosts                      string `mapstructure:"proxysql-servers"`
	ProxysqlPort                       string `mapstructure:"proxysql-port"`
	ProxysqlAdminPort                  string `mapstructure:"proxysql-admin-port"`
	ProxysqlUser                       string `mapstructure:"proxysql-user"`
	ProxysqlPassword                   string `mapstructure:"proxysql-password"`
	ProxysqlWriterHostgroup            string `mapstructure:"proxysql-writer-hostgroup"`
	ProxysqlReaderHostgroup            string `mapstructure:"proxysql-reader-hostgroup"`
	RegistryConsul                     bool   `mapstructure:"registry-consul"`
	RegistryHosts                      string `mapstructure:"registry-servers"`
	KeyPath                            string `mapstructure:"keypath"`
	Topology                           string `mapstructure:"topology"` // use by bootstrap
	GraphiteMetrics                    bool   `mapstructure:"graphite-metrics"`
	GraphiteEmbedded                   bool   `mapstructure:"graphite-embedded"`
	GraphiteCarbonHost                 string `mapstructure:"graphite-carbon-host"`
	GraphiteCarbonPort                 int    `mapstructure:"graphite-carbon-port"`
	GraphiteCarbonApiPort              int    `mapstructure:"graphite-carbon-api-port"`
	GraphiteCarbonServerPort           int    `mapstructure:"graphite-carbon-server-port"`
	GraphiteCarbonLinkPort             int    `mapstructure:"graphite-carbon-link-port"`
	GraphiteCarbonPicklePort           int    `mapstructure:"graphite-carbon-pickle-port"`
	GraphiteCarbonPprofPort            int    `mapstructure:"graphite-carbon-pprof-port"`
	SysbenchBinaryPath                 string `mapstructure:"sysbench-binary-path"`
	SysbenchTime                       int    `mapstructure:"sysbench-time"`
	SysbenchThreads                    int    `mapstructure:"sysbench-threads"`
	Arbitration                        bool   `mapstructure:"arbitration-external"`
	ArbitrationSasSecret               string `mapstructure:"arbitration-external-secret"`
	ArbitrationSasHosts                string `mapstructure:"arbitration-external-hosts"`
	ArbitrationSasUniqueId             int    `mapstructure:"arbitration-external-unique-id"`
	ArbitrationPeerHosts               string `mapstructure:"arbitration-peer-hosts"`
	FailForceGtid                      bool   //suspicious code
	Test                               bool   `mapstructure:"test"`
	TestInjectTraffic                  bool   `mapstructure:"test-inject-traffic"`
	Enterprise                         bool   //used to talk to opensvc collector
	ProvHost                           string `mapstructure:"opensvc-host"`
	ProvRegister                       bool   `mapstructure:"opensvc-register"`
	ProvAdminUser                      string `mapstructure:"opensvc-admin-user"`
	ProvUser                           string `mapstructure:"opensvc-user"`
	ProvCodeApp                        string `mapstructure:"opensvc-codeapp"`
	ProvDBServerPath                   string `mapstructure:"prov-db-localhost-binary-path"`
	ProvType                           string `mapstructure:"prov-db-service-type"`
	ProvAgents                         string `mapstructure:"prov-db-agents"`
	ProvMem                            string `mapstructure:"prov-db-memory"`
	ProvIops                           string `mapstructure:"prov-db-disk-iops"`
	ProvCores                          string `mapstructure:"prov-db-cpu-cores"`
	ProvTags                           string `mapstructure:"prov-db-tags"`
	ProvDisk                           string `mapstructure:"prov-db-disk-size"`
	ProvDiskFS                         string `mapstructure:"prov-db-disk-fs"`
	ProvDiskPool                       string `mapstructure:"prov-db-disk-pool"`
	ProvDiskDevice                     string `mapstructure:"prov-db-disk-device"`
	ProvDiskType                       string `mapstructure:"prov-db-disk-type"`
	ProvNetIface                       string `mapstructure:"prov-db-net-iface"`
	ProvNetmask                        string `mapstructure:"prov-db-net-mask"`
	ProvGateway                        string `mapstructure:"prov-db-net-gateway"`
	ProvProxType                       string `mapstructure:"prov-proxy-service-type"`
	ProvProxAgents                     string `mapstructure:"prov-proxy-agents"`
	ProvProxDisk                       string `mapstructure:"prov-proxy-disk-size"`
	ProvProxDiskFS                     string `mapstructure:"prov-proxy-disk-fs"`
	ProvProxDiskPool                   string `mapstructure:"prov-proxy-disk-pool"`
	ProvProxDiskDevice                 string `mapstructure:"prov-proxy-disk-device"`
	ProvProxDiskType                   string `mapstructure:"prov-proxy-disk-type"`
	ProvProxNetIface                   string `mapstructure:"prov-proxy-net-iface"`
	ProvProxNetmask                    string `mapstructure:"prov-proxy-net-mask"`
	ProvProxGateway                    string `mapstructure:"prov-proxy-net-gateway"`
	ProvMaxscaleVip                    string `mapstructure:"prov-proxy-net-maxscale-vip"`
	ProvMdbshardproxyVip               string `mapstructure:"prov-proxy-net-mdbsproxy-vip"`
	ProvHaproxyVip                     string `mapstructure:"prov-proxy-net-haproxy-vip"`
	ProvWebsqlproxyVip                 string `mapstructure:"prov-proxy-net-websqlproxy-vip"`
	ProvDbImg                          string `mapstructure:"prov-db-docker-img"`
	ProvProxMaxscaleImg                string `mapstructure:"prov-proxy-docker-maxscale-img"`
	ProvProxHaproxyImg                 string `mapstructure:"prov-proxy-docker-haproxy-img"`
	ProvProxProxysqlImg                string `mapstructure:"prov-proxy-docker-proxysql-img"`
	ProvProxTags                       string `mapstructure:"prov-proxy-tags"`
	APIUser                            string `mapstructure:"api-credential"`
	APIPort                            string `mapstructure:"api-port"`
	APIBind                            string `mapstructure:"api-bind"`
	AlertScript                        string `mapstructure:"alert-script"`
	ConfigFile                         string `mapstructure:"config"`
	GoOS                               string `mapstructure:"goos"`
	GoArch                             string `mapstructure:"goarch"`
}
