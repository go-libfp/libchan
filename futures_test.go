package libchan

import "testing"

func TestPromise(t *testing.T) {
	p := Promise[string]()
	s := "hello friend"
	Resolve(p, s)
	if Await(p) != s {
		t.Fail()
	}
}

func TestFutureCombinators(t *testing.T) {
	f := Future("hello")

	f1 := Map(f, func(x string) string {
		return x + " world"
	})

	f2 := Bind(f1, func(x string) chan string {
		return Future(x)
	})

	if Await(f2) != "hello world" {
		t.Fail()
	}

}
