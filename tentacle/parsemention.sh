#!/bin/sh

# parsemention 4.0.0

logmentions $1 | while read line
do
	rule=false
	rcp=false
	for word in $(printf "$line" | cut -f 3 -d :)
	do
		case $word in 
			rcp)
				rule=false
				rcp=true
			;;
			rule)
				rcp=false
				rule=true
			;;
			*)
				if $rcp
				then
					sed 's/  / /' /var/git/tns-nomic-records/rcp/$word |
						tr '\n' ' ' |
						sed 's/  /\
\
/g' |
						sed 's/[^^]	//g' | msg && continue
					msg no such rcp \"$word\"
				elif $rule
				then
					if [ "$word" = "list" ]
					then
						scripts/rulelist.sh | msg
						rule=false
					fi
					scripts/findrule.sh $word | msg && continue
					msg no such rule \"$word\"
				fi
			;;
		esac
	done
done
