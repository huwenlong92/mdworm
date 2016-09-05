# Code Syntax Highlighting

by bcli

---
## Intro ##

* Code syntax highlighted by [highlightJS](https://highlightjs.org)

* All text wrapped in `<pre><code>...</code></pre>` will be highlighted. 

* Code Syntax will be __detected automatically__, so you don't have to add Github code indicator like ````java`

* To change theme, modify `./res/html/head.html`. All available themes are under `./res/css` (except `mou.css`)


## JAVA ##

    public static void main (String[] args)
    {
        // some comment
        System.out.println("Hello World");
    }
    
## Python ##
    
    import scrapy
    
    from tutorial.items import DmozItem
    
    class DmozSpider(scrapy.Spider):
        name = "dmoz"
        allowed_domains = ["dmoz.org"]
        start_urls = [
            "http://www.dmoz.org/Computers/Programming/Languages/Python/Books/",
            "http://www.dmoz.org/Computers/Programming/Languages/Python/Resources/"
        ]
    
        def parse(self, response):
            for sel in response.xpath('//ul/li'):
                item = DmozItem()
                item['title'] = sel.xpath('a/text()').extract()
                item['link'] = sel.xpath('a/@href').extract()
                item['desc'] = sel.xpath('text()').extract()
                yield item

## GoLang ##

    import os
    
    func main() {
        i := 1
        fmt.Fprintf(os.Stdout, "Hello, the i is %d", i++)
    }


## NodeJS ##

    require('net').createServer(function(socket)
    {
        socket.on('data',function(data)
        (
            console.log(data);
        );
    }).listen(2333);

## PHP ##

    public function hello ($arg1 = 1, $arg2 = 2)
    {
        $arg1++;
        echo "Argument 1 is: ".$arg1;
    }

## Matlab ##

    a = [1 2 3 4 6 4 3 4 5]
    b = a + 2
    plot(b)
    grid on
    bar(b)
    xlabel('Sample #')
    ylabel('Pounds')
    plot(b,'*')
    axis([0 10 0 10])

## SQL ##

    SELECT * FROM my_table WHERE 'id' = 2 AND 'name' = 'Marry'

## BASH ##

    yum install httpd php -y
    service httpd start
    sudo rm -rf /

## Supported Languages ##

* Apache  
* Bash
* C#
* C++
* CSS 
* CoffeeScript
* Diff
* HTML
* XML 
* HTTP
* Ini
* JSON
* Java
* JavaScript
* Makefile
* Markdown
* Nginx
* Objective-C  
* PHP  
* Perl  
* Python 
* Ruby  
* SQL
* Clojure
* Arduino
* Lua
* Matlab
* R
* Scala
* Swift

---
bcli 2016-9-5
