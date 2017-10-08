package main

func request(reg string, inm <-chan uint64, inml <-chan uint64,
	outm chan<- uint64, outml chan<- uint64) {
	const MODEMATCH = 1
	const MODEMATCHLIST = 2

	//var sem sync.Mutex

	select {
	//case m := <-inm:
	// Call API to retrieve match
	}
}
