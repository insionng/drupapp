package model

type TaxonomyTermData struct {
	Tid      int64  `xorm:"not null pk autoincr INT(10)"`
	Vid      string `xorm:"not null index VARCHAR(32)"`
	Uuid     string `xorm:"not null unique VARCHAR(128)"`
	Langcode string `xorm:"not null VARCHAR(12)"`
}

// GetTaxonomyTermDatasCount TaxonomyTermDatas' Count
func GetTaxonomyTermDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&TaxonomyTermData{})
	return total, err
}

// GetTaxonomyTermDataCountViaTid Get TaxonomyTermData via Tid
func GetTaxonomyTermDataCountViaTid(iTid int64) int64 {
	n, _ := Engine.Where("tid = ?", iTid).Count(&TaxonomyTermData{Tid: iTid})
	return n
}

// GetTaxonomyTermDataCountViaVid Get TaxonomyTermData via Vid
func GetTaxonomyTermDataCountViaVid(iVid string) int64 {
	n, _ := Engine.Where("vid = ?", iVid).Count(&TaxonomyTermData{Vid: iVid})
	return n
}

// GetTaxonomyTermDataCountViaUuid Get TaxonomyTermData via Uuid
func GetTaxonomyTermDataCountViaUuid(iUuid string) int64 {
	n, _ := Engine.Where("uuid = ?", iUuid).Count(&TaxonomyTermData{Uuid: iUuid})
	return n
}

// GetTaxonomyTermDataCountViaLangcode Get TaxonomyTermData via Langcode
func GetTaxonomyTermDataCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&TaxonomyTermData{Langcode: iLangcode})
	return n
}

