###
echo "Clean distribution..."


###
echo "Build application for default OS..."
go build -o dist/default/hello-world-concurrent main.go 
echo "End.Build" 

###
echo "Build application for Linux OS..."
GOOS=linux GOARCH=amd64 go build -o dist/linux/hello-world-concurrent main.go 
echo "End.Build" 

###
echo "Build application for Mac OS..."
GOOS=darwin GOARCH=amd64 go build -o dist/mac/hello-world-concurrent main.go 
echo "End.Build" 

###
echo "Build application for Windows (64bit) OS..."
GOOS=windows GOARCH=amd64 go build -o dist/win/hello-world-concurrent.exe main.go 
echo "End.Build" 

### Run default application
echo "Running application:"
echo "-------------------"
echo ""
./dist/default/hello-world-concurrent