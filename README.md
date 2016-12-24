# So what's Go really good for?

Besides servers and tcp/http clients, Go is also amazing for creating command line applications.
Executable files are generated with great speed and elegance.
So *mchang13* and I decided to make a cli app that will have its own bash command.


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




*s/o to https://github.com/urfave/cli for providing a nice API for cli building*