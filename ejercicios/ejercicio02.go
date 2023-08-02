package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PedirNumero() {
	var num int
	var err error
	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número : ")
		if sc.Scan() {
			num, err = strconv.Atoi(sc.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", num, i, i*num)
	}
}
