package main

import (
	"flag"
	"fmt"

	//"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

func main() {
	flag.Parse()
	mini := Minify(flag.Arg(0))
	fmt.Printf("%v\n", mini)
}

func Minify(inputFilePath string) string {
	fname, _ := filepath.Abs(inputFilePath)
	//log.Println(outputCssPath)

	// css をがあるか調べる
	_, err := os.Stat(fname)
	if err != nil {
		// cssファイルがない
		// デフォルトのCSSを使う
		// minifyしない
		return "default"
	}
	// ファイル読み込み
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	inputCss := string(bytes)

	mediatype := "text/css"
	m := minify.New()
	m.AddFunc(mediatype, css.Minify)
	minifiedCss, _ := m.String(mediatype, inputCss)

	//log.Println(minifiedCss)
	return minifiedCss
}
