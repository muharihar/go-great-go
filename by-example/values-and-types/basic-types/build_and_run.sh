###
echo "Clean distribution..."


###
echo "Build application for default OS..."
go build -o dist/default/<...> main.go 
echo "End.Build" 

###
echo "Build application for Linux OS..."
GOOS=linux GOARCH=amd64 go build -o dist/linux/<...> main.go 
echo "End.Build" 

###
echo "Build application for Mac OS..."
GOOS=darwin GOARCH=amd64 go build -o dist/mac/<...> main.go 
echo "End.Build" 

###
echo "Build application for Windows (64bit) OS..."
GOOS=windows GOARCH=amd64 go build -o dist/win/<...>.exe main.go 
echo "End.Build" 

### Run default application
echo "Running application:"
echo "-------------------"
echo ""
./dist/default/<...>