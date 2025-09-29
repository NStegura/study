package main

type NonOptimized struct {
	ByteA byte // 1 байт
	Int   int  // 8 байт
	ByteB byte // 1 байт

}

type Optimized struct {
	Int   int  // 8 байт
	ByteA byte // 1 байт
	ByteB byte // 1 байт
}
