package model

type SearchDataset struct {
	Sid      int64  `xorm:"not null pk default 0 INT(10)"`
	Langcode string `xorm:"not null pk default '' VARCHAR(12)"`
	Type     string `xorm:"not null pk VARCHAR(64)"`
	Data     string `xorm:"not null LONGTEXT"`
	Reindex  int    `xorm:"not null default 0 INT(10)"`
}

// GetSearchDatasetsCount SearchDatasets' Count
func GetSearchDatasetsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&SearchDataset{})
	return total, err
}

// GetSearchDatasetCountViaSid Get SearchDataset via Sid
func GetSearchDatasetCountViaSid(iSid int64) int64 {
	n, _ := Engine.Where("sid = ?", iSid).Count(&SearchDataset{Sid: iSid})
	return n
}

// GetSearchDatasetCountViaLangcode Get SearchDataset via Langcode
func GetSearchDatasetCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&SearchDataset{Langcode: iLangcode})
	return n
}

// GetSearchDatasetCountViaType Get SearchDataset via Type
func GetSearchDatasetCountViaType(iType string) int64 {
	n, _ := Engine.Where("type = ?", iType).Count(&SearchDataset{Type: iType})
	return n
}

// GetSearchDatasetCountViaData Get SearchDataset via Data
func GetSearchDatasetCountViaData(iData string) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&SearchDataset{Data: iData})
	return n
}

// GetSearchDatasetCountViaReindex Get SearchDataset via Reindex
func GetSearchDatasetCountViaReindex(iReindex int) int64 {
	n, _ := Engine.Where("reindex = ?", iReindex).Count(&SearchDataset{Reindex: iReindex})
	return n
}

