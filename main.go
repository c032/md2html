package main

import (
	"bytes"
	"flag"
	"html/template"
	"io"
	"io/ioutil"
	logger "log"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var log = logger.New(os.Stderr, "", 0)

func Render(w io.Writer, r io.Reader) error {
	var (
		err         error
		rawMarkdown []byte

		tmpl *template.Template
	)

	rawMarkdown, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	unsafe := blackfriday.Run(rawMarkdown)
	safeHTML := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	tmpl, err = template.New("main").Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, struct {
		Content template.HTML
	}{
		Content: template.HTML(safeHTML),
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var flagOutputFile string

	flag.StringVar(&flagOutputFile, "o", "", "output file")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "-")
	}

	var w io.Writer
	if flagOutputFile == "" {
		w = os.Stdout
	} else {
		var (
			err  error
			outf *os.File
		)

		outf, err = os.Create(flagOutputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer outf.Close()

		w = outf
	}

	var stdinContent []byte

	for _, file := range args {
		if file == "-" {
			var err error

			if len(stdinContent) == 0 {
				stdinContent, err = ioutil.ReadAll(os.Stdin)
				if err != nil {
					log.Fatal(err)
				}
			}

			r := bytes.NewReader(stdinContent)

			err = Render(w, r)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			var (
				err error
				f   *os.File
			)

			f, err = os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			Render(w, f)
		}
	}

}
