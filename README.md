# Sistema de Registro de Empleados - REST API

## Descripción

Este proyecto consiste en el desarrollo de una aplicación basada en arquitectura **Cliente-Servidor** utilizando una **REST API** para la administración y registro de empleados.

El servidor fue desarrollado utilizando **Go + Gin + GORM + MySQL**, mientras que el cliente fue desarrollado utilizando **HTML, CSS y JavaScript**.

La aplicación permite registrar empleados, consultar información almacenada y realizar operaciones CRUD mediante comunicación en formato JSON.

---

# Integrantes

- Silva Sanchez Enrique
- Alvarez Guevara Estefania Guadalupe

---

# Tecnologías utilizadas

| Tecnología | Uso |
|------------|------|
| Go | Desarrollo del servidor |
| Gin | Framework HTTP |
| GORM | ORM para MySQL |
| MySQL | Base de datos |
| HTML | Interfaz cliente |
| CSS | Diseño |
| JavaScript | Consumo de API |
| GitHub | Control de versiones |
| GitHub Actions | Integración continua |

---

# Arquitectura del sistema

```text
Cliente Web
     ↓
REST API (Go + Gin)
     ↓
GORM
     ↓
MySQL
```

---

# Funcionalidades implementadas

## CRUD de empleados

- Crear empleado
- Obtener empleados
- Obtener empleado por ID
- Actualizar empleado
- Eliminar empleado

---

# Casos de uso implementados

## Caso 1: Registro de empleados

El usuario:

1. Selecciona su nombre
2. Selecciona puesto
3. Selecciona departamento
4. Selecciona fecha
5. Presiona registrar

El sistema:

- Envía información al servidor
- Procesa los datos
- Guarda información en MySQL

---

## Caso 2: Consulta de empleados

El usuario puede visualizar todos los empleados registrados.

---

## Caso 3: Consulta por departamento

Permite consultar únicamente empleados pertenecientes a un departamento específico.

Ruta:

```http
GET /departamento/:departamento
```

---

# Estructura del proyecto

```text
proyecto-empleados/

├── controllers/
│   └── empleadoController.go
│
├── database/
│   └── database.go
│
├── models/
│   └── empleado.go
│
├── routes/
│   └── routes.go
│
├── cliente/
│   ├── index.html
│   ├── empleados.js
│   └── estilos.css
│
├── .github/
│   └── workflows/
│       └── go-test.yml
│
├── main.go
├── go.mod
└── README.md
```

---

# Base de datos

Base utilizada:

```sql
empleados_db
```

Tabla principal:

```text
empleados
```

Campos:

- ID
- Nombre
- Puesto
- Departamento
- FechaIngreso

---

# Endpoints disponibles

## Obtener empleados

```http
GET /empleados
```

---

## Obtener empleado por ID

```http
GET /empleados/:id
```

---

## Registrar empleado

```http
POST /empleados
```

Ejemplo:

```json
{
    "nombre":"Enrique Silva Sánchez",
    "puesto":"Líder",
    "departamento":"Producción",
    "fechaIngreso":"2026-05-23T00:00:00Z"
}
```

---

## Actualizar empleado

```http
PUT /empleados/:id
```

---

## Eliminar empleado

```http
DELETE /empleados/:id
```

---

## Buscar empleados por departamento

```http
GET /departamento/:departamento
```

---

# Instalación y ejecución

## Clonar repositorio

```bash
git clone https://github.com/enriqueSilva1012/proyecto-empleados.git
```

Entrar al proyecto:

```bash
cd proyecto-empleados
```

---

## Instalar dependencias

```bash
go mod tidy
```

---

## Crear base de datos

Abrir MySQL y ejecutar:

```sql
CREATE DATABASE empleados_db;
```

---

## Ejecutar servidor

```bash
go run main.go
```

Servidor disponible en:

```text
http://localhost:8080
```

---

# Cliente

Abrir:

```text
cliente/index.html
```

La aplicación cliente consumirá automáticamente la API.

---

# GitHub Actions

Se implementó integración continua mediante GitHub Actions.

Funciones:

- Descargar proyecto
- Instalar dependencias
- Compilar aplicación
- Ejecutar pruebas automáticas

Archivo:

```text
.github/workflows/go-test.yml
```

---

# Resultados

Se desarrolló una aplicación funcional basada en arquitectura REST API con:

✅ Comunicación Cliente-Servidor

✅ Integración con MySQL

✅ Operaciones CRUD

✅ Casos de uso implementados

✅ JSON como medio de comunicación

✅ GitHub Actions

---

# Referencias

Go:

https://go.dev/

Gin:

https://gin-gonic.com/

GORM:

https://gorm.io/

MySQL:

https://www.mysql.com/

GitHub:

https://github.com/
