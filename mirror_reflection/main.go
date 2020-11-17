package main

func mirrorReflection(p int, q int) int {
    m, n := 1, 1
    for m*p != n*q {
        n++
        m = q*n/p
        // fmt.Println("m ", m, " n ", n)
    }
    
    if n%2 == 0 && m%2 == 1 {return 2}
    if n%2 == 1 && m%2 == 0 {return 0}
    if n%2 == 1 && m%2 == 1 {return 1}
    return 1
}