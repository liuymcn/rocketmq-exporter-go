package main

import (
	"flag"
	"net"

	"github.com/rocketmq-exporter-go/exporter"

	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
	// _ "net/http/pprof"
)

func toFlagString(name string, help string, value string) *string {
	flag.CommandLine.String(name, value, help) // hack around flag.Parse and glog.init flags
	return kingpin.Flag(name, help).Default(value).String()
}

func toFlagStrings(name string, help string, value string) *[]string {
	flag.CommandLine.String(name, value, help) // hack around flag.Parse and glog.init flags
	return kingpin.Flag(name, help).Default(value).Strings()
}

func toFlagBool(name string, help string, value bool, valueString string) *bool {
	flag.CommandLine.Bool(name, value, help) // hack around flag.Parse and glog.init flags
	return kingpin.Flag(name, help).Default(valueString).Bool()
}

func toFlagIntVar(name string, help string, value int, valueString string, target *int) {
	flag.CommandLine.Int(name, value, help) // hack around flag.Parse and glog.init flags
	kingpin.Flag(name, help).Default(valueString).IntVar(target)
}

func main() {

	var (
		listenAddress = toFlagString("web.listen-address", "Address to listen on for web interface and telemetry.", ":9999")
		metricsPath   = toFlagString("web.telemetry-path", "Path under which to expose metrics.", "/metrics")

		logFileEnable = toFlagBool("log.file.enable", "Log write to file enable.", false, "false")
		logFilePath   = toFlagString("log.file.path", "Path of log file.", "./logs/go.log")

		useDomain               = toFlagBool("rocketmq.nameserver.use.domian", "Address of rocketmq nameserver whether use domian.", false, "false")
		nameServerIpAddress     = toFlagStrings("rocketmq.nameserver", "Address (host:port) of rocketmq nameserver.", "127.0.0.1:9876")
		nameServerDomainAddress = toFlagString("rocketmq.nameserver.domain", "Address (domain:port) of rocketmq nameserver.", "rocketmq.namesvr:9876")

		opts = &exporter.RocketmqOptions{}
	)

	toFlagIntVar("workers", "Number of workers", 100, "100", &opts.Workers)

	kingpin.Version(version.Print("rocketmq_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	if *logFileEnable {
		rlog.SetOutputPath(*logFilePath)
	}

	if *useDomain {

		var ips, err = LookupIP(*nameServerDomainAddress)
		rlog.Fatal("rocketmq.nameserver.domain parse err", map[string]interface{}{
			"domain": *nameServerDomainAddress,
			"err":    err,
		})
		opts.NameServerAddress = ips

	} else {
		opts.NameServerAddress = *nameServerIpAddress
	}

	rlog.Info("Fly ", map[string]interface{}{
		"listenAddress": *listenAddress,
		"metricsPath":   *metricsPath,
		"opts":          *opts,
	})

	exporter.Fly(*listenAddress, *metricsPath, opts)

}

// Because rocketmq client requires ip address rather than domain address
// LookupIP looks up host
// It returns a slice of that host's IPv4 addresses.
func LookupIP(host string) ([]string, error) {
	ips, err := net.LookupIP(host)

	if err != nil {
		return nil, err
	}
	var ipList = make([]string, len(ips))
	for index, ip := range ips {
		ipList[index] = ip.To4().String()
	}

	return ipList, nil
}
