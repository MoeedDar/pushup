package main

import _ "pushup/handlers"

func main() {
	defer close()
	poll()
}
