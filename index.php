<?php
/**
* mdWorm
* v1.0.1
* A very simple markdown reader in PHP
*
* ./index.php
* List files in ./md
*
* @author bcli
* @since  v1.0
*/

// set default timezone, if the timezone settings got by 'date_default_timezone_get'
// is not correct, replace it with a timezone string ex: 'Asian/Shanghai'
// see http://www.php.net/manual/en/timezones.php for avaliable timezones
date_default_timezone_set(date_default_timezone_get());

// get page head
$head       = file_get_contents("res/html/head.html");
// get page tail
$tail       = file_get_contents("res/html/tail.html");
// store files from dir: ./md/ to array $files_raw
$files_raw 	= array();
$it         = new DirectoryIterator(dirname(__FILE__)."/md");
$it->rewind();
while ($it->valid())
{
  $files_raw[$it->getFilename()] = $it->getMTime();
  $it->next();
}
// order files by modification time DESC
arsort($files_raw);

// Get title and author
// you should write your markdown file in such fashion
// 1. First line should be title and started with '#'
// 2. Second or Third line should be author ex: 'by ***'
// For Example
// line 1: # <My md title>#
// line 2:
// line 3: by <my name>
// line 4: ---
// line 5: <md content>

//=======
// Script will stop trying to fetch title & author after first N lines
// default value is 5
$MAX_LINE = 5;
//=======
$files      = array();
$i          = 0;
$noData			= true;
foreach ($files_raw as $filename => $modi_time)
{
  // ignore dir & hidden files
  if (substr($filename, 0, 1) === '.')
  {
    continue;
  }
	// otherwise check file content
	else
	{
		$noData = false;
		$gotTitle   = false;
		$gotAuthor  = false;
		$title      = "Unkown";
		$author     = "unkown";
		// get file path
		$filePath = dirname(__FILE__)."/md/".$filename;
		// get file object
		$f = fopen($filePath, 'r');
		// check first N lines to get markdown file & author
		for ($j = 1;$j <= $MAX_LINE;$j++)
		{
			$line= fgets($f);
			if (substr(trim($line),0,1)=="#" && $gotTitle == false)
			{
				$gotTitle = true;
				$title = trim(str_replace("#","",$line));
			}
			if (substr(trim($line),0,2)=="by" && $gotAuthor == false)
			{
				$gotAuthor = true;
				$author = trim(substr(trim($line),2,15));
			}
			if ($gotTitle && $gotAuthor)
			{
				break;
			}
		}
		// close file
		fclose($f);
		// add file data to array
		$files[$i] = array(
											"name" 		=> $filename,
											"title" 	=> $title,
											"author" 	=> $author,
											"time" 		=> date('Y-m-d H:i:s',$modi_time)
										);
		$i++;
	}
}
// echo page head
echo $head;
// echo title
echo "<h1>MD Worm</h1>\n";
echo "<p>Please set md file permission to 644</p>\n";
echo "<hr>\n";
// echo markdown file list
echo "<ul>";
foreach ($files as $data)
{
    echo "<li>";
    echo $data["time"];
    echo "&nbsp;&nbsp;";
    echo "<a href='get.php?md=".$data["name"]."' title='".$data["name"]."'>";
    echo $data["title"];
    echo "</a>";
    echo "<a style='float:right'>";
    echo "by&nbsp;";
    echo $data["author"];
    echo "</a>";
    echo "</li>\n";
}
echo "</ul>\n";
if ($noData)
{
	echo "<p>No data under ./md/</p>";
}
echo "<hr>\n";
// echo copyright
echo "<p>mdWorm v1.0.1 @ bcli 2016-9-5 <a href='https://github.com/bclicn/mdWorm'>GitHub</a></p>";
// echo page tail
echo $tail;
?>
