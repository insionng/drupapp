package model

type KeyValueExpire struct {
	Collection string `xorm:"not null pk default '' index(all) VARCHAR(128)"`
	Name       string `xorm:"not null pk default '' index(all) VARCHAR(128)"`
	Value      []byte `xorm:"not null LONGBLOB"`
	Expire     int    `xorm:"not null default 2147483647 index(all) index INT(11)"`
}

// GetKeyValueExpiresCount KeyValueExpires' Count
func GetKeyValueExpiresCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&KeyValueExpire{})
	return total, err
}

// GetKeyValueExpireCountViaCollection Get KeyValueExpire via Collection
func GetKeyValueExpireCountViaCollection(iCollection string) int64 {
	n, _ := Engine.Where("collection = ?", iCollection).Count(&KeyValueExpire{Collection: iCollection})
	return n
}

// GetKeyValueExpireCountViaName Get KeyValueExpire via Name
func GetKeyValueExpireCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&KeyValueExpire{Name: iName})
	return n
}

// GetKeyValueExpireCountViaValue Get KeyValueExpire via Value
func GetKeyValueExpireCountViaValue(iValue []byte) int64 {
	n, _ := Engine.Where("value = ?", iValue).Count(&KeyValueExpire{Value: iValue})
	return n
}

// GetKeyValueExpireCountViaExpire Get KeyValueExpire via Expire
func GetKeyValueExpireCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&KeyValueExpire{Expire: iExpire})
	return n
}

