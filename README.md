# Go Version Manager (govm)
## Instalación

Para instalar govm, simplemente clona este repositorio y compila el proyecto:

```sh
git clone https://github.com/tu_usuario/govm.git
cd govm
go build
```

## Uso

El comando principal es `govm`, seguido de una subcomando y, opcionalmente, una versión de Go. Aquí están los comandos disponibles:

### Ayuda

Muestra la ayuda de uso:

```sh
govm help
```

### Listar versiones instaladas

Lista todas las versiones de Go instaladas en tu sistema:

```sh
govm list
```

### Usar una versión específica

Configura una versión específica de Go como la activa:

```sh
govm use 1.16.5
```

### Instalar una versión

Descarga e instala una versión específica de Go:

```sh
govm install 1.16.5
```

### Desinstalar una versión

Desinstala una versión específica de Go:

```sh
govm uninstall 1.16.5
```

### Mostrar la versión de govm

Muestra la versión actual de govm:

```sh
govm version
```

## Ejemplos

### Instalar una versión de Go

Para instalar la versión 1.16.5 de Go, ejecuta:

```sh
govm install 1.16.5
```

### Usar una versión instalada de Go

Para usar la versión 1.16.5 de Go, ejecuta:

```sh
govm use 1.16.5
```

### Listar todas las versiones instaladas

Para listar todas las versiones de Go instaladas, ejecuta:

```sh
govm list
```

### Desinstalar una versión de Go

Para desinstalar la versión 1.16.5 de Go, ejecuta:

```sh
govm uninstall 1.16.5
```

## Contribuir

Si deseas contribuir a este proyecto, por favor abre un issue o envía un pull request en GitHub.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.
