package model

type KeyValue struct {
	Collection string `xorm:"not null pk default '' VARCHAR(128)"`
	Name       string `xorm:"not null pk default '' VARCHAR(128)"`
	Value      []byte `xorm:"not null LONGBLOB"`
}

// GetKeyValuesCount KeyValues' Count
func GetKeyValuesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&KeyValue{})
	return total, err
}

// GetKeyValueCountViaCollection Get KeyValue via Collection
func GetKeyValueCountViaCollection(iCollection string) int64 {
	n, _ := Engine.Where("collection = ?", iCollection).Count(&KeyValue{Collection: iCollection})
	return n
}

// GetKeyValueCountViaName Get KeyValue via Name
func GetKeyValueCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&KeyValue{Name: iName})
	return n
}

// GetKeyValueCountViaValue Get KeyValue via Value
func GetKeyValueCountViaValue(iValue []byte) int64 {
	n, _ := Engine.Where("value = ?", iValue).Count(&KeyValue{Value: iValue})
	return n
}

// GetKeyValuesByCollectionAndNameAndValue Get KeyValues via CollectionAndNameAndValue
func GetKeyValuesByCollectionAndNameAndValue(offset int, limit int, Collection_ string, Name_ string, Value_ []byte) (*[]*KeyValue, error) {
	var _KeyValue = new([]*KeyValue)
	err := Engine.Table("key_value").Where("collection = ? and name = ? and value = ?", Collection_, Name_, Value_).Limit(limit, offset).Find(_KeyValue)
	return _KeyValue, err
}

// GetKeyValuesByCollectionAndName Get KeyValues via CollectionAndName
func GetKeyValuesByCollectionAndName(offset int, limit int, Collection_ string, Name_ string) (*[]*KeyValue, error) {
	var _KeyValue = new([]*KeyValue)
	err := Engine.Table("key_value").Where("collection = ? and name = ?", Collection_, Name_).Limit(limit, offset).Find(_KeyValue)
	return _KeyValue, err
}

// GetKeyValuesByCollectionAndValue Get KeyValues via CollectionAndValue
func GetKeyValuesByCollectionAndValue(offset int, limit int, Collection_ string, Value_ []byte) (*[]*KeyValue, error) {
	var _KeyValue = new([]*KeyValue)
	err := Engine.Table("key_value").Where("collection = ? and value = ?", Collection_, Value_).Limit(limit, offset).Find(_KeyValue)
	return _KeyValue, err
}

// GetKeyValuesByNameAndValue Get KeyValues via NameAndValue
func GetKeyValuesByNameAndValue(offset int, limit int, Name_ string, Value_ []byte) (*[]*KeyValue, error) {
	var _KeyValue = new([]*KeyValue)
	err := Engine.Table("key_value").Where("name = ? and value = ?", Name_, Value_).Limit(limit, offset).Find(_KeyValue)
	return _KeyValue, err
}

// GetKeyValues Get KeyValues via field
func GetKeyValues(offset int, limit int, field string) (*[]*KeyValue, error) {
	var _KeyValue = new([]*KeyValue)
	err := Engine.Table("key_value").Limit(limit, offset).Desc(field).Find(_KeyValue)
	return _KeyValue, err
}

// GetKeyValuesViaCollection Get KeyValues via Collection
func GetKeyValuesViaCollection(offset int, limit int, Collection_ string, field string) (*[]*KeyValue, error) {
	var _KeyValue = new([]*KeyValue)
	err := Engine.Table("key_value").Where("collection = ?", Collection_).Limit(limit, offset).Desc(field).Find(_KeyValue)
	return _KeyValue, err
}

// GetKeyValuesViaName Get KeyValues via Name
func GetKeyValuesViaName(offset int, limit int, Name_ string, field string) (*[]*KeyValue, error) {
	var _KeyValue = new([]*KeyValue)
	err := Engine.Table("key_value").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_KeyValue)
	return _KeyValue, err
}

// GetKeyValuesViaValue Get KeyValues via Value
func GetKeyValuesViaValue(offset int, limit int, Value_ []byte, field string) (*[]*KeyValue, error) {
	var _KeyValue = new([]*KeyValue)
	err := Engine.Table("key_value").Where("value = ?", Value_).Limit(limit, offset).Desc(field).Find(_KeyValue)
	return _KeyValue, err
}

