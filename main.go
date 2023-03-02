package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	random := rand.NewSource(time.Now().UnixMilli())
	rnd := rand.New(random)
	randomInt := rnd.Intn(100)
    
    var user int
    var input string

    fmt.Println("Main Tebak Angka Yuk!!")
    fmt.Println("Tebak angka antara 1-100")
    
    for {
        fmt.Println("Masukkan Angka: ")
        fmt.Scanln(&input)

        user, _ = strconv.Atoi(input) // konversi input pengguna dari string ke int

        if input == "exit" || input == "999" { // Jika pengguna memasukkan "exit" atau "999"  keluar dari program
            fmt.Println("Yaa Udahan ya ")
            break
            
        }
        if  user < 1 || user > 100 {
            fmt.Println("Harus Tebak 1 - 100 Dong!!")
            continue
        }

        if user < randomInt {
            fmt.Println("Terlalu Rendah, Coba Tebak Lagi")
        } else if user > randomInt {
            fmt.Println("Terlalu Tinggi, Coba Tebak Lagi")
        } else{
            fmt.Printf("Hebat.. Benar Sekali, angkanya Yaitu : %d \n", randomInt,  )
            break
        }
    }
}