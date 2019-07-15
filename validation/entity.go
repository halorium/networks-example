package validation

type Entity struct {
	Errors []Error `json:"errors"`
}

func New() *Entity {
	return &Entity{}
}

func NewFrom(message string) *Entity {
	entity := New()
	entity.Add(message)
	return entity
}
