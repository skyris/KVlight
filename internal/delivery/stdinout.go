package delivery

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/skyris/KVlight/pkg/interfaces"
)

type StdDelivery struct {
	scanner *bufio.Scanner
	output  io.Writer
}

var _ interfaces.Delivery = &StdDelivery{}

func Default() *StdDelivery {
	return &StdDelivery{
		scanner: bufio.NewScanner(os.Stdin),
		output:  os.Stdout,
	}
}

func NewStdinDelivery(in io.Reader, out io.Writer) *StdDelivery {
	return &StdDelivery{
		scanner: bufio.NewScanner(in),
		output:  out,
	}
}

func (d *StdDelivery) GetRequest(_ context.Context) (string, error) {
	log.Println("Insert command, to finish press CTRL+D:")
	if !d.scanner.Scan() {
		if err := d.scanner.Err(); err != nil {
			log.Println("Scanner error:", err)
			return "", err
		}

		return "", io.EOF
	}
	log.Println(d.scanner.Text())

	return d.scanner.Text(), nil
}

func (d *StdDelivery) SendResponse(_ context.Context, msg string, err error) error {
	switch {
	case err == nil && msg == "":
		_, errOut := fmt.Fprintf(d.output, "done\n")
		if errOut != nil {
			return errOut
		}
	case err != nil:
		_, errOut := fmt.Fprintf(d.output, "error: %v\n", err)
		if errOut != nil {
			return errOut
		}
	default:
		_, errOut := fmt.Fprintf(d.output, "msg: %s\n", msg)
		if errOut != nil {
			return errOut
		}
	}

	return nil
}
