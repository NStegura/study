package main

import "sync/atomic"

type CounterNonOpt struct {
	a int64
	b int64
}

type CounterOpt struct {
	a int64
	_ [56]byte // to new cache line
	b int64
}

type CounterAtomicNonOpt struct {
	a atomic.Int64
	b atomic.Int64
}

type CounterAtomicOpt struct {
	a atomic.Int64
	_ [56]byte // to new cache line
	b atomic.Int64
}
