package main

import( "bufio"
	"os"
	"strings"
	"net/http"
	"fmt"
	"io"
	"strconv"
	"net/url"
)

const (
        InfoColor    = "\033[1;34m%s\033[0m"
        LinkColor  = "\033[1;36m%s\033[0m"
        HeaderColor = "\033[1;33m%s\033[0m"
        ErrorColor   = "\033[1;31m%s\033[0m"
)


type Link struct {
     Href string
     Alias string
}

var links []*Link
var currURL string
var urlHist []string


func trim(str string) string {
	return strings.TrimSpace(str)
}

func isLink(str string) bool {
	return len(trim(str)) > 3 &&  trim(str)[0:3] == "=> "
}

func isHeader(str string) bool {
	strCut := trim(str)
	idx := strings.Index(strCut, " ")

	res := true

	if idx != -1 {
		hd := strCut[0:idx]
		for i := 0; i < len(hd); i++ {
			if hd[i] != '#' {
				res = false
				break
			}
		}
	} else {
		res = false
	}

	return res
}

func isVerb(str string) bool {
	return len(trim(str)) >= 3 && trim(str)[0:3] == "```"
}

func toHeader(str string) string {
	strCut := trim(str)
	idx := strings.Index(strCut, " ")

	return strCut[idx+1:]
}

func toLink(str string) Link {
     var link Link

     cutStr := trim(str[2:])
     idx := strings.Index(cutStr, " ")

     if idx == -1 {
     	link.Href = cutStr
		link.Alias = link.Href
     } else {
       link.Href = cutStr[:idx]
       link.Alias = cutStr[idx +1:]
     }

	base, _ := url.Parse(currURL)
	rel, _ := url.Parse(link.Href)

	link.Href = base.ResolveReference(rel).String()


     return link
}

func toHref(str string) string {

     link := toLink(str)

     return "=> " + link.Alias
}

func printLinks() {
     for i := range links {
     	 fmt.Printf("%d %s\n", i, links[i].Alias)
     }
}

func printHist() {
	for i := range urlHist {
		fmt.Printf("%d %s\n", i, urlHist[i])
	}
}

func wrapStr(str string, width int) string {
	if width >= len(str) || width < 1 {
		return str
	}

	res := ""
	cutStr := []rune(trim(str))

	var i int

	for i = 0; len(cutStr[i:]) > width; i += width {
		res += string(cutStr[i: i + width]) + "\n"
	}

	res += string(cutStr[i:])

	return res
}

func parse(reader io.Reader, writer *bufio.Writer) {

	scanner := bufio.NewScanner(reader)

	verb := false

	for scanner.Scan() {
		r := scanner.Text()

		if isVerb(r) {
			verb = !verb

			if verb {
				writer.WriteString("====VERBATIM BLOCK====\n")
			} else {
				writer.WriteString("==END VERBATIM BLOCK==\n")
			}
			continue
		}

		if verb {
			writer.WriteString(r + "\n")
			continue
		}


		if isLink(r) {
			writer.WriteString(fmt.Sprintf(LinkColor, toHref(r) + "\n"))
			link := toLink(r)
			links = append(links, &link)
			continue
		}

		if isHeader(r) {
			writer.WriteString(fmt.Sprintf(HeaderColor, toHeader(r) + "\n\n"))
			continue
		}

		writer.WriteString(wrapStr(r, 80) + "\n")

	}

	if verb {
		writer.WriteString("==END VERBATIM BLOCK==\n")
	}

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func getnprint(url string, refresh bool) {
    resp, err := http.Get(url)

    if err != nil {
     	fmt.Printf(ErrorColor, "Error when open url\n")
		return
    }

    defer resp.Body.Close()

    currURL = url
	urlHist = append(urlHist, currURL)

    writer := bufio.NewWriter(os.Stdout)
    parse(resp.Body, writer)
    if refresh {
        clearScreen()
    }
    writer.Flush()
}

func openLink(idx int) {
     if len(links)-1 >= idx {
     	href := links[idx].Href
	links = nil
     	getnprint(href, true)
     }
}

func openUrlHist(idx int) {
     if len(urlHist)-1 >= idx {
     	href := urlHist[idx]
		links = nil
     	getnprint(href, true)
     }
}


func main() {
     var c string
     var url string
     for c != "q" {
     	 fmt.Printf(InfoColor, "\nEnter command:")
     	 fmt.Scanf("%s", &c)

	 if c == "o" {
	    fmt.Printf(InfoColor, "\nEnter url:")
	    fmt.Scanf("%s", &url)
	    if url != "" {
	       links = nil
	       getnprint(url, true)
	    }
	    continue
	 }

	 if c == "l" {
	    printLinks()
	    continue
	 }

	 if c == "h" {
		printHist()
	 	continue
	 }

	 if len(c) > 1 && c[0] == 'h' {
	    idx, err := strconv.Atoi(c[1:])
	    if err == nil {
	       openUrlHist(idx)
	       continue
	    }

	 }

	 if len(c) > 1 && c[0] == 'l' {
	    idx, err := strconv.Atoi(c[1:])
	    if err == nil {
	       openLink(idx)
	       continue
	    }

	 }
     }
}
