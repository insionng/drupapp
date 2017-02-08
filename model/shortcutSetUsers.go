package model

type ShortcutSetUsers struct {
	Uid     int64  `xorm:"not null pk default 0 INT(10)"`
	SetName string `xorm:"not null default '' index VARCHAR(32)"`
}

// GetShortcutSetUsersesCount ShortcutSetUserss' Count
func GetShortcutSetUsersesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&ShortcutSetUsers{})
	return total, err
}

// GetShortcutSetUsersCountViaUid Get ShortcutSetUsers via Uid
func GetShortcutSetUsersCountViaUid(iUid int64) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&ShortcutSetUsers{Uid: iUid})
	return n
}

// GetShortcutSetUsersCountViaSetName Get ShortcutSetUsers via SetName
func GetShortcutSetUsersCountViaSetName(iSetName string) int64 {
	n, _ := Engine.Where("set_name = ?", iSetName).Count(&ShortcutSetUsers{SetName: iSetName})
	return n
}

// GetShortcutSetUsersesByUidAndSetName Get ShortcutSetUserses via UidAndSetName
func GetShortcutSetUsersesByUidAndSetName(offset int, limit int, Uid_ int64, SetName_ string) (*[]*ShortcutSetUsers, error) {
	var _ShortcutSetUsers = new([]*ShortcutSetUsers)
	err := Engine.Table("shortcut_set_users").Where("uid = ? and set_name = ?", Uid_, SetName_).Limit(limit, offset).Find(_ShortcutSetUsers)
	return _ShortcutSetUsers, err
}

// GetShortcutSetUserses Get ShortcutSetUserses via field
func GetShortcutSetUserses(offset int, limit int, field string) (*[]*ShortcutSetUsers, error) {
	var _ShortcutSetUsers = new([]*ShortcutSetUsers)
	err := Engine.Table("shortcut_set_users").Limit(limit, offset).Desc(field).Find(_ShortcutSetUsers)
	return _ShortcutSetUsers, err
}

// GetShortcutSetUsersesViaUid Get ShortcutSetUserss via Uid
func GetShortcutSetUsersesViaUid(offset int, limit int, Uid_ int64, field string) (*[]*ShortcutSetUsers, error) {
	var _ShortcutSetUsers = new([]*ShortcutSetUsers)
	err := Engine.Table("shortcut_set_users").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_ShortcutSetUsers)
	return _ShortcutSetUsers, err
}

// GetShortcutSetUsersesViaSetName Get ShortcutSetUserss via SetName
func GetShortcutSetUsersesViaSetName(offset int, limit int, SetName_ string, field string) (*[]*ShortcutSetUsers, error) {
	var _ShortcutSetUsers = new([]*ShortcutSetUsers)
	err := Engine.Table("shortcut_set_users").Where("set_name = ?", SetName_).Limit(limit, offset).Desc(field).Find(_ShortcutSetUsers)
	return _ShortcutSetUsers, err
}

