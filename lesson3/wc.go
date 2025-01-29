package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	C bool
	M bool
	L bool
	W bool
}

type Result struct {
	Filename string
	Count    map[string]int
}

func (r Result) Display() {
	fmt.Println(r.Count, r.Filename)
}

func ParseArgs() (Options, []string) {
	c := flag.Bool("c", false, "")
	m := flag.Bool("m", false, "")
	l := flag.Bool("l", false, "")
	w := flag.Bool("w", false, "")

	flag.Parse()

	files := flag.Args()
	return Options{
		C: *c,
		M: *m,
		L: *l,
		W: *w,
	}, files
}

func WC(filename string, options Options) (Result, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Result{}, err
	}

	res := Result{
		Count: make(map[string]int),
	}

	dataStr := string(data)

	if options.M {
		countChars := len([]rune(dataStr))
		res.Count["m"] = countChars
	}

	if options.L {
		countLines := strings.Count(dataStr, "\n")
		res.Count["l"] = countLines
	}

	if options.W {
		countWords := len(strings.Split(dataStr, " "))
		res.Count["w"] = countWords
	}

	if options.C {
		countBytes := len(data)
		res.Count["c"] = countBytes
	}

	res.Filename = filename

	return res, nil
}

func main() {
	opts, files := ParseArgs()
	if len(files) == 0 {
		fmt.Println("Please specify some files")
	}

	totalRes := Result{
		Filename: "total",
		Count:    map[string]int{},
	}

	for _, file := range files {
		res, err := WC(file, opts)
		if err != nil {
			fmt.Println("error while reading file:", err)
			continue
		}

		if opts.C {
			totalRes.Count["c"] += res.Count["c"]
		}

		if opts.L {
			totalRes.Count["l"] += res.Count["l"]
		}

		if opts.M {
			totalRes.Count["m"] += res.Count["m"]
		}

		if opts.W {
			totalRes.Count["w"] += res.Count["w"]
		}

		res.Display()
	}

	if len(files) > 1 {
		totalRes.Display()
	}
}
