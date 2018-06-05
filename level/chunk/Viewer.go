package chunk

import "github.com/google/uuid"

type Viewer interface {
	GetUUID() uuid.UUID
	GetXUID() string
}
