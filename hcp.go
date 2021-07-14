/*
 * Author(s):
 * - Mattia Cabrini <dev@mattiacabrini.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	bufferSize = 4096
)

func print_usage(n int) {
	println("Usage:")
	print(os.Args[0])
	println(" [:port] file")
	os.Exit(n)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "<html>\n")

	fmt.Fprintf(w, "<head>\n")
	fmt.Fprintf(w, "<title>HTTP Copy</title>\n")
	fmt.Fprintf(w, "</head>\n")

	fmt.Fprintf(w, "<body>\n")

	fmt.Fprintf(w, "<h1>HTTP Copy Index</h1>\n")
	fmt.Fprintf(w, "<hr />\n")

	fmt.Fprintf(w, "<ul>\n")

	if os.Args[1][:1] != ":" {
		fmt.Fprintf(w, "<li><a href=\"/%d\">%s</a></li>\n", 1, os.Args[1])
	}

	for c := 2; c < len(os.Args); c = c + 1 {
		fmt.Fprintf(w, "<li><a href=\"/%d\">%s</a></li>\n", c, os.Args[c])
	}

	fmt.Fprintf(w, "</ul>\n")

	fmt.Fprintf(w, "</body>\n")
	fmt.Fprintf(w, "</html>\n")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["file"])

	if err != nil || num < 1 || (num == 1 && os.Args[1][:1] == ":") || len(os.Args) <= num {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(os.Stdout, "BAD REQUEST\n\tfile #%d\n\terr  %s`", num, err)
		return
	}

	file := os.Args[num]
	fp, err := os.Open(file)
	count := 0

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't open file"))
		log.Fatal("Can't open file")
	}
	defer fp.Close()

	stats, statsErr := fp.Stat()
	if statsErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't stat"))
		log.Fatal("Can't stat")
	}

	var size int64 = stats.Size()
	bytes := make([]byte, bufferSize)

	w.Header().Set("Content-Length", strconv.FormatInt(size, 10))

	bufr := bufio.NewReader(fp)

	for count, err = bufr.Read(bytes); err == nil && count > 0; count, err = bufr.Read(bytes) {
		count, err = w.Write(bytes[:count])
	}
}

func main() {
	port := ":8000"
	l := len(os.Args)
	r := mux.NewRouter()

	if l > 1 {
		if os.Args[1][:1] == ":" {
			port = os.Args[1]

			if port[:1] == ":" && l < 3 {
				print_usage(1)
			}
		}
	} else {
		print_usage(1)
	}

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/{file}", fileHandler)

    fmt.Fprint(os.Stdout, "Serving on port ", port)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(port, nil))
}
