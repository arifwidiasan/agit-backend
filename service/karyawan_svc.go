package service

import (
	"agit-backend/helper"
	"agit-backend/model"

	"fmt"
)

//create karyawan service
func (s *svc) CreateKaryawanService(karyawan model.Karyawan) error {
	//check if name is not empty
	if karyawan.Nama == "" {
		return fmt.Errorf("field name cannot be empty")
	}

	//check if tanggal lahir is not empty
	if karyawan.TanggalLahir == nil {
		return fmt.Errorf("field tanggal lahir cannot be nil or empty")
	}

	//set some field
	karyawan.Umur = helper.CalculateUmur(karyawan.TanggalLahir)
	karyawan.UpdatedAt = nil
	karyawan.DeletedAt = nil

	//insert karyawan to repository
	err := s.repo.CreateKaryawan(karyawan)
	if err != nil {
		return err
	}

	return nil
}

//get all karyawan service
func (s *svc) GetAllKaryawanService() []model.Karyawan {
	//get all karyawan from repository
	karyawans := s.repo.GetAllKaryawan()

	return karyawans
}

//get karyawan by id service
func (s *svc) GetKaryawanByIDService(id int) (karyawan model.Karyawan, err error) {
	return s.repo.GetKaryawanByID(id)
}