// GetKeyValueExpiresByCollectionAndNameAndValue Get KeyValueExpires via CollectionAndNameAndValue
func GetKeyValueExpiresByCollectionAndNameAndValue(offset int, limit int, Collection_ string, Name_ string, Value_ []byte) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("collection = ? and name = ? and value = ?", Collection_, Name_, Value_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByCollectionAndNameAndExpire Get KeyValueExpires via CollectionAndNameAndExpire
func GetKeyValueExpiresByCollectionAndNameAndExpire(offset int, limit int, Collection_ string, Name_ string, Expire_ int) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("collection = ? and name = ? and expire = ?", Collection_, Name_, Expire_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByCollectionAndValueAndExpire Get KeyValueExpires via CollectionAndValueAndExpire
func GetKeyValueExpiresByCollectionAndValueAndExpire(offset int, limit int, Collection_ string, Value_ []byte, Expire_ int) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("collection = ? and value = ? and expire = ?", Collection_, Value_, Expire_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByNameAndValueAndExpire Get KeyValueExpires via NameAndValueAndExpire
func GetKeyValueExpiresByNameAndValueAndExpire(offset int, limit int, Name_ string, Value_ []byte, Expire_ int) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("name = ? and value = ? and expire = ?", Name_, Value_, Expire_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByCollectionAndName Get KeyValueExpires via CollectionAndName
func GetKeyValueExpiresByCollectionAndName(offset int, limit int, Collection_ string, Name_ string) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("collection = ? and name = ?", Collection_, Name_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByCollectionAndValue Get KeyValueExpires via CollectionAndValue
func GetKeyValueExpiresByCollectionAndValue(offset int, limit int, Collection_ string, Value_ []byte) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("collection = ? and value = ?", Collection_, Value_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByCollectionAndExpire Get KeyValueExpires via CollectionAndExpire
func GetKeyValueExpiresByCollectionAndExpire(offset int, limit int, Collection_ string, Expire_ int) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("collection = ? and expire = ?", Collection_, Expire_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByNameAndValue Get KeyValueExpires via NameAndValue
func GetKeyValueExpiresByNameAndValue(offset int, limit int, Name_ string, Value_ []byte) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("name = ? and value = ?", Name_, Value_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByNameAndExpire Get KeyValueExpires via NameAndExpire
func GetKeyValueExpiresByNameAndExpire(offset int, limit int, Name_ string, Expire_ int) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("name = ? and expire = ?", Name_, Expire_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresByValueAndExpire Get KeyValueExpires via ValueAndExpire
func GetKeyValueExpiresByValueAndExpire(offset int, limit int, Value_ []byte, Expire_ int) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("value = ? and expire = ?", Value_, Expire_).Limit(limit, offset).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpires Get KeyValueExpires via field
func GetKeyValueExpires(offset int, limit int, field string) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Limit(limit, offset).Desc(field).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresViaCollection Get KeyValueExpires via Collection
func GetKeyValueExpiresViaCollection(offset int, limit int, Collection_ string, field string) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("collection = ?", Collection_).Limit(limit, offset).Desc(field).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresViaName Get KeyValueExpires via Name
func GetKeyValueExpiresViaName(offset int, limit int, Name_ string, field string) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresViaValue Get KeyValueExpires via Value
func GetKeyValueExpiresViaValue(offset int, limit int, Value_ []byte, field string) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("value = ?", Value_).Limit(limit, offset).Desc(field).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// GetKeyValueExpiresViaExpire Get KeyValueExpires via Expire
func GetKeyValueExpiresViaExpire(offset int, limit int, Expire_ int, field string) (*[]*KeyValueExpire, error) {
	var _KeyValueExpire = new([]*KeyValueExpire)
	err := Engine.Table("key_value_expire").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_KeyValueExpire)
	return _KeyValueExpire, err
}

// HasKeyValueExpireViaCollection Has KeyValueExpire via Collection
func HasKeyValueExpireViaCollection(iCollection string) bool {
	if has, err := Engine.Where("collection = ?", iCollection).Get(new(KeyValueExpire)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasKeyValueExpireViaName Has KeyValueExpire via Name
func HasKeyValueExpireViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(KeyValueExpire)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasKeyValueExpireViaValue Has KeyValueExpire via Value
func HasKeyValueExpireViaValue(iValue []byte) bool {
	if has, err := Engine.Where("value = ?", iValue).Get(new(KeyValueExpire)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasKeyValueExpireViaExpire Has KeyValueExpire via Expire
func HasKeyValueExpireViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(KeyValueExpire)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetKeyValueExpireViaCollection Get KeyValueExpire via Collection
func GetKeyValueExpireViaCollection(iCollection string) (*KeyValueExpire, error) {
	var _KeyValueExpire = &KeyValueExpire{Collection: iCollection}
	has, err := Engine.Get(_KeyValueExpire)
	if has {
		return _KeyValueExpire, err
	} else {
		return nil, err
	}
}

// GetKeyValueExpireViaName Get KeyValueExpire via Name
func GetKeyValueExpireViaName(iName string) (*KeyValueExpire, error) {
	var _KeyValueExpire = &KeyValueExpire{Name: iName}
	has, err := Engine.Get(_KeyValueExpire)
	if has {
		return _KeyValueExpire, err
	} else {
		return nil, err
	}
}

// GetKeyValueExpireViaValue Get KeyValueExpire via Value
func GetKeyValueExpireViaValue(iValue []byte) (*KeyValueExpire, error) {
	var _KeyValueExpire = &KeyValueExpire{Value: iValue}
	has, err := Engine.Get(_KeyValueExpire)
	if has {
		return _KeyValueExpire, err
	} else {
		return nil, err
	}
}

// GetKeyValueExpireViaExpire Get KeyValueExpire via Expire
func GetKeyValueExpireViaExpire(iExpire int) (*KeyValueExpire, error) {
	var _KeyValueExpire = &KeyValueExpire{Expire: iExpire}
	has, err := Engine.Get(_KeyValueExpire)
	if has {
		return _KeyValueExpire, err
	} else {
		return nil, err
	}
}

// SetKeyValueExpireViaCollection Set KeyValueExpire via Collection
func SetKeyValueExpireViaCollection(iCollection string, key_value_expire *KeyValueExpire) (int64, error) {
	key_value_expire.Collection = iCollection
	return Engine.Insert(key_value_expire)
}

// SetKeyValueExpireViaName Set KeyValueExpire via Name
func SetKeyValueExpireViaName(iName string, key_value_expire *KeyValueExpire) (int64, error) {
	key_value_expire.Name = iName
	return Engine.Insert(key_value_expire)
}

// SetKeyValueExpireViaValue Set KeyValueExpire via Value
func SetKeyValueExpireViaValue(iValue []byte, key_value_expire *KeyValueExpire) (int64, error) {
	key_value_expire.Value = iValue
	return Engine.Insert(key_value_expire)
}

// SetKeyValueExpireViaExpire Set KeyValueExpire via Expire
func SetKeyValueExpireViaExpire(iExpire int, key_value_expire *KeyValueExpire) (int64, error) {
	key_value_expire.Expire = iExpire
	return Engine.Insert(key_value_expire)
}

// AddKeyValueExpire Add KeyValueExpire via all columns
func AddKeyValueExpire(iCollection string, iName string, iValue []byte, iExpire int) error {
	_KeyValueExpire := &KeyValueExpire{Collection: iCollection, Name: iName, Value: iValue, Expire: iExpire}
	if _, err := Engine.Insert(_KeyValueExpire); err != nil {
		return err
	} else {
		return nil
	}
}

// PostKeyValueExpire Post KeyValueExpire via iKeyValueExpire
func PostKeyValueExpire(iKeyValueExpire *KeyValueExpire) (string, error) {
	_, err := Engine.Insert(iKeyValueExpire)
	return iKeyValueExpire.Collection, err
}

// PutKeyValueExpire Put KeyValueExpire
func PutKeyValueExpire(iKeyValueExpire *KeyValueExpire) (string, error) {
	_, err := Engine.Id(iKeyValueExpire.Collection).Update(iKeyValueExpire)
	return iKeyValueExpire.Collection, err
}

// PutKeyValueExpireViaCollection Put KeyValueExpire via Collection
func PutKeyValueExpireViaCollection(Collection_ string, iKeyValueExpire *KeyValueExpire) (int64, error) {
	row, err := Engine.Update(iKeyValueExpire, &KeyValueExpire{Collection: Collection_})
	return row, err
}

// PutKeyValueExpireViaName Put KeyValueExpire via Name
func PutKeyValueExpireViaName(Name_ string, iKeyValueExpire *KeyValueExpire) (int64, error) {
	row, err := Engine.Update(iKeyValueExpire, &KeyValueExpire{Name: Name_})
	return row, err
}

// PutKeyValueExpireViaValue Put KeyValueExpire via Value
func PutKeyValueExpireViaValue(Value_ []byte, iKeyValueExpire *KeyValueExpire) (int64, error) {
	row, err := Engine.Update(iKeyValueExpire, &KeyValueExpire{Value: Value_})
	return row, err
}

// PutKeyValueExpireViaExpire Put KeyValueExpire via Expire
func PutKeyValueExpireViaExpire(Expire_ int, iKeyValueExpire *KeyValueExpire) (int64, error) {
	row, err := Engine.Update(iKeyValueExpire, &KeyValueExpire{Expire: Expire_})
	return row, err
}

// UpdateKeyValueExpireViaCollection via map[string]interface{}{}
func UpdateKeyValueExpireViaCollection(iCollection string, iKeyValueExpireMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(KeyValueExpire)).Where("collection = ?", iCollection).Update(iKeyValueExpireMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateKeyValueExpireViaName via map[string]interface{}{}
func UpdateKeyValueExpireViaName(iName string, iKeyValueExpireMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(KeyValueExpire)).Where("name = ?", iName).Update(iKeyValueExpireMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateKeyValueExpireViaValue via map[string]interface{}{}
func UpdateKeyValueExpireViaValue(iValue []byte, iKeyValueExpireMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(KeyValueExpire)).Where("value = ?", iValue).Update(iKeyValueExpireMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateKeyValueExpireViaExpire via map[string]interface{}{}
func UpdateKeyValueExpireViaExpire(iExpire int, iKeyValueExpireMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(KeyValueExpire)).Where("expire = ?", iExpire).Update(iKeyValueExpireMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteKeyValueExpire Delete KeyValueExpire
func DeleteKeyValueExpire(iCollection string) (int64, error) {
	row, err := Engine.Id(iCollection).Delete(new(KeyValueExpire))
	return row, err
}

// DeleteKeyValueExpireViaCollection Delete KeyValueExpire via Collection
func DeleteKeyValueExpireViaCollection(iCollection string) (err error) {
	var has bool
	var _KeyValueExpire = &KeyValueExpire{Collection: iCollection}
	if has, err = Engine.Get(_KeyValueExpire); (has == true) && (err == nil) {
		if row, err := Engine.Where("collection = ?", iCollection).Delete(new(KeyValueExpire)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteKeyValueExpireViaName Delete KeyValueExpire via Name
func DeleteKeyValueExpireViaName(iName string) (err error) {
	var has bool
	var _KeyValueExpire = &KeyValueExpire{Name: iName}
	if has, err = Engine.Get(_KeyValueExpire); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(KeyValueExpire)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteKeyValueExpireViaValue Delete KeyValueExpire via Value
func DeleteKeyValueExpireViaValue(iValue []byte) (err error) {
	var has bool
	var _KeyValueExpire = &KeyValueExpire{Value: iValue}
	if has, err = Engine.Get(_KeyValueExpire); (has == true) && (err == nil) {
		if row, err := Engine.Where("value = ?", iValue).Delete(new(KeyValueExpire)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteKeyValueExpireViaExpire Delete KeyValueExpire via Expire
func DeleteKeyValueExpireViaExpire(iExpire int) (err error) {
	var has bool
	var _KeyValueExpire = &KeyValueExpire{Expire: iExpire}
	if has, err = Engine.Get(_KeyValueExpire); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(KeyValueExpire)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
