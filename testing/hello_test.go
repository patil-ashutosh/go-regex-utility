package testing

import (
	"testing"

	"github.com/patil-ashutosh/RegexUtility/regex"
)

func TestHello(t *testing.T) {

	emptyResult := regex.Hello("")
	if emptyResult != "Hello Dude" {
		t.Errorf("Hello(\"\") Failed expected %v got %v", "Hello Dude", emptyResult)
	}

}

func TestHelloValidArgs(t *testing.T) {

	emptyResult := regex.Hello("Mike")
	if emptyResult != "Hello Mike!" {
		t.Errorf("Hello(\"\") Failed expected %v got %v", "Hello Dude", emptyResult)
	}

}
