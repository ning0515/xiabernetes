package util

import (
	"flag"
	"github.com/learnk8s/golang/glog"
	"log"
	"time"
)

var logFlushFreq = flag.Duration("log_flush_frequency", 5*time.Second, "Maximum number of seconds between log flushes")

func init() {
	flag.Set("logtostderr", "true")
}

type GlogWriter struct{}

func (writer GlogWriter) Write(data []byte) (n int, err error) {
	glog.Info(string(data))
	return len(data), nil
}
func InitLogs() {
	log.SetOutput(GlogWriter{})
	log.SetFlags(0)
	// The default glog flush interval is 30 seconds, which is frighteningly long.
	go Forever(glog.Flush, *logFlushFreq)
}

func FlushLogs() {
	glog.Flush()
}

func NewLogger(prefix string) *log.Logger {
	return log.New(GlogWriter{}, prefix, 0)
}
