package model

type Shortcut struct {
	Id          int64  `xorm:"not null pk autoincr INT(10)"`
	ShortcutSet string `xorm:"not null index VARCHAR(32)"`
	Uuid        string `xorm:"not null unique VARCHAR(128)"`
	Langcode    string `xorm:"not null VARCHAR(12)"`
}

// GetShortcutsCount Shortcuts' Count
func GetShortcutsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Shortcut{})
	return total, err
}

// GetShortcutCountViaId Get Shortcut via Id
func GetShortcutCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&Shortcut{Id: iId})
	return n
}

// GetShortcutCountViaShortcutSet Get Shortcut via ShortcutSet
func GetShortcutCountViaShortcutSet(iShortcutSet string) int64 {
	n, _ := Engine.Where("shortcut_set = ?", iShortcutSet).Count(&Shortcut{ShortcutSet: iShortcutSet})
	return n
}

// GetShortcutCountViaUuid Get Shortcut via Uuid
func GetShortcutCountViaUuid(iUuid string) int64 {
	n, _ := Engine.Where("uuid = ?", iUuid).Count(&Shortcut{Uuid: iUuid})
	return n
}

// GetShortcutCountViaLangcode Get Shortcut via Langcode
func GetShortcutCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Shortcut{Langcode: iLangcode})
	return n
}

