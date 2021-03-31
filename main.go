package main

import (
	"os"

	"github.com/kang2681/modtool/configs"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/sirupsen/logrus"
)

var logLevel string
var listenAddress string
var configFile string

func main() {
	// 日志级别
	kingpin.Flag("log.level", "log level output.").Default("info").StringVar(&logLevel)
	kingpin.Flag("config.file", "modtool configuration file name").Default("config.yaml").StringVar(&configFile)
	kingpin.Flag("web.listen-address", " Address to listen on for web interface and telemetry.").Default(":7777").StringVar(&listenAddress)
	kingpin.Version("0.0.1")
	kingpin.Parse()
	setLog()
	configs.LoadConfig(configFile)
}

func setLog() {
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	// set log level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
