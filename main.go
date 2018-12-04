package main

import (
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

func Render(title string, w io.Writer, markdown []byte) error {
	var (
		err  error
		tmpl *template.Template
	)

	unsafe := blackfriday.Run(markdown)
	safeHTML := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	tmpl, err = template.New("main").Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, struct {
		Title   string
		Content template.HTML
	}{
		Title:   title,
		Content: template.HTML(safeHTML),
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var (
		flagTitle      string
		flagOutputFile string
	)

	flag.StringVar(&flagTitle, "t", "", "document title")
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

	var (
		stdinContent    []byte
		documentContent []byte
	)

	for _, file := range args {
		if file == "-" {
			var err error

			if len(stdinContent) == 0 {
				stdinContent, err = ioutil.ReadAll(os.Stdin)
				if err != nil {
					log.Fatal(err)
				}
			}

			documentContent = append(documentContent, stdinContent...)
		} else {
			var (
				err error
				f   *os.File

				fileContent []byte
			)

			f, err = os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			fileContent, err = ioutil.ReadAll(f)
			if err != nil {
				log.Fatal(err)
			}

			documentContent = append(documentContent, fileContent...)
		}
	}

	err := Render(flagTitle, w, documentContent)
	if err != nil {
		log.Fatal(err)
	}
}
