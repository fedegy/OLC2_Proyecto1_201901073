package entorno

import "strings"

type Entorno struct {
	Nombre   string
	Anterior *Entorno
	Tabla    map[string]interface{}
}

func NewEntorno(nombre string, anterior *Entorno) Entorno {
	Tabla := make(map[string]interface{})
	env := Entorno{Nombre: nombre, Anterior: anterior, Tabla: Tabla}
	env.Tabla = make(map[string]interface{})
	return env
}

func (env *Entorno) ExisteSimbolo(identificador string) bool {
	ideFinal := strings.ToLower(identificador)
	for envActual := env; envActual != nil; envActual = envActual.Anterior {
		for key := range envActual.Tabla {
			if key == ideFinal {
				return true
			}
		}
	}
	return false
}

func (env *Entorno) SetSimbolo(identificador string, simbolo interface{}) {
	ideFinal := strings.ToLower(identificador)
	env.Tabla[ideFinal] = simbolo
}

func (env *Entorno) GetSimbolo(identificador string) interface{} {
	ideFinal := strings.ToLower(identificador)
	for envActual := env; envActual != nil; envActual = envActual.Anterior {
		for key, simbolElement := range envActual.Tabla {
			if key == ideFinal {
				return simbolElement
			}
		}
	}
	var simbolnil Simbolo
	return simbolnil
}
