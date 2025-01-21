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
		fmt.Println("Error reading file", filename, err)
		return
	}

	if opts.V {
		fmt.Println("cat (GNU coreutils) 9.4")
		return
	}

	rows := make([]string, 0)
	row := ""
	previousLine := ""
	countrow := 1

	if opts.E {
		row += "$"
	}

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
			row = strings.Replace(row, "\t", "^I", -1)
		}

		if opts.N {
			fmt.Printf("\t%d ", numRow+1)
		}

		if opts.E {
			row += "$"
		}

		if opts.S {
			if strings.TrimSpace(row) == "" && strings.TrimSpace(previousLine) == "" {
				continue
			}
		}

		if opts.B {
			trimResult := strings.TrimSpace(row)

			if trimResult != "" {
				fmt.Printf("\t%d ", countrow)
				countrow++
			}
		}

		previousLine = row
		fmt.Print(row)

	}
}

// TODO:
// 1. Подправить работу комбинаций флагов (сначала почекать, потом исправить) +
// 2. Подправить работу флага -e, сейчас выводит $ на новой строке
// 3. При комбинации флагов -s и -n ломается нумерация строк, починить +

func main() {
	opts, files := ParseArgs()

	if opts.B && opts.N {
		opts.N = false
	}

	if len(files) == 0 && !opts.H {
		fmt.Print("your file is empty")
	}

	if opts.H {
		fmt.Println(

			"-b, Нумерация только непустых строк\n",
			"-e, Показать символы $ в конце строк\n",
			"-n, Нумерация всех строк\n",
			"-s, Сжать последовательные пустые строки\n",
			"-t, Отображение табуляции как ^I\n",
			"-h, Показать справку\n",
			"-v, Показать версию программы",
		)
	} else {
		for _, file := range files {
			Cat(file, opts)
		}
	}

}