// HasShortcutSetUsersViaUid Has ShortcutSetUsers via Uid
func HasShortcutSetUsersViaUid(iUid int64) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(ShortcutSetUsers)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutSetUsersViaSetName Has ShortcutSetUsers via SetName
func HasShortcutSetUsersViaSetName(iSetName string) bool {
	if has, err := Engine.Where("set_name = ?", iSetName).Get(new(ShortcutSetUsers)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetShortcutSetUsersViaUid Get ShortcutSetUsers via Uid
func GetShortcutSetUsersViaUid(iUid int64) (*ShortcutSetUsers, error) {
	var _ShortcutSetUsers = &ShortcutSetUsers{Uid: iUid}
	has, err := Engine.Get(_ShortcutSetUsers)
	if has {
		return _ShortcutSetUsers, err
	} else {
		return nil, err
	}
}

// GetShortcutSetUsersViaSetName Get ShortcutSetUsers via SetName
func GetShortcutSetUsersViaSetName(iSetName string) (*ShortcutSetUsers, error) {
	var _ShortcutSetUsers = &ShortcutSetUsers{SetName: iSetName}
	has, err := Engine.Get(_ShortcutSetUsers)
	if has {
		return _ShortcutSetUsers, err
	} else {
		return nil, err
	}
}

// SetShortcutSetUsersViaUid Set ShortcutSetUsers via Uid
func SetShortcutSetUsersViaUid(iUid int64, shortcut_set_users *ShortcutSetUsers) (int64, error) {
	shortcut_set_users.Uid = iUid
	return Engine.Insert(shortcut_set_users)
}

// SetShortcutSetUsersViaSetName Set ShortcutSetUsers via SetName
func SetShortcutSetUsersViaSetName(iSetName string, shortcut_set_users *ShortcutSetUsers) (int64, error) {
	shortcut_set_users.SetName = iSetName
	return Engine.Insert(shortcut_set_users)
}

// AddShortcutSetUsers Add ShortcutSetUsers via all columns
func AddShortcutSetUsers(iUid int64, iSetName string) error {
	_ShortcutSetUsers := &ShortcutSetUsers{Uid: iUid, SetName: iSetName}
	if _, err := Engine.Insert(_ShortcutSetUsers); err != nil {
		return err
	} else {
		return nil
	}
}

// PostShortcutSetUsers Post ShortcutSetUsers via iShortcutSetUsers
func PostShortcutSetUsers(iShortcutSetUsers *ShortcutSetUsers) (int64, error) {
	_, err := Engine.Insert(iShortcutSetUsers)
	return iShortcutSetUsers.Uid, err
}

// PutShortcutSetUsers Put ShortcutSetUsers
func PutShortcutSetUsers(iShortcutSetUsers *ShortcutSetUsers) (int64, error) {
	_, err := Engine.Id(iShortcutSetUsers.Uid).Update(iShortcutSetUsers)
	return iShortcutSetUsers.Uid, err
}

// PutShortcutSetUsersViaUid Put ShortcutSetUsers via Uid
func PutShortcutSetUsersViaUid(Uid_ int64, iShortcutSetUsers *ShortcutSetUsers) (int64, error) {
	row, err := Engine.Update(iShortcutSetUsers, &ShortcutSetUsers{Uid: Uid_})
	return row, err
}

// PutShortcutSetUsersViaSetName Put ShortcutSetUsers via SetName
func PutShortcutSetUsersViaSetName(SetName_ string, iShortcutSetUsers *ShortcutSetUsers) (int64, error) {
	row, err := Engine.Update(iShortcutSetUsers, &ShortcutSetUsers{SetName: SetName_})
	return row, err
}

// UpdateShortcutSetUsersViaUid via map[string]interface{}{}
func UpdateShortcutSetUsersViaUid(iUid int64, iShortcutSetUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutSetUsers)).Where("uid = ?", iUid).Update(iShortcutSetUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutSetUsersViaSetName via map[string]interface{}{}
func UpdateShortcutSetUsersViaSetName(iSetName string, iShortcutSetUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutSetUsers)).Where("set_name = ?", iSetName).Update(iShortcutSetUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteShortcutSetUsers Delete ShortcutSetUsers
func DeleteShortcutSetUsers(iUid int64) (int64, error) {
	row, err := Engine.Id(iUid).Delete(new(ShortcutSetUsers))
	return row, err
}

// DeleteShortcutSetUsersViaUid Delete ShortcutSetUsers via Uid
func DeleteShortcutSetUsersViaUid(iUid int64) (err error) {
	var has bool
	var _ShortcutSetUsers = &ShortcutSetUsers{Uid: iUid}
	if has, err = Engine.Get(_ShortcutSetUsers); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(ShortcutSetUsers)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutSetUsersViaSetName Delete ShortcutSetUsers via SetName
func DeleteShortcutSetUsersViaSetName(iSetName string) (err error) {
	var has bool
	var _ShortcutSetUsers = &ShortcutSetUsers{SetName: iSetName}
	if has, err = Engine.Get(_ShortcutSetUsers); (has == true) && (err == nil) {
		if row, err := Engine.Where("set_name = ?", iSetName).Delete(new(ShortcutSetUsers)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
