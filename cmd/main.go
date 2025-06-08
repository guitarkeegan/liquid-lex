package main

import (
	"flag"
	"os"

	"github.com/charmbracelet/log"
)

var dbg = func() func(format string, args ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(format string, args ...any) {}
	}
	return func(format string, args ...any) {
		log.Debugf(format, args...)
	}
}()

func main() {
	var filepath string
	flag.StringVar(&filepath, "file", "", "path to .liquid file")
	flag.Parse()

	logger := log.New(os.Stdout)
	log.SetLevel(log.DebugLevel)

	dbg("Debugger Started")

	file, err := os.ReadFile(filepath)
	if err != nil {
		logger.Error("no file", "required", "use -file <path to .liquid file>", "err", err)
		os.Exit(1)
	}

	logger.Infof("%s\n%s\n\n%s", "Starting lexer", "Original Input:", string(file))

	lexer, err := NewLexer(string(file))
	if err != nil {
		logger.Error("main: newLexer: %s", err)
		os.Exit(1)
	}

	// Run the lexer and collect all tokens
	tokens := lexer.Run()

	// Display results
	logger.Info("Lexing completed. Tokens found:")
	for _, token := range tokens {
		logger.Infof("  %s: %s", token.TypeString(), token.String())
	}
}
