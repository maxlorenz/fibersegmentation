package main

func main() {
	a := []int{1, 2, 3, 4, 5}
	for _, v := range a[-1:0] {
		print(v)
	}
}