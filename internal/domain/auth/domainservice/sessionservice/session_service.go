package sessionservice

import (
	entity "getfund-api-v2/internal/domain/auth/entity/sessionentity"
)

type SessionService interface {
	//Salvar sessão
	SaveSession(session entity.Session) error
	//Criar token a partir dos dados da sessão
	BuildToken(session entity.Session) error
}
