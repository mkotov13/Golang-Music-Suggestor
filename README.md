# So what's Go really good for?

Besides servers and tcp/http clients, Go is also amazing for creating command line applications.
Executable files are generated with great speed and elegance.
So *mchang13* and I decided to make a cli app that will have its own bash command.

## First things First

The main file can be found at `src/mnm/mnm.go`


## The magic

After cloning the repo and setting up the Go environment on your machine, run
```
$ go install mnm
```
This should install our command (mnm stands for M&M: we haven't decided who's name goes first) into the `$GOPATH/bin` directory.

This way our app becomes available via
```
$ mnm
```
Magic, innit?


## Getting what you want

The app can take user input via cli flags. Please use the `-g`  flag to specify what genre of music you are interested in. For example, try
```
$ mnm -g rap
```
or
```
$ mnm -g house
```


*s/o to https://github.com/urfave/cli for providing a nice API for cli building*
