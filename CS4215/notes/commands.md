# GO Lang Build Commands
Useful commands for working with the go repository.

Consider ./ in this context to be $GOROOT.



## Project Layout
### GO Compiler
After compiling the compiler the new compiler binary will be at:
./bin/go

### Standard Lib
The standard lib files are at:
./src/{package}/{lib}
./src/crypto/sha1


## Commands
Running a build all and test all.
: ./all.bash

Make only and don't test.
: ./make.bash

Test only
: ./run.bash

Test a package
: ./bin/go test ./src/crypto/sha1 

Recompile the compiler only
: cd src
: ./bin/go install cmd/compile

