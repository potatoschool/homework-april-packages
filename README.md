# Shapes

A simple console project, written in Go, to create and render shapes.
Shape renderer is built on-top of [Ebiten](https://github.com/hajimehoshi/ebiten/).

## Build and Launch

### Step 0: Add config.json
example config:
```json
{
  "mode": "DEBUG",
  "logFolder": "./_log"
}
```
Available modes: DEBUG, PRODUCTION. Default mode is Production.
Log folder is a path where user can store the logs. It's ./_log by default.
#### Mode = DEBUG
tells the program to run in debug mode and print messages to a console.
#### Mode = PRODUCTION
tells the program to run in production mode and create only log files.


### Step 1: Build
Run `make build` or `go build -o run main.go`

This will build the `run` executable to launch the project.
### Step 2: Run
Execute `./run`

You will see an output in a console shortly.

## Usage
The programm will prompt you to enter the shape name

![Startup](https://github.com/potatoschool/shapes/blob/main/assets/startup.png?raw=true)

### Available shapes by default:
- rectangle
- circle

### Available commands by default:
- --list - Will print all available shapes with their aliases in console