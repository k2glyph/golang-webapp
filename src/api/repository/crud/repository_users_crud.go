package crud

type &repositoryUsersCRUD struct {
	db *gorm.DB
}

func (r *repositoryUsersCRUD) NewRepositoryUsersCrud(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD(db)
}

func (r *repositoryUsersCRUD) Save(users models.User)(models.User, error) {
	var error
}