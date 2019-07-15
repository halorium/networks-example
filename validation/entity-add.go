package validation

func (entity *Entity) Add(message string) {
	entity.Errors = append(entity.Errors, Error{message})
}
