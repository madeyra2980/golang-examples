package main

import "fmt"

func main() {

// Динамикалық типтілеу
	name := "Madiyar"
	age := 22


	fmt.Println(name, age)
// Статикалық типтілеу

// Сөздік жол деректер типі
    var str string;
    str = "Madeyra"

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

    fmt.Println(isAlive, str)

}