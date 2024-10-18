package main
func main() {
gotoLABEL2 := false
gotoLABEL := false
x := true
y := false
for i := 0; i < 10; i++ {
    if x {
        if y {
            x = false
            if i == 6 {
                gotoLABEL = true
            } else if true {
                gotoLABEL2 = true
            } else {
                continue
            }
            if !gotoLABEL {
                if gotoLABEL2 {
                    goto LABEL2
                }
            }
            if !gotoLABEL2 {
                if gotoLABEL2 {
                    goto LABEL2
                }
            }
        } else {
            y = true
            continue
        }
        if !gotoLABEL2 {
            if gotoLABEL {
                goto LABEL
            }
        }
    }
    if gotoLABEL2 {
        break
    }
    println(i)
}
if !gotoLABEL2 {
}
}