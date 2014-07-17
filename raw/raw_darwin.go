package raw

import "syscall"

const (
	getTermios = syscall.TIOCGETA
	setTermios = syscall.TIOCSETA
)

// Termios holds the TTY attributes. See man termios(4).
type Termios struct {
	Iflag  uint64
	Oflag  uint64
	Cflag  uint64
	Lflag  uint64
	Cc     [20]byte
	Ispeed uint64
	Ospeed uint64
}
