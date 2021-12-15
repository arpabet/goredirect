package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	Version string
	Build   string
)


var templateIndex = `
<!DOCTYPE html>
<html lang="en">

<head>
<meta charset="UTF-8">
<title>Golang</title>
<link href="https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700" rel="stylesheet">
<meta name="go-import" content="{{ .ModuleName }}/{{ .ReferenceName }} git https://{{ .GitHost }}/{{ .GitUser }}/{{ .ProjectName }}">
<meta name="go-source" content="{{ .ModuleName }}/{{ .ReferenceName }} _ https://{{ .GitHost }}/{{ .GitUser }}/{{ .ProjectName }}/tree/master{/dir} https://{{ .GitHost }}/{{ .GitUser }}/{{ .ProjectName }}/blob/master{/dir}/{file}#L{line}">
</head>

<body>
</body>

</html>
`

func main() {

	if err := doMain(); err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func doMain() error {

	log.Printf("Welcome to GoRedirect Version %s, Biuld %s", Version, Build)

	listenAddr := flag.String("listen", "localhost:7001", "address and port to serve on")
	to := flag.String("to", "github.com", "destination golang repo address")
	user := flag.String("user", "arpabet", "destination golang repo username")

	flag.Parse()
	flag.PrintDefaults()

	indexPage, err := NewRedirectPage(templateIndex, *to, *user)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/",  indexPage)

	server := &http.Server{
		Addr:         *listenAddr,
		ReadTimeout:  5 * time.Minute, // 5 min to allow for delays when 'curl' on OSx prompts for username/password
		WriteTimeout: 10 * time.Second,
		Handler:      mux}

	return server.ListenAndServe()

}
