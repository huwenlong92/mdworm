# Markdown Worm

---
![logo](res/img/logo.jpg)

一个基于 PHP, [ParseDown](http://parsedown.org)和[Mou](http://25.io/mou/)实现的非常简单的Markdown浏览器  

[English Version](README.md)

## 什么是Markdown ?

Markdown是一个轻量级标记语言，能够被轻易的转化为HTML或者PDF。Markdown经常被用来撰写readme文件，用来在论坛上发布消息，或者将简单的文本转化为复杂的表现形式。

顺便一提，你现在看到的这段话就是解析自Markdown。

## 应用场景

我每天都在用Markdown来做各种笔记，比如新技术学习，BUG解决笔记等等。但是如何简单的和同事之间共享这些资源呢？
比如，有天小刘跑过来问：“哥们儿！你以前遇到过这种问题吗？那个程序一直在报错!”

'是!' 你说, '我解决过这个问题，并且写过一个Markdown文档，你可以去http://192.168.1.111/mdWorm查看'

'一个FTP服务器?' 小刘问。

'不是!' 你说, '当你从浏览器访问Markdown文档的时候，服务器会自动帮你转化成HTML，这样你什么都不需要下载，看就好喽！'

## 展示
###1. 显示Markdown文件列表

![logo](res/img/showcase_1.jpg)

###2. 使用Parsedown进行HTML转码，Mou的CSS进行显示

![logo](res/img/showcase_2.jpg)

## 安装
请将你的Markdown文件上传到`./md`目录下，Linux用户请务必将FTP umask设置为022或者文件权限644

### Mac & Win

1. 安装 [XAMPP](https://www.apachefriends.org/index.html)

2. 下载 mdWorm & 解压到 `xampp/htdocs`
3. 在XAMPP图形化界面下启动 Apache Server
4. 打开浏览器输入链接`localhost/mdWorm`


### RHEL & CentOS

1. 安装httpd

		sudo yum install httpd

2. 安装PHP

		sudo yum install php

3. 下载 mdWorm & 解压到`/var/www/html`

		cd /var/www/html
		wget <mdWorm-release-tarball>
		tar -zxvf <mdWorm-release-tarball>

4. 启动Apache Server

		service httpd start

5. 打开浏览器输入

		localhost/mdWorm

### Debian & Ubuntu
1. 安装apache2

		sudo apt-get install apache2

2. 安装PHP

		sudo apt-get install php5 libapache2-mod-php5

3. 下载 mdWorm & 解压到`/var/www/html`

		cd /var/www/html
		wget <mdWorm-release-tarball>
		tar -zxvf <mdWorm-release-tarball>

4. 启动Apache Server

		service apache2 start

5. 打开浏览器输入

		localhost/mdWorm


### MIT License

===
Beichen Li 2016-8-29
