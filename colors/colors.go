// Copyright © 2024 GOMV
// Licensed under the MIT License. See LICENSE file for details.

package colors

import (
	"github.com/fatih/color"
)

func makeMapColor() map[string]color.Attribute {
	// Este mapa asocia nombres de colores con sus codigos correspondientes.

	// Los colores se representan mediante valores enteros definidos en el paquete `color`.
	// Por ejemplo:
	// - Codigo para rojo: 31 (color.FgRed)
	// - Codigo para verde: 32 (color.FgGreen)

	// Cada entrada en el mapa utiliza la constante `color.Attribute`, que es el tipo manejado
	// por el paquete `color` para definir colores de texto.

	mapColor := map[string]color.Attribute{
		"red":   color.FgRed,
		"green": color.FgGreen,
	}

	return mapColor
}

func SetColor(c color.Attribute, format string, a ...interface{}) {
	// SetColor aplica un color especifico al texto y lo imprime.
	// Utiliza el mapa generado por `makeMapColor` para obtener los atributos de color.
	//
	// Args:
	//   [1] c: color.Attribute, el atributo de color, por el paquete color.
	//   [2] format: Cadena de texto que queremos imprimir.
	//   [3] a: Argumentos adicionales, se pueden pasar para formatear el texto.
	//
	// Ejemplo:
	//   SetColor(color.FgRed, "Error: %v", err)

	mapC := makeMapColor()
	switch c {
	case 31:
		red := color.New(mapC["red"]).PrintfFunc()
		red(format, a...)
	case 32:
		green := color.New(mapC["green"]).PrintfFunc()
		green(format, a...)
	}
}
