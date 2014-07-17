package win

import (
	"syscall"
	"unsafe"
)

// Winsize stores the Heighty and Width of a terminal.
type Winsize struct {
	Height uint16
	Width  uint16
	x      uint16 // unused
	y      uint16 // unused
}

// GetWinsize returns the size of the given tty.
func GetWinsize(fd uintptr) (*Winsize, error) {
	ws := &Winsize{}
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(ws))); err != 0 {
		return ws, err
	}
	return ws, nil
}

// SetWinsize sets the size of the given tty.
func SetWinsize(fd uintptr, ws *Winsize) error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TIOCSWINSZ), uintptr(unsafe.Pointer(ws))); err != 0 {
		return err
	}
	return nil
}
