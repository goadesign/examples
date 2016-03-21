# adder

This is the example used in the README. Note that apart from [design/design.go](design/design.go)
the only line of code that is *not* generated is the line that does the addition and writes the
response in [operands.go](operands.go). Everything else is generated with:
```
goagen bootstrap -d github.com/goadesign/examples/adder/design
```

## Running

To run the example simply compile and run the binary:
```
git clone https://github.com/goadesign/examples
cd examples/adder
go build
./adder
```
Compile and run the client in a different terminal:
```
cd client/adder-cli
go build
./adder-cli add operands /add/1/2
2016/03/21 00:33:10 [INFO] started id=nclom9xa GET=http://localhost:8080/add/1/2
2016/03/21 00:33:10 [INFO] completed id=nclom9xa status=200 time=20.102916ms
3
```
