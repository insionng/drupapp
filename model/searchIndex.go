package model

type SearchIndex struct {
	Word     string  `xorm:"not null pk default '' VARCHAR(50)"`
	Sid      int     `xorm:"not null pk default 0 index(sid_type) INT(10)"`
	Langcode string  `xorm:"not null pk default '' index(sid_type) VARCHAR(12)"`
	Type     string  `xorm:"not null pk index(sid_type) VARCHAR(64)"`
	Score    float32 `xorm:"FLOAT"`
}

// GetSearchIndexsCount SearchIndexs' Count
func GetSearchIndexsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&SearchIndex{})
	return total, err
}

// GetSearchIndexCountViaWord Get SearchIndex via Word
func GetSearchIndexCountViaWord(iWord string) int64 {
	n, _ := Engine.Where("word = ?", iWord).Count(&SearchIndex{Word: iWord})
	return n
}

// GetSearchIndexCountViaSid Get SearchIndex via Sid
func GetSearchIndexCountViaSid(iSid int) int64 {
	n, _ := Engine.Where("sid = ?", iSid).Count(&SearchIndex{Sid: iSid})
	return n
}

// GetSearchIndexCountViaLangcode Get SearchIndex via Langcode
func GetSearchIndexCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&SearchIndex{Langcode: iLangcode})
	return n
}

// GetSearchIndexCountViaType Get SearchIndex via Type
func GetSearchIndexCountViaType(iType string) int64 {
	n, _ := Engine.Where("type = ?", iType).Count(&SearchIndex{Type: iType})
	return n
}

// GetSearchIndexCountViaScore Get SearchIndex via Score
func GetSearchIndexCountViaScore(iScore float32) int64 {
	n, _ := Engine.Where("score = ?", iScore).Count(&SearchIndex{Score: iScore})
	return n
}

