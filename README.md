# Detector de ADN Mutante

## Descripción

Este proyecto tiene como objetivo detectar si una secuencia de ADN pertenece a un mutante o a un humano. Utiliza un enfoque basado en el análisis de patrones en las secuencias de ADN para determinar la mutación.

## Requisitos

- Go 1.18 o superior
- PostgreSQL

## Instalación

1. Clona el repositorio:
   ```bash
   git clone https://github.com/NicoJCastro/meli-test-mutante.git
   cd meli-test-mutante
   ```

2. Configura la base de datos PostgreSQL:
   - Crea una base de datos llamada `test_db`
   - Actualiza la configuración de conexión en `database.go`:
     ```go
     dsn := "host=localhost user=nicolas password=nico123 dbname=test_db port=5432 sslmode=disable"
     ```

## Ejecución

Ejecuta la aplicación:
```bash
go run main.go
```

## Estructura del Proyecto

- `main.go`: Punto de entrada de la aplicación
- `controllers/`: Lógica para manejar solicitudes HTTP
- `database/`: Conexión y manejo de base de datos
- `models/`: Estructuras de datos

## Uso de la API

### Detección de Mutante

**Endpoint**: `POST /mutant/`

**Ejemplo de solicitud**:
```json
{
    "dna": [
        "ATGCGA",
        "CAGTGC",
        "TTATGT",
        "AGAAGG",
        "CCCCTA",
        "TCACTG"
    ]
}
```

### Estadísticas

**Endpoint**: `GET /stats/`

Obtiene estadísticas sobre las secuencias de ADN procesadas, incluyendo:
- Número de ADN mutante
- Número de ADN humano
- Relación entre mutantes y humanos

## Contribuciones

Contribuciones son bienvenidas:
1. Abre un issue
2. Realiza un fork del repositorio
3. Crea una nueva rama
4. Envía un pull request

## Licencia

Proyecto bajo Licencia MIT. Consulta el archivo `LICENSE` para más detalles.