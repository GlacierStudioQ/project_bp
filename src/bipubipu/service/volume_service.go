package service

import (
	"bipubipu/entity"
	"bipubipu/dao"
)

func GetVolumeAll() (volumes []entity.Volume) {

	volumes = dao.GetVolumeAll()

	return volumes
}

func GetVolumesByUserId(uid string) (volumes []entity.Volume) {

	volumes = dao.GetVolumesByUserId(uid)

	return volumes
}