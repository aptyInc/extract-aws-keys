
#!/bin/bash

set -e

env GOOS=darwin go build -o ./dist/main_darwin  main.go
env GOOS=linux go build -o ./dist/main_linux  main.go
env GOOS=windows go build -o ./dist/main_windows.exe  main.go