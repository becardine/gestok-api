package common

import (
	"fmt"
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
	if s == "" {
		return ID{}, errors.NewErrInvalidID("ID cannot be empty")
	}

	id, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}

	return ID{value: id}, nil
}

func (i ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, i.value.String())), nil
}

func (i ID) String() string {
	return i.value.String()
}

func NewIDFromUUID(uuid uuid.UUID) ID {
	return ID{value: uuid}
}

func (i ID) Value() uuid.UUID {
	return i.value
}

func (i *ID) Scan(value interface{}) error {
	id, ok := value.(uuid.UUID)
	if !ok {
		return errors.NewErrInvalidIDType("invalid ID type")
	}

	if id == uuid.Nil {
		return errors.NewErrInvalidID("invalid ID")
	}

	i.value = id

	return nil
}

func (i ID) IsEmpty() bool {
	return i.value == uuid.Nil
}
