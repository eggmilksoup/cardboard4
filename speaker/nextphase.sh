#!/bin/sh

# nextphase.sh version 4.0.0

case $(cat data/phase) in
	listen)
		echo rcp > data/phase
		;;
	rcp)
		if [ $(($(cat data/rcp) - $(cat data/topofcircuit))) -eq 
			$(ls data/players | wc -l) ]
		then
			echo flavor > data/phase
		else
			echo listen > data/phase
		fi
		;;
	flavor)
		echo vc > data/phase
		;;
	vc)
		echo listen > data/phase
		echo $((cat data/rcp - 1)) > data/topofcircuit
		;;
esac