// GetSearchDatasetsBySidAndLangcodeAndType Get SearchDatasets via SidAndLangcodeAndType
func GetSearchDatasetsBySidAndLangcodeAndType(offset int, limit int, Sid_ int64, Langcode_ string, Type_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and langcode = ? and type = ?", Sid_, Langcode_, Type_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndLangcodeAndData Get SearchDatasets via SidAndLangcodeAndData
func GetSearchDatasetsBySidAndLangcodeAndData(offset int, limit int, Sid_ int64, Langcode_ string, Data_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and langcode = ? and data = ?", Sid_, Langcode_, Data_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndLangcodeAndReindex Get SearchDatasets via SidAndLangcodeAndReindex
func GetSearchDatasetsBySidAndLangcodeAndReindex(offset int, limit int, Sid_ int64, Langcode_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and langcode = ? and reindex = ?", Sid_, Langcode_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndTypeAndData Get SearchDatasets via SidAndTypeAndData
func GetSearchDatasetsBySidAndTypeAndData(offset int, limit int, Sid_ int64, Type_ string, Data_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and type = ? and data = ?", Sid_, Type_, Data_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndTypeAndReindex Get SearchDatasets via SidAndTypeAndReindex
func GetSearchDatasetsBySidAndTypeAndReindex(offset int, limit int, Sid_ int64, Type_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and type = ? and reindex = ?", Sid_, Type_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndDataAndReindex Get SearchDatasets via SidAndDataAndReindex
func GetSearchDatasetsBySidAndDataAndReindex(offset int, limit int, Sid_ int64, Data_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and data = ? and reindex = ?", Sid_, Data_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByLangcodeAndTypeAndData Get SearchDatasets via LangcodeAndTypeAndData
func GetSearchDatasetsByLangcodeAndTypeAndData(offset int, limit int, Langcode_ string, Type_ string, Data_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("langcode = ? and type = ? and data = ?", Langcode_, Type_, Data_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByLangcodeAndTypeAndReindex Get SearchDatasets via LangcodeAndTypeAndReindex
func GetSearchDatasetsByLangcodeAndTypeAndReindex(offset int, limit int, Langcode_ string, Type_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("langcode = ? and type = ? and reindex = ?", Langcode_, Type_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByLangcodeAndDataAndReindex Get SearchDatasets via LangcodeAndDataAndReindex
func GetSearchDatasetsByLangcodeAndDataAndReindex(offset int, limit int, Langcode_ string, Data_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("langcode = ? and data = ? and reindex = ?", Langcode_, Data_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByTypeAndDataAndReindex Get SearchDatasets via TypeAndDataAndReindex
func GetSearchDatasetsByTypeAndDataAndReindex(offset int, limit int, Type_ string, Data_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("type = ? and data = ? and reindex = ?", Type_, Data_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndLangcode Get SearchDatasets via SidAndLangcode
func GetSearchDatasetsBySidAndLangcode(offset int, limit int, Sid_ int64, Langcode_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and langcode = ?", Sid_, Langcode_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndType Get SearchDatasets via SidAndType
func GetSearchDatasetsBySidAndType(offset int, limit int, Sid_ int64, Type_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and type = ?", Sid_, Type_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndData Get SearchDatasets via SidAndData
func GetSearchDatasetsBySidAndData(offset int, limit int, Sid_ int64, Data_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and data = ?", Sid_, Data_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsBySidAndReindex Get SearchDatasets via SidAndReindex
func GetSearchDatasetsBySidAndReindex(offset int, limit int, Sid_ int64, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ? and reindex = ?", Sid_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByLangcodeAndType Get SearchDatasets via LangcodeAndType
func GetSearchDatasetsByLangcodeAndType(offset int, limit int, Langcode_ string, Type_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("langcode = ? and type = ?", Langcode_, Type_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByLangcodeAndData Get SearchDatasets via LangcodeAndData
func GetSearchDatasetsByLangcodeAndData(offset int, limit int, Langcode_ string, Data_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("langcode = ? and data = ?", Langcode_, Data_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByLangcodeAndReindex Get SearchDatasets via LangcodeAndReindex
func GetSearchDatasetsByLangcodeAndReindex(offset int, limit int, Langcode_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("langcode = ? and reindex = ?", Langcode_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByTypeAndData Get SearchDatasets via TypeAndData
func GetSearchDatasetsByTypeAndData(offset int, limit int, Type_ string, Data_ string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("type = ? and data = ?", Type_, Data_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByTypeAndReindex Get SearchDatasets via TypeAndReindex
func GetSearchDatasetsByTypeAndReindex(offset int, limit int, Type_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("type = ? and reindex = ?", Type_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsByDataAndReindex Get SearchDatasets via DataAndReindex
func GetSearchDatasetsByDataAndReindex(offset int, limit int, Data_ string, Reindex_ int) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("data = ? and reindex = ?", Data_, Reindex_).Limit(limit, offset).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasets Get SearchDatasets via field
func GetSearchDatasets(offset int, limit int, field string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Limit(limit, offset).Desc(field).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsViaSid Get SearchDatasets via Sid
func GetSearchDatasetsViaSid(offset int, limit int, Sid_ int64, field string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("sid = ?", Sid_).Limit(limit, offset).Desc(field).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsViaLangcode Get SearchDatasets via Langcode
func GetSearchDatasetsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsViaType Get SearchDatasets via Type
func GetSearchDatasetsViaType(offset int, limit int, Type_ string, field string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("type = ?", Type_).Limit(limit, offset).Desc(field).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsViaData Get SearchDatasets via Data
func GetSearchDatasetsViaData(offset int, limit int, Data_ string, field string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_SearchDataset)
	return _SearchDataset, err
}

// GetSearchDatasetsViaReindex Get SearchDatasets via Reindex
func GetSearchDatasetsViaReindex(offset int, limit int, Reindex_ int, field string) (*[]*SearchDataset, error) {
	var _SearchDataset = new([]*SearchDataset)
	err := Engine.Table("search_dataset").Where("reindex = ?", Reindex_).Limit(limit, offset).Desc(field).Find(_SearchDataset)
	return _SearchDataset, err
}

// HasSearchDatasetViaSid Has SearchDataset via Sid
func HasSearchDatasetViaSid(iSid int64) bool {
	if has, err := Engine.Where("sid = ?", iSid).Get(new(SearchDataset)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchDatasetViaLangcode Has SearchDataset via Langcode
func HasSearchDatasetViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(SearchDataset)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchDatasetViaType Has SearchDataset via Type
func HasSearchDatasetViaType(iType string) bool {
	if has, err := Engine.Where("type = ?", iType).Get(new(SearchDataset)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchDatasetViaData Has SearchDataset via Data
func HasSearchDatasetViaData(iData string) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(SearchDataset)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchDatasetViaReindex Has SearchDataset via Reindex
func HasSearchDatasetViaReindex(iReindex int) bool {
	if has, err := Engine.Where("reindex = ?", iReindex).Get(new(SearchDataset)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSearchDatasetViaSid Get SearchDataset via Sid
func GetSearchDatasetViaSid(iSid int64) (*SearchDataset, error) {
	var _SearchDataset = &SearchDataset{Sid: iSid}
	has, err := Engine.Get(_SearchDataset)
	if has {
		return _SearchDataset, err
	} else {
		return nil, err
	}
}

// GetSearchDatasetViaLangcode Get SearchDataset via Langcode
func GetSearchDatasetViaLangcode(iLangcode string) (*SearchDataset, error) {
	var _SearchDataset = &SearchDataset{Langcode: iLangcode}
	has, err := Engine.Get(_SearchDataset)
	if has {
		return _SearchDataset, err
	} else {
		return nil, err
	}
}

// GetSearchDatasetViaType Get SearchDataset via Type
func GetSearchDatasetViaType(iType string) (*SearchDataset, error) {
	var _SearchDataset = &SearchDataset{Type: iType}
	has, err := Engine.Get(_SearchDataset)
	if has {
		return _SearchDataset, err
	} else {
		return nil, err
	}
}

// GetSearchDatasetViaData Get SearchDataset via Data
func GetSearchDatasetViaData(iData string) (*SearchDataset, error) {
	var _SearchDataset = &SearchDataset{Data: iData}
	has, err := Engine.Get(_SearchDataset)
	if has {
		return _SearchDataset, err
	} else {
		return nil, err
	}
}

// GetSearchDatasetViaReindex Get SearchDataset via Reindex
func GetSearchDatasetViaReindex(iReindex int) (*SearchDataset, error) {
	var _SearchDataset = &SearchDataset{Reindex: iReindex}
	has, err := Engine.Get(_SearchDataset)
	if has {
		return _SearchDataset, err
	} else {
		return nil, err
	}
}

// SetSearchDatasetViaSid Set SearchDataset via Sid
func SetSearchDatasetViaSid(iSid int64, search_dataset *SearchDataset) (int64, error) {
	search_dataset.Sid = iSid
	return Engine.Insert(search_dataset)
}

// SetSearchDatasetViaLangcode Set SearchDataset via Langcode
func SetSearchDatasetViaLangcode(iLangcode string, search_dataset *SearchDataset) (int64, error) {
	search_dataset.Langcode = iLangcode
	return Engine.Insert(search_dataset)
}

// SetSearchDatasetViaType Set SearchDataset via Type
func SetSearchDatasetViaType(iType string, search_dataset *SearchDataset) (int64, error) {
	search_dataset.Type = iType
	return Engine.Insert(search_dataset)
}

// SetSearchDatasetViaData Set SearchDataset via Data
func SetSearchDatasetViaData(iData string, search_dataset *SearchDataset) (int64, error) {
	search_dataset.Data = iData
	return Engine.Insert(search_dataset)
}

// SetSearchDatasetViaReindex Set SearchDataset via Reindex
func SetSearchDatasetViaReindex(iReindex int, search_dataset *SearchDataset) (int64, error) {
	search_dataset.Reindex = iReindex
	return Engine.Insert(search_dataset)
}

// AddSearchDataset Add SearchDataset via all columns
func AddSearchDataset(iSid int64, iLangcode string, iType string, iData string, iReindex int) error {
	_SearchDataset := &SearchDataset{Sid: iSid, Langcode: iLangcode, Type: iType, Data: iData, Reindex: iReindex}
	if _, err := Engine.Insert(_SearchDataset); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSearchDataset Post SearchDataset via iSearchDataset
func PostSearchDataset(iSearchDataset *SearchDataset) (int64, error) {
	_, err := Engine.Insert(iSearchDataset)
	return iSearchDataset.Sid, err
}

// PutSearchDataset Put SearchDataset
func PutSearchDataset(iSearchDataset *SearchDataset) (int64, error) {
	_, err := Engine.Id(iSearchDataset.Sid).Update(iSearchDataset)
	return iSearchDataset.Sid, err
}

// PutSearchDatasetViaSid Put SearchDataset via Sid
func PutSearchDatasetViaSid(Sid_ int64, iSearchDataset *SearchDataset) (int64, error) {
	row, err := Engine.Update(iSearchDataset, &SearchDataset{Sid: Sid_})
	return row, err
}

// PutSearchDatasetViaLangcode Put SearchDataset via Langcode
func PutSearchDatasetViaLangcode(Langcode_ string, iSearchDataset *SearchDataset) (int64, error) {
	row, err := Engine.Update(iSearchDataset, &SearchDataset{Langcode: Langcode_})
	return row, err
}

// PutSearchDatasetViaType Put SearchDataset via Type
func PutSearchDatasetViaType(Type_ string, iSearchDataset *SearchDataset) (int64, error) {
	row, err := Engine.Update(iSearchDataset, &SearchDataset{Type: Type_})
	return row, err
}

// PutSearchDatasetViaData Put SearchDataset via Data
func PutSearchDatasetViaData(Data_ string, iSearchDataset *SearchDataset) (int64, error) {
	row, err := Engine.Update(iSearchDataset, &SearchDataset{Data: Data_})
	return row, err
}

// PutSearchDatasetViaReindex Put SearchDataset via Reindex
func PutSearchDatasetViaReindex(Reindex_ int, iSearchDataset *SearchDataset) (int64, error) {
	row, err := Engine.Update(iSearchDataset, &SearchDataset{Reindex: Reindex_})
	return row, err
}

// UpdateSearchDatasetViaSid via map[string]interface{}{}
func UpdateSearchDatasetViaSid(iSid int64, iSearchDatasetMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchDataset)).Where("sid = ?", iSid).Update(iSearchDatasetMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchDatasetViaLangcode via map[string]interface{}{}
func UpdateSearchDatasetViaLangcode(iLangcode string, iSearchDatasetMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchDataset)).Where("langcode = ?", iLangcode).Update(iSearchDatasetMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchDatasetViaType via map[string]interface{}{}
func UpdateSearchDatasetViaType(iType string, iSearchDatasetMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchDataset)).Where("type = ?", iType).Update(iSearchDatasetMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchDatasetViaData via map[string]interface{}{}
func UpdateSearchDatasetViaData(iData string, iSearchDatasetMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchDataset)).Where("data = ?", iData).Update(iSearchDatasetMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchDatasetViaReindex via map[string]interface{}{}
func UpdateSearchDatasetViaReindex(iReindex int, iSearchDatasetMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchDataset)).Where("reindex = ?", iReindex).Update(iSearchDatasetMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSearchDataset Delete SearchDataset
func DeleteSearchDataset(iSid int64) (int64, error) {
	row, err := Engine.Id(iSid).Delete(new(SearchDataset))
	return row, err
}

// DeleteSearchDatasetViaSid Delete SearchDataset via Sid
func DeleteSearchDatasetViaSid(iSid int64) (err error) {
	var has bool
	var _SearchDataset = &SearchDataset{Sid: iSid}
	if has, err = Engine.Get(_SearchDataset); (has == true) && (err == nil) {
		if row, err := Engine.Where("sid = ?", iSid).Delete(new(SearchDataset)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchDatasetViaLangcode Delete SearchDataset via Langcode
func DeleteSearchDatasetViaLangcode(iLangcode string) (err error) {
	var has bool
	var _SearchDataset = &SearchDataset{Langcode: iLangcode}
	if has, err = Engine.Get(_SearchDataset); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(SearchDataset)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchDatasetViaType Delete SearchDataset via Type
func DeleteSearchDatasetViaType(iType string) (err error) {
	var has bool
	var _SearchDataset = &SearchDataset{Type: iType}
	if has, err = Engine.Get(_SearchDataset); (has == true) && (err == nil) {
		if row, err := Engine.Where("type = ?", iType).Delete(new(SearchDataset)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchDatasetViaData Delete SearchDataset via Data
func DeleteSearchDatasetViaData(iData string) (err error) {
	var has bool
	var _SearchDataset = &SearchDataset{Data: iData}
	if has, err = Engine.Get(_SearchDataset); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(SearchDataset)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchDatasetViaReindex Delete SearchDataset via Reindex
func DeleteSearchDatasetViaReindex(iReindex int) (err error) {
	var has bool
	var _SearchDataset = &SearchDataset{Reindex: iReindex}
	if has, err = Engine.Get(_SearchDataset); (has == true) && (err == nil) {
		if row, err := Engine.Where("reindex = ?", iReindex).Delete(new(SearchDataset)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
