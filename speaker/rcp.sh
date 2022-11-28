#!/bin/sh

# rcp.sh version 4.0.0

[ -f data/rcp-poll-msg ] || {
	channel=$rcpchan mkthread $((7 * 24 * 60)) \
		"RCP $(cat data/rcp) - $(cat data/rcp-title)" > data/rcp-thread
	{
		echo "**Official Vote Thread**"
		echo "RCP $(cat data/rcp) - $(cat data/rcp-title)"
		channel=$(cat data/rcp-draftchan) readmsg < data/rcp-messages
		echo "----------------------------------------"
		echo "$legislators"
		echo "Please react to this message with ✅ for a \"yes\" vote or ❌" \
			"for a \"no\" vote."
	} | channel=$(cat data/rcp-thread) msg > data/rcp-poll-msg &&
		rm data/rcp-title data/rcp-draftchan data/rcp-messages
}

closercp() {
	channel=$announcements msg "$legislators, enough results have come in to" \
		"close the poll!  If nobody retracts their votes in the next 15" \
		"minutes, the results will be tallied and the RCP will either be" \
		"enacted or discarded.  If you are unsure about your vote, be sure" \
		"to retract it within the next 15 minutes so the polls stay open."

	sleep $((15 * 60))

	y=$(countemoji $msg ✅)
	n=$(countemoji $msg ❌)

	export channel=$announcements

	if $unanimous
	then
		if [ $n -gt 0 ]
		then
			msg "$legislators, the vote was vetoed by one of the players." \
				"RCP $(cat data/rcp) will not go into effect."
		elif [ $y -ge $(ls data/players | wc -l) ]
		then
			msg "$legislators, the vote has passed unanimously.  RCP $(cat \
				data/rcp) will go into effect immediately."
		else
			msg "$legislators, votes were retracted, so the polls will" \
				"remain open."
			return
		fi
	else
		if [ $(($y + $n)) -ge $(ls data/players | wc -l) ]
		then
			if [ $y -gt $n ]
			then
				msg "$legislators, the vote has passed at a vote of $y-$n." \
					"RCP $(cat data/rcp) will go into effect immediately."
			elif [ $n -gt $y ]
			then
				msg "$legislators, the vote has failed at a vote of $y-$n." \
					"RCP $(cat data/rcp) will not go into effect."
			else
				msg "$legislators, the vote has tied at a vote of $y-$n." \
					"The polls will remain open in case a player wishes to" \
					"defect in order to break the tie. This is currently the" \
					"only way to resolve a tie until a rule is added that" \
					"handles ties."
				return
			fi
		else
			msg "$legislators, votes were retracted, so the polls will remain open."
			return
		fi
	fi
}

export channel=$(cat data/rcp-thread)
export msg=$(cat data/rcp-poll-msg)
export unanimous=false
[ $(cat data/rcp) -le 320 ] || [ -f data/unanimous ] && unanimous=true
while true
do
	emojicat 48h $msg | while read line
	do
		if $unanimous
		then
			( [ $(countemoji $msg ❌) -gt 0 ] ||
				[ $(countemoji $msg ✅) -ge $(ls data/players | wc -l) ]
			) && closercp
		else
			[ $(($(countemoji $msg ✅) + $(countemoji $msg ❌))) -ge 
				$(ls data/players | wc -l) ] && closercp
		fi
	done

	export channel=$(cat data/rcp-thread)

	[ -f data/rcp-poll-msg ] || break

	nonvoters="$players"
	for player in $(identify-voters $msg ✅| tr '\n' ' ') \
		$(identify-voters $msg ❌ | tr '\n' ' ')
	do
		nonvoters="$(printf "$nonvoters" |
			tr ' ' '\n' |
			grep -v $player |
			tr '\n' ' ')"
	done

	msg $announcements "The following players have not yet voted: \n" \
	"$($mention $nonvoters).\n\n" \
	"If you are one of these players, it is requested that you vote on RCP" \
	"$(cat data/rcp) soon!"

done
