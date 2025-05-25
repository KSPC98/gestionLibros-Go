# Sistema de Gestión de Libros Electrónicos en Go 📚

Plataforma para gestionar y comercializar libros electrónicos desarrollada en Go.

## Características Principales ⚡
- Publicación y búsqueda de libros electrónicos
- Carrito de compras integrado
- Gestión de usuarios y permisos
- Almacenamiento seguro en memoria

## Estructura del Proyecto 🗂️
```plaintext
Proyecto/
├── cmd/               # Punto de entrada
│   └── main.go        # Servidor principal
├── database/          # Scripts de base de datos
├── internal/          # Lógica interna
│   ├── handlers/      # Manejo de peticiones HTTP
│   ├── models/        # Estructuras de datos
│   ├── security/      # Autenticación/autorización
│   ├── storage/       # Almacenamiento persistente
│   └── utils/         # Funciones auxiliares
├── static/            # Recursos estáticos
├── .gitignore         # Archivos excluidos
