package main

func main() {
	for a := 0; a < 20000000; a++ {
		if (a & 0x8ffff) == 0x80000 {
			println(a)
		}
	}
}
