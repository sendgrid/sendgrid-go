package textprocess

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

// Html2Txt converts an html string to its polished textual representation.
func Html2Txt(_html string) (txt string) {
	out := bytes.NewBuffer(nil)
	in := strings.NewReader(_html)
	_ = StreamHtml2Txt(out, in)
	polished := bytes.NewBuffer(nil)
	_ = PolishWhiteSpace(polished, out)
	return string(polished.Bytes())
}

// SteamHtml2Txt reads the input stream as HTML and writes the plaintext conversion
// to output stream
func StreamHtml2Txt(out io.Writer, in io.Reader) error {
	tokenizer := html.NewTokenizer(in)
parseLoop:
	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			break parseLoop // End of the document,  done
		case tt == html.TextToken:
			txt := html.UnescapeString(string(tokenizer.Text()))
			if len(txt) > 0 {
				_, err := out.Write([]byte(txt))
				if err != nil {
					return errors.New("textprocess: " + err.Error())
				}
			}
		}
	}
	return nil
}

// PolishWhiteSpace removes empty lines and leading/trailing whitespace on each line
func PolishWhiteSpace(out io.Writer, in io.Reader) error {
	scan := bufio.NewScanner(in)
	isEmptyLine := func(s string) bool {
		for _, r := range s {
			if !unicode.IsSpace(r) {
				return false
			}
		}
		return true
	}
	for scan.Scan() {
		t := scan.Text()
		if isEmptyLine(t) {
			continue
		}
		_, err := fmt.Fprintln(out, strings.TrimSpace(t))
		if err != nil {
			return errors.New("textprocess: " + err.Error())
		}
	}
	return nil
}
