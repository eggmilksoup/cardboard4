#!/bin/sh

# unusual.sh version 3.0.1

curl -s https://en.wikipedia.org/wiki/Wikipedia:Unusual_articles > data/unusual.html

title=$(
	grep '<td><b><a href="' data/unusual.html |
		sed 's/<[^>]*>//g' |
		head -n $(
			/home/egg/bin/rnum $(
				grep '<td><b><a href="' data/unusual.html | wc -l
			)
		) |
		tail -n 1
)

text=$(
	head -n $(($( grep -n "title=\"$title\">$title" data/unusual.html | cut -f 1 -d :) + 2)) data/unusual.html |
		tail -n 1 |
		sed 's/<[^>]*>//g'
)

echo $title: $text
