package stream
import (
//	"runtime"
//	"sync"
)




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
		defer close(out)
		for i := 0; i < n; i++ {
			x, ok := <- c
			if !ok {
				return 
			} 
			out <- x
		}

		
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


func Filter[T any](c chan T, f func(T) bool) chan T{
	out := make(chan T)
	go func() {
		for x := range c {
			if f(x) {
				out <- x
			}
		}
		close(out)
	}()

	return out 
}



func Reduce[T, U any](c chan T, init U, f func(acc U, x T) U ) U {
	out := init

	for x := range c {
		out = f(out, x)
	}

	return out 

} 
