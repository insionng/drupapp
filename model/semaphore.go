package model

type Semaphore struct {
	Name   string  `xorm:"not null pk default '' VARCHAR(255)"`
	Value  string  `xorm:"not null default '' index VARCHAR(255)"`
	Expire float64 `xorm:"not null index DOUBLE"`
}

// GetSemaphoresCount Semaphores' Count
func GetSemaphoresCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Semaphore{})
	return total, err
}

// GetSemaphoreCountViaName Get Semaphore via Name
func GetSemaphoreCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&Semaphore{Name: iName})
	return n
}

// GetSemaphoreCountViaValue Get Semaphore via Value
func GetSemaphoreCountViaValue(iValue string) int64 {
	n, _ := Engine.Where("value = ?", iValue).Count(&Semaphore{Value: iValue})
	return n
}

// GetSemaphoreCountViaExpire Get Semaphore via Expire
func GetSemaphoreCountViaExpire(iExpire float64) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&Semaphore{Expire: iExpire})
	return n
}

// GetSemaphoresByNameAndValueAndExpire Get Semaphores via NameAndValueAndExpire
func GetSemaphoresByNameAndValueAndExpire(offset int, limit int, Name_ string, Value_ string, Expire_ float64) (*[]*Semaphore, error) {
	var _Semaphore = new([]*Semaphore)
	err := Engine.Table("semaphore").Where("name = ? and value = ? and expire = ?", Name_, Value_, Expire_).Limit(limit, offset).Find(_Semaphore)
	return _Semaphore, err
}

// GetSemaphoresByNameAndValue Get Semaphores via NameAndValue
func GetSemaphoresByNameAndValue(offset int, limit int, Name_ string, Value_ string) (*[]*Semaphore, error) {
	var _Semaphore = new([]*Semaphore)
	err := Engine.Table("semaphore").Where("name = ? and value = ?", Name_, Value_).Limit(limit, offset).Find(_Semaphore)
	return _Semaphore, err
}

// GetSemaphoresByNameAndExpire Get Semaphores via NameAndExpire
func GetSemaphoresByNameAndExpire(offset int, limit int, Name_ string, Expire_ float64) (*[]*Semaphore, error) {
	var _Semaphore = new([]*Semaphore)
	err := Engine.Table("semaphore").Where("name = ? and expire = ?", Name_, Expire_).Limit(limit, offset).Find(_Semaphore)
	return _Semaphore, err
}

// GetSemaphoresByValueAndExpire Get Semaphores via ValueAndExpire
func GetSemaphoresByValueAndExpire(offset int, limit int, Value_ string, Expire_ float64) (*[]*Semaphore, error) {
	var _Semaphore = new([]*Semaphore)
	err := Engine.Table("semaphore").Where("value = ? and expire = ?", Value_, Expire_).Limit(limit, offset).Find(_Semaphore)
	return _Semaphore, err
}

// GetSemaphores Get Semaphores via field
func GetSemaphores(offset int, limit int, field string) (*[]*Semaphore, error) {
	var _Semaphore = new([]*Semaphore)
	err := Engine.Table("semaphore").Limit(limit, offset).Desc(field).Find(_Semaphore)
	return _Semaphore, err
}

// GetSemaphoresViaName Get Semaphores via Name
func GetSemaphoresViaName(offset int, limit int, Name_ string, field string) (*[]*Semaphore, error) {
	var _Semaphore = new([]*Semaphore)
	err := Engine.Table("semaphore").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_Semaphore)
	return _Semaphore, err
}

// GetSemaphoresViaValue Get Semaphores via Value
func GetSemaphoresViaValue(offset int, limit int, Value_ string, field string) (*[]*Semaphore, error) {
	var _Semaphore = new([]*Semaphore)
	err := Engine.Table("semaphore").Where("value = ?", Value_).Limit(limit, offset).Desc(field).Find(_Semaphore)
	return _Semaphore, err
}

// GetSemaphoresViaExpire Get Semaphores via Expire
func GetSemaphoresViaExpire(offset int, limit int, Expire_ float64, field string) (*[]*Semaphore, error) {
	var _Semaphore = new([]*Semaphore)
	err := Engine.Table("semaphore").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_Semaphore)
	return _Semaphore, err
}