// GetTaxonomyTermDatasByTidAndVidAndUuid Get TaxonomyTermDatas via TidAndVidAndUuid
func GetTaxonomyTermDatasByTidAndVidAndUuid(offset int, limit int, Tid_ int64, Vid_ string, Uuid_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("tid = ? and vid = ? and uuid = ?", Tid_, Vid_, Uuid_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByTidAndVidAndLangcode Get TaxonomyTermDatas via TidAndVidAndLangcode
func GetTaxonomyTermDatasByTidAndVidAndLangcode(offset int, limit int, Tid_ int64, Vid_ string, Langcode_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("tid = ? and vid = ? and langcode = ?", Tid_, Vid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByTidAndUuidAndLangcode Get TaxonomyTermDatas via TidAndUuidAndLangcode
func GetTaxonomyTermDatasByTidAndUuidAndLangcode(offset int, limit int, Tid_ int64, Uuid_ string, Langcode_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("tid = ? and uuid = ? and langcode = ?", Tid_, Uuid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByVidAndUuidAndLangcode Get TaxonomyTermDatas via VidAndUuidAndLangcode
func GetTaxonomyTermDatasByVidAndUuidAndLangcode(offset int, limit int, Vid_ string, Uuid_ string, Langcode_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("vid = ? and uuid = ? and langcode = ?", Vid_, Uuid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByTidAndVid Get TaxonomyTermDatas via TidAndVid
func GetTaxonomyTermDatasByTidAndVid(offset int, limit int, Tid_ int64, Vid_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("tid = ? and vid = ?", Tid_, Vid_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByTidAndUuid Get TaxonomyTermDatas via TidAndUuid
func GetTaxonomyTermDatasByTidAndUuid(offset int, limit int, Tid_ int64, Uuid_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("tid = ? and uuid = ?", Tid_, Uuid_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByTidAndLangcode Get TaxonomyTermDatas via TidAndLangcode
func GetTaxonomyTermDatasByTidAndLangcode(offset int, limit int, Tid_ int64, Langcode_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("tid = ? and langcode = ?", Tid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByVidAndUuid Get TaxonomyTermDatas via VidAndUuid
func GetTaxonomyTermDatasByVidAndUuid(offset int, limit int, Vid_ string, Uuid_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("vid = ? and uuid = ?", Vid_, Uuid_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByVidAndLangcode Get TaxonomyTermDatas via VidAndLangcode
func GetTaxonomyTermDatasByVidAndLangcode(offset int, limit int, Vid_ string, Langcode_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("vid = ? and langcode = ?", Vid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasByUuidAndLangcode Get TaxonomyTermDatas via UuidAndLangcode
func GetTaxonomyTermDatasByUuidAndLangcode(offset int, limit int, Uuid_ string, Langcode_ string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("uuid = ? and langcode = ?", Uuid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatas Get TaxonomyTermDatas via field
func GetTaxonomyTermDatas(offset int, limit int, field string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Limit(limit, offset).Desc(field).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasViaTid Get TaxonomyTermDatas via Tid
func GetTaxonomyTermDatasViaTid(offset int, limit int, Tid_ int64, field string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("tid = ?", Tid_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasViaVid Get TaxonomyTermDatas via Vid
func GetTaxonomyTermDatasViaVid(offset int, limit int, Vid_ string, field string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("vid = ?", Vid_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasViaUuid Get TaxonomyTermDatas via Uuid
func GetTaxonomyTermDatasViaUuid(offset int, limit int, Uuid_ string, field string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("uuid = ?", Uuid_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// GetTaxonomyTermDatasViaLangcode Get TaxonomyTermDatas via Langcode
func GetTaxonomyTermDatasViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*TaxonomyTermData, error) {
	var _TaxonomyTermData = new([]*TaxonomyTermData)
	err := Engine.Table("taxonomy_term_data").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermData)
	return _TaxonomyTermData, err
}

// HasTaxonomyTermDataViaTid Has TaxonomyTermData via Tid
func HasTaxonomyTermDataViaTid(iTid int64) bool {
	if has, err := Engine.Where("tid = ?", iTid).Get(new(TaxonomyTermData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermDataViaVid Has TaxonomyTermData via Vid
func HasTaxonomyTermDataViaVid(iVid string) bool {
	if has, err := Engine.Where("vid = ?", iVid).Get(new(TaxonomyTermData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermDataViaUuid Has TaxonomyTermData via Uuid
func HasTaxonomyTermDataViaUuid(iUuid string) bool {
	if has, err := Engine.Where("uuid = ?", iUuid).Get(new(TaxonomyTermData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermDataViaLangcode Has TaxonomyTermData via Langcode
func HasTaxonomyTermDataViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(TaxonomyTermData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTaxonomyTermDataViaTid Get TaxonomyTermData via Tid
func GetTaxonomyTermDataViaTid(iTid int64) (*TaxonomyTermData, error) {
	var _TaxonomyTermData = &TaxonomyTermData{Tid: iTid}
	has, err := Engine.Get(_TaxonomyTermData)
	if has {
		return _TaxonomyTermData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermDataViaVid Get TaxonomyTermData via Vid
func GetTaxonomyTermDataViaVid(iVid string) (*TaxonomyTermData, error) {
	var _TaxonomyTermData = &TaxonomyTermData{Vid: iVid}
	has, err := Engine.Get(_TaxonomyTermData)
	if has {
		return _TaxonomyTermData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermDataViaUuid Get TaxonomyTermData via Uuid
func GetTaxonomyTermDataViaUuid(iUuid string) (*TaxonomyTermData, error) {
	var _TaxonomyTermData = &TaxonomyTermData{Uuid: iUuid}
	has, err := Engine.Get(_TaxonomyTermData)
	if has {
		return _TaxonomyTermData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermDataViaLangcode Get TaxonomyTermData via Langcode
func GetTaxonomyTermDataViaLangcode(iLangcode string) (*TaxonomyTermData, error) {
	var _TaxonomyTermData = &TaxonomyTermData{Langcode: iLangcode}
	has, err := Engine.Get(_TaxonomyTermData)
	if has {
		return _TaxonomyTermData, err
	} else {
		return nil, err
	}
}

// SetTaxonomyTermDataViaTid Set TaxonomyTermData via Tid
func SetTaxonomyTermDataViaTid(iTid int64, taxonomy_term_data *TaxonomyTermData) (int64, error) {
	taxonomy_term_data.Tid = iTid
	return Engine.Insert(taxonomy_term_data)
}

// SetTaxonomyTermDataViaVid Set TaxonomyTermData via Vid
func SetTaxonomyTermDataViaVid(iVid string, taxonomy_term_data *TaxonomyTermData) (int64, error) {
	taxonomy_term_data.Vid = iVid
	return Engine.Insert(taxonomy_term_data)
}

// SetTaxonomyTermDataViaUuid Set TaxonomyTermData via Uuid
func SetTaxonomyTermDataViaUuid(iUuid string, taxonomy_term_data *TaxonomyTermData) (int64, error) {
	taxonomy_term_data.Uuid = iUuid
	return Engine.Insert(taxonomy_term_data)
}

// SetTaxonomyTermDataViaLangcode Set TaxonomyTermData via Langcode
func SetTaxonomyTermDataViaLangcode(iLangcode string, taxonomy_term_data *TaxonomyTermData) (int64, error) {
	taxonomy_term_data.Langcode = iLangcode
	return Engine.Insert(taxonomy_term_data)
}

// AddTaxonomyTermData Add TaxonomyTermData via all columns
func AddTaxonomyTermData(iTid int64, iVid string, iUuid string, iLangcode string) error {
	_TaxonomyTermData := &TaxonomyTermData{Tid: iTid, Vid: iVid, Uuid: iUuid, Langcode: iLangcode}
	if _, err := Engine.Insert(_TaxonomyTermData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostTaxonomyTermData Post TaxonomyTermData via iTaxonomyTermData
func PostTaxonomyTermData(iTaxonomyTermData *TaxonomyTermData) (int64, error) {
	_, err := Engine.Insert(iTaxonomyTermData)
	return iTaxonomyTermData.Tid, err
}

// PutTaxonomyTermData Put TaxonomyTermData
func PutTaxonomyTermData(iTaxonomyTermData *TaxonomyTermData) (int64, error) {
	_, err := Engine.Id(iTaxonomyTermData.Tid).Update(iTaxonomyTermData)
	return iTaxonomyTermData.Tid, err
}

// PutTaxonomyTermDataViaTid Put TaxonomyTermData via Tid
func PutTaxonomyTermDataViaTid(Tid_ int64, iTaxonomyTermData *TaxonomyTermData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermData, &TaxonomyTermData{Tid: Tid_})
	return row, err
}

// PutTaxonomyTermDataViaVid Put TaxonomyTermData via Vid
func PutTaxonomyTermDataViaVid(Vid_ string, iTaxonomyTermData *TaxonomyTermData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermData, &TaxonomyTermData{Vid: Vid_})
	return row, err
}

// PutTaxonomyTermDataViaUuid Put TaxonomyTermData via Uuid
func PutTaxonomyTermDataViaUuid(Uuid_ string, iTaxonomyTermData *TaxonomyTermData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermData, &TaxonomyTermData{Uuid: Uuid_})
	return row, err
}

// PutTaxonomyTermDataViaLangcode Put TaxonomyTermData via Langcode
func PutTaxonomyTermDataViaLangcode(Langcode_ string, iTaxonomyTermData *TaxonomyTermData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermData, &TaxonomyTermData{Langcode: Langcode_})
	return row, err
}

// UpdateTaxonomyTermDataViaTid via map[string]interface{}{}
func UpdateTaxonomyTermDataViaTid(iTid int64, iTaxonomyTermDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermData)).Where("tid = ?", iTid).Update(iTaxonomyTermDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermDataViaVid via map[string]interface{}{}
func UpdateTaxonomyTermDataViaVid(iVid string, iTaxonomyTermDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermData)).Where("vid = ?", iVid).Update(iTaxonomyTermDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermDataViaUuid via map[string]interface{}{}
func UpdateTaxonomyTermDataViaUuid(iUuid string, iTaxonomyTermDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermData)).Where("uuid = ?", iUuid).Update(iTaxonomyTermDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermDataViaLangcode via map[string]interface{}{}
func UpdateTaxonomyTermDataViaLangcode(iLangcode string, iTaxonomyTermDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermData)).Where("langcode = ?", iLangcode).Update(iTaxonomyTermDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteTaxonomyTermData Delete TaxonomyTermData
func DeleteTaxonomyTermData(iTid int64) (int64, error) {
	row, err := Engine.Id(iTid).Delete(new(TaxonomyTermData))
	return row, err
}

// DeleteTaxonomyTermDataViaTid Delete TaxonomyTermData via Tid
func DeleteTaxonomyTermDataViaTid(iTid int64) (err error) {
	var has bool
	var _TaxonomyTermData = &TaxonomyTermData{Tid: iTid}
	if has, err = Engine.Get(_TaxonomyTermData); (has == true) && (err == nil) {
		if row, err := Engine.Where("tid = ?", iTid).Delete(new(TaxonomyTermData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermDataViaVid Delete TaxonomyTermData via Vid
func DeleteTaxonomyTermDataViaVid(iVid string) (err error) {
	var has bool
	var _TaxonomyTermData = &TaxonomyTermData{Vid: iVid}
	if has, err = Engine.Get(_TaxonomyTermData); (has == true) && (err == nil) {
		if row, err := Engine.Where("vid = ?", iVid).Delete(new(TaxonomyTermData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermDataViaUuid Delete TaxonomyTermData via Uuid
func DeleteTaxonomyTermDataViaUuid(iUuid string) (err error) {
	var has bool
	var _TaxonomyTermData = &TaxonomyTermData{Uuid: iUuid}
	if has, err = Engine.Get(_TaxonomyTermData); (has == true) && (err == nil) {
		if row, err := Engine.Where("uuid = ?", iUuid).Delete(new(TaxonomyTermData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermDataViaLangcode Delete TaxonomyTermData via Langcode
func DeleteTaxonomyTermDataViaLangcode(iLangcode string) (err error) {
	var has bool
	var _TaxonomyTermData = &TaxonomyTermData{Langcode: iLangcode}
	if has, err = Engine.Get(_TaxonomyTermData); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(TaxonomyTermData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
