# lapis
This is a Mirakurun version Golang port of [Antennas](https://github.com/TheJF/antennas) .
An API bridge server whitch translate Mirakurun API to HDhomerun API.
This enables to use [PlexDVR feature](https://www.plex.tv/ja/live-tv-dvr/) .

## How to use
- edit `.env` file and set your configuration.
- build and run. `docker-compose build && docker-compose up -d`
- log in Plex with account whitch has **enabled PLEXPASS** and set up `Live TV & DVR`
- Have fun!!

## Known issues (?)
- Live TV : guide from Gracenote : GR only but very convenient. guide from `/epg.xml` : GR/BS/CS available.
- Program guide : provided by Plex (GR only). `/epg.xml` is not fully functional EPG, but GR/BS/CS available.
- Recording : It works.