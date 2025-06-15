package models

import "errors"

// Libro representa un libro en el sistema (estructura principal)
type Libro struct {
	id          string
	titulo      string
	autor       string
	precio      float64
	descripcion string
	stock       int
}

// Constructor para crear un nuevo libro
func NewLibro(id, titulo, autor string, precio float64, descripcion string, stock int) (*Libro, error) {
	if id == "" {
		return nil, errors.New("ID no puede estar vacío")
	}
	if precio <= 0 {
		return nil, errors.New("precio debe ser positivo")
	}
	return &Libro{
		id:          id,
		titulo:      titulo,
		autor:       autor,
		precio:      precio,
		descripcion: descripcion,
		stock:       stock,
	}, nil
}

// Getters para acceder a campos privados (Encapsulación)
func (l *Libro) ID() string {
	return l.id
}

func (l *Libro) Titulo() string {
	return l.titulo
}

func (l *Libro) Autor() string {
	return l.autor
}

func (l *Libro) Precio() float64 {
	return l.precio
}

func (l *Libro) Descripcion() string {
	return l.descripcion
}

func (l *Libro) Stock() int {
	return l.stock
}

// Setters con validación (Encapsulación)
func (l *Libro) SetPrecio(nuevoPrecio float64) error {
	if nuevoPrecio <= 0 {
		return errors.New("precio debe ser positivo")
	}
	l.precio = nuevoPrecio
	return nil
}

func (l *Libro) SetStock(nuevoStock int) error {
	if nuevoStock < 0 {
		return errors.New("stock no puede ser negativo")
	}
	l.stock = nuevoStock
	return nil
}
