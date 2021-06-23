# Basic golang app

Just main function and go modules to define app dependencies.
Steps done:
- created main.go
- Initialize go modules (our package and dependencies)
```
go mod init github.com/holdedhub/basics
```
Take a look on go.mod, you will find the package definition

Compile the code doing:
```
go build
```
This will generate a binary file called basics (as the package has defined)

Ensure execution rights on the created binary (chomd +x basics), you can run it doing:
```
./basics
```

## Dependency inclusion
To add an external package to our app, we need to include our modules dependency, in the example we will use Spew package (var_dump kind of):
```
go get github.com/davecgh/go-spew/spew
```
Check go.mod, you will see spew dependency included. To use it on main.go:
```bigquery
import(
    "github.com/davecgh/go-spew/spew"
)
```
And then in any part of the code (main function in our example)
```
spew.Dump("fooo")
```

Run the app again:
```
go run main.go
```
You will see spew in action