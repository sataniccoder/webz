# webz
webz is a simple website runner i built in golang, it's desinged to be small and fast handling mulitple client's whilst you just have to run the program    
it's desing'd to be as simple as possible and as small as possible

## Install
to install webz you can either go to the realses or download and compile from source   
###### Source

```
git clone https://github.com/sataniccoder/webz.git
cd webz
go build webz.go
```
once that's done you will have the compiled binary for webz

### Feturs
1) simple file hosting allowing other to view or download them with a simple command-line argument
2) allows for live edit's so once you have set webz up with the options you wan't you can leave it alonw
3) auto-update webz can auto-update it's self and restart the program with your curren't config (if nothings changed)

#### Modes
webz has two modes, the defualt mode allows users to view primalry html files, this mode requires there be a index.hmlt in the directory you have selecte or where you run the program from    
the second mode is called 'list' mode wich allows client's to view files and read them in the browser and allowing users to download them onto there mechined


#### Next Update
the next update will come with the 'list' mode added along with the '--silent' commadn line argument to work   
rember webz is ment to be small and fast so as time goes on and if the project expands more it might need two dirrent versions, one 'slim' and another 'expaned' version   
