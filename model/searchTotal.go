package model

type SearchTotal struct {
	Word  string  `xorm:"not null pk default '' VARCHAR(50)"`
	Count float32 `xorm:"FLOAT"`
}

// GetSearchTotalsCount SearchTotals' Count
func GetSearchTotalsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&SearchTotal{})
	return total, err
}

// GetSearchTotalCountViaWord Get SearchTotal via Word
func GetSearchTotalCountViaWord(iWord string) int64 {
	n, _ := Engine.Where("word = ?", iWord).Count(&SearchTotal{Word: iWord})
	return n
}

// GetSearchTotalCountViaCount Get SearchTotal via Count
func GetSearchTotalCountViaCount(iCount float32) int64 {
	n, _ := Engine.Where("count = ?", iCount).Count(&SearchTotal{Count: iCount})
	return n
}

// GetSearchTotalsByWordAndCount Get SearchTotals via WordAndCount
func GetSearchTotalsByWordAndCount(offset int, limit int, Word_ string, Count_ float32) (*[]*SearchTotal, error) {
	var _SearchTotal = new([]*SearchTotal)
	err := Engine.Table("search_total").Where("word = ? and count = ?", Word_, Count_).Limit(limit, offset).Find(_SearchTotal)
	return _SearchTotal, err
}

// GetSearchTotals Get SearchTotals via field
func GetSearchTotals(offset int, limit int, field string) (*[]*SearchTotal, error) {
	var _SearchTotal = new([]*SearchTotal)
	err := Engine.Table("search_total").Limit(limit, offset).Desc(field).Find(_SearchTotal)
	return _SearchTotal, err
}

// GetSearchTotalsViaWord Get SearchTotals via Word
func GetSearchTotalsViaWord(offset int, limit int, Word_ string, field string) (*[]*SearchTotal, error) {
	var _SearchTotal = new([]*SearchTotal)
	err := Engine.Table("search_total").Where("word = ?", Word_).Limit(limit, offset).Desc(field).Find(_SearchTotal)
	return _SearchTotal, err
}

// GetSearchTotalsViaCount Get SearchTotals via Count
func GetSearchTotalsViaCount(offset int, limit int, Count_ float32, field string) (*[]*SearchTotal, error) {
	var _SearchTotal = new([]*SearchTotal)
	err := Engine.Table("search_total").Where("count = ?", Count_).Limit(limit, offset).Desc(field).Find(_SearchTotal)
	return _SearchTotal, err
}

// HasSearchTotalViaWord Has SearchTotal via Word
func HasSearchTotalViaWord(iWord string) bool {
	if has, err := Engine.Where("word = ?", iWord).Get(new(SearchTotal)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSearchTotalViaCount Has SearchTotal via Count
func HasSearchTotalViaCount(iCount float32) bool {
	if has, err := Engine.Where("count = ?", iCount).Get(new(SearchTotal)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSearchTotalViaWord Get SearchTotal via Word
func GetSearchTotalViaWord(iWord string) (*SearchTotal, error) {
	var _SearchTotal = &SearchTotal{Word: iWord}
	has, err := Engine.Get(_SearchTotal)
	if has {
		return _SearchTotal, err
	} else {
		return nil, err
	}
}

// GetSearchTotalViaCount Get SearchTotal via Count
func GetSearchTotalViaCount(iCount float32) (*SearchTotal, error) {
	var _SearchTotal = &SearchTotal{Count: iCount}
	has, err := Engine.Get(_SearchTotal)
	if has {
		return _SearchTotal, err
	} else {
		return nil, err
	}
}

// SetSearchTotalViaWord Set SearchTotal via Word
func SetSearchTotalViaWord(iWord string, search_total *SearchTotal) (int64, error) {
	search_total.Word = iWord
	return Engine.Insert(search_total)
}

// SetSearchTotalViaCount Set SearchTotal via Count
func SetSearchTotalViaCount(iCount float32, search_total *SearchTotal) (int64, error) {
	search_total.Count = iCount
	return Engine.Insert(search_total)
}

// AddSearchTotal Add SearchTotal via all columns
func AddSearchTotal(iWord string, iCount float32) error {
	_SearchTotal := &SearchTotal{Word: iWord, Count: iCount}
	if _, err := Engine.Insert(_SearchTotal); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSearchTotal Post SearchTotal via iSearchTotal
func PostSearchTotal(iSearchTotal *SearchTotal) (string, error) {
	_, err := Engine.Insert(iSearchTotal)
	return iSearchTotal.Word, err
}

// PutSearchTotal Put SearchTotal
func PutSearchTotal(iSearchTotal *SearchTotal) (string, error) {
	_, err := Engine.Id(iSearchTotal.Word).Update(iSearchTotal)
	return iSearchTotal.Word, err
}

// PutSearchTotalViaWord Put SearchTotal via Word
func PutSearchTotalViaWord(Word_ string, iSearchTotal *SearchTotal) (int64, error) {
	row, err := Engine.Update(iSearchTotal, &SearchTotal{Word: Word_})
	return row, err
}

// PutSearchTotalViaCount Put SearchTotal via Count
func PutSearchTotalViaCount(Count_ float32, iSearchTotal *SearchTotal) (int64, error) {
	row, err := Engine.Update(iSearchTotal, &SearchTotal{Count: Count_})
	return row, err
}

// UpdateSearchTotalViaWord via map[string]interface{}{}
func UpdateSearchTotalViaWord(iWord string, iSearchTotalMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchTotal)).Where("word = ?", iWord).Update(iSearchTotalMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSearchTotalViaCount via map[string]interface{}{}
func UpdateSearchTotalViaCount(iCount float32, iSearchTotalMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(SearchTotal)).Where("count = ?", iCount).Update(iSearchTotalMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSearchTotal Delete SearchTotal
func DeleteSearchTotal(iWord string) (int64, error) {
	row, err := Engine.Id(iWord).Delete(new(SearchTotal))
	return row, err
}

// DeleteSearchTotalViaWord Delete SearchTotal via Word
func DeleteSearchTotalViaWord(iWord string) (err error) {
	var has bool
	var _SearchTotal = &SearchTotal{Word: iWord}
	if has, err = Engine.Get(_SearchTotal); (has == true) && (err == nil) {
		if row, err := Engine.Where("word = ?", iWord).Delete(new(SearchTotal)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSearchTotalViaCount Delete SearchTotal via Count
func DeleteSearchTotalViaCount(iCount float32) (err error) {
	var has bool
	var _SearchTotal = &SearchTotal{Count: iCount}
	if has, err = Engine.Get(_SearchTotal); (has == true) && (err == nil) {
		if row, err := Engine.Where("count = ?", iCount).Delete(new(SearchTotal)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
