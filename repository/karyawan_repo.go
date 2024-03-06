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

//update karyawan
func (r *repositoryMysqlLayer) UpdateKaryawanByID(id int, karyawan model.Karyawan) error {
	//update karyawan
	res := r.DB.Where("id = ?", id).UpdateColumns(&karyawan)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update karyawan")
	}
	return nil
}

//delete karyawan
func (r *repositoryMysqlLayer) DeleteKaryawanByID(id int) error {
	//delete karyawan
	res := r.DB.Where("id = ?", id).Delete(&model.Karyawan{})
	if res.RowsAffected < 1 {
		return fmt.Errorf("karyawan not found")
	}
	return nil
}

//get all soft delete karyawan
func (r *repositoryMysqlLayer) GetAllSoftDeleteKaryawan() []model.Karyawan {
	//get all karyawan
	var karyawans []model.Karyawan
	r.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&karyawans)

	return karyawans
}

//restore soft delete karyawan
func (r *repositoryMysqlLayer) RestoreSoftDeleteKaryawan(id int) error {
	//restore karyawan
	res := r.DB.Unscoped().Model(&model.Karyawan{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NULL"))
	if res.RowsAffected < 1 {
		return fmt.Errorf("karyawan not found")
	}
	return nil
}
