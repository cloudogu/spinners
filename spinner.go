package spinners

import (
	"fmt"
	"io"
	"time"

	"sync"
)

//go:generate go run gen/gen.go
//go:generate gofmt -w spinners.go

// Configuration describes the spinner
type Configuration struct {
	Interval int
	Frames   []string
}

// NewSpinner creates a new spinner with the given configuration
func NewSpinner(configuration Configuration, writer io.Writer) *Spinner {
	stop := make(chan interface{})

	return &Spinner{
		configuration: configuration,
		writer:        writer,
		stop:          stop,
		lock:          &sync.Mutex{},
	}
}

// Spinner is able to print a waiting spinner to a writer
type Spinner struct {
	configuration Configuration
	writer        io.Writer
	stop          chan interface{}
	lock          *sync.Mutex
}

// Start starts the spinner in a separate go routine
func (spinner *Spinner) Start(message string) {
	// TODO what about the error returned from spin?
	go spinner.spinEndless(message)
}

func (spinner *Spinner) spinEndless(message string) error {
	for {
		err := spinner.spinOnce(message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (spinner *Spinner) spinOnce(message string) error {
	for _, frame := range spinner.configuration.Frames {
		select {
		case <-spinner.stop:
			return nil
		default:
			err := spinner.writeFrame(frame, message)
			if err != nil {
				return err
			}
			spinner.sleep()
		}
	}
	return nil
}

func (spinner *Spinner) writeFrame(frame string, message string) error {
	spinner.lock.Lock()
	defer spinner.lock.Unlock()

	_, err := fmt.Fprintf(spinner.writer, "\r%s %s", frame, message)
	if err != nil {
		return fmt.Errorf("failed to write frame to writer: %v", err)
	}
	return err
}

func (spinner *Spinner) sleep() {
	time.Sleep(time.Duration(spinner.configuration.Interval) * time.Millisecond)
}

// Stop stops the spinner and clear the terminal line
func (spinner *Spinner) Stop() error {
	spinner.stop <- true
	return spinner.resetLine()
}

func (spinner *Spinner) resetLine() error {
	spinner.lock.Lock()
	defer spinner.lock.Unlock()

	_, err := spinner.writer.Write([]byte("\r\033[K"))
	if err != nil {
		return fmt.Errorf("failed to write line reset characters to writer: %v", err)
	}
	return nil
}
