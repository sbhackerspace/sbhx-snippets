// Steve Phillips / elimisteve
// 2012.01.01

package main

func PingPonger(opponent chan string, msg string) {
	opponent <- msg
	println(msg)
}

func main() {
	toPinger := make(chan string)
	toPonger := make(chan string)

	for {
		pinger := PingPonger()
		ponger := func(opponentCh chan string, msg string) { opponentCh <- msg }()
	}
}