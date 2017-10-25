package plaintext

import (
    "bytes"
    "io"
    "strings"
    "golang.org/x/net/html"
)

func FromHTMLString(input string) (string, error) {
    buffer := bytes.NewBuffer(nil)
    err := FromHTML(strings.NewReader(input), buffer)
    return buffer.String(), err
}

func FromHTML(input io.Reader, output io.Writer) error {
    z := html.NewTokenizer(input)

    inTags := []string{}
    inScript := false
    inStyle := false
    inPre := false
    inTextarea := false
    aHref := ""
    lastOut := ""

    for {
        tt := z.Next()

        switch tt {
        case html.StartTagToken:
            tagName, _ := z.TagName()
            tag := string(tagName)

            if tag == "script" {
                inScript = true
                lastOut = ""
            } else if tag == "style" {
                inStyle = true
                lastOut = ""
            } else if tag == "pre" {
                inPre = true
                lastOut = ""
            } else if tag == "textarea" {
                inTextarea = true
                lastOut = ""
            }

            inTags = append(inTags, tag)

            if !inScript {
                if tag == "hr" {
                    if (lastOut != "\n") {
                        output.Write([]byte("\n"))
                    }
                    output.Write([]byte("------------------------------------------------------------\n"))
                    lastOut = "\n"
                } else if lastOut != "\n" {
                    if tag == "br" || tag == "p" || tag == "div" || tag == "h1" || tag == "h2" || tag == "h3" || tag == "h4" || tag == "h5" || tag == "h6" {
                        output.Write([]byte("\n\n"))
                        lastOut = "\n"
                    } else if tag == "dd" || tag == "li" {
                        output.Write([]byte("\n"))
                        lastOut = "\n"
                    }
                } else if tag == "a" {
                    for {
                        key, val, moreAttr := z.TagAttr()
                        if string(key) == "href" {
                            if len(val) > 0 && val[0] != '#' {
                                aHref = string(val)
                                if !strings.HasPrefix(aHref, "http") {
                                    if strings.HasPrefix(aHref, "//") {
                                        aHref = "http:" + aHref
                                    } else if strings.HasPrefix(aHref, "/") || strings.HasPrefix(aHref, ".") {
                                        aHref = "" // ignore internal/relative links
                                    }
                                }
                            }
                            break
                        }
                        if !moreAttr {
                            break
                        }
                    }
                }
            }

        case html.SelfClosingTagToken:
            tagName, _ := z.TagName()
            tag := string(tagName)

            if tag == "hr" {
                if (lastOut != "\n") {
                    output.Write([]byte("\n"))
                }
                output.Write([]byte("------------------------------------------------------------\n"))
                lastOut = "\n"
            } else if lastOut != "\n" {
                if tag == "br" || tag == "p" || tag == "div" || tag == "h1" || tag == "h2" || tag == "h3" || tag == "h4" || tag == "h5" || tag == "h6" {
                    output.Write([]byte("\n\n"))
                    lastOut = "\n"
                }
            }

        case html.EndTagToken:
            tagName, _ := z.TagName()
            tag := string(tagName)

            index := findTag(inTags, tag)
            if index >= 0 {
                inTags = inTags[0:index]
            }

            stillInTag := findTag(inTags, tag) >= 0
            if (!stillInTag) {
                if tag == "script" {
                    inScript = false
                } else if tag == "style" {
                    inStyle = false
                } else if tag == "pre" {
                    inPre = false
                } else if tag == "textarea" {
                    inTextarea = false
                }
            }

            if !inScript {
                if lastOut != "\n" {
                    if tag == "br" || tag == "p" || tag == "div" || tag == "h1" || tag == "h2" || tag == "h3" || tag == "h4" || tag == "h5" || tag == "h6" {
                        output.Write([]byte("\n\n"))
                        lastOut = "\n"
                    } else if tag == "dd" || tag == "li" {
                        output.Write([]byte("\n"))
                        lastOut = "\n"
                    }
                }

                if tag == "a" {
                    if aHref != "" {
                        output.Write([]byte(" (" + aHref + ") "))
                        aHref = ""
                    }
                }
            }

        case html.TextToken:
            if !inScript && !inStyle {
                data := z.Text()

                if inPre || inTextarea {
                    output.Write(data)
                } else if len(data) > 0 {
                    text := string(data)

                    trimmed := strings.Trim(text, " \t\n")
                    if trimmed == "" && lastOut != " " {
                        output.Write([]byte(" "))
                        lastOut = " "
                    } else {
                        i := strings.Index(text, trimmed)
                        if i > 0 {
                            output.Write([]byte(" "))
                        }
                        l := len(trimmed)
                        output.Write(data[i:i+l])
                        if i+l < len(data) {
                            output.Write([]byte(" "))
                            lastOut = " "
                        } else {
                            lastOut = string(data[i:i+l])
                        }
                    }
                }
            }

        case html.ErrorToken:
            if err := z.Err(); err == io.EOF {
                return nil
            } else {
                return err
            }
        }
    }
    return nil
}

func findTag(tags []string, tag string) int {
    for i := len(tags) - 1; i >= 0; i-- {
        if tags[i] == tag {
            return i
        }
    }
    return -1
}
