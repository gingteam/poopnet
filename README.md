# poopnet
Usage:
- You need 2 port: 12345 and 54321 opened.
- Using raw TCP connection to connect to the CNC, you can use nc command or use Putty with RAW method
- CNC port: 12345, Bot server port: 54321
- Demo video: https://streamable.com/53y3sb
- I just made it last night so It only has the remote shell execution feature, I'll try to add more functions soon XD
- Change the SERVER_IP and SERVER_PORT in client/configs.go and SERVER_IP, SERVER_BOT, BOT_PORT in server/configs.go then build the source.
