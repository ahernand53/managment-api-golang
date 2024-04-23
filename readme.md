### Proyecto de API Management con Golang, MySQL y JWT Autenticado

Este proyecto implementa una API RESTful para la gestión de recursos utilizando Golang, MySQL y autenticación JWT.

**Características:**

* **Autenticación JWT:** Los usuarios se autentican mediante tokens JWT, lo que garantiza un acceso seguro a los recursos.
* **Gestión de tareas:** Permite crear, leer, actualizar y eliminar tareas.
* **Validación de datos:** Se valida la entrada del usuario para garantizar datos consistentes y seguros.
* **Docker Image** Imagen de docker para lanzar el aplicativo a produccion

**Requisitos:**

* Go 1.18 o superior
* MySQL 8.0 o superior
* Paquetes Golang:
    * `github.com/gorilla/mux`
    * `github.com/jinzhu/gorm`
    * `github.com/dgrijalva/jwt-go`
**Configuración:**

1. Clonar el repositorio.
2. Modificar el archivo `config.go` para editar las variables de entorno
```go
func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password_secret"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "projectmanager"),
		JWTSecret:  getEnv("JWT_SECRET", "randomjwtsercretkey"),
	}
}
```
3. Ejecutar `go mod tidy` para descargar los modulos de golang


4. Ejecutar `make run` para iniciar el servidor API.

**Uso de la API:**

**Ejemplo de uso:**

Para crear una nueva tarea, envíe una solicitud POST a `/api/tasks`:

```json
{
  "name": "Creating the task service",
  "status": "Active",
  "project_id": 1,
  "assigned_to_id": 1
}
```

Para obtener una lista de todos las tareas, envíe una solicitud GET a `/api/tasks`:

```
GET /api/tasks
```

Para obtener una tarea específica, envíe una solicitud GET a `/api/tasks/{id}`, donde `{id}` es el ID de la tarea:

```
GET /api/tasks/1
```

Para actualizar un recurso, envíe una solicitud PUT a `/api/resources/{id}` con los datos actualizados:

```json
{
  "name": "Creating the task service",
  "status": "Active",
  "project_id": 1,
  "assigned_to_id": 1
}
PUT /api/tasks/1
```

Para eliminar un recurso, envíe una solicitud DELETE a `/api/tasks/{id}`:

```
DELETE /api/tasks/1
```

**Contribuciones:**

Se agradecen las contribuciones a este proyecto. Si desea contribuir, cree una solicitud de extracción con sus cambios y pruebas.

**Licencia:**

Este proyecto está licenciado bajo la licencia MIT.
