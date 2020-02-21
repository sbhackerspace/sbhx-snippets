// Steve Phillips / elimisteve
// 2020.02.21

// Inspired by this great Python concurrency talk by David Beazley:
// https://youtu.be/MCs5OvhV9S4

module main

import net

fn main() {
	port := 25000
	println("Listening on port $port")
	ln := net.listen(port)?

	for {
		conn := ln.accept() or {
			eprintln('Error accepting: $err')
			continue
		}

		// In V, to launch a function into a goroutine (cheap
		// thread), just put the 'go' keyword in front of it!

		fib_handler(conn) // <-- blocking version
		// go fib_handler(conn) // <-- non-blocking version(!)
	}
}

fn fib_handler(conn net.Socket) {
	defer { conn.close() or {} }

	for {
		num := conn.read_line().trim_right('\r\n').int()
		if num == 0 {
			return
		}

		answer := fib(num)
		conn.send_string('$answer\n') or {
			eprintln('Error sending response: $err')
			return
		}
	}
}

fn fib(num int) int {
	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}
