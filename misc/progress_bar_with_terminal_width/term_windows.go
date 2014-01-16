package main

import (
	"log"
	"syscall"
	"unsafe"
)

func main() {
	x, y, err := getConWinSize()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("W=%d, H=%d\n", x, y)
}

func getConWinSize() (x, y int, err error) {
	hCon, err := syscall.Open("CONOUT$", syscall.O_RDONLY, 0)
	if err != nil {
		return
	}
	defer syscall.Close(hCon)

	sb, err := getConsoleScreenBufferInfo(hCon)
	if err != nil {
		return
	}
	x = int(sb.size.x)
	y = int(sb.size.y)
	return
}

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetConScrBufInfo = modkernel32.NewProc("GetConsoleScreenBufferInfo")
)

type coord struct {
	x int16
	y int16
}

type smallRect struct {
	left   int16
	top    int16
	right  int16
	bottom int16
}

type consoleScreenBuffer struct {
	size       coord
	cursorPos  coord
	attrs      int32
	window     smallRect
	maxWinSize coord
}

func getConsoleScreenBufferInfo(hCon syscall.Handle) (sb consoleScreenBuffer, err error) {
	rc, _, ec := syscall.Syscall(procGetConScrBufInfo.Addr(), 2,
		uintptr(hCon), uintptr(unsafe.Pointer(&sb)), 0)
	if rc == 0 {
		err = syscall.Errno(ec)
	}
	return
}