// HasKeyValueViaCollection Has KeyValue via Collection
func HasKeyValueViaCollection(iCollection string) bool {
	if has, err := Engine.Where("collection = ?", iCollection).Get(new(KeyValue)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasKeyValueViaName Has KeyValue via Name
func HasKeyValueViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(KeyValue)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasKeyValueViaValue Has KeyValue via Value
func HasKeyValueViaValue(iValue []byte) bool {
	if has, err := Engine.Where("value = ?", iValue).Get(new(KeyValue)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetKeyValueViaCollection Get KeyValue via Collection
func GetKeyValueViaCollection(iCollection string) (*KeyValue, error) {
	var _KeyValue = &KeyValue{Collection: iCollection}
	has, err := Engine.Get(_KeyValue)
	if has {
		return _KeyValue, err
	} else {
		return nil, err
	}
}

// GetKeyValueViaName Get KeyValue via Name
func GetKeyValueViaName(iName string) (*KeyValue, error) {
	var _KeyValue = &KeyValue{Name: iName}
	has, err := Engine.Get(_KeyValue)
	if has {
		return _KeyValue, err
	} else {
		return nil, err
	}
}

// GetKeyValueViaValue Get KeyValue via Value
func GetKeyValueViaValue(iValue []byte) (*KeyValue, error) {
	var _KeyValue = &KeyValue{Value: iValue}
	has, err := Engine.Get(_KeyValue)
	if has {
		return _KeyValue, err
	} else {
		return nil, err
	}
}

// SetKeyValueViaCollection Set KeyValue via Collection
func SetKeyValueViaCollection(iCollection string, key_value *KeyValue) (int64, error) {
	key_value.Collection = iCollection
	return Engine.Insert(key_value)
}

// SetKeyValueViaName Set KeyValue via Name
func SetKeyValueViaName(iName string, key_value *KeyValue) (int64, error) {
	key_value.Name = iName
	return Engine.Insert(key_value)
}

// SetKeyValueViaValue Set KeyValue via Value
func SetKeyValueViaValue(iValue []byte, key_value *KeyValue) (int64, error) {
	key_value.Value = iValue
	return Engine.Insert(key_value)
}

// AddKeyValue Add KeyValue via all columns
func AddKeyValue(iCollection string, iName string, iValue []byte) error {
	_KeyValue := &KeyValue{Collection: iCollection, Name: iName, Value: iValue}
	if _, err := Engine.Insert(_KeyValue); err != nil {
		return err
	} else {
		return nil
	}
}

// PostKeyValue Post KeyValue via iKeyValue
func PostKeyValue(iKeyValue *KeyValue) (string, error) {
	_, err := Engine.Insert(iKeyValue)
	return iKeyValue.Collection, err
}

// PutKeyValue Put KeyValue
func PutKeyValue(iKeyValue *KeyValue) (string, error) {
	_, err := Engine.Id(iKeyValue.Collection).Update(iKeyValue)
	return iKeyValue.Collection, err
}

// PutKeyValueViaCollection Put KeyValue via Collection
func PutKeyValueViaCollection(Collection_ string, iKeyValue *KeyValue) (int64, error) {
	row, err := Engine.Update(iKeyValue, &KeyValue{Collection: Collection_})
	return row, err
}

// PutKeyValueViaName Put KeyValue via Name
func PutKeyValueViaName(Name_ string, iKeyValue *KeyValue) (int64, error) {
	row, err := Engine.Update(iKeyValue, &KeyValue{Name: Name_})
	return row, err
}

// PutKeyValueViaValue Put KeyValue via Value
func PutKeyValueViaValue(Value_ []byte, iKeyValue *KeyValue) (int64, error) {
	row, err := Engine.Update(iKeyValue, &KeyValue{Value: Value_})
	return row, err
}

// UpdateKeyValueViaCollection via map[string]interface{}{}
func UpdateKeyValueViaCollection(iCollection string, iKeyValueMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(KeyValue)).Where("collection = ?", iCollection).Update(iKeyValueMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateKeyValueViaName via map[string]interface{}{}
func UpdateKeyValueViaName(iName string, iKeyValueMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(KeyValue)).Where("name = ?", iName).Update(iKeyValueMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateKeyValueViaValue via map[string]interface{}{}
func UpdateKeyValueViaValue(iValue []byte, iKeyValueMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(KeyValue)).Where("value = ?", iValue).Update(iKeyValueMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteKeyValue Delete KeyValue
func DeleteKeyValue(iCollection string) (int64, error) {
	row, err := Engine.Id(iCollection).Delete(new(KeyValue))
	return row, err
}

// DeleteKeyValueViaCollection Delete KeyValue via Collection
func DeleteKeyValueViaCollection(iCollection string) (err error) {
	var has bool
	var _KeyValue = &KeyValue{Collection: iCollection}
	if has, err = Engine.Get(_KeyValue); (has == true) && (err == nil) {
		if row, err := Engine.Where("collection = ?", iCollection).Delete(new(KeyValue)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteKeyValueViaName Delete KeyValue via Name
func DeleteKeyValueViaName(iName string) (err error) {
	var has bool
	var _KeyValue = &KeyValue{Name: iName}
	if has, err = Engine.Get(_KeyValue); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(KeyValue)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteKeyValueViaValue Delete KeyValue via Value
func DeleteKeyValueViaValue(iValue []byte) (err error) {
	var has bool
	var _KeyValue = &KeyValue{Value: iValue}
	if has, err = Engine.Get(_KeyValue); (has == true) && (err == nil) {
		if row, err := Engine.Where("value = ?", iValue).Delete(new(KeyValue)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
