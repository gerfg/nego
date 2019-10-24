package department

import (
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/jinzhu/gorm"
)

type postgresDepartmentRepository struct {
	db *gorm.DB
}

// NewDepartmentRepository returns an interface abstraction for CRUD methods.
func NewDepartmentRepository(db *gorm.DB) repository.DepartmentRepository {
	return &postgresDepartmentRepository{
		db: db,
	}
}

func (r *postgresDepartmentRepository) Create(dep *model.Department) (id int, err error) {
	if err = r.db.Create(&dep).Error; err != nil {
		return
	}
	return dep.ID, nil
}

func (r *postgresDepartmentRepository) Fetch(limit int) (deps []*model.Department, err error) {
	if err = r.db.Find(&deps).Error; err != nil {
		return
	}
	return
}