package main

import (
	"github.com/golang-commonmark/markdown"	// go get github.com/golang-commonmark/markdown
	"fmt"
	"io/ioutil"
	"strconv"
	"net/http"
	"time"
	"strings"
	"os"
	"bufio"
)

// Markdown Worm
// v2.0.0
// Beichen Li, bclicn@gmail.com. relidin@126.com, 2016
// A simple markdown viewer for OSX and Linux
// usage: $./mdworm [-p <port>] [-t <theme>]

const (
	TITLE	    =   `Markdown Worm`
	FOOTER	    =	`<hr>` +
				`<p>Markdown Worm v2.0.0, @ bcli 2016, <a href="https://github.com/bclicn/mdWorm">GitHub</a></p>` +
				`<p>Render: {$cost}</p>` +
				`<h5 style="text-align:right"><a href="#" onclick="window.history.back()">Back</a><h5>`

	LS_TEMPLATE  =	`<!DOCTYPE html>` +
			`<head>`+
				`<meta charset="utf8">` +
				`<link rel="icon" href="/public/img/favicon.ico">` +
				`<title>{$title}</title>` +
				`<link href="/public/css/mou.css" rel="stylesheet">` +
			`</head>` +
			`<body>` +
					`<h5 style="text-align:right"><a href="#" onclick="window.history.back()">Back</a><h5>` +
					`<h1>{$title}</h1><hr>` +
					`{$list}` +
			`{$footer}` +
			`</body>` +
			`</html>`

	MD_TEMPLATE =	`<!DOCTYPE html>` +
			`<head>`+
				`<meta charset="utf8">` +
				`<link rel="icon" href="/public/img/favicon.ico">` +
				`<title>{$title}</title>` +
				`<link href="/public/css/mou.css" rel="stylesheet">` +
				`<link href="/public/css/themes/{$theme}.css" rel="stylesheet">` +
				`<script src="/public/js/highlight.pack.js"></script>` +
				`<script>hljs.initHighlightingOnLoad();</script>` +
			`</head>` +
			`<body>` +
				`<h5 style="text-align:right"><a href="#" onclick="window.history.back()">Back</a><h5>` +
				`{$markdown}` +
			`{$footer}` +
			`</body>` +
			`</html>`
)

var (
	codeTheme = "default"
)

func main(){
	var port int = 2333
	var intParseErr error
	var cmd, themePath string
	var cmdDeclared bool = false
	if len(os.Args) > 1{
		for key, val := range os.Args{
			if key == 0 {continue}
			if cmdDeclared{
				switch cmd{
				case "-p":
					port, intParseErr = strconv.Atoi(val)
					if intParseErr != nil{
						fmt.Println("Invalid port: " + val)
						os.Exit(1)
					}else{
						cmdDeclared = false
					}
				case "-t":
					themePath = "public/css/themes/" + val +".css"
					if _, err := os.Stat(themePath); os.IsNotExist(err) {
						fmt.Println("Couldn't find code theme: " + themePath)
						os.Exit(1)
					}else{
						codeTheme = val
						cmdDeclared = false
					}
				default:
					fmt.Println("Unkown command: " + cmd)
					os.Exit(1)
				}
			}else{
				switch val{
				case "-p": cmd = "-p"; cmdDeclared = true; break
				case "-t": cmd = "-t"; cmdDeclared = true; break
				default: fmt.Println("Unkown command: " + val); os.Exit(1)
				}
			}
		}
		if cmdDeclared{
			fmt.Println("Please assign a value for command " + cmd)
			os.Exit(1)
		}

	}
	fmt.Print("Markdown Worm will be launched on port [")
	fmt.Print(port)
	fmt.Println("] by code theme [" + codeTheme + "]")
	server(port)
}

