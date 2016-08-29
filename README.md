# MD Worm

---
![logo](res/img/logo.jpg)

A very simple markdown viewer implemented by PHP, [ParseDown](http://parsedown.org) and [Mou](http://25.io/mou/)

[中文版说明点这里](README_CN.md)

## Showcase
###1. Markdown File List

![logo](res/img/showcase_1.jpg)

###2. Converted by Parsedown and displayed by CSS from Mou

![logo](res/img/showcase_2.jpg)

## Install

### Mac & Win

1. Install [XAMPP](https://www.apachefriends.org/index.html)

2. Download mdWorm & Unzip it to `xampp/htdocs`
3. under XAMPP GUI, start Apache Server
4. Open browser and type `localhost/mdWorm`

### RHEL & CentOS

1. Install httpd

		sudo yum install httpd
	
2. Install PHP

		sudo yum install php

3. Download mdWorm & Unzip it to `/var/www/html`

		cd /var/www/html
		wget <mdWorm-release-tarball>
		tar -zxvf <mdWorm-release-tarball>
		
4. Start Apache Server

		service httpd start
		
5. Open browser and type

		redHatServerIP/mdWorm
	
6. (optional) It will be better if you also install vsftpd for markdown upload
		
### Debian & Ubuntu
1. Install apache2

		sudo apt-get install apache2

2. Install PHP

		sudo apt-get install php5 libapache2-mod-php5

3. Download mdWorm & and unzip to`/var/www/html`

		cd /var/www/html
		wget <mdWorm-release-tarball>
		tar -zxvf <mdWorm-release-tarball>

4. Start Apache Server

		service apache2 start

5. Open browser and type

		localhost/mdWorm
	
		
### MIT License

===
Beichen Li 2016-8-29	