// GetSearchIndexsByWordAndSidAndLangcode Get SearchIndexs via WordAndSidAndLangcode
func GetSearchIndexsByWordAndSidAndLangcode(offset int, limit int, Word_ string, Sid_ int, Langcode_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and sid = ? and langcode = ?", Word_, Sid_, Langcode_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndSidAndType Get SearchIndexs via WordAndSidAndType
func GetSearchIndexsByWordAndSidAndType(offset int, limit int, Word_ string, Sid_ int, Type_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and sid = ? and type = ?", Word_, Sid_, Type_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndSidAndScore Get SearchIndexs via WordAndSidAndScore
func GetSearchIndexsByWordAndSidAndScore(offset int, limit int, Word_ string, Sid_ int, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and sid = ? and score = ?", Word_, Sid_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndLangcodeAndType Get SearchIndexs via WordAndLangcodeAndType
func GetSearchIndexsByWordAndLangcodeAndType(offset int, limit int, Word_ string, Langcode_ string, Type_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and langcode = ? and type = ?", Word_, Langcode_, Type_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndLangcodeAndScore Get SearchIndexs via WordAndLangcodeAndScore
func GetSearchIndexsByWordAndLangcodeAndScore(offset int, limit int, Word_ string, Langcode_ string, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and langcode = ? and score = ?", Word_, Langcode_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndTypeAndScore Get SearchIndexs via WordAndTypeAndScore
func GetSearchIndexsByWordAndTypeAndScore(offset int, limit int, Word_ string, Type_ string, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and type = ? and score = ?", Word_, Type_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsBySidAndLangcodeAndType Get SearchIndexs via SidAndLangcodeAndType
func GetSearchIndexsBySidAndLangcodeAndType(offset int, limit int, Sid_ int, Langcode_ string, Type_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("sid = ? and langcode = ? and type = ?", Sid_, Langcode_, Type_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsBySidAndLangcodeAndScore Get SearchIndexs via SidAndLangcodeAndScore
func GetSearchIndexsBySidAndLangcodeAndScore(offset int, limit int, Sid_ int, Langcode_ string, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("sid = ? and langcode = ? and score = ?", Sid_, Langcode_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsBySidAndTypeAndScore Get SearchIndexs via SidAndTypeAndScore
func GetSearchIndexsBySidAndTypeAndScore(offset int, limit int, Sid_ int, Type_ string, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("sid = ? and type = ? and score = ?", Sid_, Type_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByLangcodeAndTypeAndScore Get SearchIndexs via LangcodeAndTypeAndScore
func GetSearchIndexsByLangcodeAndTypeAndScore(offset int, limit int, Langcode_ string, Type_ string, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("langcode = ? and type = ? and score = ?", Langcode_, Type_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndSid Get SearchIndexs via WordAndSid
func GetSearchIndexsByWordAndSid(offset int, limit int, Word_ string, Sid_ int) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and sid = ?", Word_, Sid_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndLangcode Get SearchIndexs via WordAndLangcode
func GetSearchIndexsByWordAndLangcode(offset int, limit int, Word_ string, Langcode_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and langcode = ?", Word_, Langcode_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndType Get SearchIndexs via WordAndType
func GetSearchIndexsByWordAndType(offset int, limit int, Word_ string, Type_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and type = ?", Word_, Type_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByWordAndScore Get SearchIndexs via WordAndScore
func GetSearchIndexsByWordAndScore(offset int, limit int, Word_ string, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ? and score = ?", Word_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsBySidAndLangcode Get SearchIndexs via SidAndLangcode
func GetSearchIndexsBySidAndLangcode(offset int, limit int, Sid_ int, Langcode_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("sid = ? and langcode = ?", Sid_, Langcode_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsBySidAndType Get SearchIndexs via SidAndType
func GetSearchIndexsBySidAndType(offset int, limit int, Sid_ int, Type_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("sid = ? and type = ?", Sid_, Type_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsBySidAndScore Get SearchIndexs via SidAndScore
func GetSearchIndexsBySidAndScore(offset int, limit int, Sid_ int, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("sid = ? and score = ?", Sid_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByLangcodeAndType Get SearchIndexs via LangcodeAndType
func GetSearchIndexsByLangcodeAndType(offset int, limit int, Langcode_ string, Type_ string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("langcode = ? and type = ?", Langcode_, Type_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByLangcodeAndScore Get SearchIndexs via LangcodeAndScore
func GetSearchIndexsByLangcodeAndScore(offset int, limit int, Langcode_ string, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("langcode = ? and score = ?", Langcode_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsByTypeAndScore Get SearchIndexs via TypeAndScore
func GetSearchIndexsByTypeAndScore(offset int, limit int, Type_ string, Score_ float32) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("type = ? and score = ?", Type_, Score_).Limit(limit, offset).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexs Get SearchIndexs via field
func GetSearchIndexs(offset int, limit int, field string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Limit(limit, offset).Desc(field).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsViaWord Get SearchIndexs via Word
func GetSearchIndexsViaWord(offset int, limit int, Word_ string, field string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("word = ?", Word_).Limit(limit, offset).Desc(field).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsViaSid Get SearchIndexs via Sid
func GetSearchIndexsViaSid(offset int, limit int, Sid_ int, field string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("sid = ?", Sid_).Limit(limit, offset).Desc(field).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsViaLangcode Get SearchIndexs via Langcode
func GetSearchIndexsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsViaType Get SearchIndexs via Type
func GetSearchIndexsViaType(offset int, limit int, Type_ string, field string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("type = ?", Type_).Limit(limit, offset).Desc(field).Find(_SearchIndex)
	return _SearchIndex, err
}

// GetSearchIndexsViaScore Get SearchIndexs via Score
func GetSearchIndexsViaScore(offset int, limit int, Score_ float32, field string) (*[]*SearchIndex, error) {
	var _SearchIndex = new([]*SearchIndex)
	err := Engine.Table("search_index").Where("score = ?", Score_).Limit(limit, offset).Desc(field).Find(_SearchIndex)
	return _SearchIndex, err
}

// HasSearchIndexViaWord Has SearchIndex via Word
func HasSearchIndexViaWord(iWord string) bool {
	if has, err := Engine.Where("word = ?", iWord).Get(new(SearchIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchIndexViaSid Has SearchIndex via Sid
func HasSearchIndexViaSid(iSid int) bool {
	if has, err := Engine.Where("sid = ?", iSid).Get(new(SearchIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchIndexViaLangcode Has SearchIndex via Langcode
func HasSearchIndexViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(SearchIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchIndexViaType Has SearchIndex via Type
func HasSearchIndexViaType(iType string) bool {
	if has, err := Engine.Where("type = ?", iType).Get(new(SearchIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchIndexViaScore Has SearchIndex via Score
func HasSearchIndexViaScore(iScore float32) bool {
	if has, err := Engine.Where("score = ?", iScore).Get(new(SearchIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSearchIndexViaWord Get SearchIndex via Word
func GetSearchIndexViaWord(iWord string) (*SearchIndex, error) {
	var _SearchIndex = &SearchIndex{Word: iWord}
	has, err := Engine.Get(_SearchIndex)
	if has {
		return _SearchIndex, err
	} else {
		return nil, err
	}
}

// GetSearchIndexViaSid Get SearchIndex via Sid
func GetSearchIndexViaSid(iSid int) (*SearchIndex, error) {
	var _SearchIndex = &SearchIndex{Sid: iSid}
	has, err := Engine.Get(_SearchIndex)
	if has {
		return _SearchIndex, err
	} else {
		return nil, err
	}
}

// GetSearchIndexViaLangcode Get SearchIndex via Langcode
func GetSearchIndexViaLangcode(iLangcode string) (*SearchIndex, error) {
	var _SearchIndex = &SearchIndex{Langcode: iLangcode}
	has, err := Engine.Get(_SearchIndex)
	if has {
		return _SearchIndex, err
	} else {
		return nil, err
	}
}

// GetSearchIndexViaType Get SearchIndex via Type
func GetSearchIndexViaType(iType string) (*SearchIndex, error) {
	var _SearchIndex = &SearchIndex{Type: iType}
	has, err := Engine.Get(_SearchIndex)
	if has {
		return _SearchIndex, err
	} else {
		return nil, err
	}
}

// GetSearchIndexViaScore Get SearchIndex via Score
func GetSearchIndexViaScore(iScore float32) (*SearchIndex, error) {
	var _SearchIndex = &SearchIndex{Score: iScore}
	has, err := Engine.Get(_SearchIndex)
	if has {
		return _SearchIndex, err
	} else {
		return nil, err
	}
}

// SetSearchIndexViaWord Set SearchIndex via Word
func SetSearchIndexViaWord(iWord string, search_index *SearchIndex) (int64, error) {
	search_index.Word = iWord
	return Engine.Insert(search_index)
}

// SetSearchIndexViaSid Set SearchIndex via Sid
func SetSearchIndexViaSid(iSid int, search_index *SearchIndex) (int64, error) {
	search_index.Sid = iSid
	return Engine.Insert(search_index)
}

// SetSearchIndexViaLangcode Set SearchIndex via Langcode
func SetSearchIndexViaLangcode(iLangcode string, search_index *SearchIndex) (int64, error) {
	search_index.Langcode = iLangcode
	return Engine.Insert(search_index)
}

// SetSearchIndexViaType Set SearchIndex via Type
func SetSearchIndexViaType(iType string, search_index *SearchIndex) (int64, error) {
	search_index.Type = iType
	return Engine.Insert(search_index)
}

// SetSearchIndexViaScore Set SearchIndex via Score
func SetSearchIndexViaScore(iScore float32, search_index *SearchIndex) (int64, error) {
	search_index.Score = iScore
	return Engine.Insert(search_index)
}

// AddSearchIndex Add SearchIndex via all columns
func AddSearchIndex(iWord string, iSid int, iLangcode string, iType string, iScore float32) error {
	_SearchIndex := &SearchIndex{Word: iWord, Sid: iSid, Langcode: iLangcode, Type: iType, Score: iScore}
	if _, err := Engine.Insert(_SearchIndex); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSearchIndex Post SearchIndex via iSearchIndex
func PostSearchIndex(iSearchIndex *SearchIndex) (string, error) {
	_, err := Engine.Insert(iSearchIndex)
	return iSearchIndex.Word, err
}

// PutSearchIndex Put SearchIndex
func PutSearchIndex(iSearchIndex *SearchIndex) (string, error) {
	_, err := Engine.Id(iSearchIndex.Word).Update(iSearchIndex)
	return iSearchIndex.Word, err
}

// PutSearchIndexViaWord Put SearchIndex via Word
func PutSearchIndexViaWord(Word_ string, iSearchIndex *SearchIndex) (int64, error) {
	row, err := Engine.Update(iSearchIndex, &SearchIndex{Word: Word_})
	return row, err
}

// PutSearchIndexViaSid Put SearchIndex via Sid
func PutSearchIndexViaSid(Sid_ int, iSearchIndex *SearchIndex) (int64, error) {
	row, err := Engine.Update(iSearchIndex, &SearchIndex{Sid: Sid_})
	return row, err
}

// PutSearchIndexViaLangcode Put SearchIndex via Langcode
func PutSearchIndexViaLangcode(Langcode_ string, iSearchIndex *SearchIndex) (int64, error) {
	row, err := Engine.Update(iSearchIndex, &SearchIndex{Langcode: Langcode_})
	return row, err
}

// PutSearchIndexViaType Put SearchIndex via Type
func PutSearchIndexViaType(Type_ string, iSearchIndex *SearchIndex) (int64, error) {
	row, err := Engine.Update(iSearchIndex, &SearchIndex{Type: Type_})
	return row, err
}

// PutSearchIndexViaScore Put SearchIndex via Score
func PutSearchIndexViaScore(Score_ float32, iSearchIndex *SearchIndex) (int64, error) {
	row, err := Engine.Update(iSearchIndex, &SearchIndex{Score: Score_})
	return row, err
}

// UpdateSearchIndexViaWord via map[string]interface{}{}
func UpdateSearchIndexViaWord(iWord string, iSearchIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchIndex)).Where("word = ?", iWord).Update(iSearchIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchIndexViaSid via map[string]interface{}{}
func UpdateSearchIndexViaSid(iSid int, iSearchIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchIndex)).Where("sid = ?", iSid).Update(iSearchIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchIndexViaLangcode via map[string]interface{}{}
func UpdateSearchIndexViaLangcode(iLangcode string, iSearchIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchIndex)).Where("langcode = ?", iLangcode).Update(iSearchIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchIndexViaType via map[string]interface{}{}
func UpdateSearchIndexViaType(iType string, iSearchIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchIndex)).Where("type = ?", iType).Update(iSearchIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchIndexViaScore via map[string]interface{}{}
func UpdateSearchIndexViaScore(iScore float32, iSearchIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchIndex)).Where("score = ?", iScore).Update(iSearchIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSearchIndex Delete SearchIndex
func DeleteSearchIndex(iWord string) (int64, error) {
	row, err := Engine.Id(iWord).Delete(new(SearchIndex))
	return row, err
}

// DeleteSearchIndexViaWord Delete SearchIndex via Word
func DeleteSearchIndexViaWord(iWord string) (err error) {
	var has bool
	var _SearchIndex = &SearchIndex{Word: iWord}
	if has, err = Engine.Get(_SearchIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("word = ?", iWord).Delete(new(SearchIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchIndexViaSid Delete SearchIndex via Sid
func DeleteSearchIndexViaSid(iSid int) (err error) {
	var has bool
	var _SearchIndex = &SearchIndex{Sid: iSid}
	if has, err = Engine.Get(_SearchIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("sid = ?", iSid).Delete(new(SearchIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchIndexViaLangcode Delete SearchIndex via Langcode
func DeleteSearchIndexViaLangcode(iLangcode string) (err error) {
	var has bool
	var _SearchIndex = &SearchIndex{Langcode: iLangcode}
	if has, err = Engine.Get(_SearchIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(SearchIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchIndexViaType Delete SearchIndex via Type
func DeleteSearchIndexViaType(iType string) (err error) {
	var has bool
	var _SearchIndex = &SearchIndex{Type: iType}
	if has, err = Engine.Get(_SearchIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("type = ?", iType).Delete(new(SearchIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchIndexViaScore Delete SearchIndex via Score
func DeleteSearchIndexViaScore(iScore float32) (err error) {
	var has bool
	var _SearchIndex = &SearchIndex{Score: iScore}
	if has, err = Engine.Get(_SearchIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("score = ?", iScore).Delete(new(SearchIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
