package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetTaxonomyTermDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetTaxonomyTermDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["taxonomy_term_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDataCountViaTidHandler(self *macross.Context) error {
	Tid_ := self.Args("tid").MustInt64()
	_int64 := model.GetTaxonomyTermDataCountViaTid(Tid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_dataCount"] = 0
	}
	m["taxonomy_term_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermDataCountViaVidHandler(self *macross.Context) error {
	Vid_ := self.Args("vid").String()
	_int64 := model.GetTaxonomyTermDataCountViaVid(Vid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_dataCount"] = 0
	}
	m["taxonomy_term_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermDataCountViaUuidHandler(self *macross.Context) error {
	Uuid_ := self.Args("uuid").String()
	_int64 := model.GetTaxonomyTermDataCountViaUuid(Uuid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_dataCount"] = 0
	}
	m["taxonomy_term_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermDataCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetTaxonomyTermDataCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_dataCount"] = 0
	}
	m["taxonomy_term_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermDatasViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTid := self.Args("tid").MustInt64()
	if (offset > 0) && helper.IsHas(iTid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasViaTid(offset, limit, iTid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iVid := self.Args("vid").String()
	if (offset > 0) && helper.IsHas(iVid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasViaVid(offset, limit, iVid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUuid := self.Args("uuid").String()
	if (offset > 0) && helper.IsHas(iUuid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasViaUuid(offset, limit, iUuid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByTidAndVidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByTidAndVidAndUuid(offset, limit, iTid,iVid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByTidAndVidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByTidAndVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByTidAndVidAndLangcode(offset, limit, iTid,iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByTidAndVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByTidAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByTidAndUuidAndLangcode(offset, limit, iTid,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByTidAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByVidAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByVidAndUuidAndLangcode(offset, limit, iVid,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByVidAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByTidAndVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByTidAndVid(offset, limit, iTid,iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByTidAndVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByTidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByTidAndUuid(offset, limit, iTid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByTidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByTidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByTidAndLangcode(offset, limit, iTid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByTidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByVidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByVidAndUuid(offset, limit, iVid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByVidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByVidAndLangcode(offset, limit, iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasByUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUuid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatasByUuidAndLangcode(offset, limit, iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatasByUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTid := self.Args("tid").MustInt64()
	if helper.IsHas(iTid) {
		_TaxonomyTermData := model.HasTaxonomyTermDataViaTid(iTid)
		var m = map[string]interface{}{}
		m["taxonomy_term_data"] = _TaxonomyTermData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermDataViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").String()
	if helper.IsHas(iVid) {
		_TaxonomyTermData := model.HasTaxonomyTermDataViaVid(iVid)
		var m = map[string]interface{}{}
		m["taxonomy_term_data"] = _TaxonomyTermData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermDataViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_TaxonomyTermData := model.HasTaxonomyTermDataViaUuid(iUuid)
		var m = map[string]interface{}{}
		m["taxonomy_term_data"] = _TaxonomyTermData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermDataViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_TaxonomyTermData := model.HasTaxonomyTermDataViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["taxonomy_term_data"] = _TaxonomyTermData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTid := self.Args("tid").MustInt64()
	if helper.IsHas(iTid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDataViaTid(iTid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDataViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").String()
	if helper.IsHas(iVid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDataViaVid(iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDataViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDataViaUuid(iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDataViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_TaxonomyTermData, _error := model.GetTaxonomyTermDataViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	if helper.IsHas(Tid_) {
		var iTaxonomyTermData model.TaxonomyTermData
		self.Bind(&iTaxonomyTermData)
		_TaxonomyTermData, _error := model.SetTaxonomyTermDataViaTid(Tid_, &iTaxonomyTermData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermDataViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").String()
	if helper.IsHas(Vid_) {
		var iTaxonomyTermData model.TaxonomyTermData
		self.Bind(&iTaxonomyTermData)
		_TaxonomyTermData, _error := model.SetTaxonomyTermDataViaVid(Vid_, &iTaxonomyTermData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermDataViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	if helper.IsHas(Uuid_) {
		var iTaxonomyTermData model.TaxonomyTermData
		self.Bind(&iTaxonomyTermData)
		_TaxonomyTermData, _error := model.SetTaxonomyTermDataViaUuid(Uuid_, &iTaxonomyTermData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermDataViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iTaxonomyTermData model.TaxonomyTermData
		self.Bind(&iTaxonomyTermData)
		_TaxonomyTermData, _error := model.SetTaxonomyTermDataViaLangcode(Langcode_, &iTaxonomyTermData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddTaxonomyTermDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	Vid_ := self.Args("vid").String()
	Uuid_ := self.Args("uuid").String()
	Langcode_ := self.Args("langcode").String()

	if helper.IsHas(Tid_) {
		_error := model.AddTaxonomyTermData(Tid_,Vid_,Uuid_,Langcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddTaxonomyTermData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostTaxonomyTermDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	_int64, _error := model.PostTaxonomyTermData(&iTaxonomyTermData)
	if (helper.IsHas(_int64)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["created"] = _int64
		return self.JSON(m, macross.StatusCreated)
	}
	return self.JSON(herr)
}

func PutTaxonomyTermDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	_int64, _error := model.PutTaxonomyTermData(&iTaxonomyTermData)
	if (helper.IsHas(_int64)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["updated"] = _int64
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func PutTaxonomyTermDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	_int64, _error := model.PutTaxonomyTermDataViaTid(Tid_, &iTaxonomyTermData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").String()
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	_int64, _error := model.PutTaxonomyTermDataViaVid(Vid_, &iTaxonomyTermData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermDataViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	_int64, _error := model.PutTaxonomyTermDataViaUuid(Uuid_, &iTaxonomyTermData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	_int64, _error := model.PutTaxonomyTermDataViaLangcode(Langcode_, &iTaxonomyTermData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	var iMap = helper.StructToMap(iTaxonomyTermData)
	_error := model.UpdateTaxonomyTermDataViaTid(Tid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").String()
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	var iMap = helper.StructToMap(iTaxonomyTermData)
	_error := model.UpdateTaxonomyTermDataViaVid(Vid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermDataViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	var iMap = helper.StructToMap(iTaxonomyTermData)
	_error := model.UpdateTaxonomyTermDataViaUuid(Uuid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iTaxonomyTermData model.TaxonomyTermData
	self.Bind(&iTaxonomyTermData)
	var iMap = helper.StructToMap(iTaxonomyTermData)
	_error := model.UpdateTaxonomyTermDataViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	_int64, _error := model.DeleteTaxonomyTermData(Tid_)
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

func DeleteTaxonomyTermDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	_error := model.DeleteTaxonomyTermDataViaTid(Tid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").String()
	_error := model.DeleteTaxonomyTermDataViaVid(Vid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermDataViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	_error := model.DeleteTaxonomyTermDataViaUuid(Uuid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteTaxonomyTermDataViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
