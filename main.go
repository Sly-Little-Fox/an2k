package main

import (
	"encoding/base32"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/anko/parser"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

func Must[T any](result T, err error) T {
	if err != nil {
		MaybeFatal(err, "Assertion failed:")
	}
	return result
}

func PrintError(err error, message string) {
	fmt.Println(color.RedString("Error |"), message, err)
}

func MaybeFatal(err error, message string) {
	if err != nil {
		PrintError(err, message)
		os.Exit(1)
	}
}

func PrintWarning(message string, args ...any) {
	fmt.Println(color.YellowString("Warn |"), fmt.Sprintf(message, args...))
}

var Name = flag.StringP("name", "n", "", "Ставит имя программы / Sets the program name")
var W = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
var VarStore []string
var Source []byte

func main() {
	flag.Parse()
	filename := ""
	if len(flag.Args()) < 1 {
		filename = "/dev/stdin"
	} else {
		filename = flag.Arg(0)
	}
	file, err := os.Open(filename)
	MaybeFatal(err, "Failed to open "+filename+":")
	Source, err = io.ReadAll(file)
	MaybeFatal(err, "Failed to read data from file:")
	stmt, err := parser.ParseSrc(string(regexp.MustCompile(`for (.+) times`).
		ReplaceAllStringFunc(string(Source), func(s string) string {
			return `for ______` + base32.HexEncoding.
				WithPadding(base32.NoPadding).EncodeToString([]byte(s[4:len(s)-6]))
		})))
	MaybeFatal(err, "Failed to parse Anko source (visit https://github.com/mattn/anko for docs):")
	output := ConvertStmt(stmt, &VarStore)
	for k, v := range DirReplaceMap {
		output = strings.ReplaceAll(output, k, v)
	}
	output = regexp.MustCompile(`(^использовать .+)`).ReplaceAllString(output, "$1\nалг "+*Name+"\nнач")
	output = regexp.MustCompile(`пока ______(.+)`).ReplaceAllStringFunc(output, func(s string) string {
		return string(Must(base32.HexEncoding.WithPadding(base32.NoPadding).DecodeString(s[15:]))) + " раз"
	})
	output += "\n" + ProgramEnd
	if !strings.Contains(output, "использовать ") {
		output = "алг " + *Name + "\nнач\n" + output
	}
	fmt.Println(strings.TrimSpace(strings.ReplaceAll(output, "\n\n", "\n")))
}
