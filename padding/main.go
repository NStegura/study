package main

import (
	"fmt"
	"reflect"
)

func PrintOffsets(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%-10s offset=%2d, size=%2d\n", f.Name, f.Offset, f.Type.Size())
	}
	fmt.Printf("Total size: %d\n", t.Size())
}

func main() {
	fmt.Println("NonOptimized struct layout:")
	PrintOffsets(reflect.TypeOf(NonOptimized{}))
	fmt.Println()

	fmt.Println("Optimized struct layout:")
	PrintOffsets(reflect.TypeOf(Optimized{}))
	fmt.Println()

	fmt.Println("MessageNonOptimized struct layout:")
	PrintOffsets(reflect.TypeOf(MessageNonOptimized{}))
	fmt.Println()

	fmt.Println("MessageOptimized struct layout:")
	PrintOffsets(reflect.TypeOf(MessageOptimized{}))
	fmt.Println()

	fmt.Println("ConcurrentSimpleNonOptimized struct layout:")
	PrintOffsets(reflect.TypeOf(CounterNonOpt{}))
	fmt.Println()

	fmt.Println("ConcurrentSimpleOptimized struct layout:")
	PrintOffsets(reflect.TypeOf(CounterOpt{}))
	fmt.Println()

	fmt.Println("ConcurrentAtomicNonOptimized struct layout:")
	PrintOffsets(reflect.TypeOf(CounterAtomicNonOpt{}))
	fmt.Println()

	fmt.Println("ConcurrentAtomicOptimized struct layout:")
	PrintOffsets(reflect.TypeOf(CounterAtomicOpt{}))
	fmt.Println()
}
