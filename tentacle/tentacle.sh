#!/bin/sh

# tentacle.sh version 4.0.0

export PATH=`pwd`/bin:$PATH
export key=$(cat data/tentacle-key)
export channel=$(cat data/channels/agora)

nickname "Eldritch Tentacles 4.0"

while true
do
	msg Collecting eldritch data...
	tentacle/unusual.sh | msg
	tentacle/parsemention.sh $(($(rnum 65) + 6))h$(rnum 60)m
done
