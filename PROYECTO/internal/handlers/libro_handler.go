package handlers

import (
	"bufio"
	"fmt"
	"gestion-libros/internal/models"
	"gestion-libros/internal/storage"
	"os"
	"strconv"
	"strings"
)

// Helper para leer entrada de usuario
func leerEntrada(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	return strings.TrimSpace(entrada)
}

// Agregar un nuevo libro
func AgregarLibroHandler(store storage.BookStorage) {
	fmt.Println("\n--- AGREGAR NUEVO LIBRO ---")

	id := leerEntrada("ID: ")
	titulo := leerEntrada("Título: ")
	autor := leerEntrada("Autor: ")

	precioStr := leerEntrada("Precio: ")
	precio, err := strconv.ParseFloat(precioStr, 64)
	if err != nil {
		fmt.Println("Error en precio:", err)
		return
	}

	descripcion := leerEntrada("Descripción: ")

	stockStr := leerEntrada("Stock: ")
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		fmt.Println("Error en stock:", err)
		return
	}

	// Usamos el constructor del modelo
	libro, err := models.NewLibro(id, titulo, autor, precio, descripcion, stock)
	if err != nil {
		fmt.Println("Error creando libro:", err)
		return
	}

	// Guardamos en el almacenamiento
	if err := store.AgregarLibro(libro); err != nil {
		fmt.Println("Error guardando libro:", err)
	} else {
		fmt.Println("¡Libro agregado exitosamente!")
	}
}

// Buscar libro por ID
func BuscarLibroHandler(store storage.BookStorage) {
	fmt.Println("\n--- BUSCAR LIBRO POR ID ---")
	id := leerEntrada("Ingrese ID del libro: ")

	libro, err := store.BuscarPorID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nLibro encontrado:\nID: %s\nTítulo: %s\nAutor: %s\nPrecio: %.2f\nStock: %d\nDescripción: %s\n",
		libro.ID(), libro.Titulo(), libro.Autor(), libro.Precio(), libro.Stock(), libro.Descripcion())
}

// Actualizar un libro existente
func ModificarLibroHandler(store storage.BookStorage) {
	fmt.Println("\n--- MODIFICAR LIBRO ---")
	id := leerEntrada("Ingrese ID del libro a modificar: ")

	// Buscamos el libro existente
	libro, err := store.BuscarPorID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Deje en blanco los campos que no desea cambiar")

	nuevoPrecioStr := leerEntrada(fmt.Sprintf("Nuevo precio [actual: %.2f]: ", libro.Precio()))
	if nuevoPrecioStr != "" {
		nuevoPrecio, err := strconv.ParseFloat(nuevoPrecioStr, 64)
		if err != nil {
			fmt.Println("Error en precio:", err)
			return
		}
		if err := libro.SetPrecio(nuevoPrecio); err != nil {
			fmt.Println("Error actualizando precio:", err)
			return
		}
	}

	nuevoStockStr := leerEntrada(fmt.Sprintf("Nuevo stock [actual: %d]: ", libro.Stock()))
	if nuevoStockStr != "" {
		nuevoStock, err := strconv.Atoi(nuevoStockStr)
		if err != nil {
			fmt.Println("Error en stock:", err)
			return
		}
		if err := libro.SetStock(nuevoStock); err != nil {
			fmt.Println("Error actualizando stock:", err)
			return
		}
	}

	// Actualizamos en el almacenamiento
	if err := store.ActualizarLibro(libro); err != nil {
		fmt.Println("Error actualizando libro:", err)
	} else {
		fmt.Println("¡Libro actualizado exitosamente!")
	}
}

// Eliminar un libro
func EliminarLibroHandler(store storage.BookStorage) {
	fmt.Println("\n--- ELIMINAR LIBRO ---")
	id := leerEntrada("Ingrese ID del libro a eliminar: ")

	if err := store.EliminarLibro(id); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("¡Libro eliminado exitosamente!")
	}
}

// Listar todos los libros
func ListarLibrosHandler(store storage.BookStorage) {
	fmt.Println("\n--- LISTA DE TODOS LOS LIBROS ---")
	libros := store.ListarTodos()

	if len(libros) == 0 {
		fmt.Println("No hay libros registrados")
		return
	}

	for i, libro := range libros {
		fmt.Printf("\n%d. [ID: %s] %s - %s (%.2f USD) Stock: %d",
			i+1, libro.ID(), libro.Titulo(), libro.Autor(), libro.Precio(), libro.Stock())
	}
	fmt.Println("\n---------------------------------")
	fmt.Printf("Total de libros: %d\n", len(libros))
}
