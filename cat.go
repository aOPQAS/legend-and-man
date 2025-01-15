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
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла", filename, err)
		return
	}

	if opts.V {
		fmt.Println("cat (GNU coreutils) 9.4")
		return
	}

	if opts.H {
		fmt.Println(

			"-b, Нумерация только непустых строк",
			"-e, Показать символы $ в конце строк",
			"-n, Нумерация всех строк",
			"-s, Сжать последовательные пустые строки",
			"-t, Отображение табуляции как ^I",
			"-h, Показать справку",
			"-v, Показать версию программы",
		)
	}
	rows := make([]string, 0)
	row := ""
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
			row += "S"
		}
	}
}

func main() {
	opts, files := ParseArgs()
	for _, file := range files {
		Cat(file, opts)
	}

}
