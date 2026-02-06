# Portfolio Backend API 🚀

Este repositorio contiene el backend para mi Portafolio Profesional. Es una API RESTful construida con **Go (Golang)** utilizando el framework **Gin** y **MongoDB** como base de datos.

El proyecto ha sido diseñado siguiendo principios de **Arquitectura Limpia (Clean Architecture)** y el estándar de diseño de proyectos en Go, garantizando escalabilidad, mantenibilidad y desacoplamiento entre capas.

---

## 🛠️ Tech Stack

*   **Lenguaje:** Go 1.25+
*   **Framework Web:** [Gin Gonic](https://github.com/gin-gonic/gin) (Alto rendimiento y facilidad de uso).
*   **Base de Datos:** MongoDB Atlas (NoSQL).
*   **Driver:** [MongoDB Go Driver](https://go.mongodb.org/mongo-driver).
*   **Configuración:** [Godotenv](https://github.com/joho/godotenv) para manejo de variables de entorno.

---

## 🏗️ Arquitectura del Proyecto

El código está organizado siguiendo el **Standard Go Project Layout**, separando responsabilidades claramente:

```text
portfolio-backend/
├── cmd/
│   └── api/
│       └── main.go          # 🏁 Punto de entrada (Entry Point). Inicializa DB, rutas y servidor.
├── internal/
│   ├── database/            # 🔌 Conexión a MongoDB (Singleton pattern).
│   ├── handlers/            # 🎮 Capa de Controladores (HTTP). Valida inputs y llama al repositorio.
│   ├── models/              # 📦 Definición de Estructuras (Structs) y Tags (JSON/BSON).
│   └── repository/          # 💾 Capa de Acceso a Datos (Repository Pattern). Abstrae la lógica de DB.
├── .env                     # 🔐 Variables de entorno (No subir al repo).
├── requests.http            # 🧪 Archivo de pruebas HTTP (para VS Code / IntelliJ).
└── go.mod                   # 📋 Gestión de dependencias.
```

### Decisiones de Diseño
1.  **Repository Pattern:** La lógica de base de datos está aislada en `internal/repository`. Esto permite cambiar la base de datos en el futuro o realizar pruebas unitarias (Mocking) sin afectar a los controladores.
2.  **Inyección de Dependencias:** El repositorio se inyecta en los handlers, y los handlers se inyectan en el router. Evitamos el uso de variables globales.
3.  **Contextos:** Todas las operaciones a base de datos utilizan `context.WithTimeout` para asegurar la resiliencia del sistema y evitar bloqueos.

---

## 🚀 Instalación y Ejecución

### 1. Prerrequisitos
*   Tener instalado Go.
*   Tener una cuenta en MongoDB Atlas (o una instancia local).

### 2. Clonar y Dependencias
```bash
git clone <url-del-repo>
cd portfolio-backend
go mod tidy
```

### 3. Configuración (.env)
Crea un archivo `.env` en la raíz del proyecto basándote en `.env.example`:

```ini
MONGO_URI=mongodb+srv://<usuario>:<password>@cluster.mongodb.net/?retryWrites=true&w=majority
DB_NAME=portfolio
PORT=8080
```

### 4. Ejecutar el Servidor
⚠️ **Importante:** El punto de entrada está en `cmd/api`.

```bash
go run cmd/api/main.go
```
El servidor iniciará en `http://localhost:8080`.

---

## 📡 Endpoints de la API

### 💻 Tecnologías (`/tech`)
Gestión dinámica de las habilidades técnicas.

| Método | Endpoint     | Descripción                          |
| :---   | :---         | :---                                 |
| `POST` | `/tech`      | Crear una nueva tecnología.          |
| `GET`  | `/tech`      | Listar todas las tecnologías.        |
| `PUT`  | `/tech/:id`  | Actualizar una tecnología existente. |
| `DELETE`| `/tech/:id` | Eliminar una tecnología.             |

### 💼 Experiencia (`/experience`)
Gestión del historial laboral.

| Método | Endpoint           | Descripción                          |
| :---   | :---               | :---                                 |
| `POST` | `/experience`      | Agregar una experiencia laboral.     |
| `GET`  | `/experience`      | Listar historial completo.           |
| `PUT`  | `/experience/:id`  | Actualizar información laboral.      |
| `DELETE`| `/experience/:id` | Eliminar registro de experiencia.    |

### 📬 Contacto (`/contact`)
Recepción de mensajes desde el formulario del portafolio.

| Método | Endpoint   | Descripción                                      |
| :---   | :---       | :---                                             |
| `POST` | `/contact` | Recibe nombre, email y mensaje. Guarda en DB.    |

---

## 🧪 Pruebas Manuales

El proyecto incluye un archivo `requests.http` en la raíz. Puedes utilizarlo con:
*   **IntelliJ IDEA / GoLand:** Cliente HTTP integrado.
*   **VS Code:** Extensión "REST Client".

Esto te permite probar todos los endpoints (CRUD completo) directamente desde el editor sin necesidad de Postman.

---

## 👤 Autor

Desarrollado por **Juan Fernando** como backend para su portafolio profesional.
