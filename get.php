<?php
/**
* mdWorm
* v1.0
* A very simple markdown reader in PHP
*
* ./get.php
* Read & Parse markdown file
*
* This script includes Parsedown.php, which is created by Emanuil Rusev at http://parsedown.org
* @author bcli
* @since  v1.0
*/

// include markdown parser class
include("Parsedown.php");
// get user request
$filename = $_GET["md"];
// assemble file path
$filepath = "md/".$filename;
// check if file exists
$markdown2html = "";
if (!file_exists($filepath))
{
  $markdown2html = "File not exists";
}
else
{
  $content = file_get_contents($filepath);
  if ($content == false)
  {
    $markdown2html = "Cannot read file, did you set file permission to 644 ?";
  }
  else
  {
		// parse markdown file
    $Parsedown = new Parsedown();
    $markdown2html = $Parsedown->text($content);
  }
}
// define page head
$head     = file_get_contents("res/html/head.html");
// define page tail
$tail     = file_get_contents("res/html/tail.html");
// echo page head
echo $head;
// echo parsed markdown content
echo $markdown2html;
// echo page tail
echo $tail;
?>
