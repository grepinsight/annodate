// annodate prepends/appends day of week to stdin or file
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func procLine(line string, delim string, field int, last bool) string {
	fields := strings.Split(line, delim)

	dateField := fields[field-1]

	// change formt
	t, _ := time.Parse("2006-01-02", dateField)

	if last {
		return fmt.Sprintf("%s%s%s", line, delim, t.Format("Monday"))
	} else {
		return fmt.Sprintf("%s%s%s", t.Format("Monday"), delim, line)
	}

}
func annodate(r io.Reader, delim string, field int, last bool) {

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text() // Println will add back the final '\n'
		fmt.Printf("%s\n", procLine(line, delim, field, last))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func stdinOrFile() io.Reader {

	fi, err := os.Stdin.Stat()
	check(err)

	if fi.Mode()&os.ModeNamedPipe != 0 {
		return os.Stdin
	}
	// If nothing is given produce help message
	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Assume os.Args[1] is a file
	file, err := os.Open(os.Args[1])
	check(err)
	return file

}

func main() {

	// func String(name string, value string, usage string) *string {
	delimPtr := flag.String("d", " ", "Delimiter")
	fieldPtr := flag.Int("f", 1, "Field that contains date data")
	lastColPtr := flag.Bool("e", false, "Add column at the end?")
	// formatPtr := flag.String("t", "2006-01-02", "Format")
	flag.Parse()

	annodate(stdinOrFile(), *delimPtr, *fieldPtr, *lastColPtr)
}
