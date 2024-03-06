package service

import (
	"agit-backend/helper"
	"agit-backend/model"
	"time"

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

//update karyawan by id service
func (s *svc) UpdateKaryawanByIDService(id int, karyawan model.Karyawan) error {
	//check if tanggal lahir is not empty
	if karyawan.TanggalLahir != nil {
		//set new umur
		karyawan.Umur = helper.CalculateUmur(karyawan.TanggalLahir)
	}

	//set updated at
	now := time.Now()
	karyawan.UpdatedAt = &now

	//update karyawan to repository
	err := s.repo.UpdateKaryawanByID(id, karyawan)
	if err != nil {
		return err
	}

	return nil
}

//delete karyawan by id service
func (s *svc) DeleteKaryawanByIDService(id int) error {
	return s.repo.DeleteKaryawanByID(id)
}
