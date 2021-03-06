# yikesbot (Abandoned because of IP-Ban?)

**yikesbot** is a Destiny.gg chat bot that has an integrated "yikes" counter that increases every time it detects a variation of the word "yikes" in a chat message. It is heavily inspired by the [RandomFerret](https://github.com/voloshink/FerretBot) bot made by **Polecat** in chat. Most of his existing code was used in creating this bot.

## How does it work

yikesbot listenes for a defined variation of the word "yikes" in a chat message and increases its internal counter by 10. The internal counter decreases by 1 every 2 seconds until it reaches 0.

Expected variations of "yikes" the bot currently processes:


- Yikes
- Y I K E S
- Yikerz
- Yikers

Capitalization is not important. Any of the above words can be spelled as you wish.

### New since Version 1.5

- omegayikes

Typing omegayikes increases the current amount by 50 but has a cooldown of 60 seconds before it can be used again.

### New since Version 1.7.1

- !graphyikes links to a graphical overview that displays the yikes amount over time.

## Commands

## **UPDATE**: Yikesbot got banned and is now only operating in stealth mode... No more !yikes

Currently everyone in chat can get the current yikes count by typing the following commands:

- !yikes
- !YIKES
- ! yikes
- ! YIKES
- !yikers
- !YIKERS

Cooldown of the command is currently set to 10 seconds

*Note that the command just needs to start with one of the above words. You can append any word behind a command: i.e. !yikes**BOI***

## Admin Commands

Currently theneedyguy and LeEpicMemeXd are admins of the bot and can thus execute the following commands:

| Command         | Effect                                                                  |
| ----------      |:----------------------------------------------------------------------- |
| **!ver**        | Displays current version of yikesbot                                    |
| **!reset**      | Resets the yikes counter                                                |
| **!sleep**      | Toggles Sleep Mode (Bot stops responding to !yikes)                     |
| **!ping**       | Sends back a pong (Check if bot is alive)                               |
| **!ipban**      | Increases the counter by 100 (Triggered only when Destiny bans someone) |
| **!topyikes**   | Shows the highest yikes level ever achieved since starting the bot      |

## Metrics

Since version 1.7 yikesbot exposes Prometheus time series data on port 9998. This is only relevant for people that know what Prometheus is and what it is used for.
Metric label:

- dgg_yikes_level

## Changelog

[Changelogs are here](https://github.com/theneedyguy/yikesbot/blob/master/CHANGELOG.md)

## Feature plans

- Some kind of "buy" function for yikes. Will require a database or just a local json that gets updated. We'll see what I can come up with.
The buy function would come with a cooldown. The first person to buy yikes would get them. A leaderboard would be a nice meme.
