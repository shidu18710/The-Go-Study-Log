package main

import (
	"fmt"
	"go02/src/g18/tempconv"
	"os"
	"strconv"
)

func main() {
	// 包和文件
	fmt.Printf("%v\n", tempconv.CToF(212.2))
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s=%s,%s=%s\n", f, tempconv.FToc(f), c, tempconv.CToF(c))
	}
}
