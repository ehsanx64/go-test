PROGRAM=hello-world
build-win64:
	GOOS=windows GOARCH=amd64 go build -o $(PROGRAM)-amd64.exe $(PROGRAM).go
build-win386:
	GOOS=windows GOARCH=386 go build -o $(PROGRAM)-i386.exe $(PROGRAM).go
clean:
	rm *.exe
	rm $(PROGRAM)
