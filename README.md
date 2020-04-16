# vrb
Convenient verbose output library.

Inspired by paulvollmer/go-verbose, but trying to keep an usage pattern similar to fmt.

## How to use

Initialize the package once at the start of your command:

```
import "github.com/telegrapher/vrb"

...

// Enable the verbose messages if verboseBoolean == true
// Write the verbose messages to os.Stderr
vrb.Init(verboseBoolean, os.Stderr)

```

Use "vrb" the same way you would use "fmt":


```

fmt.Println("This text goes to STDOUT")
vrb.Println("This is verbose-only text and goes to STDERR")

```

