package service

import (
	"bipubipu/entity"
	"bipubipu/dao"
)

func GetVolumeAll() (volumes []entity.Volume) {

	volumes = dao.GetVolumeAll()

	return volumes
}
