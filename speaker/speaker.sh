#!/bin/sh

# speaker.sh version 4.0.0

export avian=$(cat data/players/avian)
export cardinal=$(cat data/players/cardinal)
export cshainyr=$(cat data/players/cshainyr)
export ghost=$(cat data/players/ghost)
export hieratum=$(cat data/players/hieratum)
export maymun=$(cat data/players/maymun)
export olrora=$(cat data/players/olrora)
export qitiano=$(cat data/players/qitiano)
export x_and=$(cat data/players/x-and)
export yngling=$(cat data/players/yngling)

export players="$avian $cardinal $cshainyr $ghost $hieratum $maymun $olrora $qitiano $x_and $yngling"

export PATH=`pwd`/bin:$PATH
export key=$(cat data/speaker-key)
export testchan=$(cat data/channels/testchan)
export rcpchan=$(cat data/channels/rcpchan)
export announcements=$(cat data/channels/announcements)
export election=$(cat data/channels/election)
export legislators="<@&1001702350571974716>"
export mention=mention

## debug overrides ##
export announcements=$testchan
export rcpchan=$testchan
export election=$testchan
export legislators="@Eldritch Legislators"
export mention=username


while true
do
	sh speaker/$(cat data/phase).sh
	sh speaker/nextphase.sh
done
