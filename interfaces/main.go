package main

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	sq := square{5}
	tr := triangle{3, 5}

	printArea(sq)
	printArea(tr)
}
