# Expense Tracker API

API REST para el registro y seguimiento de gastos personales, construida en Go con el framework Echo. **Proyecto en fase de MVP** — funcionalidades core implementadas, sujeto a cambios e iteraciones.

## 📋 Descripción

Expense Tracker API permite a los usuarios registrarse, autenticarse y gestionar sus gastos personales de forma segura. Cada usuario solo puede ver y modificar sus propios gastos, protegidos mediante autenticación JWT.

## ✨ Features

- **Registro de usuarios** — Sign up como nuevo usuario.
- **Autenticación** — Generación y validación de JWTs para el manejo de sesión.
- **Listado gastos** — Listado de todos los gastos.
- **Alta de gastos** — Agregar un nuevo gasto.

## 🛠️ Stack Tecnológico

| Módulo | Tecnología | Conceptos |
|---|---|---|
| **API REST** | Go + Echo | Routing, Middlewares |
| **Base de datos** | SQLite + GORM | Manejo de datos |
| **Autenticación** | JWT + bcrypt | Seguridad, hashing de contraseñas |

## 🏗️ Arquitectura

El proyecto sigue una arquitectura en capas:

```
handler → repository → database
```

```
expense-tracker-api/
├── internal/          # Raíz del proyecto
  ├── config/          # Carga y validación de configuración/env vars
  ├── database/         # Conexión y migraciones de PostgreSQL
  ├── models/           # Entidades del dominio
  ├── handlers/          # Controladores HTTP (Echo)
  ├── middleware/       # JWT auth, logging, etc.
  ├── router/            # Definición de rutas
├── .env               # Variables de entorno
├── tmp/               # Archivos temporales (logs, etc.)
```

## 🔐 Autenticación

- Los endpoints de auth (`sign up`, `login`) son **públicos**.
- Los endpoints de gastos (CRUD) están **protegidos** y requieren un JWT válido en el header `Authorization: Bearer <token>`.
- Las contraseñas se almacenan hasheadas con **bcrypt**.

## 🚀 Cómo correr el proyecto

### Requisitos previos
- Go 1.21+ instalado.

### Pasos

1. Cloná el repositorio.
2. Configurá las variables de entorno (ver `.env.example` si existe, o las variables requeridas en `config/`).
3. La API va a estar disponible en `http://localhost:PORT` (ver configuración).
4. Las migraciones del schema se aplican automáticamente.

## 📌 Endpoints principales

| Método | Endpoint | Descripción | Auth |
|---|---|---|---|
| POST | `v1/auth/signup` | Registrar nuevo usuario | No |
| POST | `v1/auth/login` | Iniciar sesión y obtener JWT | No |
| GET | `v1/expenses` | Listar gastos (con filtros de fecha) | Sí |
| POST | `/v1/expenses` | Crear un nuevo gasto | Sí |

> ⚠️ Los nombres exactos de rutas y payloads pueden variar según la implementación final en `router/` y `dto/`.

## 🗓️ Estado del proyecto

**MVP** — En desarrollo activo. Próximos pasos:
- [ ] Tests unitarios y de integración
- [ ] Documentación de la API (Swagger/OpenAPI)
- [ ] Filtrado de gastos
- [ ] Paginación en el listado de gastos
- [ ] Categorías de gastos
- [ ] Roles y permisos

## 📄 Licencia

MIT
