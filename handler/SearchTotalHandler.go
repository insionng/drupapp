package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetSearchTotalsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSearchTotalsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["search_totalsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSearchTotalsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchTotalCountViaWordHandler(self *macross.Context) error {
	Word_ := self.Args("word").String()
	_int64 := model.GetSearchTotalCountViaWord(Word_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_totalCount"] = 0
	}
	m["search_totalCount"] = _int64
	return self.JSON(m)
}

func GetSearchTotalCountViaCountHandler(self *macross.Context) error {
	Count_ := self.Args("count").MustFloat32()
	_int64 := model.GetSearchTotalCountViaCount(Count_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_totalCount"] = 0
	}
	m["search_totalCount"] = _int64
	return self.JSON(m)
}

func GetSearchTotalsViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iWord := self.Args("word").String()
	if (offset > 0) && helper.IsHas(iWord) {
		_SearchTotal, _error := model.GetSearchTotalsViaWord(offset, limit, iWord, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchTotal)
	}
	herr.Message = "Can't get to the GetSearchTotalsViaWord's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchTotalsViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCount := self.Args("count").MustFloat32()
	if (offset > 0) && helper.IsHas(iCount) {
		_SearchTotal, _error := model.GetSearchTotalsViaCount(offset, limit, iCount, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchTotal)
	}
	herr.Message = "Can't get to the GetSearchTotalsViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchTotalsByWordAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iCount := self.Args("count").MustFloat32()

	if helper.IsHas(iWord) {
		_SearchTotal, _error := model.GetSearchTotalsByWordAndCount(offset, limit, iWord,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchTotal)
	}
	herr.Message = "Can't get to the GetSearchTotalsByWordAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchTotalsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_SearchTotal, _error := model.GetSearchTotals(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchTotal)
	}
	herr.Message = "Can't get to the GetSearchTotals' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchTotalViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWord := self.Args("word").String()
	if helper.IsHas(iWord) {
		_SearchTotal := model.HasSearchTotalViaWord(iWord)
		var m = map[string]interface{}{}
		m["search_total"] = _SearchTotal
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchTotalViaWord's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchTotalViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCount := self.Args("count").MustFloat32()
	if helper.IsHas(iCount) {
		_SearchTotal := model.HasSearchTotalViaCount(iCount)
		var m = map[string]interface{}{}
		m["search_total"] = _SearchTotal
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchTotalViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchTotalViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWord := self.Args("word").String()
	if helper.IsHas(iWord) {
		_SearchTotal, _error := model.GetSearchTotalViaWord(iWord)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchTotal)
	}
	herr.Message = "Can't get to the GetSearchTotalViaWord's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchTotalViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCount := self.Args("count").MustFloat32()
	if helper.IsHas(iCount) {
		_SearchTotal, _error := model.GetSearchTotalViaCount(iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchTotal)
	}
	herr.Message = "Can't get to the GetSearchTotalViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchTotalViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	if helper.IsHas(Word_) {
		var iSearchTotal model.SearchTotal
		self.Bind(&iSearchTotal)
		_SearchTotal, _error := model.SetSearchTotalViaWord(Word_, &iSearchTotal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchTotal)
	}
	herr.Message = "Can't get to the SetSearchTotalViaWord's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchTotalViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustFloat32()
	if helper.IsHas(Count_) {
		var iSearchTotal model.SearchTotal
		self.Bind(&iSearchTotal)
		_SearchTotal, _error := model.SetSearchTotalViaCount(Count_, &iSearchTotal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchTotal)
	}
	herr.Message = "Can't get to the SetSearchTotalViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSearchTotalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	Count_ := self.Args("count").MustFloat32()

	if helper.IsHas(Word_) {
		_error := model.AddSearchTotal(Word_,Count_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSearchTotal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSearchTotalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSearchTotal model.SearchTotal
	self.Bind(&iSearchTotal)
	_string, _error := model.PostSearchTotal(&iSearchTotal)
	if (helper.IsHas(_string)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["created"] = _string
		return self.JSON(m, macross.StatusCreated)
	}
	return self.JSON(herr)
}

func PutSearchTotalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSearchTotal model.SearchTotal
	self.Bind(&iSearchTotal)
	_string, _error := model.PutSearchTotal(&iSearchTotal)
	if (helper.IsHas(_string)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["updated"] = _string
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func PutSearchTotalViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	var iSearchTotal model.SearchTotal
	self.Bind(&iSearchTotal)
	_int64, _error := model.PutSearchTotalViaWord(Word_, &iSearchTotal)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchTotalViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustFloat32()
	var iSearchTotal model.SearchTotal
	self.Bind(&iSearchTotal)
	_int64, _error := model.PutSearchTotalViaCount(Count_, &iSearchTotal)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchTotalViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	var iSearchTotal model.SearchTotal
	self.Bind(&iSearchTotal)
	var iMap = helper.StructToMap(iSearchTotal)
	_error := model.UpdateSearchTotalViaWord(Word_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchTotalViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustFloat32()
	var iSearchTotal model.SearchTotal
	self.Bind(&iSearchTotal)
	var iMap = helper.StructToMap(iSearchTotal)
	_error := model.UpdateSearchTotalViaCount(Count_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchTotalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	_int64, _error := model.DeleteSearchTotal(Word_)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["deleted"] = _int64
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func DeleteSearchTotalViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	_error := model.DeleteSearchTotalViaWord(Word_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchTotalViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustFloat32()
	_error := model.DeleteSearchTotalViaCount(Count_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
