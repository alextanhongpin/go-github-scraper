package profiler

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// MakeCPU sets the cpu profiling if the output filename is provided
func MakeCPU(file string) {
	if file != "" {
		f, err := os.Create(file)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
}

// MakeMemory sets the memory profiling if the output filename is provided
func MakeMemory(file string) {
	if file != "" {
		f, err := os.Create(file)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
