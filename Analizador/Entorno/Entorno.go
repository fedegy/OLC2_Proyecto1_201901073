package entorno

type Entorno struct {
	Nombre   string
	Anterior *Entorno
	Tabla    map[string]interface{}
}

func SetEntorno(nombre string, anterior *Entorno) Entorno {
	env := Entorno{Nombre: nombre, Anterior: anterior}
	env.Tabla = make(map[string]interface{})
	return env
}
