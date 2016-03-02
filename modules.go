package main

import (
	_ "github.com/3scale/logspout/adapters/raw"
	_ "github.com/3scale/logspout/adapters/syslog"
	_ "github.com/3scale/logspout/httpstream"
	_ "github.com/3scale/logspout/routesapi"
	_ "github.com/3scale/logspout/transports/tcp"
	_ "github.com/3scale/logspout/transports/udp"
	_ "github.com/3scale/logspout/transports/tls"
	_ "github.com/3scale/logspout-redis-logstash"
)
