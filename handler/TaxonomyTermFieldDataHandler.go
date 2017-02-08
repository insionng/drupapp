package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetTaxonomyTermFieldDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetTaxonomyTermFieldDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["taxonomy_term_field_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataCountViaTidHandler(self *macross.Context) error {
	Tid_ := self.Args("tid").MustInt64()
	_int64 := model.GetTaxonomyTermFieldDataCountViaTid(Tid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDataCountViaVidHandler(self *macross.Context) error {
	Vid_ := self.Args("vid").String()
	_int64 := model.GetTaxonomyTermFieldDataCountViaVid(Vid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDataCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetTaxonomyTermFieldDataCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDataCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetTaxonomyTermFieldDataCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDataCountViaDescription_valueHandler(self *macross.Context) error {
	Description_value_ := self.Args("description__value").String()
	_int64 := model.GetTaxonomyTermFieldDataCountViaDescription_value(Description_value_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDataCountViaDescription_formatHandler(self *macross.Context) error {
	Description_format_ := self.Args("description__format").String()
	_int64 := model.GetTaxonomyTermFieldDataCountViaDescription_format(Description_format_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDataCountViaWeightHandler(self *macross.Context) error {
	Weight_ := self.Args("weight").MustInt()
	_int64 := model.GetTaxonomyTermFieldDataCountViaWeight(Weight_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDataCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetTaxonomyTermFieldDataCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDataCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetTaxonomyTermFieldDataCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_field_dataCount"] = 0
	}
	m["taxonomy_term_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermFieldDatasViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTid := self.Args("tid").MustInt64()
	if (offset > 0) && helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaTid(offset, limit, iTid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iVid := self.Args("vid").String()
	if (offset > 0) && helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaVid(offset, limit, iVid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasViaDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDescription_value := self.Args("description__value").String()
	if (offset > 0) && helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaDescription_value(offset, limit, iDescription_value, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasViaDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDescription_format := self.Args("description__format").String()
	if (offset > 0) && helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaDescription_format(offset, limit, iDescription_format, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iWeight := self.Args("weight").MustInt()
	if (offset > 0) && helper.IsHas(iWeight) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaWeight(offset, limit, iWeight, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndVidAndLangcode(offset, limit, iTid,iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndVidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iName := self.Args("name").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndVidAndName(offset, limit, iTid,iVid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndVidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndVidAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndVidAndDescription_value(offset, limit, iTid,iVid,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndVidAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndVidAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndVidAndDescription_format(offset, limit, iTid,iVid,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndVidAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndVidAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndVidAndWeight(offset, limit, iTid,iVid,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndVidAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndVidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndVidAndChanged(offset, limit, iTid,iVid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndVidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndVidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndVidAndDefaultLangcode(offset, limit, iTid,iVid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndVidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndLangcodeAndName(offset, limit, iTid,iLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_value(offset, limit, iTid,iLangcode,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_format(offset, limit, iTid,iLangcode,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndLangcodeAndWeight(offset, limit, iTid,iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndLangcodeAndChanged(offset, limit, iTid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndLangcodeAndDefaultLangcode(offset, limit, iTid,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndNameAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iName := self.Args("name").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndNameAndDescription_value(offset, limit, iTid,iName,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndNameAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndNameAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iName := self.Args("name").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndNameAndDescription_format(offset, limit, iTid,iName,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndNameAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iName := self.Args("name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndNameAndWeight(offset, limit, iTid,iName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndNameAndChanged(offset, limit, iTid,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndNameAndDefaultLangcode(offset, limit, iTid,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_value := self.Args("description_value").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDescription_format(offset, limit, iTid,iDescription_value,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_valueAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_value := self.Args("description_value").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_valueAndWeight(offset, limit, iTid,iDescription_value,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_valueAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_valueAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_value := self.Args("description_value").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_valueAndChanged(offset, limit, iTid,iDescription_value,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_valueAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_value := self.Args("description_value").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDefaultLangcode(offset, limit, iTid,iDescription_value,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_formatAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_format := self.Args("description_format").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_formatAndWeight(offset, limit, iTid,iDescription_format,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_formatAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_formatAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_format := self.Args("description_format").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_formatAndChanged(offset, limit, iTid,iDescription_format,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_formatAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_formatAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_format := self.Args("description_format").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_formatAndDefaultLangcode(offset, limit, iTid,iDescription_format,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_formatAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndWeightAndChanged(offset, limit, iTid,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndWeightAndDefaultLangcode(offset, limit, iTid,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndChangedAndDefaultLangcode(offset, limit, iTid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndLangcodeAndName(offset, limit, iVid,iLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_value(offset, limit, iVid,iLangcode,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_format(offset, limit, iVid,iLangcode,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndLangcodeAndWeight(offset, limit, iVid,iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndLangcodeAndChanged(offset, limit, iVid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndLangcodeAndDefaultLangcode(offset, limit, iVid,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndNameAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iName := self.Args("name").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndNameAndDescription_value(offset, limit, iVid,iName,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndNameAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndNameAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iName := self.Args("name").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndNameAndDescription_format(offset, limit, iVid,iName,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndNameAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iName := self.Args("name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndNameAndWeight(offset, limit, iVid,iName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndNameAndChanged(offset, limit, iVid,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndNameAndDefaultLangcode(offset, limit, iVid,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_value := self.Args("description_value").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDescription_format(offset, limit, iVid,iDescription_value,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_valueAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_value := self.Args("description_value").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_valueAndWeight(offset, limit, iVid,iDescription_value,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_valueAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_valueAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_value := self.Args("description_value").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_valueAndChanged(offset, limit, iVid,iDescription_value,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_valueAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_value := self.Args("description_value").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDefaultLangcode(offset, limit, iVid,iDescription_value,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_formatAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_format := self.Args("description_format").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_formatAndWeight(offset, limit, iVid,iDescription_format,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_formatAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_formatAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_format := self.Args("description_format").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_formatAndChanged(offset, limit, iVid,iDescription_format,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_formatAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_formatAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_format := self.Args("description_format").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_formatAndDefaultLangcode(offset, limit, iVid,iDescription_format,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_formatAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndWeightAndChanged(offset, limit, iVid,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndWeightAndDefaultLangcode(offset, limit, iVid,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndChangedAndDefaultLangcode(offset, limit, iVid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_value(offset, limit, iLangcode,iName,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_format(offset, limit, iLangcode,iName,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndNameAndWeight(offset, limit, iLangcode,iName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndNameAndChanged(offset, limit, iLangcode,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndNameAndDefaultLangcode(offset, limit, iLangcode,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_value := self.Args("description_value").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDescription_format(offset, limit, iLangcode,iDescription_value,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_value := self.Args("description_value").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndWeight(offset, limit, iLangcode,iDescription_value,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_value := self.Args("description_value").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndChanged(offset, limit, iLangcode,iDescription_value,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_value := self.Args("description_value").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDefaultLangcode(offset, limit, iLangcode,iDescription_value,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_format := self.Args("description_format").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndWeight(offset, limit, iLangcode,iDescription_format,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_format := self.Args("description_format").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndChanged(offset, limit, iLangcode,iDescription_format,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_format := self.Args("description_format").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndDefaultLangcode(offset, limit, iLangcode,iDescription_format,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndWeightAndChanged(offset, limit, iLangcode,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndWeightAndDefaultLangcode(offset, limit, iLangcode,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset, limit, iLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_value := self.Args("description_value").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDescription_format(offset, limit, iName,iDescription_value,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_valueAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_value := self.Args("description_value").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_valueAndWeight(offset, limit, iName,iDescription_value,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_valueAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_valueAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_value := self.Args("description_value").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_valueAndChanged(offset, limit, iName,iDescription_value,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_valueAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_value := self.Args("description_value").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDefaultLangcode(offset, limit, iName,iDescription_value,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_formatAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_format := self.Args("description_format").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_formatAndWeight(offset, limit, iName,iDescription_format,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_formatAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_formatAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_format := self.Args("description_format").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_formatAndChanged(offset, limit, iName,iDescription_format,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_formatAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_formatAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_format := self.Args("description_format").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_formatAndDefaultLangcode(offset, limit, iName,iDescription_format,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_formatAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndWeightAndChanged(offset, limit, iName,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndWeightAndDefaultLangcode(offset, limit, iName,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndChangedAndDefaultLangcode(offset, limit, iName,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iDescription_format := self.Args("description_format").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndWeight(offset, limit, iDescription_value,iDescription_format,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iDescription_format := self.Args("description_format").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndChanged(offset, limit, iDescription_value,iDescription_format,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iDescription_format := self.Args("description_format").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndDefaultLangcode(offset, limit, iDescription_value,iDescription_format,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndChanged(offset, limit, iDescription_value,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndDefaultLangcode(offset, limit, iDescription_value,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndChangedAndDefaultLangcode(offset, limit, iDescription_value,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_format := self.Args("description_format").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndChanged(offset, limit, iDescription_format,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_format := self.Args("description_format").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndDefaultLangcode(offset, limit, iDescription_format,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_formatAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_format := self.Args("description_format").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_formatAndChangedAndDefaultLangcode(offset, limit, iDescription_format,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_formatAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByWeightAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByWeightAndChangedAndDefaultLangcode(offset, limit, iWeight,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByWeightAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iVid := self.Args("vid").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndVid(offset, limit, iTid,iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndLangcode(offset, limit, iTid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iName := self.Args("name").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndName(offset, limit, iTid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_value(offset, limit, iTid,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDescription_format(offset, limit, iTid,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndWeight(offset, limit, iTid,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndChanged(offset, limit, iTid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByTidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByTidAndDefaultLangcode(offset, limit, iTid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByTidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndLangcode(offset, limit, iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iName := self.Args("name").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndName(offset, limit, iVid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_value(offset, limit, iVid,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDescription_format(offset, limit, iVid,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndWeight(offset, limit, iVid,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndChanged(offset, limit, iVid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByVidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByVidAndDefaultLangcode(offset, limit, iVid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByVidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndName(offset, limit, iLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_value(offset, limit, iLangcode,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDescription_format(offset, limit, iLangcode,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndWeight(offset, limit, iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_value := self.Args("description_value").String()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_value(offset, limit, iName,iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDescription_format(offset, limit, iName,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndWeight(offset, limit, iName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndChanged(offset, limit, iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByNameAndDefaultLangcode(offset, limit, iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iDescription_format := self.Args("description_format").String()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndDescription_format(offset, limit, iDescription_value,iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndWeight(offset, limit, iDescription_value,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndChanged(offset, limit, iDescription_value,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_valueAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_value := self.Args("description_value").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_valueAndDefaultLangcode(offset, limit, iDescription_value,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_valueAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_formatAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_format := self.Args("description_format").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_formatAndWeight(offset, limit, iDescription_format,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_formatAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_formatAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_format := self.Args("description_format").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_formatAndChanged(offset, limit, iDescription_format,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_formatAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByDescription_formatAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription_format := self.Args("description_format").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByDescription_formatAndDefaultLangcode(offset, limit, iDescription_format,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByDescription_formatAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iWeight) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByWeightAndChanged(offset, limit, iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByWeightAndDefaultLangcode(offset, limit, iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasByChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatasByChangedAndDefaultLangcode(offset, limit, iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatasByChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTid := self.Args("tid").MustInt64()
	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaTid(iTid)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").String()
	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaVid(iVid)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaName(iName)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription_value := self.Args("description__value").String()
	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaDescription_value(iDescription_value)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription_format := self.Args("description__format").String()
	if helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaDescription_format(iDescription_format)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWeight := self.Args("weight").MustInt()
	if helper.IsHas(iWeight) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaWeight(iWeight)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_TaxonomyTermFieldData := model.HasTaxonomyTermFieldDataViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["taxonomy_term_field_data"] = _TaxonomyTermFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTid := self.Args("tid").MustInt64()
	if helper.IsHas(iTid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaTid(iTid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").String()
	if helper.IsHas(iVid) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaVid(iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription_value := self.Args("description__value").String()
	if helper.IsHas(iDescription_value) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaDescription_value(iDescription_value)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription_format := self.Args("description__format").String()
	if helper.IsHas(iDescription_format) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaDescription_format(iDescription_format)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWeight := self.Args("weight").MustInt()
	if helper.IsHas(iWeight) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaWeight(iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_TaxonomyTermFieldData, _error := model.GetTaxonomyTermFieldDataViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the GetTaxonomyTermFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	if helper.IsHas(Tid_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaTid(Tid_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").String()
	if helper.IsHas(Vid_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaVid(Vid_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaLangcode(Langcode_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaName(Name_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_value_ := self.Args("description__value").String()
	if helper.IsHas(Description_value_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaDescription_value(Description_value_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaDescription_value's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_format_ := self.Args("description__format").String()
	if helper.IsHas(Description_format_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaDescription_format(Description_format_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaDescription_format's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	if helper.IsHas(Weight_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaWeight(Weight_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaChanged(Changed_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iTaxonomyTermFieldData model.TaxonomyTermFieldData
		self.Bind(&iTaxonomyTermFieldData)
		_TaxonomyTermFieldData, _error := model.SetTaxonomyTermFieldDataViaDefaultLangcode(DefaultLangcode_, &iTaxonomyTermFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermFieldData)
	}
	herr.Message = "Can't get to the SetTaxonomyTermFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddTaxonomyTermFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	Vid_ := self.Args("vid").String()
	Langcode_ := self.Args("langcode").String()
	Name_ := self.Args("name").String()
	Description_value_ := self.Args("description__value").String()
	Description_format_ := self.Args("description__format").String()
	Weight_ := self.Args("weight").MustInt()
	Changed_ := self.Args("changed").MustInt()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Tid_) {
		_error := model.AddTaxonomyTermFieldData(Tid_,Vid_,Langcode_,Name_,Description_value_,Description_format_,Weight_,Changed_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddTaxonomyTermFieldData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostTaxonomyTermFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PostTaxonomyTermFieldData(&iTaxonomyTermFieldData)
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

func PutTaxonomyTermFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldData(&iTaxonomyTermFieldData)
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

func PutTaxonomyTermFieldDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaTid(Tid_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaVid(Vid_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaLangcode(Langcode_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaName(Name_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermFieldDataViaDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_value_ := self.Args("description__value").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaDescription_value(Description_value_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermFieldDataViaDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_format_ := self.Args("description__format").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaDescription_format(Description_format_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaWeight(Weight_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaChanged(Changed_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	_int64, _error := model.PutTaxonomyTermFieldDataViaDefaultLangcode(DefaultLangcode_, &iTaxonomyTermFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaTid(Tid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaVid(Vid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_value_ := self.Args("description__value").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaDescription_value(Description_value_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_format_ := self.Args("description__format").String()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaDescription_format(Description_format_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaWeight(Weight_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iTaxonomyTermFieldData model.TaxonomyTermFieldData
	self.Bind(&iTaxonomyTermFieldData)
	var iMap = helper.StructToMap(iTaxonomyTermFieldData)
	_error := model.UpdateTaxonomyTermFieldDataViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	_int64, _error := model.DeleteTaxonomyTermFieldData(Tid_)
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

func DeleteTaxonomyTermFieldDataViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	_error := model.DeleteTaxonomyTermFieldDataViaTid(Tid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").String()
	_error := model.DeleteTaxonomyTermFieldDataViaVid(Vid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteTaxonomyTermFieldDataViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteTaxonomyTermFieldDataViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataViaDescription_valueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_value_ := self.Args("description__value").String()
	_error := model.DeleteTaxonomyTermFieldDataViaDescription_value(Description_value_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataViaDescription_formatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_format_ := self.Args("description__format").String()
	_error := model.DeleteTaxonomyTermFieldDataViaDescription_format(Description_format_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	_error := model.DeleteTaxonomyTermFieldDataViaWeight(Weight_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteTaxonomyTermFieldDataViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteTaxonomyTermFieldDataViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
