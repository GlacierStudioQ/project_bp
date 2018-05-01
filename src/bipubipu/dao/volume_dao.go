package dao

import (
	"bipubipu/entity"
)

const (
	entityName = "tab_volume"
)

// 获取全部volume
func GetVolumeAll() (volumes []entity.Volume) {
	
	volumes = make([]entity.Volume, 0)

	result, _ := DbGet().Query("SELECT * FROM " + entityName)
	defer result.Close()

	for result.Next() {
		var volume entity.Volume
		
		result.Scan(&volume.Vid, &volume.Name, &volume.Cover, 
					&volume.CreateTime, &volume.Desc, &volume.Status, &volume.Uid)

		volumes = append(volumes, volume)
	}

	return volumes
}
