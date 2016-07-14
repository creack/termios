// +build darwin freebsd dragonfly openbsd netbsd

package raw

import "syscall"

const (
	getTermios = syscall.TIOCGETA
	setTermios = syscall.TIOCSETA
)
