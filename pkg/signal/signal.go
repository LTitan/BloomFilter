package signal

import (
	"os"
	"os/signal"
	"syscall"
)

// ExitBeautiful .
func ExitBeautiful(exit func()) {
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			exit()
			os.Exit(0)
		default:
		}
	}
}
