package main

import "testing"

func TestHello(t *testing.T) {
	//test with empty args
	withEmptyName := hello("")
	if withEmptyName != "Hello Dude!" {
		t.Errorf("hello(\"\") failed, expected=%v, got=%v", "Hello Dude!", withEmptyName)
	} else {
		t.Logf("hello(\"\") success, expected=%v, got=%v", "Hello Dude!", withEmptyName)
	}
	//test with some name
	testWithName := hello("Rajesh")
	if testWithName != "Hello Rajesh!" {
		t.Errorf("hello(\"Rajesh\") failed, expected=%v, got=%v", "Hello Rajesh!", testWithName)
	} else {
		t.Logf("hello(\"Rajesh\") success, expected=%v, got=%v", "Hello Rajesh!", testWithName)
	}
}
