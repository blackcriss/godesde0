package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TabladeMultiplicar() string {
	var num int
	var err error
	var texto string
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
		texto += fmt.Sprintf("%d x %d = %d \n", num, i, i*num)
	}

	return texto
}
