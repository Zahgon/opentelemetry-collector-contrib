// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package procx // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension/internal/procx"

const (
	ApacheProcessIdentifier             ProcessIdentifier = "apache"
	Apache2ProcessIdentifier            ProcessIdentifier = "apache2"
	HttpdProcessIdentifier              ProcessIdentifier = "httpd"
	DockerProcessIdentifier             ProcessIdentifier = "docker"
	ElasticsearchProcessIdentifier      ProcessIdentifier = "elasticsearch"
	MysqlServerProcessIdentifier        ProcessIdentifier = "mysql-server"
	MysqldProcessIdentifier             ProcessIdentifier = "mysqld"
	NginxProcessIdentifier              ProcessIdentifier = "nginx"
	PostgresProcessIdentifier           ProcessIdentifier = "postgresql"
	Postgres95ProcessIdentifier         ProcessIdentifier = "postgresql-9.5"
	RabbitmqServerProcessIdentifier     ProcessIdentifier = "rabbitmq-server"
	RedisProcessIdentifier              ProcessIdentifier = "redis"
	TomcatProcessIdentifier             ProcessIdentifier = "tomcat"
	KafkaServerStartShProcessIdentifier ProcessIdentifier = "kafka-server-start.sh"
	RedisServerProcessIdentifier        ProcessIdentifier = "redis-server"
	MongodProcessIdentifier             ProcessIdentifier = "mongod"
	CassandraProcessIdentifier          ProcessIdentifier = "cassandra"
	JmxProcessIdentifier                ProcessIdentifier = "jmx"
	ActiveMQProcessIdentifier           ProcessIdentifier = "activemq"
	MemcachedProcessIdentifier          ProcessIdentifier = "memcached"
	HaproxyProcessIdentifier            ProcessIdentifier = "haproxy"
	DockerdProcessIdentifier            ProcessIdentifier = "dockerd"
	DockerDesktopJavaProcessIdentifier  ProcessIdentifier = "com.docker.backend"
	SqlservrProcessIdentifier           ProcessIdentifier = "sqlservr"
	// Java Process Identifiers
	JavaProcessIdentifier          ProcessIdentifier = "java"
	CassandraJavaProcessIdentifier ProcessIdentifier = "org.apache.cassandra.service.CassandraDaemon"
	JmxJavaProcessIdentifier       ProcessIdentifier = "com.sun.management.jmxremote"
	ActiveMQJavaProcessIdentifier  ProcessIdentifier = "activemq.jar"
)

var sumoAppProcesses = map[ProcessIdentifier]SumoTag{
	ApacheProcessIdentifier:             ApacheTag,
	Apache2ProcessIdentifier:            ApacheTag,
	HttpdProcessIdentifier:              ApacheTag,
	DockerProcessIdentifier:             DockerTag, // docker cli
	ElasticsearchProcessIdentifier:      ElasticsearchTag,
	MysqlServerProcessIdentifier:        MysqlTag,
	MysqldProcessIdentifier:             MysqlTag,
	NginxProcessIdentifier:              NginxTag,
	PostgresProcessIdentifier:           PostgresTag,
	Postgres95ProcessIdentifier:         PostgresTag,
	RabbitmqServerProcessIdentifier:     RabbitmqTag,
	RedisProcessIdentifier:              RedisTag,
	TomcatProcessIdentifier:             TomcatTag,
	KafkaServerStartShProcessIdentifier: KafkaTag, // Need to test this, most common shell wrapper.
	RedisServerProcessIdentifier:        RedisTag,
	MongodProcessIdentifier:             MongoDBTag,
	CassandraProcessIdentifier:          CassandraTag,
	JmxProcessIdentifier:                JmxTag,
	ActiveMQProcessIdentifier:           ActiveMQTag,
	MemcachedProcessIdentifier:          MemcachedTag,
	HaproxyProcessIdentifier:            HaproxyTag,
	DockerdProcessIdentifier:            DockerCETag, // docker engine, for when process runs natively
	DockerDesktopJavaProcessIdentifier:  DockerCETag, // docker daemon runs on a VM in Docker Desktop, process doesn't show on mac
	SqlservrProcessIdentifier:           MssqlTag,    // linux SQL Server process
	// Java Process Tags
	CassandraJavaProcessIdentifier: CassandraTag,
	JmxJavaProcessIdentifier:       JmxTag,
	ActiveMQJavaProcessIdentifier:  ActiveMQTag,
}

func GetSumoTag(processName string) (SumoTag, bool) {
	_ = "STUB: not implemented"
	return *new(SumoTag), false
}

func (procx *Procx) FilteredProcessList() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handling Java background processes

func (procx *Procx) getProcessName(process Process) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// If we can't get a process name, it may be a zombie process.
// We do not want to error out here, as it's not worth disrupting
// the startup process of the collector.

func (procx *Procx) getJavaProcessName(processName string, process Process) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
