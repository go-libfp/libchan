package stream
import "testing"
import "fmt"
import "math/rand"
import "strings"


func Generator[T any](f func() T) chan T {
	c := make(chan T)

	rec := func() {
		if r := recover(); r != nil {
			return 
		}
	}

	go func() { 
		defer rec() 
		for {
			c <- f()
		}}()

	return c
} 

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}




func TestCombinators(t *testing.T) {

	f := func() string {return randStr(30)}
	mapFunc := func(s string) string {
		return fmt.Sprintf("%s:%s", s, s) 
	}


	toSlice :=  func(acc []string, x string) []string  {
		return append(acc,x) 
	}

	gen := Generator(f)
	

	genM := Map(gen, mapFunc)

	x := <- genM 

	if !strings.Contains(x, ":")  { 
		t.Fail()
	}


	genT := Take(genM, 5)
	r := Reduce(genT, []string{}, toSlice)



	if len(r) != 5 {
		t.Fail() 
	
	}
	
	genT1 := Take(genM, 25)
	genF := Filter(genT1, func(s string) bool {return strings.Contains(s, "a") })

	ForEach(genF, func(x string) {fmt.Println(x)} ) 


}   
