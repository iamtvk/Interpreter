package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/interpreter/evaluator"
	"github.com/interpreter/lexer"
	"github.com/interpreter/object"
	"github.com/interpreter/parser"
)

const PROMPT = "--> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}

}

func printParserErrors(out io.Writer, errors []string) {
	// errArt := `
	// 	  ___ _ __ _ __ ___  _ __
	// 	 / _ \ '__| '__/ _ \| '__|
	// 	|  __/ |  | | | (_) | |
	// 	 \___|_|  |_|  \___/|_|
	// `
	// io.WriteString(out, errArt)
	for _, msg := range errors {
		io.WriteString(out, "ERROR: "+msg+"\n")
	}
}
