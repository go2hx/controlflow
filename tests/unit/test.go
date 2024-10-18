package main

func main() {
	x := true
	y := false
LABEL:
LABEL2:
	for i := 0; i < 10; i++ {
		if x {
			if y {
				x = false
				if i == 6 {
					goto LABEL
				} else if true {
					goto LABEL2
				} else {
					continue
				}
			} else {
				y = true
				continue
			}
		}
		println(i)
	}
	// LABEL2:
}
