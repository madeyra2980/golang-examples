package main

import "fmt"

func main() {

    // Бүтін санның дерек типі
    var num float32 
    num = 12;

    // Қалдық санның дерек типі
    var celcies float32
    celcies = 32.5

    fmt.Println(celcies + num) 

    // Логикалық деректің типі
    var isAlive bool;

    var res bool;
    res = celcies > num

    isAlive = res; 

    fmt.Println(isAlive)

}



