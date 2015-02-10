package console

import (
	"fmt"
	"io"
	"os"
)

var Default = Console{Stdout: os.Stdout, Stderr: os.Stderr}

func Infof(format string, a ...interface{}) (n int, err error) {
	return Default.Infof(format, a...)
}

func Infoln(a ...interface{}) (n int, err error) {
	return Default.Infoln(a...)
}

func Errorf(format string, a ...interface{}) (n int, err error) {
	return Default.Errorf(format, a...)
}

func Errorln(a ...interface{}) (n int, err error) {
	return Default.Errorln(a...)
}

type Console struct {
	Stdout io.Writer
	Stderr io.Writer
}

func (c Console) Infof(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.Stdout, format, a...)
}

func (c Console) Infoln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.Stdout, a...)
}

func (c Console) Errorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.Stderr, format, a...)
}

func (c Console) Errorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.Stderr, a...)
}
