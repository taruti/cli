// cli is a library of cli program helpers.
package cli

import (
	"flag"
	"io"
	"log"

	"github.com/chzyer/readline"
)

// Program is the cli program.
type Program interface {
	HandleCmdLine([]string) error
	HandleCliLine(string) error
}

// Main works as a main for a cli program.
func Main(p Program) {
	flag.Parse()
	err := mainWork(p)
	if err != nil {
		log.Fatal("ERROR", err)
	}
}

func mainWork(p Program) error {
	args := flag.Args()

	if len(args) == 0 {
		return console(p)
	}

	return p.HandleCmdLine(args)
}

// console displays a console.
func console(p Program) error {
	rl, err := readline.New("> ")
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if line == "" {
			continue
		}
		err = p.HandleCliLine(line)
		if err != nil { // io.EOF
			return err
		}
	}
}
