package main
func main() {
n := 2
m := 2
foo := true
for i := 0; i < n; i++ {
    println("I: ", i)
    for j := 0; j < m; j++ {
        switch foo {
        case true:
            println(foo)
            break OuterLoop
        case false:
            println(foo)
        }
    }
}
}