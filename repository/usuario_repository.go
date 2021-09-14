package repository

import (
	"errors"
	"strconv"

	"www.github.com/miprimerproyectoengo/entities"
)

type UsuarioRepository interface {
	FindAll() []entities.Usuario
	FindOne(id string) (entities.Usuario, error)
	Save(entities.Usuario) entities.Usuario
	Update(id string, usuario entities.Usuario) entities.Usuario
}
type usuarioRepository struct {
	usuarios []entities.Usuario
}

func NewUsuarioRepository() UsuarioRepository {
	usuario1 := entities.Usuario{
		Id:       1,
		Nombre:   "Carlos",
		Apellido: "Perez",
		Edad:     "50",
		Estado:   true,
	}
	usuario2 := entities.Usuario{
		Id:       2,
		Nombre:   "Yeison",
		Apellido: "Sanchez",
		Edad:     "40",
		Estado:   true,
	}
	usuario3 := entities.Usuario{
		Id:       3,
		Nombre:   "Juan",
		Apellido: "Salcedo",
		Edad:     "20",
		Estado:   true,
	}

	usuarios := []entities.Usuario{}
	usuarios = append(usuarios, usuario1)
	usuarios = append(usuarios, usuario2)
	usuarios = append(usuarios, usuario3)

	return &usuarioRepository{usuarios: usuarios}
}

func (r *usuarioRepository) FindAll() []entities.Usuario {
	return r.usuarios
}
func (r *usuarioRepository) FindOne(id string) (entities.Usuario, error) {
	idEntero, _ := strconv.Atoi(id)

	for _, usuario := range r.usuarios {
		if idEntero == usuario.Id {
			return usuario, nil
		}
	}
	return entities.Usuario{}, errors.New("no se encontro el usuario...")
}
func (r *usuarioRepository) Save(usuario entities.Usuario) entities.Usuario {
	//Seteando ID
	r.usuarios = append(r.usuarios, usuario)
	return usuario
}
func (r *usuarioRepository) Update(id string, usuario entities.Usuario) entities.Usuario {
	idEntero, _ := strconv.Atoi(id)
	usuariosFinal := []entities.Usuario{}

	for _, usuarioOriginal := range r.usuarios {
		if idEntero == usuario.Id {
			usuariosFinal = append(usuariosFinal, usuario)
		} else {
			usuariosFinal = append(usuariosFinal, usuarioOriginal)
		}
	}
	r.usuarios = usuariosFinal

	return usuario
}
