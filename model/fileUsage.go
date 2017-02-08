package model

type FileUsage struct {
	Fid    int64  `xorm:"not null pk index(fid_count) index(fid_module) INT(10)"`
	Module string `xorm:"not null pk default '' index(fid_module) VARCHAR(50)"`
	Type   string `xorm:"not null pk default '' index(type_id) VARCHAR(64)"`
	Id     string `xorm:"not null pk default '0' index(type_id) VARCHAR(64)"`
	Count  int    `xorm:"not null default 0 index(fid_count) INT(10)"`
}

// GetFileUsagesCount FileUsages' Count
func GetFileUsagesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&FileUsage{})
	return total, err
}

// GetFileUsageCountViaFid Get FileUsage via Fid
func GetFileUsageCountViaFid(iFid int64) int64 {
	n, _ := Engine.Where("fid = ?", iFid).Count(&FileUsage{Fid: iFid})
	return n
}

// GetFileUsageCountViaModule Get FileUsage via Module
func GetFileUsageCountViaModule(iModule string) int64 {
	n, _ := Engine.Where("module = ?", iModule).Count(&FileUsage{Module: iModule})
	return n
}

// GetFileUsageCountViaType Get FileUsage via Type
func GetFileUsageCountViaType(iType string) int64 {
	n, _ := Engine.Where("type = ?", iType).Count(&FileUsage{Type: iType})
	return n
}

// GetFileUsageCountViaId Get FileUsage via Id
func GetFileUsageCountViaId(iId string) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&FileUsage{Id: iId})
	return n
}

// GetFileUsageCountViaCount Get FileUsage via Count
func GetFileUsageCountViaCount(iCount int) int64 {
	n, _ := Engine.Where("count = ?", iCount).Count(&FileUsage{Count: iCount})
	return n
}

