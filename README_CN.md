# Markdown Worm #

---
![logo](res/img/logo.jpg)

__v1.0.1__

一个基于 PHP, [ParseDown](http://parsedown.org),[Mou](http://25.io/mou/)和[HighlightJS](https://highlightjs.org/)实现的非常简单的Markdown浏览器  

[English Version](README.md)

## 什么是Markdown ?

Markdown是一个轻量级标记语言，能够被轻易的转化为HTML或者PDF。Markdown经常被用来撰写readme文件，用来在论坛上发布消息，或者将简单的文本转化为复杂的表现形式。

## 使用

建立一个能够执行PH脚本的HTTP服务器(比如阿帕奇),将mdWorm放入网站目录,将你的Markdown文件上传到`./md`文件夹下.
mdWorm会自动列出`./md`文件夹下的所有Markdown文件,并在你点击链接时自动将其解析为HTML,自动检测代码类型并高亮(无需在Markdown中
指定代码类型,比如````java`).

## 展示
###1. 显示Markdown文件列表

![logo](res/img/showcase_1.jpg)

###2. 使用Parsedown进行HTML转码，Mou的CSS进行显示

![logo](res/img/showcase_2.jpg)

###3. 自动检测语言类型并高亮代码

![logo](res/img/showcase_3.jpg)

## 安装
请将你的Markdown文件上传到`./md`目录下，Linux用户请务必将FTP umask设置为022或者markdown文件权限644

### Mac & Win

1. 安装 [XAMPP](https://www.apachefriends.org/index.html)
2. 下载 mdWorm & 解压到 `xampp/htdocs`
3. 在XAMPP图形化界面下启动 Apache Server
4. 打开浏览器输入链接`localhost/mdWorm`


### RHEL & CentOS

1. 安装httpd

		sudo yum install httpd -y

2. 安装PHP

		sudo yum install php -y

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
Beichen Li 2016-9-5
