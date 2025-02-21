package campaign

type Repository interface {
	Save(*Campaign) error
}