func server(port int){
	http.HandleFunc("/",		rootHandler)
	http.HandleFunc("/ls/",		lsHandler)
	http.HandleFunc("/md/", 	mdHandler)
	http.Handle("/public/", 	http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	fmt.Println(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/ls/md", 301)
}

func lsHandler(w http.ResponseWriter, r *http.Request){
	var title string = TITLE
	var list, name, mdPath string
	var req time.Time = time.Now()

	dir := strings.Replace(r.URL.Path,"/ls/","",1)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		list = `<a style="color:red">[ERROR] ` + dir + ` is not a directory or it doesn't exist</a>`
		fmt.Fprintf(w,page("ls", "Error", list, time.Since(req)))
	}
	list = "<ul>\n"
	for _, file := range files {
		name = file.Name()
		if file.IsDir(){
			list += "<li>"
			list += file.ModTime().Format("2006-01-01 15:04:05") + "&nbsp;&nbsp;"
			list += "<a href=\"/ls/" + dir + "/" + name + "\"><strong>"+ name + "/</strong></a>"
			list += "</li>\n"
		}else{
			if strings.HasSuffix(name,".md") || strings.HasSuffix(name,".markdown"){
				list += "<li>"
				list += file.ModTime().Format("2006-01-01 15:04:05") + "&nbsp;&nbsp;"
				mdPath = dir + "/" + name
				detail := getTitleAndAuthor(mdPath)
				list +="<a href=\"/md/" + dir + "/" + name + "\">" + detail[0] + "</a>"
				list += `<a style="float:right">` + detail[1] + "</a>"
				list += "</li>\n"
			}else{
				continue
			}
		}
	}
	list +="</ul>\n"
	if list == "<ul>\n</ul>\n"{
		list = `<ul><li><h5><a href="#" onclick="window.history.back()">Empty Directory</a><h5></li></ul>`
	}
	fmt.Fprintf(w,page("ls", title, list, time.Since(req)))
}

func mdHandler(w http.ResponseWriter, r *http.Request){
	var content, title string
	var req time.Time = time.Now()
	mdPath := strings.Replace(r.URL.Path,"/md/","",1)
	file, err := ioutil.ReadFile(mdPath)
	detail := getTitleAndAuthor(mdPath)
	title = detail[0]
	if err != nil{
		content = `<a style="color:red">[ERROR]` + mdPath + ` doesn't exist</a>`
		fmt.Fprintf(w,page("md","Error",content,time.Since(req)))
	}
	md := markdown.New(markdown.XHTMLOutput(true), markdown.Nofollow(true))
	content = md.RenderToString([]byte(file))
	fmt.Fprintf(w,page("md",title,string(content),time.Since(req) ))
}

func page (template string, title string, content string, cost time.Duration) string{
	var page string
	switch template {
	case "ls" :
		page = LS_TEMPLATE
		page = strings.Replace(page, "{$title}", 	title, 		-1)
		page = strings.Replace(page, "{$list}",		content, 	-1)
		page = strings.Replace(page, "{$footer}", 	footer(cost), 	-1)
		return page
	case "md" :
		page = MD_TEMPLATE
		page = strings.Replace(page, "{$title}", 	title, 		-1)
		page = strings.Replace(page, "{$markdown}",	content, 	-1)
		page = strings.Replace(page, "{$theme}",	codeTheme,	-1)
		page = strings.Replace(page, "{$footer}", 	footer(cost), 	-1)
		return page
	default:
		page = LS_TEMPLATE
		page = strings.Replace(page, "{$title}", 	title, 		-1)
		page = strings.Replace(page, "{$markdown}",	"[ERROR] Unkown template: " + template, -1)
		page = strings.Replace(page, "{$footer}", 	footer(cost), 	-1)
		return page
	}
	return ""
}

func footer (cost time.Duration) string{
	footer 	:= FOOTER
	return strings.Replace(footer,"{$cost}",cost.String(),1)
}

func getTitleAndAuthor (mdPath string) []string{
	var title string = ""
	var author string = ""
	var line string
	file, err := os.Open(mdPath)
	if err != nil{
		title = "<i>" + mdPath + " Permission denined</i>"
		author = "by Unkown"
		result := []string{title,author}
		return result
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			line = strings.Replace(line, "#", "", 2)
			title = line
		} else if strings.HasPrefix(line, "by"){
			author = line
		}
		if title != "" && author != ""{
			break
		}
	}
	result := []string{title,author}
	return result
}