package future




func Ok[T any](x T) chan T {
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




func Map[T, U any](c chan T, f func(x T) U)  chan U {
	p := Promise[U]()
	go func() {
		x := <- c 
		y := f(x)
		Resolve(p, y)

	}()
	return p 
}
