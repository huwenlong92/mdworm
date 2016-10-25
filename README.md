# Markdown Worm #

---

![logo](public/img/logo.jpg)

__v2.0.0__

GoLang Markdown viewer, based on [golang-commonmark/markdown](https://github.com/golang-commonmark/markdown), 
[Mou](http://25.io/mou/), and [HighlightJS](https://highlightjs.org/)

[中文版说明](README_CN.md)

## What is Markdown ?

Markdown is a lightweight markup language with plain text formatting syntax designed so that it can be converted to HTML and many other formats using a tool by the same name. Markdown is often used to format readme files, for writing messages in online discussion forums, and to create rich text using a plain text editor.

## Finish your targets ~~quietly~~ LIKE A BOSS

![boss](public/img/like-a-boss.jpg)

Markdown Worm act as a simple markdown file server. Which is useful on storing & sharing Markdown notes & documents, especially on a local network with
you colleagues (which I'm doing) 

## Changes since v1.0.1

1. Completely rewrite GoLang, speed up & don't need anything to run compiled code
2. It can now recognize folders under `./md`
3. No windows support by now

## Showcase
###1. Show File & Folders under `./md`

![logo](public/img/showcase-1.jpg)

###2. [golang-commonmark/markdown](https://github.com/golang-commonmark/markdown) convert it
to HTML

![logo](public/img/showcase-2.png)

###3. [Highlight.js](https://highlightjs.org/) detect code type and highlight it for you (theme darcula)

![logo](public/img/showcase-3.jpg)

## Notice

`golang-commonmark` has a weird feature, which __requires__ you to add a space between
title text and hash tags, for example `#title#` won't be parsed correctly, but
`# title #` will work fine

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

## Arguments

You can use `-p` to change the default port (`2333`), `-t` to change default code theme (`./public/css/themes/default.css`).
A full list of highlights theme can be found under `./public/css/themes/`, 

here are some command line examples:

    # listen on port 6666
    ./mdworm -p 6666
    # change default theme to darcula
    ./mdworm -t darcula
    # listen on port 6666 & change default theme to darcula
    ./mdworm -p 6666 -t darcula

    # to run in background just add &
    ./mdworm -p 6666 &
    
## MIT License

===
Beichen Li 2016-10-25
