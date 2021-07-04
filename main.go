package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func buildIndexHandler(paths []string) func(http.ResponseWriter, *http.Request) {
	index_tmpl := template.Must(template.New("index").Funcs(template.FuncMap{"makeRouteName": makeRouteName}).Parse(
		`
<html>
	<head>
		<meta charset="UTF-8">
		<title>Folders</title>
	</head>
	<body>
		<h2>Folders</h2>
		{{range $index, $element := .}}<a href="{{makeRouteName $element}}">{{$element}}</a></br>{{end}}
	</body>
</html>
`,
	))

	return func(w http.ResponseWriter, r *http.Request) {
		index_tmpl.Execute(w, paths)
	}
}
func makeRouteName(foldername string) string {
	stripped := filepath.Base(foldername)
	return fmt.Sprintf("/%s/", stripped)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	port := flag.String("p", "8080", "port to run server")
	flag.Parse()

	folders := flag.Args()
	if len(folders) == 0 {
		printUsage()
		return
	}

	var filtered_folders []string
	for _, f := range folders {
		if !checkIsDir(f) {
			continue
		}
		if f == "." {
			//Handles a pointer to the current working dir
			abs_path, _ := filepath.Abs(f)
			f = fmt.Sprintf("../%s", filepath.Base(abs_path))
		}
		filtered_folders = append(filtered_folders, f)
	}

	if len(filtered_folders) == 0 {
		printUsage()
		return
	}

	for _, f := range filtered_folders {
		fs := http.FileServer(http.Dir(f))
		routename := makeRouteName(f)
		http.Handle(routename, http.StripPrefix(routename, fs))
	}

	http.HandleFunc("/", buildIndexHandler(filtered_folders))

	fmt.Println("Starting server at :8080")
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil).Error())
}
