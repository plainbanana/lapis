# lapis
This is a Mirakurun version Golang port of [Antennas](https://github.com/TheJF/antennas) .
An API bridge server whitch translate Mirakurun API to HDhomerun API.
This enables to use [PlexDVR feature](https://www.plex.tv/ja/live-tv-dvr/) .

## How to use
- edit `.env` file and set your configuration.
- build and run
- log in Plex with account whitch has **enabled PLEXPASS** and set up `Live TV & DVR`
- Have fun!!

## Known issues (?)
- Live TV : GR only
- Program guide : provided by Plex (GR only)
- Recording : It may works but can not watch at web console. It seems that CM removal does not works
