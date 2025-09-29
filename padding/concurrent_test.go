package main

import (
	"runtime"
	"runtime/debug"
	"sync"
	"testing"
)

const (
	iterations = 1_000_000
)

func BenchmarkAllocSimpleFalseSharing(b *testing.B) {
	var (
		before, after runtime.MemStats
		counter       CounterNonOpt
		wg            sync.WaitGroup
	)
	{
		b.Run("SimpleFalseSharing", func(b *testing.B) {
			runtime.GC()
			debug.FreeOSMemory()
			runtime.ReadMemStats(&before)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				wg.Add(2)
				go func() {
					for i := 0; i < iterations; i++ {
						counter.a++
					}
					wg.Done()
				}()
				go func() {
					for i := 0; i < iterations; i++ {
						counter.b++
					}
					wg.Done()
				}()
				wg.Wait()
			}

			b.StopTimer()
			runtime.ReadMemStats(&after)
		})
		printMemDiff(&before, &after)
	}
}

func BenchmarkAllocSimpleNoFalseSharing(b *testing.B) {
	var (
		before, after runtime.MemStats
		counter       CounterOpt
		wg            sync.WaitGroup
	)
	{
		b.Run("SimpleNoFalseSharing", func(b *testing.B) {
			runtime.GC()
			debug.FreeOSMemory()
			runtime.ReadMemStats(&before)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				wg.Add(2)
				go func() {
					for i := 0; i < iterations; i++ {
						counter.a++
					}
					wg.Done()
				}()
				go func() {
					for i := 0; i < iterations; i++ {
						counter.b++
					}
					wg.Done()
				}()
				wg.Wait()
			}

			b.StopTimer()
			runtime.ReadMemStats(&after)
		})
		printMemDiff(&before, &after)
	}
}

func BenchmarkAllocAtomicFalseSharing(b *testing.B) {
	var (
		before, after runtime.MemStats
		counter       CounterAtomicNonOpt
		wg            sync.WaitGroup
	)
	{
		b.Run("AtomicFalseSharing", func(b *testing.B) {
			runtime.GC()
			debug.FreeOSMemory()
			runtime.ReadMemStats(&before)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				wg.Add(2)
				go func() {
					for i := 0; i < iterations; i++ {
						counter.a.Add(1)
					}
					wg.Done()
				}()
				go func() {
					for i := 0; i < iterations; i++ {
						counter.b.Add(1)
					}
					wg.Done()
				}()
				wg.Wait()
			}

			b.StopTimer()
			runtime.ReadMemStats(&after)
		})
		printMemDiff(&before, &after)
	}
}

func BenchmarkAllocAtomicNoFalseSharing(b *testing.B) {
	var (
		before, after runtime.MemStats
		counter       CounterAtomicOpt
		wg            sync.WaitGroup
	)
	{
		b.Run("AtomicNoFalseSharing", func(b *testing.B) {
			runtime.GC()
			debug.FreeOSMemory()
			runtime.ReadMemStats(&before)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				wg.Add(2)
				go func() {
					for i := 0; i < iterations; i++ {
						counter.a.Add(1)
					}
					wg.Done()
				}()
				go func() {
					for i := 0; i < iterations; i++ {
						counter.b.Add(1)
					}
					wg.Done()
				}()
				wg.Wait()
			}

			b.StopTimer()
			runtime.ReadMemStats(&after)
		})
		printMemDiff(&before, &after)
	}
}
