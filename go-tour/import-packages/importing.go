package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	/*Em Go, um nome é exportado se ele começa com uma letra maiúscula. Por exemplo,
	Pizza é um nome exportado, assim como Pi, que é exportado do pacote math.
	pizza e pi não começam com uma letra maiúscula, logo eles não são exportados.
	Ao importar um pacote, você pode referenciar apenas seus nomes exportados.
	Todos os nomes "não exportados" não são acessíveis de fora do pacote.*/

	fmt.Println(math.Pi)
}
