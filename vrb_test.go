package vrb_test

import (
	"os"
	"testing"

	"github.com/telegrapher/vrb"
)

func Example() {
	vrb.Init(true, os.Stderr)
	vrb.Println("Verbose message in stderr.")
}

func Test_Verbose(t *testing.T) {

	vrb.Init(true, os.Stderr)

	n, err := vrb.Print("Test Print\n")
	if err != nil {
		t.Error("vrb.Print failed")
	}
	if n != 12 {
		t.Error("vrb.Print returned n not equal")
	}

	n, err = vrb.Printf("Test Printf integer %v\n", 1337)
	if err != nil {
		t.Error("vrb.Printf failed")
	}
	if n != 26 {
		t.Error("vrb.Printf returned n not equal")
	}

	n, err = vrb.Println("Test Println")
	if err != nil {
		t.Error("vrb.Println failed")
	}
	if n != 14 {
		t.Error("vrb.Println returned n not equal")
	}

}
