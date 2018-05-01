package dao

import (
	"bipubipu/entity"
	"bipubipu/util"
	"database/sql"
)

const (
	ENTITY_TAB_NAME = "tab_volume"
)

func buildVolumesByRst(result *sql.Rows)(volumes []entity.Volume){
	
	volumes = make([]entity.Volume, 0)
	
	for result.Next() {
		var volume entity.Volume

		// 这里字段的顺序必须和数据库查出来的顺序保持一致，否则数据取不出来
		result.Scan(&volume.Vid, &volume.Name, &volume.Cover, 
					&volume.Desc, &volume.Uid, &volume.Status, &volume.CreateTime)

		volumes = append(volumes, volume)
	}
	
	return volumes
}

//-----------------------------------------------------------------------------
// 内部使用函数，使用common dao的公共接口进行扩展
func getAllVolume() (result *sql.Rows) {
	return GetAll(ENTITY_TAB_NAME);
}
func getVolumeBy1Field(structFieldName string, valName string) (result *sql.Rows) {
	tabFieldName := util.GetStructTagName(&entity.Volume{}, structFieldName, entity.TAB_FIELD_KEY)
	return GetBy1Field(ENTITY_TAB_NAME, tabFieldName, valName)
}
//-----------------------------------------------------------------------------

// 获取全部volume
func GetVolumeAll() (volumes []entity.Volume) {
	

	result := getAllVolume();
	defer result.Close()

	volumes = buildVolumesByRst(result)

	return volumes
}

// 按照user id获取全部volume
func GetVolumesByUserId(uid string) (volumes []entity.Volume) {
	
	result := getVolumeBy1Field("Uid", uid)
	defer result.Close()

	volumes = buildVolumesByRst(result)

	return volumes
}