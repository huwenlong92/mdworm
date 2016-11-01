# Markdown Worm #

---

![logo](public/img/logo.jpg)

__v2.0.0__

A simple GoLang Markdown viewer, based on [golang-commonmark/markdown](https://github.com/golang-commonmark/markdown), 
[Mou](http://25.io/mou/), and [HighlightJS](https://highlightjs.org/)

[中文版说明](README_CN.md)

## Intro

Markdown Worm is a simple GoLang file server for storing & sharing Markdown notes.
It's useful for you and your teammates to share technical documents and thoughts. 
One can easily write a markdown doc and upload it to mdworm by a FTP client. 
You can see the content of Markdown file as a rendered html on your web browser.

## Usage

1. download mdworm binary release and place its folder under a path which can be accessed by FTP client.
2. start mdworm by `$./mdworm -p <port>`, also make sure that port is allowed by firewall
3. put some markdown files (or new folders) under `./md`
4. open your browser and visit `http://<mdworm-host>:<mdworm-port>`

## Changes since v1.0.1

The first version of mdworm is written in PHP, which is pretty slow and old fashioned, so at
v2.0.0, it was rewritten by GoLang. Plus mdworm can now recognize folders under `./md`.
Plus plus this time windows is not supported.

## Showcase
###1. Show File & Folders under `./md`

![logo](public/img/showcase-1.jpg)

###2. [golang-markdown](https://github.com/golang-commonmark/markdown) will convert md to HTML

![logo](public/img/showcase-2.jpg)

###3. [Highlight.js](https://highlightjs.org/) highlights your code

![logo](public/img/showcase-3.jpg)

## Notice

The new md parser has a weird feature, which __requires__ you to add a space between
title and its surrounding hash tags, for example `#title#` won't be parsed correctly, you
will write `# title #` instead

## Install

### Mac & Linux

1. Download & Unzip

        cd /to/dir/
        # For Mac
        wget https://github.com/bclicn/MarkdownWorm/releases/download/2.0.0/mdworm-2.0.0-darwin-bin.tar.gz
        tar zxvf mdworm-2.0.0-darwin.tar.gz
    
        # For Linux
        https://github.com/bclicn/MarkdownWorm/releases/download/2.0.0/mdworm-2.0.0-linux-bin.tar.gz
        tar zxvf mdworm-2.0.0-linux.tar.gz
    
2. Run
    
        cd mdworm-2.0.0-*
        ./mdworm
   
3. Open your browser and type `localhost:2333`
4. You can also point out the port by `-p`, code theme by `-t`, like bellow,
A full list of highlights themes can be found under `./public/css/themes/`

        # listen on port 6666
        ./mdworm -p 6666
        
        # change highlight theme to darcula
        ./mdworm -t darcula
        
        # listen on port 6666 & change highlight theme to darcula
        ./mdworm -p 6666 -t darcula
        
5. If you want to (usually will) run in daemon mode, just add `&` at the end

## MIT License

===
Beichen Li 2016-11-1
