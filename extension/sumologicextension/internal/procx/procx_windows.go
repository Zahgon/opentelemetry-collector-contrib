// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package procx // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension/internal/procx"

const (
	ApacheProcessIdentifier   ProcessIdentifier = "apache.exe"
	HttpdProcessIdentifier    ProcessIdentifier = "httpd.exe"
	DockerProcessIdentifier   ProcessIdentifier = "docker.exe"
	MysqldProcessIdentifier   ProcessIdentifier = "mysqld.exe"
	NginxProcessIdentifier    ProcessIdentifier = "nginx.exe"
	PostgresProcessIdentifier ProcessIdentifier = "postgresql.exe"
	TomcatProcessIdentifier   ProcessIdentifier = "tomcat.exe"
	Tomcat9ProcessIdentifier  ProcessIdentifier = "tomcat9.exe"
	MongodProcessIdentifier   ProcessIdentifier = "mongod.exe"
	DockerdProcessIdentifier  ProcessIdentifier = "dockerd.exe"
	SqlservrProcessIdentifier ProcessIdentifier = "sqlservr.exe"
	// Java Process Identifiers
	JavaProcessIdentifier          ProcessIdentifier = "java.exe"
	ElasticsearchProcessIdentifier ProcessIdentifier = "elasticsearch" // cmdline args does not have exe
	CasandraJavaProcessIdentifier  ProcessIdentifier = "org.apache.cassandra.service.CassandraDaemon"
	JmxJavaProcessIdentifier       ProcessIdentifier = "com.sun.management.jmxremote"
	ActiveMQJavaProcessIdentifier  ProcessIdentifier = "activemq.jar"
	// Erlang Process Identifiers
	ErlangProcessIdentifier         ProcessIdentifier = "erl.exe"
	RabbitmqServerProcessIdentifier ProcessIdentifier = "rabbit"
)

var sumoAppProcesses = map[ProcessIdentifier]SumoTag{
	ApacheProcessIdentifier:   ApacheTag,
	HttpdProcessIdentifier:    ApacheTag,
	DockerProcessIdentifier:   DockerTag, // docker cli
	MysqldProcessIdentifier:   MysqlTag,
	NginxProcessIdentifier:    NginxTag,
	PostgresProcessIdentifier: PostgresTag,
	TomcatProcessIdentifier:   TomcatTag,
	Tomcat9ProcessIdentifier:  TomcatTag,
	MongodProcessIdentifier:   MongoDBTag,
	DockerdProcessIdentifier:  DockerCETag, // docker engine, for when process runs natively
	SqlservrProcessIdentifier: MssqlTag,
	// Java Process Tags
	ElasticsearchProcessIdentifier: ElasticsearchTag,
	CasandraJavaProcessIdentifier:  CassandraTag,
	JmxJavaProcessIdentifier:       JmxTag,
	ActiveMQJavaProcessIdentifier:  ActiveMQTag,
	// Erlang Process Identifiers
	RabbitmqServerProcessIdentifier: RabbitmqTag,
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

// handling erlang processes

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

func (procx *Procx) getErlangProcessName(processName string, process Process) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
