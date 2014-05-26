package models

type Owner interface {
	CanAccessRepository(repository *Repository) error
}
