###
echo "Clean distribution..."


###
echo "Build application for default OS..."
go build -o dist/default/hello-world-main main.go 
echo "End.Build" 

### Run default application
echo "Running application:"
echo "-------------------"
echo ""
./dist/default/hello-world-main