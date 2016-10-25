# Markdown Worm #

---

![logo](public/img/logo.jpg)

__v2.0.0__

GoLang Markdown浏览器, 基于[golang-commonmark/markdown](https://github.com/golang-commonmark/markdown), 
[Mou](http://25.io/mou/), 和 [HighlightJS](https://highlightjs.org/)

[English Version](README.md)

## 什么是Markdown ?

一种标记语言,写文档专用.作为码农不会用Markdown写笔记的,砍头!

## 像个男人一样~~十秒就射~~ 分享你的文件！

![boss](public/img/like-a-boss.jpg)

Markdown Worm一般部署在Linux服务器上，用于在公司内网或者小范围内分享各种Markdown文档。它显示时可以自动吧Markdown转换成HTML，并且用JavaScript做代码高亮，方便你的浏览。

## 自v1.0.1的改动

1. 用GoLang代替PHP，现在你什么都不需要安装就可以直接部署__release tarball__了
2. Markdown Worm现在可以识别__md__下的文件夹了
3. 由于Windows的路径和Mac.Linux不太一样，所以没有放出Windows编译版本。

## 传销
###1. 文档列表

![logo](public/img/showcase-1.jpg)

###2. Markdown文件转换成HTML[golang-commonmark/markdown](https://github.com/golang-commonmark/markdown)


![logo](public/img/showcase-2.jpg)

###3. HighLightJS语法高亮 [Highlight.js](https://highlightjs.org/)

![logo](public/img/showcase-3.jpg)

## 注意

由于更换了新的Markdown解析器`golang-commonmark`，现在你必须把`#title#`改成`# title #`或者`# title`否则无法解析

##  安装

### Mac & Linux

1. 下载解压

    	cd /to/dir/
    	# For Mac
    	wget https://github.com/bclicn/MarkdownWorm/releases/download/v2.0.0/mdworm-2.0.0-darwin.tar.gz
    	tar zxvf mdworm-2.0.0-darwin.tar.gz
    
    	# For Linux
   		wget https://github.com/bclicn/MarkdownWorm/releases/download/v2.0.0/mdworm-2.0.0-linux.tar.gz
    	tar zxvf mdworm-2.0.0-linux.tar.gz
    
    
2. 运行
    
   		cd mdworm-2.0.0-*
   		./mdworm
   
3. 打开浏览器输入 `localhost:2333`

## 参数

要修改默认服务器端口（2333），使用 `-p`, 要修改默认代码高亮样式，使用`-t` 。全部代码高亮样式见`./public/css/`,下面是一些例子

    # 更改服务器端口为6666
    ./mdworm -p 6666
    # 更改代码高亮样式为 darcula
    ./mdworm -t darcula
    # 更改服务器端口为6666 并 更改代码高亮样式为 darcula
    ./mdworm -p 6666 -t darcula

一般你会想让这玩意作为背景线程运行的，在命令结尾加上`&`就好了，比如
	
	./mdworm &

## Github真TMD无聊

感觉我基本都是在给自己发`release tarbar`，这么下去我还发个球？又不是所有码农都能达到大神级水平的，难道非得我写一个绅士资源下载器才有人看吗（大雾）？
对于你们这些人，老子只想说一句：

__大爷右上角点颗星儿吧！小的求您了还不行吗？__

__您看过了项目Issue拦留个凸(-_-)凸也行啊！__

## MIT License

===
Beichen Li 2016-10-24
