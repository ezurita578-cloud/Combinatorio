// Esteban Zurita Ramos #25760105
package main

import (
	"fmt"
)

// Producto representa un artículo del catálogo de la cafetería.
type Producto struct {
	Nombre      string
	Precio      float64
	Categoria   string
	Descripcion string
}

// catalogo contiene todos los productos disponibles.
var catalogo = []Producto{
	{Nombre: "Agua mineral", Precio: 20.00, Categoria: "Bebida", Descripcion: "Botella 600ml sin gas"},
	{Nombre: "Té verde", Precio: 30.00, Categoria: "Bebida", Descripcion: "Infusión caliente"},
	{Nombre: "Café Americano", Precio: 35.00, Categoria: "Bebida", Descripcion: "Café negro doble shot"},
	{Nombre: "Jugo de naranja", Precio: 45.00, Categoria: "Bebida", Descripcion: "Natural exprimido 355ml"},
	{Nombre: "Pastel de chocolate", Precio: 55.00, Categoria: "Postre", Descripcion: "Rebanada individual 120g"},
	{Nombre: "Burrito de res", Precio: 75.00, Categoria: "Comida", Descripcion: "Tortilla, carne, frijoles"},
	{Nombre: "Sandwich de pollo", Precio: 85.00, Categoria: "Comida", Descripcion: "Pan integral, pollo, verduras"},
	{Nombre: "Ensalada César", Precio: 95.00, Categoria: "Comida", Descripcion: "Lechuga romana, crutones"},
	{Nombre: "Pizza personal", Precio: 110.00, Categoria: "Comida", Descripcion: "4 rebanadas, queso y jitomate"},
	{Nombre: "Combo del día", Precio: 130.00, Categoria: "Combo", Descripcion: "Plato fuerte + bebida + postre"},
}

// encontrarCombinaciones usa backtracking para encontrar todas las
// combinaciones de productos que no excedan el presupuesto.
// El índice de inicio avanza en cada llamada para evitar duplicados.
func encontrarCombinaciones(productos []Producto, presupuesto float64) [][]Producto {
	var resultado [][]Producto
	var actual []Producto

	// Función recursiva que explora cada combinación posible
	var backtrack func(inicio int, restante float64)
	backtrack = func(inicio int, restante float64) {
		// Si hay productos en la combinación actual, la guardamos
		if len(actual) > 0 {
			copia := make([]Producto, len(actual))
			copy(copia, actual)
			resultado = append(resultado, copia)
		}

		// Exploramos los productos desde el índice actual en adelante
		for i := inicio; i < len(productos); i++ {
			// Solo agregamos si el producto cabe en el presupuesto restante
			if productos[i].Precio <= restante {
				actual = append(actual, productos[i])
				backtrack(i+1, restante-productos[i].Precio)
				// Quitamos el último producto para probar la siguiente opción
				actual = actual[:len(actual)-1]
			}
		}
	}

	backtrack(0, presupuesto)
	return resultado
}

// imprimirResultados muestra el resumen completo de combinaciones encontradas.
func imprimirResultados(combis [][]Producto, presupuesto float64) {
	if len(combis) == 0 {
		fmt.Println("No se encontraron combinaciones con ese presupuesto.")
		return
	}

	fmt.Printf("\nPresupuesto: $%.2f\n", presupuesto)
	fmt.Printf("Total de combinaciones: %d\n", len(combis))

	// Agrupar por cantidad de productos
	porCantidad := make(map[int]int)
	for _, c := range combis {
		porCantidad[len(c)]++
	}

	fmt.Println("\nPor cantidad de productos:")
	for i := 1; i <= len(catalogo); i++ {
		if count, ok := porCantidad[i]; ok {
			if i == 1 {
				fmt.Printf("  %d producto:  %d combinación(es)\n", i, count)
			} else {
				fmt.Printf("  %d productos: %d combinación(es)\n", i, count)
			}
		}
	}

	fmt.Println()

	// Detalle de cada combinación
	var mejorCombi []Producto
	mejorTotal := 0.0

	for idx, combi := range combis {
		total := 0.0
		for _, p := range combi {
			total += p.Precio
		}
		cambio := presupuesto - total

		fmt.Printf("[%d] %d producto(s) — Total: $%.2f — Cambio: $%.2f\n",
			idx+1, len(combi), total, cambio)
		for _, p := range combi {
			fmt.Printf("     • %-25s $%.2f\n", p.Nombre, p.Precio)
		}

		// Guardar la combinación de mayor valor
		if total > mejorTotal {
			mejorTotal = total
			mejorCombi = combi
		}
	}

	// Mostrar la mejor combinación
	fmt.Println("\nMejor combinación (mayor gasto):")
	for _, p := range mejorCombi {
		fmt.Printf("     • %-25s $%.2f\n", p.Nombre, p.Precio)
	}
	fmt.Printf("     Total: $%.2f  Cambio: $%.2f\n", mejorTotal, presupuesto-mejorTotal)
}

func main() {
	var presupuesto float64
	fmt.Print("Ingresa tu presupuesto: $")
	fmt.Scan(&presupuesto)

	// Validar que el presupuesto sea positivo
	if presupuesto <= 0 {
		fmt.Println("El presupuesto debe ser mayor a $0.")
		return
	}

	combis := encontrarCombinaciones(catalogo, presupuesto)
	imprimirResultados(combis, presupuesto)
}
