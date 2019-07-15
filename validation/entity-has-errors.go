package validation

func (entity *Entity) HasErrors() bool {
	return len(entity.Errors) > 0
}
