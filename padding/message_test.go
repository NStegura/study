package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
	"unsafe"
)

const (
	RPS           = 10_000
	ObjectsPerReq = 100
)

func printTestsInfo() {
	fmt.Printf(
		"RPS=%v\nObjectsPerReq=%v\n\n",
		RPS, ObjectsPerReq)
	fmt.Println("Sizeof Optimized:", unsafe.Sizeof(MessageOptimized{}))
	fmt.Println("Sizeof NonOptimized:", unsafe.Sizeof(MessageNonOptimized{}))
}

func printMemDiff(before, after *runtime.MemStats) {
	fmt.Printf("HeapAlloc = %d KB\n", (after.HeapAlloc-before.HeapAlloc)/1024)
	fmt.Printf("TotalAlloc = %d KB\n", (after.TotalAlloc-before.TotalAlloc)/1024)
	fmt.Printf("NumGC = %d\n", after.NumGC-before.NumGC)
}

func BenchmarkAllocNonOptimized(b *testing.B) {
	var (
		before, after runtime.MemStats
		nonOptimized  []MessageNonOptimized
	)
	printTestsInfo()
	{
		b.Run("NonOptimized", func(b *testing.B) {
			runtime.GC()
			debug.FreeOSMemory()
			runtime.ReadMemStats(&before)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				for j := 0; j < RPS; j++ {
					nonOptimized = make([]MessageNonOptimized, ObjectsPerReq)
					for n := range nonOptimized {
						nonOptimized[n].WithAction = nonOptimized[n].IsReply
					}
				}
			}

			b.StopTimer()
			runtime.ReadMemStats(&after)
		})
		printMemDiff(&before, &after)
	}
}

func BenchmarkAllocOptimized(b *testing.B) {
	var (
		before, after runtime.MemStats
		optimized     []MessageOptimized
	)
	printTestsInfo()
	{
		b.Run("Optimized", func(b *testing.B) {
			runtime.GC()
			debug.FreeOSMemory()
			runtime.ReadMemStats(&before)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				for j := 0; j < RPS; j++ {
					optimized = make([]MessageOptimized, ObjectsPerReq)
					for n := range optimized {
						optimized[n].WithAction = optimized[n].IsReply
					}
				}
			}
			b.StopTimer()
			runtime.ReadMemStats(&after)
		})
		printMemDiff(&before, &after)
	}
}
