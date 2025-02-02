package majordomain

import "github.com/google/uuid"

type Major struct {
	id   uuid.UUID
	name string
}

func (a *Major) GetID() uuid.UUID {
	return a.id
}

func (a *Major) GetName() string {
	return a.name
}

func NewMajor(id uuid.UUID, name string) (*Major, error) {
	return &Major{
		id:   id,
		name: name,
	}, nil
}
