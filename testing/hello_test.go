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
	result := regex.Hello("Mike")
	if result != "Hello Mike!" {
		t.Errorf("Hello(\"\") Failed expected %v got %v", "Hello Dude", result)
	}

}
