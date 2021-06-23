# Cobra package 

 Cobra package generates all boilerplate for create console commands.
```
https://github.com/spf13/cobra
```

 Once cobra is installed, we need to initialize it doing:
```
     cobra init --pkg-name=github.com/holdedhub/basics
```

 Add one cobra command (in our example is cli command):
 ```
cobra add cli
```

 After that step we can give it a try:
 ```
go run main.go --help

A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  basics [command]

Available Commands:
  cli         A brief description of your command
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.basics.yaml)
  -h, --help            help for basics
  -t, --toggle          Help message for toggle

Use "basics [command] --help" for more information about a command.

```
From Scratch cobra does a few things, take a look on root.go
- On the bootstrap process defines a config loader using Viper package, adds required flags to change config path.
```
https://github.com/spf13/viper
```

Digging a little bit more, take a look on cmd/cli.go and cmd/root.go, you will see how commands defined to allow that "help" info.
 
As a result, we can invoke our cli command doing:
```
go run main.go cli
```

# Conclusion
 Cobra speeds up all command development, allowing easy nested commands as:
 ```
main cli foo bar
```
Root command is able to define common flags to all commands

