## Install [cobra-cli](https://github.com/spf13/cobra/tree/master/cobra)
```bash
go install github.com/spf13/cobra-cli@latest
```

## Init cobra project
```bash
cd $GOPATH/src/github/
mkdir joker
cd joker
cobra-cli init

Your Cobra application is ready at
/Users/xxx/go/src/github/[GITHUB_USER_NAME]/joker

 tree .
.
├── LICENSE
├── README.md
├── cmd
│   └── root.go
└── main.go

1 directory, 4 files
```

## Init go mod
```bash
go mod init github.com/[GITHUB_USER_NAME]/joker

go: creating new go.mod: module github.com/[GITHUB_USER_NAME]/joker
go: to add module requirements and sums:
        go mod tidy


go mod tidy
go: finding module for package github.com/spf13/viper
go: finding module for package github.com/spf13/cobra
go: found github.com/spf13/cobra in github.com/spf13/cobra v1.2.1
go: found github.com/spf13/viper in github.com/spf13/viper v1.8.1
go: downloading github.com/smartystreets/goconvey v1.6.4
go: downloading github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d
go: downloading github.com/jtolds/gls v4.20.0+incompatible
go: downloading github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1
```

### Add new command by use `cobra add [command name]`
```bash
cobra add random

random created at /Users/neilguan/go/src/github.com/[GITHUB_USER_NAME]/joker

 tree .
.
├── LICENSE
├── README.md
├── cmd
│   ├── random.go   # new command random file.
│   └── root.go
├── go.mod
├── go.sum
└── main.go

1 directory, 7 files
```

## First take a look.
```bash
 go run main.go
A longer description that joker Cmd. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  joker [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  random      A brief description of your command

Flags:
      --config string   config file (default is $HOME/.joker.yaml)
  -h, --help            help for joker
  -t, --toggle          Help message for toggle

Use "joker [command] --help" for more information about a command.
```

## joke random demo.
```bash
 ./joker random
This is Joke Command: ...
How many kids with ADD does it take to change a lightbulb? Let's go ride bikes!


# Request 6 jokes at the one time.
./joker random --count 6
This is Joke Command: ...
What's the best thing about elevator jokes? They work on so many levels.
Guy told me today he did not know what cloning is. I told him, "that makes 2 of us."
What biscuit does a short person like? Shortbread. 
Why was Santa's little helper feeling depressed? Because he has low elf esteem.

Why was Pavlov's beard so soft?  Because he conditioned it.
For Valentine's day, I decided to get my wife some beads for an abacus.  It's the little things that count.
```
