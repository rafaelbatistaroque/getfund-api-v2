package sessionentityfixture

import entity "getfund-api-v2/internal/domain/auth/entity/sessionentity"

var (
	FAKE_ID           = "fake-id"
	FAKE_FIRSTNAME    = "fake-first-name"
	FAKE_TOKEN        = "fake-first-name"
	EMPTY_STRING      = ""
	FAKE_ROLE         = 0
	FAKE_INVALID_ROLE = -1
)

func NewSessionWithInvalidId() (entity.Session, error) {
	return entity.New(EMPTY_STRING, FAKE_FIRSTNAME, FAKE_ROLE)
}

func NewSessionWithInvalidFirstName() (entity.Session, error) {
	return entity.New(FAKE_ID, EMPTY_STRING, FAKE_ROLE)
}

func NewSessionWithInvalidRole() (entity.Session, error) {
	return entity.New(FAKE_ID, FAKE_FIRSTNAME, FAKE_INVALID_ROLE)
}

func NewValidSession() (entity.Session, error) {
	return entity.New(FAKE_ID, FAKE_FIRSTNAME, FAKE_ROLE)
}
