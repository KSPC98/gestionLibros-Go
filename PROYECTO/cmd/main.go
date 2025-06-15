package main

import (
	"fmt"
	"gestion-libros/internal/handlers"
	"gestion-libros/internal/storage"
	"os"
)

func main() {
	// Inicializamos el almacenamiento en memoria
	bookStore := storage.NewMemoryBookStorage()

	// Mensaje de bienvenida
	fmt.Println("=== SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS ===")
	fmt.Println("Desarrollado por Kevin Pérez - UIDE 2025")

	// Menú principal
	for {
		fmt.Println("\nMENÚ PRINCIPAL")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Buscar libro por ID")
		fmt.Println("3. Modificar libro")
		fmt.Println("4. Eliminar libro")
		fmt.Println("5. Listar todos los libros")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error leyendo opción:", err)
			continue
		}

		switch opcion {
		case 1:
			handlers.AgregarLibroHandler(bookStore)
		case 2:
			handlers.BuscarLibroHandler(bookStore)
		case 3:
			handlers.ModificarLibroHandler(bookStore)
		case 4:
			handlers.EliminarLibroHandler(bookStore)
		case 5:
			handlers.ListarLibrosHandler(bookStore)
		case 6:
			fmt.Println("Gracias por usar el sistema. ¡Hasta pronto!")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida. Por favor seleccione 1-6")
		}
	}
}
