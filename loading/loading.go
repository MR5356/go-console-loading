package loading

import (
	"fmt"
	"github.com/MR5356/go-console-loading/config"
	"os"
	"os/signal"
	"strings"
	"time"
)

type Loading struct {
	spinner string
	running bool
}

func New(spinner string) *Loading {
	return &Loading{spinner: spinner}
}

func (l *Loading) Start(format string, a ...any) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			// signal caught, cleanup
			l.Stop()
			os.Exit(1)
		}()
		l.running = true
		nl := len(strings.Split(format, "\n"))
		format = strings.ReplaceAll(format, "\n", "\033[K\n")
		for {
			for _, v := range config.Resources[l.spinner].Frames {
				if l.running {
					fmt.Printf("\033[?25l\r%s %s\033[K", v, fmt.Sprintf(format, a...))
					if nl > 1 {
						fmt.Printf("\033[%dF", nl-1)
					}
					time.Sleep(time.Duration(config.Resources[l.spinner].Interval) * time.Millisecond)
				}
			}
		}
	}()
}

func (l *Loading) Stop() {
	fmt.Printf("\033[?25h")
	l.running = false
}
