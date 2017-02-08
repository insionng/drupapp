package model

type Cachetags struct {
	Tag           string `xorm:"not null pk default '' VARCHAR(255)"`
	Invalidations int    `xorm:"not null default 0 INT(11)"`
}

// GetCachetagsesCount Cachetagss' Count
func GetCachetagsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Cachetags{})
	return total, err
}

// GetCachetagsCountViaTag Get Cachetags via Tag
func GetCachetagsCountViaTag(iTag string) int64 {
	n, _ := Engine.Where("tag = ?", iTag).Count(&Cachetags{Tag: iTag})
	return n
}

// GetCachetagsCountViaInvalidations Get Cachetags via Invalidations
func GetCachetagsCountViaInvalidations(iInvalidations int) int64 {
	n, _ := Engine.Where("invalidations = ?", iInvalidations).Count(&Cachetags{Invalidations: iInvalidations})
	return n
}

// GetCachetagsesByTagAndInvalidations Get Cachetagses via TagAndInvalidations
func GetCachetagsesByTagAndInvalidations(offset int, limit int, Tag_ string, Invalidations_ int) (*[]*Cachetags, error) {
	var _Cachetags = new([]*Cachetags)
	err := Engine.Table("cachetags").Where("tag = ? and invalidations = ?", Tag_, Invalidations_).Limit(limit, offset).Find(_Cachetags)
	return _Cachetags, err
}

// GetCachetagses Get Cachetagses via field
func GetCachetagses(offset int, limit int, field string) (*[]*Cachetags, error) {
	var _Cachetags = new([]*Cachetags)
	err := Engine.Table("cachetags").Limit(limit, offset).Desc(field).Find(_Cachetags)
	return _Cachetags, err
}

// GetCachetagsesViaTag Get Cachetagss via Tag
func GetCachetagsesViaTag(offset int, limit int, Tag_ string, field string) (*[]*Cachetags, error) {
	var _Cachetags = new([]*Cachetags)
	err := Engine.Table("cachetags").Where("tag = ?", Tag_).Limit(limit, offset).Desc(field).Find(_Cachetags)
	return _Cachetags, err
}

// GetCachetagsesViaInvalidations Get Cachetagss via Invalidations
func GetCachetagsesViaInvalidations(offset int, limit int, Invalidations_ int, field string) (*[]*Cachetags, error) {
	var _Cachetags = new([]*Cachetags)
	err := Engine.Table("cachetags").Where("invalidations = ?", Invalidations_).Limit(limit, offset).Desc(field).Find(_Cachetags)
	return _Cachetags, err
}

// HasCachetagsViaTag Has Cachetags via Tag
func HasCachetagsViaTag(iTag string) bool {
	if has, err := Engine.Where("tag = ?", iTag).Get(new(Cachetags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCachetagsViaInvalidations Has Cachetags via Invalidations
func HasCachetagsViaInvalidations(iInvalidations int) bool {
	if has, err := Engine.Where("invalidations = ?", iInvalidations).Get(new(Cachetags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCachetagsViaTag Get Cachetags via Tag
func GetCachetagsViaTag(iTag string) (*Cachetags, error) {
	var _Cachetags = &Cachetags{Tag: iTag}
	has, err := Engine.Get(_Cachetags)
	if has {
		return _Cachetags, err
	} else {
		return nil, err
	}
}

// GetCachetagsViaInvalidations Get Cachetags via Invalidations
func GetCachetagsViaInvalidations(iInvalidations int) (*Cachetags, error) {
	var _Cachetags = &Cachetags{Invalidations: iInvalidations}
	has, err := Engine.Get(_Cachetags)
	if has {
		return _Cachetags, err
	} else {
		return nil, err
	}
}

// SetCachetagsViaTag Set Cachetags via Tag
func SetCachetagsViaTag(iTag string, cachetags *Cachetags) (int64, error) {
	cachetags.Tag = iTag
	return Engine.Insert(cachetags)
}

// SetCachetagsViaInvalidations Set Cachetags via Invalidations
func SetCachetagsViaInvalidations(iInvalidations int, cachetags *Cachetags) (int64, error) {
	cachetags.Invalidations = iInvalidations
	return Engine.Insert(cachetags)
}

// AddCachetags Add Cachetags via all columns
func AddCachetags(iTag string, iInvalidations int) error {
	_Cachetags := &Cachetags{Tag: iTag, Invalidations: iInvalidations}
	if _, err := Engine.Insert(_Cachetags); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCachetags Post Cachetags via iCachetags
func PostCachetags(iCachetags *Cachetags) (string, error) {
	_, err := Engine.Insert(iCachetags)
	return iCachetags.Tag, err
}

// PutCachetags Put Cachetags
func PutCachetags(iCachetags *Cachetags) (string, error) {
	_, err := Engine.Id(iCachetags.Tag).Update(iCachetags)
	return iCachetags.Tag, err
}

// PutCachetagsViaTag Put Cachetags via Tag
func PutCachetagsViaTag(Tag_ string, iCachetags *Cachetags) (int64, error) {
	row, err := Engine.Update(iCachetags, &Cachetags{Tag: Tag_})
	return row, err
}

// PutCachetagsViaInvalidations Put Cachetags via Invalidations
func PutCachetagsViaInvalidations(Invalidations_ int, iCachetags *Cachetags) (int64, error) {
	row, err := Engine.Update(iCachetags, &Cachetags{Invalidations: Invalidations_})
	return row, err
}

// UpdateCachetagsViaTag via map[string]interface{}{}
func UpdateCachetagsViaTag(iTag string, iCachetagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Cachetags)).Where("tag = ?", iTag).Update(iCachetagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCachetagsViaInvalidations via map[string]interface{}{}
func UpdateCachetagsViaInvalidations(iInvalidations int, iCachetagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Cachetags)).Where("invalidations = ?", iInvalidations).Update(iCachetagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCachetags Delete Cachetags
func DeleteCachetags(iTag string) (int64, error) {
	row, err := Engine.Id(iTag).Delete(new(Cachetags))
	return row, err
}

// DeleteCachetagsViaTag Delete Cachetags via Tag
func DeleteCachetagsViaTag(iTag string) (err error) {
	var has bool
	var _Cachetags = &Cachetags{Tag: iTag}
	if has, err = Engine.Get(_Cachetags); (has == true) && (err == nil) {
		if row, err := Engine.Where("tag = ?", iTag).Delete(new(Cachetags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCachetagsViaInvalidations Delete Cachetags via Invalidations
func DeleteCachetagsViaInvalidations(iInvalidations int) (err error) {
	var has bool
	var _Cachetags = &Cachetags{Invalidations: iInvalidations}
	if has, err = Engine.Get(_Cachetags); (has == true) && (err == nil) {
		if row, err := Engine.Where("invalidations = ?", iInvalidations).Delete(new(Cachetags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
