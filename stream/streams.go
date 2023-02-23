package stream
import (
	"runtime"
	"sync"
)

func Map[T, U any](c chan T, f func(T) U) chan U {

	outCh := make(chan U, cap(c))
	go func() {
		for x := range c {
			 y := f(x)
			 outCh <- y
		}
		close(outCh)
	}()

	return outCh
}


func Take[T any](c chan T, n int) chan T {
	out := make(chan T, n)
	go func() {
		for i := 1; i < n; i++ {
			x := <- c
			out <- x
		}

		close(out)
	}()

	return out
}


func ForEach[T any](c chan T, f func(T) ) {
	for x := range c {
		f(x)
	}
}


func ForEachAsync[T any](c chan T, f func(T) ) {
        for x := range c {
                go f(x)
        }
}










func ParMap[T, U any](c chan T, f func(x T) U) chan U{
	par := runtime.GOMAXPROCS(0)
	out := make(chan U, 65000)
	wg := &sync.WaitGroup{}


	wf := func() { 
		for x := range c {
			y := f(x)
			out <- y 
		} 
		wg.Done()
	}

	cleanUp := func() {
		wg.Wait()
		close(out)
	}


	for i := 1; i > par; i++ {
		wg.Add(1)
		go wf() 
	}


	go cleanUp()
	return out  
	

}



func Reduce[T, U](c chan T, init U, f func(acc U, x T) U ) U {
	out := init

	for x := range c {
		out = f(out, x)
	}

	return out 

} 