// GetFileUsagesByFidAndModuleAndType Get FileUsages via FidAndModuleAndType
func GetFileUsagesByFidAndModuleAndType(offset int, limit int, Fid_ int64, Module_ string, Type_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and module = ? and type = ?", Fid_, Module_, Type_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndModuleAndId Get FileUsages via FidAndModuleAndId
func GetFileUsagesByFidAndModuleAndId(offset int, limit int, Fid_ int64, Module_ string, Id_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and module = ? and id = ?", Fid_, Module_, Id_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndModuleAndCount Get FileUsages via FidAndModuleAndCount
func GetFileUsagesByFidAndModuleAndCount(offset int, limit int, Fid_ int64, Module_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and module = ? and count = ?", Fid_, Module_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndTypeAndId Get FileUsages via FidAndTypeAndId
func GetFileUsagesByFidAndTypeAndId(offset int, limit int, Fid_ int64, Type_ string, Id_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and type = ? and id = ?", Fid_, Type_, Id_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndTypeAndCount Get FileUsages via FidAndTypeAndCount
func GetFileUsagesByFidAndTypeAndCount(offset int, limit int, Fid_ int64, Type_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and type = ? and count = ?", Fid_, Type_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndIdAndCount Get FileUsages via FidAndIdAndCount
func GetFileUsagesByFidAndIdAndCount(offset int, limit int, Fid_ int64, Id_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and id = ? and count = ?", Fid_, Id_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByModuleAndTypeAndId Get FileUsages via ModuleAndTypeAndId
func GetFileUsagesByModuleAndTypeAndId(offset int, limit int, Module_ string, Type_ string, Id_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("module = ? and type = ? and id = ?", Module_, Type_, Id_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByModuleAndTypeAndCount Get FileUsages via ModuleAndTypeAndCount
func GetFileUsagesByModuleAndTypeAndCount(offset int, limit int, Module_ string, Type_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("module = ? and type = ? and count = ?", Module_, Type_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByModuleAndIdAndCount Get FileUsages via ModuleAndIdAndCount
func GetFileUsagesByModuleAndIdAndCount(offset int, limit int, Module_ string, Id_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("module = ? and id = ? and count = ?", Module_, Id_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByTypeAndIdAndCount Get FileUsages via TypeAndIdAndCount
func GetFileUsagesByTypeAndIdAndCount(offset int, limit int, Type_ string, Id_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("type = ? and id = ? and count = ?", Type_, Id_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndModule Get FileUsages via FidAndModule
func GetFileUsagesByFidAndModule(offset int, limit int, Fid_ int64, Module_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and module = ?", Fid_, Module_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndType Get FileUsages via FidAndType
func GetFileUsagesByFidAndType(offset int, limit int, Fid_ int64, Type_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and type = ?", Fid_, Type_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndId Get FileUsages via FidAndId
func GetFileUsagesByFidAndId(offset int, limit int, Fid_ int64, Id_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and id = ?", Fid_, Id_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByFidAndCount Get FileUsages via FidAndCount
func GetFileUsagesByFidAndCount(offset int, limit int, Fid_ int64, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ? and count = ?", Fid_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByModuleAndType Get FileUsages via ModuleAndType
func GetFileUsagesByModuleAndType(offset int, limit int, Module_ string, Type_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("module = ? and type = ?", Module_, Type_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByModuleAndId Get FileUsages via ModuleAndId
func GetFileUsagesByModuleAndId(offset int, limit int, Module_ string, Id_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("module = ? and id = ?", Module_, Id_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByModuleAndCount Get FileUsages via ModuleAndCount
func GetFileUsagesByModuleAndCount(offset int, limit int, Module_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("module = ? and count = ?", Module_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByTypeAndId Get FileUsages via TypeAndId
func GetFileUsagesByTypeAndId(offset int, limit int, Type_ string, Id_ string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("type = ? and id = ?", Type_, Id_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByTypeAndCount Get FileUsages via TypeAndCount
func GetFileUsagesByTypeAndCount(offset int, limit int, Type_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("type = ? and count = ?", Type_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesByIdAndCount Get FileUsages via IdAndCount
func GetFileUsagesByIdAndCount(offset int, limit int, Id_ string, Count_ int) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("id = ? and count = ?", Id_, Count_).Limit(limit, offset).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsages Get FileUsages via field
func GetFileUsages(offset int, limit int, field string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Limit(limit, offset).Desc(field).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesViaFid Get FileUsages via Fid
func GetFileUsagesViaFid(offset int, limit int, Fid_ int64, field string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("fid = ?", Fid_).Limit(limit, offset).Desc(field).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesViaModule Get FileUsages via Module
func GetFileUsagesViaModule(offset int, limit int, Module_ string, field string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("module = ?", Module_).Limit(limit, offset).Desc(field).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesViaType Get FileUsages via Type
func GetFileUsagesViaType(offset int, limit int, Type_ string, field string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("type = ?", Type_).Limit(limit, offset).Desc(field).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesViaId Get FileUsages via Id
func GetFileUsagesViaId(offset int, limit int, Id_ string, field string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_FileUsage)
	return _FileUsage, err
}

// GetFileUsagesViaCount Get FileUsages via Count
func GetFileUsagesViaCount(offset int, limit int, Count_ int, field string) (*[]*FileUsage, error) {
	var _FileUsage = new([]*FileUsage)
	err := Engine.Table("file_usage").Where("count = ?", Count_).Limit(limit, offset).Desc(field).Find(_FileUsage)
	return _FileUsage, err
}

// HasFileUsageViaFid Has FileUsage via Fid
func HasFileUsageViaFid(iFid int64) bool {
	if has, err := Engine.Where("fid = ?", iFid).Get(new(FileUsage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileUsageViaModule Has FileUsage via Module
func HasFileUsageViaModule(iModule string) bool {
	if has, err := Engine.Where("module = ?", iModule).Get(new(FileUsage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileUsageViaType Has FileUsage via Type
func HasFileUsageViaType(iType string) bool {
	if has, err := Engine.Where("type = ?", iType).Get(new(FileUsage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileUsageViaId Has FileUsage via Id
func HasFileUsageViaId(iId string) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(FileUsage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileUsageViaCount Has FileUsage via Count
func HasFileUsageViaCount(iCount int) bool {
	if has, err := Engine.Where("count = ?", iCount).Get(new(FileUsage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetFileUsageViaFid Get FileUsage via Fid
func GetFileUsageViaFid(iFid int64) (*FileUsage, error) {
	var _FileUsage = &FileUsage{Fid: iFid}
	has, err := Engine.Get(_FileUsage)
	if has {
		return _FileUsage, err
	} else {
		return nil, err
	}
}

// GetFileUsageViaModule Get FileUsage via Module
func GetFileUsageViaModule(iModule string) (*FileUsage, error) {
	var _FileUsage = &FileUsage{Module: iModule}
	has, err := Engine.Get(_FileUsage)
	if has {
		return _FileUsage, err
	} else {
		return nil, err
	}
}

// GetFileUsageViaType Get FileUsage via Type
func GetFileUsageViaType(iType string) (*FileUsage, error) {
	var _FileUsage = &FileUsage{Type: iType}
	has, err := Engine.Get(_FileUsage)
	if has {
		return _FileUsage, err
	} else {
		return nil, err
	}
}

// GetFileUsageViaId Get FileUsage via Id
func GetFileUsageViaId(iId string) (*FileUsage, error) {
	var _FileUsage = &FileUsage{Id: iId}
	has, err := Engine.Get(_FileUsage)
	if has {
		return _FileUsage, err
	} else {
		return nil, err
	}
}

// GetFileUsageViaCount Get FileUsage via Count
func GetFileUsageViaCount(iCount int) (*FileUsage, error) {
	var _FileUsage = &FileUsage{Count: iCount}
	has, err := Engine.Get(_FileUsage)
	if has {
		return _FileUsage, err
	} else {
		return nil, err
	}
}

// SetFileUsageViaFid Set FileUsage via Fid
func SetFileUsageViaFid(iFid int64, file_usage *FileUsage) (int64, error) {
	file_usage.Fid = iFid
	return Engine.Insert(file_usage)
}

// SetFileUsageViaModule Set FileUsage via Module
func SetFileUsageViaModule(iModule string, file_usage *FileUsage) (int64, error) {
	file_usage.Module = iModule
	return Engine.Insert(file_usage)
}

// SetFileUsageViaType Set FileUsage via Type
func SetFileUsageViaType(iType string, file_usage *FileUsage) (int64, error) {
	file_usage.Type = iType
	return Engine.Insert(file_usage)
}

// SetFileUsageViaId Set FileUsage via Id
func SetFileUsageViaId(iId string, file_usage *FileUsage) (int64, error) {
	file_usage.Id = iId
	return Engine.Insert(file_usage)
}

// SetFileUsageViaCount Set FileUsage via Count
func SetFileUsageViaCount(iCount int, file_usage *FileUsage) (int64, error) {
	file_usage.Count = iCount
	return Engine.Insert(file_usage)
}

// AddFileUsage Add FileUsage via all columns
func AddFileUsage(iFid int64, iModule string, iType string, iId string, iCount int) error {
	_FileUsage := &FileUsage{Fid: iFid, Module: iModule, Type: iType, Id: iId, Count: iCount}
	if _, err := Engine.Insert(_FileUsage); err != nil {
		return err
	} else {
		return nil
	}
}

// PostFileUsage Post FileUsage via iFileUsage
func PostFileUsage(iFileUsage *FileUsage) (int64, error) {
	_, err := Engine.Insert(iFileUsage)
	return iFileUsage.Fid, err
}

// PutFileUsage Put FileUsage
func PutFileUsage(iFileUsage *FileUsage) (int64, error) {
	_, err := Engine.Id(iFileUsage.Fid).Update(iFileUsage)
	return iFileUsage.Fid, err
}

// PutFileUsageViaFid Put FileUsage via Fid
func PutFileUsageViaFid(Fid_ int64, iFileUsage *FileUsage) (int64, error) {
	row, err := Engine.Update(iFileUsage, &FileUsage{Fid: Fid_})
	return row, err
}

// PutFileUsageViaModule Put FileUsage via Module
func PutFileUsageViaModule(Module_ string, iFileUsage *FileUsage) (int64, error) {
	row, err := Engine.Update(iFileUsage, &FileUsage{Module: Module_})
	return row, err
}

// PutFileUsageViaType Put FileUsage via Type
func PutFileUsageViaType(Type_ string, iFileUsage *FileUsage) (int64, error) {
	row, err := Engine.Update(iFileUsage, &FileUsage{Type: Type_})
	return row, err
}

// PutFileUsageViaId Put FileUsage via Id
func PutFileUsageViaId(Id_ string, iFileUsage *FileUsage) (int64, error) {
	row, err := Engine.Update(iFileUsage, &FileUsage{Id: Id_})
	return row, err
}

// PutFileUsageViaCount Put FileUsage via Count
func PutFileUsageViaCount(Count_ int, iFileUsage *FileUsage) (int64, error) {
	row, err := Engine.Update(iFileUsage, &FileUsage{Count: Count_})
	return row, err
}

// UpdateFileUsageViaFid via map[string]interface{}{}
func UpdateFileUsageViaFid(iFid int64, iFileUsageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileUsage)).Where("fid = ?", iFid).Update(iFileUsageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileUsageViaModule via map[string]interface{}{}
func UpdateFileUsageViaModule(iModule string, iFileUsageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileUsage)).Where("module = ?", iModule).Update(iFileUsageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileUsageViaType via map[string]interface{}{}
func UpdateFileUsageViaType(iType string, iFileUsageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileUsage)).Where("type = ?", iType).Update(iFileUsageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileUsageViaId via map[string]interface{}{}
func UpdateFileUsageViaId(iId string, iFileUsageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileUsage)).Where("id = ?", iId).Update(iFileUsageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileUsageViaCount via map[string]interface{}{}
func UpdateFileUsageViaCount(iCount int, iFileUsageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileUsage)).Where("count = ?", iCount).Update(iFileUsageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteFileUsage Delete FileUsage
func DeleteFileUsage(iFid int64) (int64, error) {
	row, err := Engine.Id(iFid).Delete(new(FileUsage))
	return row, err
}

// DeleteFileUsageViaFid Delete FileUsage via Fid
func DeleteFileUsageViaFid(iFid int64) (err error) {
	var has bool
	var _FileUsage = &FileUsage{Fid: iFid}
	if has, err = Engine.Get(_FileUsage); (has == true) && (err == nil) {
		if row, err := Engine.Where("fid = ?", iFid).Delete(new(FileUsage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileUsageViaModule Delete FileUsage via Module
func DeleteFileUsageViaModule(iModule string) (err error) {
	var has bool
	var _FileUsage = &FileUsage{Module: iModule}
	if has, err = Engine.Get(_FileUsage); (has == true) && (err == nil) {
		if row, err := Engine.Where("module = ?", iModule).Delete(new(FileUsage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileUsageViaType Delete FileUsage via Type
func DeleteFileUsageViaType(iType string) (err error) {
	var has bool
	var _FileUsage = &FileUsage{Type: iType}
	if has, err = Engine.Get(_FileUsage); (has == true) && (err == nil) {
		if row, err := Engine.Where("type = ?", iType).Delete(new(FileUsage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileUsageViaId Delete FileUsage via Id
func DeleteFileUsageViaId(iId string) (err error) {
	var has bool
	var _FileUsage = &FileUsage{Id: iId}
	if has, err = Engine.Get(_FileUsage); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(FileUsage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileUsageViaCount Delete FileUsage via Count
func DeleteFileUsageViaCount(iCount int) (err error) {
	var has bool
	var _FileUsage = &FileUsage{Count: iCount}
	if has, err = Engine.Get(_FileUsage); (has == true) && (err == nil) {
		if row, err := Engine.Where("count = ?", iCount).Delete(new(FileUsage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
