package raw

import "syscall"

const (
	getTermios = syscall.TCGETS
	setTermios = syscall.TCSETS
)

// Termios holds the TTY attributes. See man termios(4).
type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]byte
	Ispeed uint32
	Ospeed uint32
}
