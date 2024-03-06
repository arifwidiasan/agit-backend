package repository

import (
	"fmt"

	"agit-backend/model"

	"gorm.io/gorm"
)

//create karyawan
func (r *repositoryMysqlLayer) CreateKaryawan(karyawan model.Karyawan) error {
	//insert karyawan
	res := r.DB.Create(&karyawan).UpdateColumn("updated_at", gorm.Expr("NULL"))
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create karyawan")
	}
	return nil
}

//get all karyawan
func (r *repositoryMysqlLayer) GetAllKaryawan() []model.Karyawan {
	//get all karyawan
	var karyawans []model.Karyawan
	r.DB.Find(&karyawans)

	return karyawans
}

//get karyawan by id
func (r *repositoryMysqlLayer) GetKaryawanByID(id int) (karyawan model.Karyawan, err error) {
	//get karyawan by id
	res := r.DB.Where("id = ?", id).Find(&karyawan)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("user not found")
	}

	return
}
