package main

func main() {
	x, y := 15, 10
	sum, diff := sumAndDiff(x, y)
	print("SUM : ", sum, " DIFF : ", diff)
}

func sumAndDiff(x int, y int) (int, int) {
	return x + y, x - y
}
