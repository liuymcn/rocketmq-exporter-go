package main

import (
	"flag"

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

func toFlagBool(name string, help string, value bool, valueString string) *bool {
	flag.CommandLine.Bool(name, value, help) // hack around flag.Parse and glog.init flags
	return kingpin.Flag(name, help).Default(valueString).Bool()
}

func toFlagStringsVar(name string, help string, value string, target *[]string) {
	flag.CommandLine.String(name, value, help) // hack around flag.Parse and glog.init flags
	kingpin.Flag(name, help).Default(value).StringsVar(target)
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

		opts = &exporter.RocketmqOptions{}
	)

	toFlagStringsVar("rocketmq.nameserver", "Address (host:port) of rocketmq nameserver.", "rocketmq:9876", &opts.NameServerAddress)
	toFlagIntVar("workers", "Number of workers", 100, "100", &opts.Workers)

	kingpin.Version(version.Print("rocketmq_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	if *logFileEnable {
		rlog.SetOutputPath(*logFilePath)
	}

	rlog.Info("Fly ", map[string]interface{}{
		"listenAddress": *listenAddress,
		"metricsPath":   *metricsPath,
		"opts":          *opts,
	})

	exporter.Fly(*listenAddress, *metricsPath, opts)

}