// HasSemaphoreViaName Has Semaphore via Name
func HasSemaphoreViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(Semaphore)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSemaphoreViaValue Has Semaphore via Value
func HasSemaphoreViaValue(iValue string) bool {
	if has, err := Engine.Where("value = ?", iValue).Get(new(Semaphore)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSemaphoreViaExpire Has Semaphore via Expire
func HasSemaphoreViaExpire(iExpire float64) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(Semaphore)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSemaphoreViaName Get Semaphore via Name
func GetSemaphoreViaName(iName string) (*Semaphore, error) {
	var _Semaphore = &Semaphore{Name: iName}
	has, err := Engine.Get(_Semaphore)
	if has {
		return _Semaphore, err
	} else {
		return nil, err
	}
}

// GetSemaphoreViaValue Get Semaphore via Value
func GetSemaphoreViaValue(iValue string) (*Semaphore, error) {
	var _Semaphore = &Semaphore{Value: iValue}
	has, err := Engine.Get(_Semaphore)
	if has {
		return _Semaphore, err
	} else {
		return nil, err
	}
}

// GetSemaphoreViaExpire Get Semaphore via Expire
func GetSemaphoreViaExpire(iExpire float64) (*Semaphore, error) {
	var _Semaphore = &Semaphore{Expire: iExpire}
	has, err := Engine.Get(_Semaphore)
	if has {
		return _Semaphore, err
	} else {
		return nil, err
	}
}

// SetSemaphoreViaName Set Semaphore via Name
func SetSemaphoreViaName(iName string, semaphore *Semaphore) (int64, error) {
	semaphore.Name = iName
	return Engine.Insert(semaphore)
}

// SetSemaphoreViaValue Set Semaphore via Value
func SetSemaphoreViaValue(iValue string, semaphore *Semaphore) (int64, error) {
	semaphore.Value = iValue
	return Engine.Insert(semaphore)
}

// SetSemaphoreViaExpire Set Semaphore via Expire
func SetSemaphoreViaExpire(iExpire float64, semaphore *Semaphore) (int64, error) {
	semaphore.Expire = iExpire
	return Engine.Insert(semaphore)
}

// AddSemaphore Add Semaphore via all columns
func AddSemaphore(iName string, iValue string, iExpire float64) error {
	_Semaphore := &Semaphore{Name: iName, Value: iValue, Expire: iExpire}
	if _, err := Engine.Insert(_Semaphore); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSemaphore Post Semaphore via iSemaphore
func PostSemaphore(iSemaphore *Semaphore) (string, error) {
	_, err := Engine.Insert(iSemaphore)
	return iSemaphore.Name, err
}

// PutSemaphore Put Semaphore
func PutSemaphore(iSemaphore *Semaphore) (string, error) {
	_, err := Engine.Id(iSemaphore.Name).Update(iSemaphore)
	return iSemaphore.Name, err
}

// PutSemaphoreViaName Put Semaphore via Name
func PutSemaphoreViaName(Name_ string, iSemaphore *Semaphore) (int64, error) {
	row, err := Engine.Update(iSemaphore, &Semaphore{Name: Name_})
	return row, err
}

// PutSemaphoreViaValue Put Semaphore via Value
func PutSemaphoreViaValue(Value_ string, iSemaphore *Semaphore) (int64, error) {
	row, err := Engine.Update(iSemaphore, &Semaphore{Value: Value_})
	return row, err
}

// PutSemaphoreViaExpire Put Semaphore via Expire
func PutSemaphoreViaExpire(Expire_ float64, iSemaphore *Semaphore) (int64, error) {
	row, err := Engine.Update(iSemaphore, &Semaphore{Expire: Expire_})
	return row, err
}

// UpdateSemaphoreViaName via map[string]interface{}{}
func UpdateSemaphoreViaName(iName string, iSemaphoreMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Semaphore)).Where("name = ?", iName).Update(iSemaphoreMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSemaphoreViaValue via map[string]interface{}{}
func UpdateSemaphoreViaValue(iValue string, iSemaphoreMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Semaphore)).Where("value = ?", iValue).Update(iSemaphoreMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSemaphoreViaExpire via map[string]interface{}{}
func UpdateSemaphoreViaExpire(iExpire float64, iSemaphoreMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Semaphore)).Where("expire = ?", iExpire).Update(iSemaphoreMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSemaphore Delete Semaphore
func DeleteSemaphore(iName string) (int64, error) {
	row, err := Engine.Id(iName).Delete(new(Semaphore))
	return row, err
}

// DeleteSemaphoreViaName Delete Semaphore via Name
func DeleteSemaphoreViaName(iName string) (err error) {
	var has bool
	var _Semaphore = &Semaphore{Name: iName}
	if has, err = Engine.Get(_Semaphore); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(Semaphore)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSemaphoreViaValue Delete Semaphore via Value
func DeleteSemaphoreViaValue(iValue string) (err error) {
	var has bool
	var _Semaphore = &Semaphore{Value: iValue}
	if has, err = Engine.Get(_Semaphore); (has == true) && (err == nil) {
		if row, err := Engine.Where("value = ?", iValue).Delete(new(Semaphore)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSemaphoreViaExpire Delete Semaphore via Expire
func DeleteSemaphoreViaExpire(iExpire float64) (err error) {
	var has bool
	var _Semaphore = &Semaphore{Expire: iExpire}
	if has, err = Engine.Get(_Semaphore); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(Semaphore)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
