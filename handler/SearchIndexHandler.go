package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetSearchIndexsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSearchIndexsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["search_indexsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSearchIndexsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexCountViaWordHandler(self *macross.Context) error {
	Word_ := self.Args("word").String()
	_int64 := model.GetSearchIndexCountViaWord(Word_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_indexCount"] = 0
	}
	m["search_indexCount"] = _int64
	return self.JSON(m)
}

func GetSearchIndexCountViaSidHandler(self *macross.Context) error {
	Sid_ := self.Args("sid").MustInt()
	_int64 := model.GetSearchIndexCountViaSid(Sid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_indexCount"] = 0
	}
	m["search_indexCount"] = _int64
	return self.JSON(m)
}

func GetSearchIndexCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetSearchIndexCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_indexCount"] = 0
	}
	m["search_indexCount"] = _int64
	return self.JSON(m)
}

func GetSearchIndexCountViaTypeHandler(self *macross.Context) error {
	Type_ := self.Args("type").String()
	_int64 := model.GetSearchIndexCountViaType(Type_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_indexCount"] = 0
	}
	m["search_indexCount"] = _int64
	return self.JSON(m)
}

func GetSearchIndexCountViaScoreHandler(self *macross.Context) error {
	Score_ := self.Args("score").MustFloat32()
	_int64 := model.GetSearchIndexCountViaScore(Score_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_indexCount"] = 0
	}
	m["search_indexCount"] = _int64
	return self.JSON(m)
}

func GetSearchIndexsViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iWord := self.Args("word").String()
	if (offset > 0) && helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsViaWord(offset, limit, iWord, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsViaWord's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSid := self.Args("sid").MustInt()
	if (offset > 0) && helper.IsHas(iSid) {
		_SearchIndex, _error := model.GetSearchIndexsViaSid(offset, limit, iSid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_SearchIndex, _error := model.GetSearchIndexsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iType := self.Args("type").String()
	if (offset > 0) && helper.IsHas(iType) {
		_SearchIndex, _error := model.GetSearchIndexsViaType(offset, limit, iType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsViaScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iScore := self.Args("score").MustFloat32()
	if (offset > 0) && helper.IsHas(iScore) {
		_SearchIndex, _error := model.GetSearchIndexsViaScore(offset, limit, iScore, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsViaScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndSidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iSid := self.Args("sid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndSidAndLangcode(offset, limit, iWord,iSid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndSidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndSidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iSid := self.Args("sid").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndSidAndType(offset, limit, iWord,iSid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndSidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndSidAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iSid := self.Args("sid").MustInt()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndSidAndScore(offset, limit, iWord,iSid,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndSidAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndLangcodeAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iLangcode := self.Args("langcode").String()
	iType := self.Args("type").String()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndLangcodeAndType(offset, limit, iWord,iLangcode,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndLangcodeAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndLangcodeAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iLangcode := self.Args("langcode").String()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndLangcodeAndScore(offset, limit, iWord,iLangcode,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndLangcodeAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndTypeAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iType := self.Args("type").String()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndTypeAndScore(offset, limit, iWord,iType,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndTypeAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsBySidAndLangcodeAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt()
	iLangcode := self.Args("langcode").String()
	iType := self.Args("type").String()

	if helper.IsHas(iSid) {
		_SearchIndex, _error := model.GetSearchIndexsBySidAndLangcodeAndType(offset, limit, iSid,iLangcode,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsBySidAndLangcodeAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsBySidAndLangcodeAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt()
	iLangcode := self.Args("langcode").String()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iSid) {
		_SearchIndex, _error := model.GetSearchIndexsBySidAndLangcodeAndScore(offset, limit, iSid,iLangcode,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsBySidAndLangcodeAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsBySidAndTypeAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt()
	iType := self.Args("type").String()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iSid) {
		_SearchIndex, _error := model.GetSearchIndexsBySidAndTypeAndScore(offset, limit, iSid,iType,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsBySidAndTypeAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByLangcodeAndTypeAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iType := self.Args("type").String()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iLangcode) {
		_SearchIndex, _error := model.GetSearchIndexsByLangcodeAndTypeAndScore(offset, limit, iLangcode,iType,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByLangcodeAndTypeAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iSid := self.Args("sid").MustInt()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndSid(offset, limit, iWord,iSid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndLangcode(offset, limit, iWord,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iType := self.Args("type").String()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndType(offset, limit, iWord,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByWordAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWord := self.Args("word").String()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexsByWordAndScore(offset, limit, iWord,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByWordAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsBySidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iSid) {
		_SearchIndex, _error := model.GetSearchIndexsBySidAndLangcode(offset, limit, iSid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsBySidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsBySidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iSid) {
		_SearchIndex, _error := model.GetSearchIndexsBySidAndType(offset, limit, iSid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsBySidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsBySidAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iSid) {
		_SearchIndex, _error := model.GetSearchIndexsBySidAndScore(offset, limit, iSid,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsBySidAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByLangcodeAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iType := self.Args("type").String()

	if helper.IsHas(iLangcode) {
		_SearchIndex, _error := model.GetSearchIndexsByLangcodeAndType(offset, limit, iLangcode,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByLangcodeAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByLangcodeAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iLangcode) {
		_SearchIndex, _error := model.GetSearchIndexsByLangcodeAndScore(offset, limit, iLangcode,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByLangcodeAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsByTypeAndScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iScore := self.Args("score").MustFloat32()

	if helper.IsHas(iType) {
		_SearchIndex, _error := model.GetSearchIndexsByTypeAndScore(offset, limit, iType,iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexsByTypeAndScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_SearchIndex, _error := model.GetSearchIndexs(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexs' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchIndexViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWord := self.Args("word").String()
	if helper.IsHas(iWord) {
		_SearchIndex := model.HasSearchIndexViaWord(iWord)
		var m = map[string]interface{}{}
		m["search_index"] = _SearchIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchIndexViaWord's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchIndexViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSid := self.Args("sid").MustInt()
	if helper.IsHas(iSid) {
		_SearchIndex := model.HasSearchIndexViaSid(iSid)
		var m = map[string]interface{}{}
		m["search_index"] = _SearchIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchIndexViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchIndexViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_SearchIndex := model.HasSearchIndexViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["search_index"] = _SearchIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchIndexViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchIndexViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_SearchIndex := model.HasSearchIndexViaType(iType)
		var m = map[string]interface{}{}
		m["search_index"] = _SearchIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchIndexViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchIndexViaScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iScore := self.Args("score").MustFloat32()
	if helper.IsHas(iScore) {
		_SearchIndex := model.HasSearchIndexViaScore(iScore)
		var m = map[string]interface{}{}
		m["search_index"] = _SearchIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchIndexViaScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWord := self.Args("word").String()
	if helper.IsHas(iWord) {
		_SearchIndex, _error := model.GetSearchIndexViaWord(iWord)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexViaWord's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSid := self.Args("sid").MustInt()
	if helper.IsHas(iSid) {
		_SearchIndex, _error := model.GetSearchIndexViaSid(iSid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_SearchIndex, _error := model.GetSearchIndexViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_SearchIndex, _error := model.GetSearchIndexViaType(iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchIndexViaScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iScore := self.Args("score").MustFloat32()
	if helper.IsHas(iScore) {
		_SearchIndex, _error := model.GetSearchIndexViaScore(iScore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the GetSearchIndexViaScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchIndexViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	if helper.IsHas(Word_) {
		var iSearchIndex model.SearchIndex
		self.Bind(&iSearchIndex)
		_SearchIndex, _error := model.SetSearchIndexViaWord(Word_, &iSearchIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the SetSearchIndexViaWord's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchIndexViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt()
	if helper.IsHas(Sid_) {
		var iSearchIndex model.SearchIndex
		self.Bind(&iSearchIndex)
		_SearchIndex, _error := model.SetSearchIndexViaSid(Sid_, &iSearchIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the SetSearchIndexViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchIndexViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iSearchIndex model.SearchIndex
		self.Bind(&iSearchIndex)
		_SearchIndex, _error := model.SetSearchIndexViaLangcode(Langcode_, &iSearchIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the SetSearchIndexViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchIndexViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	if helper.IsHas(Type_) {
		var iSearchIndex model.SearchIndex
		self.Bind(&iSearchIndex)
		_SearchIndex, _error := model.SetSearchIndexViaType(Type_, &iSearchIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the SetSearchIndexViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchIndexViaScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Score_ := self.Args("score").MustFloat32()
	if helper.IsHas(Score_) {
		var iSearchIndex model.SearchIndex
		self.Bind(&iSearchIndex)
		_SearchIndex, _error := model.SetSearchIndexViaScore(Score_, &iSearchIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchIndex)
	}
	herr.Message = "Can't get to the SetSearchIndexViaScore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSearchIndexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	Sid_ := self.Args("sid").MustInt()
	Langcode_ := self.Args("langcode").String()
	Type_ := self.Args("type").String()
	Score_ := self.Args("score").MustFloat32()

	if helper.IsHas(Word_) {
		_error := model.AddSearchIndex(Word_,Sid_,Langcode_,Type_,Score_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSearchIndex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSearchIndexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	_string, _error := model.PostSearchIndex(&iSearchIndex)
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

func PutSearchIndexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	_string, _error := model.PutSearchIndex(&iSearchIndex)
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

func PutSearchIndexViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	_int64, _error := model.PutSearchIndexViaWord(Word_, &iSearchIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchIndexViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	_int64, _error := model.PutSearchIndexViaSid(Sid_, &iSearchIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchIndexViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	_int64, _error := model.PutSearchIndexViaLangcode(Langcode_, &iSearchIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchIndexViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	_int64, _error := model.PutSearchIndexViaType(Type_, &iSearchIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchIndexViaScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Score_ := self.Args("score").MustFloat32()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	_int64, _error := model.PutSearchIndexViaScore(Score_, &iSearchIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchIndexViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	var iMap = helper.StructToMap(iSearchIndex)
	_error := model.UpdateSearchIndexViaWord(Word_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchIndexViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	var iMap = helper.StructToMap(iSearchIndex)
	_error := model.UpdateSearchIndexViaSid(Sid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchIndexViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	var iMap = helper.StructToMap(iSearchIndex)
	_error := model.UpdateSearchIndexViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchIndexViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	var iMap = helper.StructToMap(iSearchIndex)
	_error := model.UpdateSearchIndexViaType(Type_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchIndexViaScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Score_ := self.Args("score").MustFloat32()
	var iSearchIndex model.SearchIndex
	self.Bind(&iSearchIndex)
	var iMap = helper.StructToMap(iSearchIndex)
	_error := model.UpdateSearchIndexViaScore(Score_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchIndexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	_int64, _error := model.DeleteSearchIndex(Word_)
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

func DeleteSearchIndexViaWordHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Word_ := self.Args("word").String()
	_error := model.DeleteSearchIndexViaWord(Word_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchIndexViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt()
	_error := model.DeleteSearchIndexViaSid(Sid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchIndexViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteSearchIndexViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchIndexViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	_error := model.DeleteSearchIndexViaType(Type_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchIndexViaScoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Score_ := self.Args("score").MustFloat32()
	_error := model.DeleteSearchIndexViaScore(Score_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
