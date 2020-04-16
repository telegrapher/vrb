# vrb
Convenient verbose output library.

Inspired by paulvollmer/go-verbose, but trying to keep an usage pattern similar to fmt.

## How to use

Initialize the package once at the start of your command:

```go
import "github.com/telegrapher/vrb"

...

// Write the verbose messages to os.Stderr
// Without any initialization, they'll be discarded
if verbose {
    vrb.Init(os.Stderr)
}

```

Use "vrb" the same way you would use "fmt":


```go

fmt.Println("This text goes to STDOUT")
vrb.Println("This is verbose-only text and goes to STDERR")

```

