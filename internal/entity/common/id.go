package common

import (
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func NewIDFromString(s string) (ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}

	return ID{value: id}, nil
}

func (i ID) String() string {
	return i.value.String()
}

func (i ID) Value() uuid.UUID {
	return i.value
}

func (i *ID) Scan(value interface{}) error {
	id, ok := value.(uuid.UUID)
	if !ok {
		return errors.ErrInvalidID
	}

	i.value = id

	return nil
}
