package libchan 


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


func Future[T any](x T) chan T {
	c := make(chan T, 1)
	c <- x 
	return c
} 

func Promise[T any]() chan T {
	return make(chan T, 1)
}

func Await[T any](c chan T) T {
	x := <- c
	return x 
} 

func Resolve[T any](c chan T, t T) {
	c <- t 
	close(c)
}



// Spawns a function asynchronously and returns a promise for completion
func Spawn[T any](f func() T) chan T {
	outCh := make(chan T, 1)

	go func() {
		outCh <- f()
		close(outCh)
	}()

	return outCh 
}



func Bind[T, U any](c chan T, f func(T) chan U) chan U {
	x := <- c
	return f(x)
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


