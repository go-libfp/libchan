package broadcast 
import (
	"sync"
)


type Broadcast[T any] struct {
	in chan T
	receivers map[chan T]interface{} 
	lock sync.RWMutex 
}  



func (b *Broadcast[T]) broadcast(msg T) {
	defer b.lock.RUnlock() 
	b.lock.RLock() 

	for c, _ := range b.receivers {
		c <- msg
	}

}  


func (b *Broadcast[T]) run() {
	for x := range b.in {
		b.broadcast(x)
	}
}

func (b *Broadcast[T]) Submit(msg T)  {
	b.in <- msg
}


func (b *Broadcast[T]) Subscribe() chan T {
	c := make(chan T, cap(b.in))
	b.lock.Lock()
	b.receivers[c] = nil
	b.lock.Unlock()
	return c
} 


func (b *Broadcast[T]) Remove(c chan T) {
	b.lock.Lock()
	delete(b.receivers, c)
	b.lock.Unlock()
} 

func (b *Broadcast[T]) Close() {
	defer b.lock.Unlock()
	b.lock.Lock()
	close(b.in)
}



func New[T any](s int) Broadcast[T] {
	c := make(chan T, s)
	return FromCh(c)
}  


func FromCh[T any](ch chan T) Broadcast[T] {
 	var lock sync.RWMutex 
	rcvrs := make(map[chan T]interface{})
	b := Broadcast[T]{ch, rcvrs, lock}
	go b.run() 
	return b
}


