package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/linnv/logx"
	"github.com/linnv/simdog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type mid struct {
}

func (m *mid) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.DefaultServeMux.ServeHTTP(w, r)
	simdog.IncQps()
}

func main() {
	simdog.ReadBuildInfo()

	appName := "demo"
	prometheus.MustRegister(simdog.NewCpuMetric(appName))
	prometheus.MustRegister(simdog.NewNetMetric(appName))
	prometheus.MustRegister(simdog.NewQpsMetric(appName))
	prometheus.MustRegister(simdog.NewVersionCollector(appName))

	m := &mid{}
	p := flag.Int64("p", 8019, "port")
	flag.Parse()

	var port = fmt.Sprintf(":%d", *p)
	http.Handle("/metrics", promhttp.Handler())

	//@TODO  mutex
	gb := []byte(`good job`)
	http.Handle("/a", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(gb)
	}))

	srv := &http.Server{
		Handler: m,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	cmd := "curl http://127.0.0.1:8019/metrics"
	logx.Debugf("Beginning to serve on port %s try\n%s\n", port, cmd)
	go srv.ListenAndServe()

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	log.Print("use c-c to exit: \n")
	<-sigChan
}
