package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	clientmode = flag.String("clientmode", "dev", "Serve dev or build")
	port       = flag.String("p", "5555", "Set proxy Port")
	clientPath = flag.String("clientPath", "../client", "Set proxy Port")
)

func main() {

	remote, err := url.Parse("http://localhost:8181")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.Handle("/api/", &ProxyHandler{proxy})

	if *clientmode == "build" {
		// gebuildete Version ausliefern
		fs := http.FileServer(http.Dir("../client/build/esm-unbundled"))
		http.Handle("/", fs)
	}

	if *clientmode == "dev" {
		http.HandleFunc("/", serveDevClient)
	}

	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		panic(err)
	}

}

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	r.Header.Add("api-base-url", "http://localhost:5555/api")
	// rewrite /api/resource... to /resource...
	r.URL.Path = r.URL.Path[4:len(r.URL.Path)]
	ph.p.ServeHTTP(w, r)
}

func serveDevClient(w http.ResponseWriter, r *http.Request) {

	FilePath := filepath.Join("../client", filepath.Clean(r.URL.Path))
	fileExtension := path.Ext(FilePath)
	info, err := os.Stat(FilePath)

	// regular 404 for files with extensions, deep links should not have a file extension
	if err != nil {
		if os.IsNotExist(err) && fileExtension != "" {
			http.NotFound(w, r)
			return
		}
		if os.IsNotExist(err) && fileExtension == "" {
			// deep link
			FilePath = filepath.Join("../client/index.html")
			fileExtension = path.Ext(FilePath)
			info, _ = os.Stat(FilePath)
		}
	}

	if info.IsDir() {
		// Return index on /
		if FilePath == "../client" {
			FilePath = filepath.Join("../client/index.html")
			fileExtension = path.Ext(FilePath)
			info, _ = os.Stat(FilePath)

		} else {
			// Return a 404 if the request is  a directory
			http.NotFound(w, r)
			return
		}
	}

	//First of check if Get is set in the URL

	if FilePath == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}

	//Check if file exists and open
	Openfile, err := os.Open(FilePath)
	defer Openfile.Close() //Close after function return

	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}

	//Send the headers
	FileContentType := DetermineMimeType(fileExtension)
	w.Header().Set("Content-Type", FileContentType)

	//todo Rewrite Imports in JS files to relative Imports <- ist das wirklich nötig?
	if fileExtension == ".js" {

		input, err := ioutil.ReadFile(FilePath)
		if err != nil {
			log.Fatalln(err)
		}

		lines := strings.Split(string(input), "\n")

		for i, line := range lines {

			if strings.Contains(line, "import ") || strings.Contains(line, "export ") {
				lines[i] = replaceImport(lines[i], FilePath)
			}
		}
		output := strings.Join(lines, "\n")
		var r io.Reader
		r = strings.NewReader(output)
		io.Copy(w, r)
		return
	}

	//Send the file
	io.Copy(w, Openfile) //'Copy' the file to the client
	return

}

// replaces node imports with browser compatible imports
func replaceImport(line string, currentFile string) string {
	currentPath := path.Dir(currentFile)
	regex := regexp.MustCompile(`(import|export) *(\{.*\})?.*?.*['"](.*)['"]`)
	matches := regex.FindStringSubmatch(line)
	if len(matches) > 0 && matches[3] != "" {
		//modules := matches[2]
		importSegment := matches[3]
		// remove .js
		importSegment = strings.Replace(importSegment, ".js", "", 1)
		originalImportSegment := matches[3]

		// leave relative paths as is, all other shows to node_modules
		if len(importSegment) > 1 && importSegment[0:1] != "." {
			importSegment = "/node_modules/" + importSegment
			currentPath = "../client" //set to "root"
		}

		// Simplest check: file existence
		cleanPath := filepath.Clean(currentPath + "/" + importSegment)
		_, err := os.Stat(cleanPath + ".js")
		if err == nil {
			// file exists
			importSegment = importSegment + ".js"
			return strings.Replace(line, originalImportSegment, importSegment, 1)
		}

		// directory check
		// ./@furo/fbp ==> ./@furo/fbp/fbp.js
		info, err := os.Stat(cleanPath)
		if err == nil {
			if info.IsDir() {
				// try with directory name as filename
				dirs := strings.Split(importSegment, "/")
				_, err = os.Stat(cleanPath + "/" + dirs[len(dirs)-1] + ".js")
				if err == nil {
					// file exists
					importSegment = importSegment + "/" + dirs[len(dirs)-1] + ".js"
					return strings.Replace(line, originalImportSegment, importSegment, 1)
				}
			}
		}

		// go for main from package.json
		if matches[1] == "import" {

			// open package.json
			packageJson := "../client/" + importSegment + "/package.json"

			// Open our jsonFile
			jsonFile, err := os.Open(packageJson)
			defer jsonFile.Close()
			// if we os.Open returns an error then handle it
			if err == nil {
				byteValue, _ := ioutil.ReadAll(jsonFile)
				var result map[string]interface{}
				json.Unmarshal([]byte(byteValue), &result)
				importSegment = importSegment + "/" + result["main"].(string)
				return strings.Replace(line, originalImportSegment, importSegment, 1)

			}

		}
	}

	return line
}

func DetermineMimeType(fileExtension string) string {
	return mime.TypeByExtension(fileExtension)
}
