package model

type UsersData struct {
	Uid        int64  `xorm:"not null pk default 0 INT(10)"`
	Module     string `xorm:"not null pk default '' index VARCHAR(50)"`
	Name       string `xorm:"not null pk default '' index VARCHAR(128)"`
	Value      []byte `xorm:"LONGBLOB"`
	Serialized int    `xorm:"default 0 TINYINT(3)"`
}

// GetUsersDatasCount UsersDatas' Count
func GetUsersDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&UsersData{})
	return total, err
}

// GetUsersDataCountViaUid Get UsersData via Uid
func GetUsersDataCountViaUid(iUid int64) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&UsersData{Uid: iUid})
	return n
}

// GetUsersDataCountViaModule Get UsersData via Module
func GetUsersDataCountViaModule(iModule string) int64 {
	n, _ := Engine.Where("module = ?", iModule).Count(&UsersData{Module: iModule})
	return n
}

// GetUsersDataCountViaName Get UsersData via Name
func GetUsersDataCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&UsersData{Name: iName})
	return n
}

// GetUsersDataCountViaValue Get UsersData via Value
func GetUsersDataCountViaValue(iValue []byte) int64 {
	n, _ := Engine.Where("value = ?", iValue).Count(&UsersData{Value: iValue})
	return n
}

// GetUsersDataCountViaSerialized Get UsersData via Serialized
func GetUsersDataCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&UsersData{Serialized: iSerialized})
	return n
}

