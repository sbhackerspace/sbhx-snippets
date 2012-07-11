package main

/* Note that is not considered idiomatic Go to import curses this way */
import . "code.google.com/p/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()
    
    stdscr.Print("Hello, World!!!")
    stdscr.Refresh()
    stdscr.GetChar()
}