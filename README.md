# phrase-counter-go

Language: Golang-1.18

## Compile and run
```
# Compile the code
go build

# accepts as arguments a list of one or more file paths
# Linux/Mac
./counter pg2009.txt pg2010.txt pg2011.txt

# Windows 
.\counter.exe .\pg2009.txt .\pg2010.txt .\pg2011.txt


# input on stdin
# Linux/Mac
cat pg2009.txt | ./counter

# Windows 
cat .\pg2009.txt | .\counter.exe
```

## Run the code directly without compiling
```
# accepts as arguments a list of one or more file paths
# Linux/Mac
go run main.go pg2009.txt pg2010.txt pg2011.txt

# Windows 
go run .\main.go .\pg2009.txt .\pg2010.txt .\pg2011.txt


# input on stdin
# Linux/Mac
cat pg2009.txt | go run main.go

# Windows 
cat .\pg2009.txt | go run .\main.go
```