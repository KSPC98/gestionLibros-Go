package storage

import (
	"errors"
	"gestion-libros/internal/models" // Importamos nuestro paquete de modelos
	"sync"
)

// Interfaz para abstraer el almacenamiento
type BookStorage interface {
	AgregarLibro(*models.Libro) error
	BuscarPorID(string) (*models.Libro, error)
	ActualizarLibro(*models.Libro) error
	EliminarLibro(string) error
	ListarTodos() []*models.Libro
}

// Implementación concreta usando map y mutex
type MemoryBookStorage struct {
	libros map[string]*models.Libro // Mapa para almacenar libros por ID
	mutex  sync.RWMutex             // Mutex para control de concurrencia
}

// Constructor del almacenamiento
func NewMemoryBookStorage() *MemoryBookStorage {
	return &MemoryBookStorage{
		libros: make(map[string]*models.Libro),
	}
}

// Agregar un libro al almacenamiento
func (s *MemoryBookStorage) AgregarLibro(libro *models.Libro) error {
	s.mutex.Lock()         // Bloqueo para escritura
	defer s.mutex.Unlock() // Aseguramos desbloqueo al final

	id := libro.ID()
	if _, existe := s.libros[id]; existe {
		return errors.New("el libro ya existe")
	}

	s.libros[id] = libro
	return nil
}

// Buscar libro por ID
func (s *MemoryBookStorage) BuscarPorID(id string) (*models.Libro, error) {
	s.mutex.RLock()         // Bloqueo para lectura
	defer s.mutex.RUnlock() // Aseguramos desbloqueo

	libro, existe := s.libros[id]
	if !existe {
		return nil, errors.New("libro no encontrado")
	}
	return libro, nil
}

// Actualizar un libro existente
func (s *MemoryBookStorage) ActualizarLibro(libro *models.Libro) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := libro.ID()
	if _, existe := s.libros[id]; !existe {
		return errors.New("libro no existe, no se puede actualizar")
	}

	s.libros[id] = libro
	return nil
}

// Eliminar un libro por ID
func (s *MemoryBookStorage) EliminarLibro(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, existe := s.libros[id]; !existe {
		return errors.New("libro no encontrado")
	}

	delete(s.libros, id)
	return nil
}

// Listar todos los libros (retorna slice de punteros)
func (s *MemoryBookStorage) ListarTodos() []*models.Libro {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	lista := make([]*models.Libro, 0, len(s.libros))
	for _, libro := range s.libros {
		lista = append(lista, libro)
	}
	return lista
}
