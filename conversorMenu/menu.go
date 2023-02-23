package conversormenu
import (
	"fmt"
)

func Menu() {
	println("Qual conversão deseja realizar?")

	println("1 - Celsius para Fahrenheit")
	println("2 - Fahrenheit para Celsius")
	println("3 - Metros para pés")
	println("4 - Pés para metros")
	println("5 - Centimetros para covados")
	println("6 - Covados para centimetros")
	println("7 - Centimetros para metros")
	println("8 - Metros para centimetros")
	println("9 - Kg para libras")
	println("10- Libras para Kg")

}

func LeComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Println("Opção digita foi: ", comando)
	return comando
}


func ConversorCparaF() {

	var celsius float64
	println("Entre com a temperatura em graus celsius")
	fmt.Scan(&(celsius))
	var fahrenheit = celsius*1.8 + 32
	fmt.Printf("A temperatura de %g", celsius)
	fmt.Printf("Cº em fahrenheit é de %g", fahrenheit)
}

func ConversorFparaC() {
	var fahrenheit float64
	println("Entre com a temperatura em fahrenheit")
	fmt.Scan(fahrenheit)
	var celsius = (fahrenheit - 32) / 1.8
	fmt.Printf("A temperatura de %g", fahrenheit)
	fmt.Printf( "Fº em celsius é de %g", celsius)
}
func ConversorComprimento() {
	var metro float64
	println("Entre com o valor em metros")
	fmt.Scan(&(metro))
	var pes float64 = metro * 3.281
	fmt.Printf("O valor de %g", metro)
	fmt.Printf(" em metros convertendo para pés fica de %g", pes)
}
func ConversorMetro() {
	var pes float64
	println("Digite o valor em pés")
	fmt.Scan(&(pes))
	var metro float64 = pes / 3.281
	fmt.Printf("O valor em pés de %g", pes)
	fmt.Printf("convertido para metros fica %g", metro)
}
func ConversorCovado() {
	var centimetro float64
	println("Digite o valor em centimetros")
	fmt.Scan(&(centimetro))
	var covado = centimetro / 45.72
	fmt.Printf("O valor %g", centimetro)
	fmt.Printf("em centimetros convertidos para covados fica %g", covado)
}
func ConversorCentimetro() {
	var covado float64
	println("Digite o valor em covados")
	fmt.Scan(&(covado))
	var centimetro = covado * 45.72
	fmt.Printf(" O valor %g", covado)
	fmt.Printf(" convertido para centimetros fica %g", centimetro)
}
func ConversorCentParaMetro() {
	var centimetro float64
	println("Digite o valor em centimetros")
	fmt.Scan(&(centimetro))
	var metro = centimetro / 100
	fmt.Printf("O valor de %g", centimetro)
	fmt.Printf("convertido para metros fica %g", metro)
}
func ConversorMetroParaCent() {
	var metro float64
	println("Digite o valor em metros")
	fmt.Scan(&(metro))
	var centimetro = metro * 100
	fmt.Printf("O valor de %g", metro)
	fmt.Printf(" para centimetros é de %g", centimetro)

}
func ConversorLibra() {
	var libra float64
	println("Digite o valor em Libras")
	fmt.Scan(&(libra))
	var kg float64 = libra / 2.205
	fmt.Printf("O peso de %g", libra)
	fmt.Printf("em libra convertido para kg fica de %g ", kg)
}
func ConversorKG () {
	var kg float64
	println("Digite o valor em kg")
	fmt.Scan(&(kg))
	var libra = kg * 2.205
	fmt.Printf("O peso de %g ", kg)
	fmt.Printf("em kg convertido para libra fica de %g", libra)
}