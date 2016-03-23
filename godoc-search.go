/**
 * godoc-search is a command line tool that queries godoc.org for packages
 *
 * author: mparaiso <mparaiso@online.fr>
 * license: gpl-3.0
 *
 * usage :
 *
 *			godoc-search expression
 * example:
 *          godoc-search mysql
 *
 */
package main

import "flag"

import "net/http"
import "strings"
import "encoding/json"
import "text/template"
import "os"

// QueryResults represents a list of results from a search query
type QueryResults struct {
	Results []struct {
		Path     string
		Synopsis string
	}
}

const resultTemplate = `
{{range $index,$result := .Results}}
{{$result.Path}}
	{{$result.Synopsis}}

{{end}}
`

func main() {
	const apiURL = "http://api.godoc.org/search?&q="
	flag.Parse()
	requestURL := apiURL + strings.Join(flag.Args(), " ")
	response, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	jsonDecoder := json.NewDecoder(response.Body)
	queryResult := &QueryResults{}
	err = jsonDecoder.Decode(queryResult)
	if err != nil {
		panic(err)
	}
	stringTemplate, err := template.New("master").Parse(resultTemplate)
	if err != nil {
		panic(err)
	}
	err = stringTemplate.Execute(os.Stdout, queryResult)
	if err != nil {
		panic(err)
	}
}
