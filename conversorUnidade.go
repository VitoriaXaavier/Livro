package main

import "exLivro/conversorMenu"
func main() {

	conversormenu.Menu()

	comando := conversormenu.LeComando()


	switch comando {
	case 1:
		conversormenu.ConversorCparaF()
	case 2:
		conversormenu.ConversorFparaC()
	case 3:
		conversormenu.ConversorComprimento()
	case 4:
		conversormenu.ConversorMetro()
	case 5:
		conversormenu.ConversorCovado()
	case 6:
		conversormenu.ConversorCentimetro()
	case 7:
		conversormenu.ConversorCentParaMetro()
	case 8:
		conversormenu.ConversorMetroParaCent()
	case 9:
		conversormenu.ConversorLibra()
	case 10:
		conversormenu.ConversorKG()
	default:
		println("Opção invalida")

	}

}