// GetUsersDatasByUidAndModuleAndName Get UsersDatas via UidAndModuleAndName
func GetUsersDatasByUidAndModuleAndName(offset int, limit int, Uid_ int64, Module_ string, Name_ string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and module = ? and name = ?", Uid_, Module_, Name_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndModuleAndValue Get UsersDatas via UidAndModuleAndValue
func GetUsersDatasByUidAndModuleAndValue(offset int, limit int, Uid_ int64, Module_ string, Value_ []byte) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and module = ? and value = ?", Uid_, Module_, Value_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndModuleAndSerialized Get UsersDatas via UidAndModuleAndSerialized
func GetUsersDatasByUidAndModuleAndSerialized(offset int, limit int, Uid_ int64, Module_ string, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and module = ? and serialized = ?", Uid_, Module_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndNameAndValue Get UsersDatas via UidAndNameAndValue
func GetUsersDatasByUidAndNameAndValue(offset int, limit int, Uid_ int64, Name_ string, Value_ []byte) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and name = ? and value = ?", Uid_, Name_, Value_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndNameAndSerialized Get UsersDatas via UidAndNameAndSerialized
func GetUsersDatasByUidAndNameAndSerialized(offset int, limit int, Uid_ int64, Name_ string, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and name = ? and serialized = ?", Uid_, Name_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndValueAndSerialized Get UsersDatas via UidAndValueAndSerialized
func GetUsersDatasByUidAndValueAndSerialized(offset int, limit int, Uid_ int64, Value_ []byte, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and value = ? and serialized = ?", Uid_, Value_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByModuleAndNameAndValue Get UsersDatas via ModuleAndNameAndValue
func GetUsersDatasByModuleAndNameAndValue(offset int, limit int, Module_ string, Name_ string, Value_ []byte) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("module = ? and name = ? and value = ?", Module_, Name_, Value_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByModuleAndNameAndSerialized Get UsersDatas via ModuleAndNameAndSerialized
func GetUsersDatasByModuleAndNameAndSerialized(offset int, limit int, Module_ string, Name_ string, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("module = ? and name = ? and serialized = ?", Module_, Name_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByModuleAndValueAndSerialized Get UsersDatas via ModuleAndValueAndSerialized
func GetUsersDatasByModuleAndValueAndSerialized(offset int, limit int, Module_ string, Value_ []byte, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("module = ? and value = ? and serialized = ?", Module_, Value_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByNameAndValueAndSerialized Get UsersDatas via NameAndValueAndSerialized
func GetUsersDatasByNameAndValueAndSerialized(offset int, limit int, Name_ string, Value_ []byte, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("name = ? and value = ? and serialized = ?", Name_, Value_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndModule Get UsersDatas via UidAndModule
func GetUsersDatasByUidAndModule(offset int, limit int, Uid_ int64, Module_ string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and module = ?", Uid_, Module_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndName Get UsersDatas via UidAndName
func GetUsersDatasByUidAndName(offset int, limit int, Uid_ int64, Name_ string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and name = ?", Uid_, Name_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndValue Get UsersDatas via UidAndValue
func GetUsersDatasByUidAndValue(offset int, limit int, Uid_ int64, Value_ []byte) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and value = ?", Uid_, Value_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByUidAndSerialized Get UsersDatas via UidAndSerialized
func GetUsersDatasByUidAndSerialized(offset int, limit int, Uid_ int64, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ? and serialized = ?", Uid_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByModuleAndName Get UsersDatas via ModuleAndName
func GetUsersDatasByModuleAndName(offset int, limit int, Module_ string, Name_ string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("module = ? and name = ?", Module_, Name_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByModuleAndValue Get UsersDatas via ModuleAndValue
func GetUsersDatasByModuleAndValue(offset int, limit int, Module_ string, Value_ []byte) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("module = ? and value = ?", Module_, Value_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByModuleAndSerialized Get UsersDatas via ModuleAndSerialized
func GetUsersDatasByModuleAndSerialized(offset int, limit int, Module_ string, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("module = ? and serialized = ?", Module_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByNameAndValue Get UsersDatas via NameAndValue
func GetUsersDatasByNameAndValue(offset int, limit int, Name_ string, Value_ []byte) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("name = ? and value = ?", Name_, Value_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByNameAndSerialized Get UsersDatas via NameAndSerialized
func GetUsersDatasByNameAndSerialized(offset int, limit int, Name_ string, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("name = ? and serialized = ?", Name_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasByValueAndSerialized Get UsersDatas via ValueAndSerialized
func GetUsersDatasByValueAndSerialized(offset int, limit int, Value_ []byte, Serialized_ int) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("value = ? and serialized = ?", Value_, Serialized_).Limit(limit, offset).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatas Get UsersDatas via field
func GetUsersDatas(offset int, limit int, field string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Limit(limit, offset).Desc(field).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasViaUid Get UsersDatas via Uid
func GetUsersDatasViaUid(offset int, limit int, Uid_ int64, field string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasViaModule Get UsersDatas via Module
func GetUsersDatasViaModule(offset int, limit int, Module_ string, field string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("module = ?", Module_).Limit(limit, offset).Desc(field).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasViaName Get UsersDatas via Name
func GetUsersDatasViaName(offset int, limit int, Name_ string, field string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasViaValue Get UsersDatas via Value
func GetUsersDatasViaValue(offset int, limit int, Value_ []byte, field string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("value = ?", Value_).Limit(limit, offset).Desc(field).Find(_UsersData)
	return _UsersData, err
}

// GetUsersDatasViaSerialized Get UsersDatas via Serialized
func GetUsersDatasViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*UsersData, error) {
	var _UsersData = new([]*UsersData)
	err := Engine.Table("users_data").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_UsersData)
	return _UsersData, err
}

// HasUsersDataViaUid Has UsersData via Uid
func HasUsersDataViaUid(iUid int64) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(UsersData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersDataViaModule Has UsersData via Module
func HasUsersDataViaModule(iModule string) bool {
	if has, err := Engine.Where("module = ?", iModule).Get(new(UsersData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersDataViaName Has UsersData via Name
func HasUsersDataViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(UsersData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersDataViaValue Has UsersData via Value
func HasUsersDataViaValue(iValue []byte) bool {
	if has, err := Engine.Where("value = ?", iValue).Get(new(UsersData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersDataViaSerialized Has UsersData via Serialized
func HasUsersDataViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(UsersData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUsersDataViaUid Get UsersData via Uid
func GetUsersDataViaUid(iUid int64) (*UsersData, error) {
	var _UsersData = &UsersData{Uid: iUid}
	has, err := Engine.Get(_UsersData)
	if has {
		return _UsersData, err
	} else {
		return nil, err
	}
}

// GetUsersDataViaModule Get UsersData via Module
func GetUsersDataViaModule(iModule string) (*UsersData, error) {
	var _UsersData = &UsersData{Module: iModule}
	has, err := Engine.Get(_UsersData)
	if has {
		return _UsersData, err
	} else {
		return nil, err
	}
}

// GetUsersDataViaName Get UsersData via Name
func GetUsersDataViaName(iName string) (*UsersData, error) {
	var _UsersData = &UsersData{Name: iName}
	has, err := Engine.Get(_UsersData)
	if has {
		return _UsersData, err
	} else {
		return nil, err
	}
}

// GetUsersDataViaValue Get UsersData via Value
func GetUsersDataViaValue(iValue []byte) (*UsersData, error) {
	var _UsersData = &UsersData{Value: iValue}
	has, err := Engine.Get(_UsersData)
	if has {
		return _UsersData, err
	} else {
		return nil, err
	}
}

// GetUsersDataViaSerialized Get UsersData via Serialized
func GetUsersDataViaSerialized(iSerialized int) (*UsersData, error) {
	var _UsersData = &UsersData{Serialized: iSerialized}
	has, err := Engine.Get(_UsersData)
	if has {
		return _UsersData, err
	} else {
		return nil, err
	}
}

// SetUsersDataViaUid Set UsersData via Uid
func SetUsersDataViaUid(iUid int64, users_data *UsersData) (int64, error) {
	users_data.Uid = iUid
	return Engine.Insert(users_data)
}

// SetUsersDataViaModule Set UsersData via Module
func SetUsersDataViaModule(iModule string, users_data *UsersData) (int64, error) {
	users_data.Module = iModule
	return Engine.Insert(users_data)
}

// SetUsersDataViaName Set UsersData via Name
func SetUsersDataViaName(iName string, users_data *UsersData) (int64, error) {
	users_data.Name = iName
	return Engine.Insert(users_data)
}

// SetUsersDataViaValue Set UsersData via Value
func SetUsersDataViaValue(iValue []byte, users_data *UsersData) (int64, error) {
	users_data.Value = iValue
	return Engine.Insert(users_data)
}

// SetUsersDataViaSerialized Set UsersData via Serialized
func SetUsersDataViaSerialized(iSerialized int, users_data *UsersData) (int64, error) {
	users_data.Serialized = iSerialized
	return Engine.Insert(users_data)
}

// AddUsersData Add UsersData via all columns
func AddUsersData(iUid int64, iModule string, iName string, iValue []byte, iSerialized int) error {
	_UsersData := &UsersData{Uid: iUid, Module: iModule, Name: iName, Value: iValue, Serialized: iSerialized}
	if _, err := Engine.Insert(_UsersData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostUsersData Post UsersData via iUsersData
func PostUsersData(iUsersData *UsersData) (int64, error) {
	_, err := Engine.Insert(iUsersData)
	return iUsersData.Uid, err
}

// PutUsersData Put UsersData
func PutUsersData(iUsersData *UsersData) (int64, error) {
	_, err := Engine.Id(iUsersData.Uid).Update(iUsersData)
	return iUsersData.Uid, err
}

// PutUsersDataViaUid Put UsersData via Uid
func PutUsersDataViaUid(Uid_ int64, iUsersData *UsersData) (int64, error) {
	row, err := Engine.Update(iUsersData, &UsersData{Uid: Uid_})
	return row, err
}

// PutUsersDataViaModule Put UsersData via Module
func PutUsersDataViaModule(Module_ string, iUsersData *UsersData) (int64, error) {
	row, err := Engine.Update(iUsersData, &UsersData{Module: Module_})
	return row, err
}

// PutUsersDataViaName Put UsersData via Name
func PutUsersDataViaName(Name_ string, iUsersData *UsersData) (int64, error) {
	row, err := Engine.Update(iUsersData, &UsersData{Name: Name_})
	return row, err
}

// PutUsersDataViaValue Put UsersData via Value
func PutUsersDataViaValue(Value_ []byte, iUsersData *UsersData) (int64, error) {
	row, err := Engine.Update(iUsersData, &UsersData{Value: Value_})
	return row, err
}

// PutUsersDataViaSerialized Put UsersData via Serialized
func PutUsersDataViaSerialized(Serialized_ int, iUsersData *UsersData) (int64, error) {
	row, err := Engine.Update(iUsersData, &UsersData{Serialized: Serialized_})
	return row, err
}

// UpdateUsersDataViaUid via map[string]interface{}{}
func UpdateUsersDataViaUid(iUid int64, iUsersDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersData)).Where("uid = ?", iUid).Update(iUsersDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersDataViaModule via map[string]interface{}{}
func UpdateUsersDataViaModule(iModule string, iUsersDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersData)).Where("module = ?", iModule).Update(iUsersDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersDataViaName via map[string]interface{}{}
func UpdateUsersDataViaName(iName string, iUsersDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersData)).Where("name = ?", iName).Update(iUsersDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersDataViaValue via map[string]interface{}{}
func UpdateUsersDataViaValue(iValue []byte, iUsersDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersData)).Where("value = ?", iValue).Update(iUsersDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersDataViaSerialized via map[string]interface{}{}
func UpdateUsersDataViaSerialized(iSerialized int, iUsersDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersData)).Where("serialized = ?", iSerialized).Update(iUsersDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteUsersData Delete UsersData
func DeleteUsersData(iUid int64) (int64, error) {
	row, err := Engine.Id(iUid).Delete(new(UsersData))
	return row, err
}

// DeleteUsersDataViaUid Delete UsersData via Uid
func DeleteUsersDataViaUid(iUid int64) (err error) {
	var has bool
	var _UsersData = &UsersData{Uid: iUid}
	if has, err = Engine.Get(_UsersData); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(UsersData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersDataViaModule Delete UsersData via Module
func DeleteUsersDataViaModule(iModule string) (err error) {
	var has bool
	var _UsersData = &UsersData{Module: iModule}
	if has, err = Engine.Get(_UsersData); (has == true) && (err == nil) {
		if row, err := Engine.Where("module = ?", iModule).Delete(new(UsersData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersDataViaName Delete UsersData via Name
func DeleteUsersDataViaName(iName string) (err error) {
	var has bool
	var _UsersData = &UsersData{Name: iName}
	if has, err = Engine.Get(_UsersData); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(UsersData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersDataViaValue Delete UsersData via Value
func DeleteUsersDataViaValue(iValue []byte) (err error) {
	var has bool
	var _UsersData = &UsersData{Value: iValue}
	if has, err = Engine.Get(_UsersData); (has == true) && (err == nil) {
		if row, err := Engine.Where("value = ?", iValue).Delete(new(UsersData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersDataViaSerialized Delete UsersData via Serialized
func DeleteUsersDataViaSerialized(iSerialized int) (err error) {
	var has bool
	var _UsersData = &UsersData{Serialized: iSerialized}
	if has, err = Engine.Get(_UsersData); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(UsersData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
