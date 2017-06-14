package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/hashicorp/terraform/builtin/providers/aws"
	"github.com/hashicorp/terraform/plugin"
)

func init() {
	runtime.MemProfileRate = 1024
}

func stats() {
	statFile, err := os.Create(fmt.Sprintf("/tmp/stats-%d.log", time.Now().UnixNano()))
	if err != nil {
		log.Fatal("error creating stat log:", err)
	}

	var memStats runtime.MemStats

	for {
		runtime.ReadMemStats(&memStats)

		fmt.Fprintf(statFile, "NumGoroutine:%d MemTotal:%d HeapSys:%d HeapAlloc:%d HeapObjects:%d StackInuse:%d StackSys:%d NextGC:%d\n",
			runtime.NumGoroutine(),
			memStats.TotalAlloc-memStats.HeapReleased,
			memStats.HeapSys,
			memStats.HeapAlloc,
			memStats.HeapObjects,
			memStats.StackInuse,
			memStats.StackSys,
			memStats.NextGC,
		)

		statFile.Sync()
		time.Sleep(time.Second)
	}
}

func memprofiler() {
	for {
		time.Sleep(20 * time.Second)
		writeMemProfile()
	}
}

func writeMemProfile() {
	f, err := os.Create(fmt.Sprintf("/tmp/memprofile-%d", time.Now().UnixNano()))
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}

func main() {
	go stats()
	go memprofiler()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aws.Provider,
	})
	writeMemProfile()
}
