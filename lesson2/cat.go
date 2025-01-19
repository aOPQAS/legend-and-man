package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	B bool
	E bool
	N bool
	S bool
	T bool
	H bool
	V bool
}

func ParseArgs() (Options, []string) {
	b := flag.Bool("b", false, "")
	n := flag.Bool("n", false, "")
	e := flag.Bool("e", false, "")
	s := flag.Bool("s", false, "")
	t := flag.Bool("t", false, "")
	h := flag.Bool("h", false, "")
	v := flag.Bool("v", false, "")
	flag.Parse()

	files := flag.Args()
	return Options{
		B: *b,
		E: *e,
		N: *n,
		S: *s,
		T: *t,
		H: *h,
		V: *v,
	}, files
}

func Cat(filename string, opts Options) {
	var language string
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", filename, err)
		return
	}

	if opts.V {
		fmt.Println("cat (GNU coreutils) 9.4")
		return
	}

	if opts.H {
		fmt.Println("Which language would be more convenient for you? (English/Russian)")
		if language == "Русский" {
			fmt.Println(

				"-b, Нумерация только непустых строк",
				"-e, Показать символы $ в конце строк",
				"-n, Нумерация всех строк",
				"-s, Сжать последовательные пустые строки",
				"-t, Отображение табуляции как ^I",
				"-h, Показать справку",
				"-v, Показать версию программы",
			)
		} else {
			fmt.Println(

				"-b, Numbering only non-empty lines",
				"-e, Show the $ characters at the end of the lines",
				"-n, Numbering of all lines",
				"-s, Compress consecutive blank lines",
				"-t, Tabulation display as ^I",
				"-h, Show the help",
				"-v, Show the program version",
			)
		}
		return
	}

	rows := make([]string, 0)
	row := ""
	var previousLine string

	for _, char := range data {
		if char == 10 {
			row += string(char)
			rows = append(rows, row)
			row = ""
		} else {
			row += string(char)
		}
	}

	for numRow, row := range rows {
		if opts.T {
			row = strings.Replace(row, "\t", "^I")
		}

		if opts.N {
			fmt.Printf("\t%d ", numRow+1)
		}

		fmt.Print(row)

		if opts.E {
			row += "$"
		}

		if opts.B {
			trimResult := strings.TrimSpace(row)
			if trimResult != "" {
				fmt.Printf("\t%d ", numRow+1)
			}
		}

		if opts.S {
			if strings.TrimSpace(row) == "" && strings.TrimSpace(previousLine) == "" {
				continue
			}
		}
		previousLine = row

	}
}

func main() {
	opts, files := ParseArgs()
	if len(files) == 0 {
		fmt.Print("your file is empty")
	}
	for _, file := range files {
		Cat(file, opts)
	}

}
