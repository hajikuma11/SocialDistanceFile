package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("ファイルを１つだけ指定してください(Please specify only one file)")
		return
	}

	filepath := flag.Arg(0)

	_, fname := path.Split(filepath)

	rfp, err := os.Open(fname)
	defer rfp.Close()
	if os.IsNotExist(err) {
		fmt.Println("ファイルが存在しません。(file does not exist.)")
		return
	}

	wfp, err := os.Create("sdf-" + fname)
	defer wfp.Close()
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	lines := bufio.NewScanner(rfp)

	for lines.Scan() {
		line := lines.Text()

		io.WriteString(wfp, (line + "\n"))

		for j := 0; j < 571; j++ {
			io.WriteString(wfp, ("\n"))
		}
	}

	if err != nil {
		log.Fatal(err)
		return
	}

}
