package fooer

import "testing"

// normal test case
func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Expected Foo, got %s", result)
	}
}

// table driven test cases
func TestFooerTable(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"Foo", 3, "Foo"},
		{"Not Foo", 4, "Not Foo"},
		{"Foo", 6, "Foo"},
		{"Not Foo", 7, "7"},
		{"Foo", 9, "Foo"},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Fooer(test.input)
			if result != test.want {
				t.Errorf("Expected %s, got %s", test.want, result)
			}
		})
	}
}

func TestFooerParallel(t *testing.T) {
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(3)
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	})
	t.Run("Test 7 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(7)
		if result != "7" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "7")
		}
	})
}

func TestFooerSkiped(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func FuzzFooer(f *testing.F) {
	f.Add(3)
	f.Fuzz(func(t *testing.T, a int) {
		Fooer(a)
	})
}
