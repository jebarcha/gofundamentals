# go lang 

## Questions
1. How do we run the code in our project?

```
go build main.go -> compiles go source code
go run main.go -> compiles and execute one or two files
go fmt -> formats all the code
go install -> compiles and installs a package
go get -> downloads the raw source code of someone else's package
go test -> Runs any tests associated with the current project
```

2. What does 'package main' mean?
When we have package main it means we are creating an executable package
any other name whatsoever means we are making a reusable or dependency type package.
The executable package always Must have a func called **'main'**


3. What does 'import "fmt"' mean?
Is used to give our package access to some code that is written inside of another package.
fmt stands for the word: **format**

There are other packages like these standard libraries:  main, fmt, debug, math, encoding, crypto, io
reusable packages examples: calculator, uploader, 

reference: https://pkg.go.dev/std


4. What's that 'func' thing?

func is a short for function


5. How is the main.go file organized?

package main -> import "fmt" -> func main() { fmt.Println("hi there!")}
