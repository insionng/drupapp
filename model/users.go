package model

type Users struct {
	Uid      int64  `xorm:"not null pk INT(10)"`
	Uuid     string `xorm:"not null unique VARCHAR(128)"`
	Langcode string `xorm:"not null VARCHAR(12)"`
}

// GetUsersesCount Userss' Count
func GetUsersesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Users{})
	return total, err
}

// GetUsersCountViaUid Get Users via Uid
func GetUsersCountViaUid(iUid int64) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&Users{Uid: iUid})
	return n
}

// GetUsersCountViaUuid Get Users via Uuid
func GetUsersCountViaUuid(iUuid string) int64 {
	n, _ := Engine.Where("uuid = ?", iUuid).Count(&Users{Uuid: iUuid})
	return n
}

// GetUsersCountViaLangcode Get Users via Langcode
func GetUsersCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Users{Langcode: iLangcode})
	return n
}

// GetUsersesByUidAndUuidAndLangcode Get Userses via UidAndUuidAndLangcode
func GetUsersesByUidAndUuidAndLangcode(offset int, limit int, Uid_ int64, Uuid_ string, Langcode_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("uid = ? and uuid = ? and langcode = ?", Uid_, Uuid_, Langcode_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUidAndUuid Get Userses via UidAndUuid
func GetUsersesByUidAndUuid(offset int, limit int, Uid_ int64, Uuid_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("uid = ? and uuid = ?", Uid_, Uuid_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUidAndLangcode Get Userses via UidAndLangcode
func GetUsersesByUidAndLangcode(offset int, limit int, Uid_ int64, Langcode_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("uid = ? and langcode = ?", Uid_, Langcode_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUuidAndLangcode Get Userses via UuidAndLangcode
func GetUsersesByUuidAndLangcode(offset int, limit int, Uuid_ string, Langcode_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("uuid = ? and langcode = ?", Uuid_, Langcode_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUserses Get Userses via field
func GetUserses(offset int, limit int, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUid Get Userss via Uid
func GetUsersesViaUid(offset int, limit int, Uid_ int64, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUuid Get Userss via Uuid
func GetUsersesViaUuid(offset int, limit int, Uuid_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("uuid = ?", Uuid_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaLangcode Get Userss via Langcode
func GetUsersesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// HasUsersViaUid Has Users via Uid
func HasUsersViaUid(iUid int64) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUuid Has Users via Uuid
func HasUsersViaUuid(iUuid string) bool {
	if has, err := Engine.Where("uuid = ?", iUuid).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaLangcode Has Users via Langcode
func HasUsersViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUsersViaUid Get Users via Uid
func GetUsersViaUid(iUid int64) (*Users, error) {
	var _Users = &Users{Uid: iUid}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUuid Get Users via Uuid
func GetUsersViaUuid(iUuid string) (*Users, error) {
	var _Users = &Users{Uuid: iUuid}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaLangcode Get Users via Langcode
func GetUsersViaLangcode(iLangcode string) (*Users, error) {
	var _Users = &Users{Langcode: iLangcode}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// SetUsersViaUid Set Users via Uid
func SetUsersViaUid(iUid int64, users *Users) (int64, error) {
	users.Uid = iUid
	return Engine.Insert(users)
}

// SetUsersViaUuid Set Users via Uuid
func SetUsersViaUuid(iUuid string, users *Users) (int64, error) {
	users.Uuid = iUuid
	return Engine.Insert(users)
}

// SetUsersViaLangcode Set Users via Langcode
func SetUsersViaLangcode(iLangcode string, users *Users) (int64, error) {
	users.Langcode = iLangcode
	return Engine.Insert(users)
}

// AddUsers Add Users via all columns
func AddUsers(iUid int64, iUuid string, iLangcode string) error {
	_Users := &Users{Uid: iUid, Uuid: iUuid, Langcode: iLangcode}
	if _, err := Engine.Insert(_Users); err != nil {
		return err
	} else {
		return nil
	}
}

// PostUsers Post Users via iUsers
func PostUsers(iUsers *Users) (int64, error) {
	_, err := Engine.Insert(iUsers)
	return iUsers.Uid, err
}

// PutUsers Put Users
func PutUsers(iUsers *Users) (int64, error) {
	_, err := Engine.Id(iUsers.Uid).Update(iUsers)
	return iUsers.Uid, err
}

// PutUsersViaUid Put Users via Uid
func PutUsersViaUid(Uid_ int64, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{Uid: Uid_})
	return row, err
}

// PutUsersViaUuid Put Users via Uuid
func PutUsersViaUuid(Uuid_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{Uuid: Uuid_})
	return row, err
}

// PutUsersViaLangcode Put Users via Langcode
func PutUsersViaLangcode(Langcode_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{Langcode: Langcode_})
	return row, err
}

// UpdateUsersViaUid via map[string]interface{}{}
func UpdateUsersViaUid(iUid int64, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("uid = ?", iUid).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUuid via map[string]interface{}{}
func UpdateUsersViaUuid(iUuid string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("uuid = ?", iUuid).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaLangcode via map[string]interface{}{}
func UpdateUsersViaLangcode(iLangcode string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("langcode = ?", iLangcode).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteUsers Delete Users
func DeleteUsers(iUid int64) (int64, error) {
	row, err := Engine.Id(iUid).Delete(new(Users))
	return row, err
}

// DeleteUsersViaUid Delete Users via Uid
func DeleteUsersViaUid(iUid int64) (err error) {
	var has bool
	var _Users = &Users{Uid: iUid}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUuid Delete Users via Uuid
func DeleteUsersViaUuid(iUuid string) (err error) {
	var has bool
	var _Users = &Users{Uuid: iUuid}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("uuid = ?", iUuid).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaLangcode Delete Users via Langcode
func DeleteUsersViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Users = &Users{Langcode: iLangcode}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