// GetShortcutsByIdAndShortcutSetAndUuid Get Shortcuts via IdAndShortcutSetAndUuid
func GetShortcutsByIdAndShortcutSetAndUuid(offset int, limit int, Id_ int64, ShortcutSet_ string, Uuid_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("id = ? and shortcut_set = ? and uuid = ?", Id_, ShortcutSet_, Uuid_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByIdAndShortcutSetAndLangcode Get Shortcuts via IdAndShortcutSetAndLangcode
func GetShortcutsByIdAndShortcutSetAndLangcode(offset int, limit int, Id_ int64, ShortcutSet_ string, Langcode_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("id = ? and shortcut_set = ? and langcode = ?", Id_, ShortcutSet_, Langcode_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByIdAndUuidAndLangcode Get Shortcuts via IdAndUuidAndLangcode
func GetShortcutsByIdAndUuidAndLangcode(offset int, limit int, Id_ int64, Uuid_ string, Langcode_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("id = ? and uuid = ? and langcode = ?", Id_, Uuid_, Langcode_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByShortcutSetAndUuidAndLangcode Get Shortcuts via ShortcutSetAndUuidAndLangcode
func GetShortcutsByShortcutSetAndUuidAndLangcode(offset int, limit int, ShortcutSet_ string, Uuid_ string, Langcode_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("shortcut_set = ? and uuid = ? and langcode = ?", ShortcutSet_, Uuid_, Langcode_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByIdAndShortcutSet Get Shortcuts via IdAndShortcutSet
func GetShortcutsByIdAndShortcutSet(offset int, limit int, Id_ int64, ShortcutSet_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("id = ? and shortcut_set = ?", Id_, ShortcutSet_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByIdAndUuid Get Shortcuts via IdAndUuid
func GetShortcutsByIdAndUuid(offset int, limit int, Id_ int64, Uuid_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("id = ? and uuid = ?", Id_, Uuid_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByIdAndLangcode Get Shortcuts via IdAndLangcode
func GetShortcutsByIdAndLangcode(offset int, limit int, Id_ int64, Langcode_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("id = ? and langcode = ?", Id_, Langcode_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByShortcutSetAndUuid Get Shortcuts via ShortcutSetAndUuid
func GetShortcutsByShortcutSetAndUuid(offset int, limit int, ShortcutSet_ string, Uuid_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("shortcut_set = ? and uuid = ?", ShortcutSet_, Uuid_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByShortcutSetAndLangcode Get Shortcuts via ShortcutSetAndLangcode
func GetShortcutsByShortcutSetAndLangcode(offset int, limit int, ShortcutSet_ string, Langcode_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("shortcut_set = ? and langcode = ?", ShortcutSet_, Langcode_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsByUuidAndLangcode Get Shortcuts via UuidAndLangcode
func GetShortcutsByUuidAndLangcode(offset int, limit int, Uuid_ string, Langcode_ string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("uuid = ? and langcode = ?", Uuid_, Langcode_).Limit(limit, offset).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcuts Get Shortcuts via field
func GetShortcuts(offset int, limit int, field string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Limit(limit, offset).Desc(field).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsViaId Get Shortcuts via Id
func GetShortcutsViaId(offset int, limit int, Id_ int64, field string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsViaShortcutSet Get Shortcuts via ShortcutSet
func GetShortcutsViaShortcutSet(offset int, limit int, ShortcutSet_ string, field string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("shortcut_set = ?", ShortcutSet_).Limit(limit, offset).Desc(field).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsViaUuid Get Shortcuts via Uuid
func GetShortcutsViaUuid(offset int, limit int, Uuid_ string, field string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("uuid = ?", Uuid_).Limit(limit, offset).Desc(field).Find(_Shortcut)
	return _Shortcut, err
}

// GetShortcutsViaLangcode Get Shortcuts via Langcode
func GetShortcutsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Shortcut, error) {
	var _Shortcut = new([]*Shortcut)
	err := Engine.Table("shortcut").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Shortcut)
	return _Shortcut, err
}

// HasShortcutViaId Has Shortcut via Id
func HasShortcutViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(Shortcut)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutViaShortcutSet Has Shortcut via ShortcutSet
func HasShortcutViaShortcutSet(iShortcutSet string) bool {
	if has, err := Engine.Where("shortcut_set = ?", iShortcutSet).Get(new(Shortcut)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutViaUuid Has Shortcut via Uuid
func HasShortcutViaUuid(iUuid string) bool {
	if has, err := Engine.Where("uuid = ?", iUuid).Get(new(Shortcut)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutViaLangcode Has Shortcut via Langcode
func HasShortcutViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Shortcut)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetShortcutViaId Get Shortcut via Id
func GetShortcutViaId(iId int64) (*Shortcut, error) {
	var _Shortcut = &Shortcut{Id: iId}
	has, err := Engine.Get(_Shortcut)
	if has {
		return _Shortcut, err
	} else {
		return nil, err
	}
}

// GetShortcutViaShortcutSet Get Shortcut via ShortcutSet
func GetShortcutViaShortcutSet(iShortcutSet string) (*Shortcut, error) {
	var _Shortcut = &Shortcut{ShortcutSet: iShortcutSet}
	has, err := Engine.Get(_Shortcut)
	if has {
		return _Shortcut, err
	} else {
		return nil, err
	}
}

// GetShortcutViaUuid Get Shortcut via Uuid
func GetShortcutViaUuid(iUuid string) (*Shortcut, error) {
	var _Shortcut = &Shortcut{Uuid: iUuid}
	has, err := Engine.Get(_Shortcut)
	if has {
		return _Shortcut, err
	} else {
		return nil, err
	}
}

// GetShortcutViaLangcode Get Shortcut via Langcode
func GetShortcutViaLangcode(iLangcode string) (*Shortcut, error) {
	var _Shortcut = &Shortcut{Langcode: iLangcode}
	has, err := Engine.Get(_Shortcut)
	if has {
		return _Shortcut, err
	} else {
		return nil, err
	}
}

// SetShortcutViaId Set Shortcut via Id
func SetShortcutViaId(iId int64, shortcut *Shortcut) (int64, error) {
	shortcut.Id = iId
	return Engine.Insert(shortcut)
}

// SetShortcutViaShortcutSet Set Shortcut via ShortcutSet
func SetShortcutViaShortcutSet(iShortcutSet string, shortcut *Shortcut) (int64, error) {
	shortcut.ShortcutSet = iShortcutSet
	return Engine.Insert(shortcut)
}

// SetShortcutViaUuid Set Shortcut via Uuid
func SetShortcutViaUuid(iUuid string, shortcut *Shortcut) (int64, error) {
	shortcut.Uuid = iUuid
	return Engine.Insert(shortcut)
}

// SetShortcutViaLangcode Set Shortcut via Langcode
func SetShortcutViaLangcode(iLangcode string, shortcut *Shortcut) (int64, error) {
	shortcut.Langcode = iLangcode
	return Engine.Insert(shortcut)
}

// AddShortcut Add Shortcut via all columns
func AddShortcut(iId int64, iShortcutSet string, iUuid string, iLangcode string) error {
	_Shortcut := &Shortcut{Id: iId, ShortcutSet: iShortcutSet, Uuid: iUuid, Langcode: iLangcode}
	if _, err := Engine.Insert(_Shortcut); err != nil {
		return err
	} else {
		return nil
	}
}

// PostShortcut Post Shortcut via iShortcut
func PostShortcut(iShortcut *Shortcut) (int64, error) {
	_, err := Engine.Insert(iShortcut)
	return iShortcut.Id, err
}

// PutShortcut Put Shortcut
func PutShortcut(iShortcut *Shortcut) (int64, error) {
	_, err := Engine.Id(iShortcut.Id).Update(iShortcut)
	return iShortcut.Id, err
}

// PutShortcutViaId Put Shortcut via Id
func PutShortcutViaId(Id_ int64, iShortcut *Shortcut) (int64, error) {
	row, err := Engine.Update(iShortcut, &Shortcut{Id: Id_})
	return row, err
}

// PutShortcutViaShortcutSet Put Shortcut via ShortcutSet
func PutShortcutViaShortcutSet(ShortcutSet_ string, iShortcut *Shortcut) (int64, error) {
	row, err := Engine.Update(iShortcut, &Shortcut{ShortcutSet: ShortcutSet_})
	return row, err
}

// PutShortcutViaUuid Put Shortcut via Uuid
func PutShortcutViaUuid(Uuid_ string, iShortcut *Shortcut) (int64, error) {
	row, err := Engine.Update(iShortcut, &Shortcut{Uuid: Uuid_})
	return row, err
}

// PutShortcutViaLangcode Put Shortcut via Langcode
func PutShortcutViaLangcode(Langcode_ string, iShortcut *Shortcut) (int64, error) {
	row, err := Engine.Update(iShortcut, &Shortcut{Langcode: Langcode_})
	return row, err
}

// UpdateShortcutViaId via map[string]interface{}{}
func UpdateShortcutViaId(iId int64, iShortcutMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Shortcut)).Where("id = ?", iId).Update(iShortcutMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutViaShortcutSet via map[string]interface{}{}
func UpdateShortcutViaShortcutSet(iShortcutSet string, iShortcutMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Shortcut)).Where("shortcut_set = ?", iShortcutSet).Update(iShortcutMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutViaUuid via map[string]interface{}{}
func UpdateShortcutViaUuid(iUuid string, iShortcutMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Shortcut)).Where("uuid = ?", iUuid).Update(iShortcutMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutViaLangcode via map[string]interface{}{}
func UpdateShortcutViaLangcode(iLangcode string, iShortcutMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Shortcut)).Where("langcode = ?", iLangcode).Update(iShortcutMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteShortcut Delete Shortcut
func DeleteShortcut(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(Shortcut))
	return row, err
}

// DeleteShortcutViaId Delete Shortcut via Id
func DeleteShortcutViaId(iId int64) (err error) {
	var has bool
	var _Shortcut = &Shortcut{Id: iId}
	if has, err = Engine.Get(_Shortcut); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(Shortcut)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutViaShortcutSet Delete Shortcut via ShortcutSet
func DeleteShortcutViaShortcutSet(iShortcutSet string) (err error) {
	var has bool
	var _Shortcut = &Shortcut{ShortcutSet: iShortcutSet}
	if has, err = Engine.Get(_Shortcut); (has == true) && (err == nil) {
		if row, err := Engine.Where("shortcut_set = ?", iShortcutSet).Delete(new(Shortcut)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutViaUuid Delete Shortcut via Uuid
func DeleteShortcutViaUuid(iUuid string) (err error) {
	var has bool
	var _Shortcut = &Shortcut{Uuid: iUuid}
	if has, err = Engine.Get(_Shortcut); (has == true) && (err == nil) {
		if row, err := Engine.Where("uuid = ?", iUuid).Delete(new(Shortcut)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutViaLangcode Delete Shortcut via Langcode
func DeleteShortcutViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Shortcut = &Shortcut{Langcode: iLangcode}
	if has, err = Engine.Get(_Shortcut); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Shortcut)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
