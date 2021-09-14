package services

import (
	"www.github.com/miprimerproyectoengo/entities"
	"www.github.com/miprimerproyectoengo/repository"
)

type UsuarioService interface {
	FindAll() []entities.Usuario
	FindOne(id string) (entities.Usuario, error)
	Save(entities.Usuario) entities.Usuario
	Update(id string, usuario entities.Usuario) entities.Usuario
}

type usuarioService struct {
	repository repository.UsuarioRepository
}

func NewUsuarioService() UsuarioService {
	return &usuarioService{repository: repository.NewUsuarioRepository()}
}
func (r *usuarioService) FindAll() []entities.Usuario {
	return r.repository.FindAll()
}
func (r *usuarioService) FindOne(id string) (entities.Usuario, error) {
	return r.repository.FindOne(id)
}
func (r *usuarioService) Save(usuario entities.Usuario) entities.Usuario {
	return r.repository.Save(usuario)
}
func (r *usuarioService) Update(id string, usuario entities.Usuario) entities.Usuario {
	return r.repository.Update(id, usuario)
}
