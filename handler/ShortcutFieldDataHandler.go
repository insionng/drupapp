package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetShortcutFieldDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetShortcutFieldDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["shortcut_field_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetShortcutFieldDataCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDataCountViaShortcutSetHandler(self *macross.Context) error {
	ShortcutSet_ := self.Args("shortcut_set").String()
	_int64 := model.GetShortcutFieldDataCountViaShortcutSet(ShortcutSet_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDataCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetShortcutFieldDataCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDataCountViaTitleHandler(self *macross.Context) error {
	Title_ := self.Args("title").String()
	_int64 := model.GetShortcutFieldDataCountViaTitle(Title_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDataCountViaWeightHandler(self *macross.Context) error {
	Weight_ := self.Args("weight").MustInt()
	_int64 := model.GetShortcutFieldDataCountViaWeight(Weight_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDataCountViaLink_uriHandler(self *macross.Context) error {
	Link_uri_ := self.Args("link__uri").String()
	_int64 := model.GetShortcutFieldDataCountViaLink_uri(Link_uri_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDataCountViaLink_titleHandler(self *macross.Context) error {
	Link_title_ := self.Args("link__title").String()
	_int64 := model.GetShortcutFieldDataCountViaLink_title(Link_title_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDataCountViaLink_optionsHandler(self *macross.Context) error {
	Link_options_ := self.Args("link__options").Bytes()
	_int64 := model.GetShortcutFieldDataCountViaLink_options(Link_options_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDataCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetShortcutFieldDataCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_field_dataCount"] = 0
	}
	m["shortcut_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetShortcutFieldDatasViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iShortcutSet := self.Args("shortcut_set").String()
	if (offset > 0) && helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaShortcutSet(offset, limit, iShortcutSet, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTitle := self.Args("title").String()
	if (offset > 0) && helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaTitle(offset, limit, iTitle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iWeight := self.Args("weight").MustInt()
	if (offset > 0) && helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaWeight(offset, limit, iWeight, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLink_uri := self.Args("link__uri").String()
	if (offset > 0) && helper.IsHas(iLink_uri) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaLink_uri(offset, limit, iLink_uri, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLink_title := self.Args("link__title").String()
	if (offset > 0) && helper.IsHas(iLink_title) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaLink_title(offset, limit, iLink_title, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLink_options := self.Args("link__options").Bytes()
	if (offset > 0) && helper.IsHas(iLink_options) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaLink_options(offset, limit, iLink_options, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndShortcutSetAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndShortcutSetAndLangcode(offset, limit, iId,iShortcutSet,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndShortcutSetAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndShortcutSetAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndShortcutSetAndTitle(offset, limit, iId,iShortcutSet,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndShortcutSetAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndShortcutSetAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndShortcutSetAndWeight(offset, limit, iId,iShortcutSet,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndShortcutSetAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndShortcutSetAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndShortcutSetAndLink_uri(offset, limit, iId,iShortcutSet,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndShortcutSetAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndShortcutSetAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndShortcutSetAndLink_title(offset, limit, iId,iShortcutSet,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndShortcutSetAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndShortcutSetAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndShortcutSetAndLink_options(offset, limit, iId,iShortcutSet,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndShortcutSetAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndShortcutSetAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndShortcutSetAndDefaultLangcode(offset, limit, iId,iShortcutSet,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndShortcutSetAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLangcodeAndTitle(offset, limit, iId,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLangcodeAndWeight(offset, limit, iId,iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLangcodeAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLangcodeAndLink_uri(offset, limit, iId,iLangcode,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLangcodeAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLangcodeAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLangcodeAndLink_title(offset, limit, iId,iLangcode,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLangcodeAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLangcodeAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLangcodeAndLink_options(offset, limit, iId,iLangcode,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLangcodeAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLangcodeAndDefaultLangcode(offset, limit, iId,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndTitleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndTitleAndWeight(offset, limit, iId,iTitle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndTitleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndTitleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndTitleAndLink_uri(offset, limit, iId,iTitle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndTitleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndTitleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndTitleAndLink_title(offset, limit, iId,iTitle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndTitleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndTitleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndTitleAndLink_options(offset, limit, iId,iTitle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndTitleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndTitleAndDefaultLangcode(offset, limit, iId,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndWeightAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndWeightAndLink_uri(offset, limit, iId,iWeight,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndWeightAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndWeightAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndWeightAndLink_title(offset, limit, iId,iWeight,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndWeightAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndWeightAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndWeightAndLink_options(offset, limit, iId,iWeight,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndWeightAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndWeightAndDefaultLangcode(offset, limit, iId,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_uriAndLink_title(offset, limit, iId,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_uriAndLink_options(offset, limit, iId,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_uriAndDefaultLangcode(offset, limit, iId,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_titleAndLink_options(offset, limit, iId,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_titleAndDefaultLangcode(offset, limit, iId,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_optionsAndDefaultLangcode(offset, limit, iId,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLangcodeAndTitle(offset, limit, iShortcutSet,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLangcodeAndWeight(offset, limit, iShortcutSet,iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_uri(offset, limit, iShortcutSet,iLangcode,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_title(offset, limit, iShortcutSet,iLangcode,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_options(offset, limit, iShortcutSet,iLangcode,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLangcodeAndDefaultLangcode(offset, limit, iShortcutSet,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndTitleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndTitleAndWeight(offset, limit, iShortcutSet,iTitle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndTitleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndTitleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndTitleAndLink_uri(offset, limit, iShortcutSet,iTitle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndTitleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndTitleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndTitleAndLink_title(offset, limit, iShortcutSet,iTitle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndTitleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndTitleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndTitleAndLink_options(offset, limit, iShortcutSet,iTitle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndTitleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndTitleAndDefaultLangcode(offset, limit, iShortcutSet,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndWeightAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iWeight := self.Args("weight").MustInt()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndWeightAndLink_uri(offset, limit, iShortcutSet,iWeight,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndWeightAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndWeightAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iWeight := self.Args("weight").MustInt()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndWeightAndLink_title(offset, limit, iShortcutSet,iWeight,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndWeightAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndWeightAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iWeight := self.Args("weight").MustInt()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndWeightAndLink_options(offset, limit, iShortcutSet,iWeight,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndWeightAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndWeightAndDefaultLangcode(offset, limit, iShortcutSet,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_title(offset, limit, iShortcutSet,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_options(offset, limit, iShortcutSet,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_uriAndDefaultLangcode(offset, limit, iShortcutSet,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_titleAndLink_options(offset, limit, iShortcutSet,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_titleAndDefaultLangcode(offset, limit, iShortcutSet,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_optionsAndDefaultLangcode(offset, limit, iShortcutSet,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndTitleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndTitleAndWeight(offset, limit, iLangcode,iTitle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndTitleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndTitleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndTitleAndLink_uri(offset, limit, iLangcode,iTitle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndTitleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndTitleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndTitleAndLink_title(offset, limit, iLangcode,iTitle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndTitleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndTitleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndTitleAndLink_options(offset, limit, iLangcode,iTitle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndTitleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndTitleAndDefaultLangcode(offset, limit, iLangcode,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndWeightAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndWeightAndLink_uri(offset, limit, iLangcode,iWeight,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndWeightAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndWeightAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndWeightAndLink_title(offset, limit, iLangcode,iWeight,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndWeightAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndWeightAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndWeightAndLink_options(offset, limit, iLangcode,iWeight,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndWeightAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndWeightAndDefaultLangcode(offset, limit, iLangcode,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_title(offset, limit, iLangcode,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_options(offset, limit, iLangcode,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_uriAndDefaultLangcode(offset, limit, iLangcode,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_titleAndLink_options(offset, limit, iLangcode,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_titleAndDefaultLangcode(offset, limit, iLangcode,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_optionsAndDefaultLangcode(offset, limit, iLangcode,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndWeightAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndWeightAndLink_uri(offset, limit, iTitle,iWeight,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndWeightAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndWeightAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndWeightAndLink_title(offset, limit, iTitle,iWeight,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndWeightAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndWeightAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndWeightAndLink_options(offset, limit, iTitle,iWeight,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndWeightAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndWeightAndDefaultLangcode(offset, limit, iTitle,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_uriAndLink_title(offset, limit, iTitle,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_uriAndLink_options(offset, limit, iTitle,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_uriAndDefaultLangcode(offset, limit, iTitle,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_titleAndLink_options(offset, limit, iTitle,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_titleAndDefaultLangcode(offset, limit, iTitle,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_optionsAndDefaultLangcode(offset, limit, iTitle,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_uriAndLink_title(offset, limit, iWeight,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_uriAndLink_options(offset, limit, iWeight,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_uriAndDefaultLangcode(offset, limit, iWeight,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_titleAndLink_options(offset, limit, iWeight,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_titleAndDefaultLangcode(offset, limit, iWeight,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_optionsAndDefaultLangcode(offset, limit, iWeight,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_uriAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLink_uri) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_uriAndLink_titleAndLink_options(offset, limit, iLink_uri,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_uriAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_uriAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_uriAndLink_titleAndDefaultLangcode(offset, limit, iLink_uri,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_uriAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_uriAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_uriAndLink_optionsAndDefaultLangcode(offset, limit, iLink_uri,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_uriAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_titleAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_titleAndLink_optionsAndDefaultLangcode(offset, limit, iLink_title,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_titleAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndShortcutSet(offset, limit, iId,iShortcutSet)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLangcode(offset, limit, iId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndTitle(offset, limit, iId,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndWeight(offset, limit, iId,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_uri(offset, limit, iId,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_title(offset, limit, iId,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndLink_options(offset, limit, iId,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByIdAndDefaultLangcode(offset, limit, iId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLangcode(offset, limit, iShortcutSet,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndTitle(offset, limit, iShortcutSet,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndWeight(offset, limit, iShortcutSet,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_uri(offset, limit, iShortcutSet,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_title(offset, limit, iShortcutSet,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndLink_options(offset, limit, iShortcutSet,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByShortcutSetAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByShortcutSetAndDefaultLangcode(offset, limit, iShortcutSet,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByShortcutSetAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndTitle(offset, limit, iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndWeight(offset, limit, iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_uri(offset, limit, iLangcode,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_title(offset, limit, iLangcode,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndLink_options(offset, limit, iLangcode,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndWeight(offset, limit, iTitle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_uri(offset, limit, iTitle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_title(offset, limit, iTitle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndLink_options(offset, limit, iTitle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByTitleAndDefaultLangcode(offset, limit, iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_uri(offset, limit, iWeight,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_title(offset, limit, iWeight,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndLink_options(offset, limit, iWeight,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByWeightAndDefaultLangcode(offset, limit, iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLink_uri) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_uriAndLink_title(offset, limit, iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLink_uri) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_uriAndLink_options(offset, limit, iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_uriAndDefaultLangcode(offset, limit, iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLink_title) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_titleAndLink_options(offset, limit, iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_titleAndDefaultLangcode(offset, limit, iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasByLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatasByLink_optionsAndDefaultLangcode(offset, limit, iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatasByLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaId(iId)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iShortcutSet := self.Args("shortcut_set").String()
	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaShortcutSet(iShortcutSet)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaTitle(iTitle)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWeight := self.Args("weight").MustInt()
	if helper.IsHas(iWeight) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaWeight(iWeight)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_uri := self.Args("link__uri").String()
	if helper.IsHas(iLink_uri) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaLink_uri(iLink_uri)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_title := self.Args("link__title").String()
	if helper.IsHas(iLink_title) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaLink_title(iLink_title)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_options := self.Args("link__options").Bytes()
	if helper.IsHas(iLink_options) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaLink_options(iLink_options)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_ShortcutFieldData := model.HasShortcutFieldDataViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["shortcut_field_data"] = _ShortcutFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iShortcutSet := self.Args("shortcut_set").String()
	if helper.IsHas(iShortcutSet) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaShortcutSet(iShortcutSet)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaTitle(iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWeight := self.Args("weight").MustInt()
	if helper.IsHas(iWeight) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaWeight(iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_uri := self.Args("link__uri").String()
	if helper.IsHas(iLink_uri) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaLink_uri(iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_title := self.Args("link__title").String()
	if helper.IsHas(iLink_title) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaLink_title(iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_options := self.Args("link__options").Bytes()
	if helper.IsHas(iLink_options) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaLink_options(iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_ShortcutFieldData, _error := model.GetShortcutFieldDataViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the GetShortcutFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaId(Id_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ShortcutSet_ := self.Args("shortcut_set").String()
	if helper.IsHas(ShortcutSet_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaShortcutSet(ShortcutSet_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaLangcode(Langcode_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	if helper.IsHas(Title_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaTitle(Title_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	if helper.IsHas(Weight_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaWeight(Weight_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_uri_ := self.Args("link__uri").String()
	if helper.IsHas(Link_uri_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaLink_uri(Link_uri_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_title_ := self.Args("link__title").String()
	if helper.IsHas(Link_title_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaLink_title(Link_title_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_options_ := self.Args("link__options").Bytes()
	if helper.IsHas(Link_options_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaLink_options(Link_options_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iShortcutFieldData model.ShortcutFieldData
		self.Bind(&iShortcutFieldData)
		_ShortcutFieldData, _error := model.SetShortcutFieldDataViaDefaultLangcode(DefaultLangcode_, &iShortcutFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutFieldData)
	}
	herr.Message = "Can't get to the SetShortcutFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddShortcutFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	ShortcutSet_ := self.Args("shortcut_set").String()
	Langcode_ := self.Args("langcode").String()
	Title_ := self.Args("title").String()
	Weight_ := self.Args("weight").MustInt()
	Link_uri_ := self.Args("link__uri").String()
	Link_title_ := self.Args("link__title").String()
	Link_options_ := self.Args("link__options").Bytes()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Id_) {
		_error := model.AddShortcutFieldData(Id_,ShortcutSet_,Langcode_,Title_,Weight_,Link_uri_,Link_title_,Link_options_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddShortcutFieldData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostShortcutFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PostShortcutFieldData(&iShortcutFieldData)
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

func PutShortcutFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldData(&iShortcutFieldData)
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

func PutShortcutFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaId(Id_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutFieldDataViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ShortcutSet_ := self.Args("shortcut_set").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaShortcutSet(ShortcutSet_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaLangcode(Langcode_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaTitle(Title_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaWeight(Weight_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutFieldDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_uri_ := self.Args("link__uri").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaLink_uri(Link_uri_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutFieldDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_title_ := self.Args("link__title").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaLink_title(Link_title_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutFieldDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_options_ := self.Args("link__options").Bytes()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaLink_options(Link_options_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	_int64, _error := model.PutShortcutFieldDataViaDefaultLangcode(DefaultLangcode_, &iShortcutFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ShortcutSet_ := self.Args("shortcut_set").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaShortcutSet(ShortcutSet_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaTitle(Title_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaWeight(Weight_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_uri_ := self.Args("link__uri").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaLink_uri(Link_uri_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_title_ := self.Args("link__title").String()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaLink_title(Link_title_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_options_ := self.Args("link__options").Bytes()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaLink_options(Link_options_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iShortcutFieldData model.ShortcutFieldData
	self.Bind(&iShortcutFieldData)
	var iMap = helper.StructToMap(iShortcutFieldData)
	_error := model.UpdateShortcutFieldDataViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteShortcutFieldData(Id_)
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

func DeleteShortcutFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteShortcutFieldDataViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ShortcutSet_ := self.Args("shortcut_set").String()
	_error := model.DeleteShortcutFieldDataViaShortcutSet(ShortcutSet_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteShortcutFieldDataViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	_error := model.DeleteShortcutFieldDataViaTitle(Title_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	_error := model.DeleteShortcutFieldDataViaWeight(Weight_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_uri_ := self.Args("link__uri").String()
	_error := model.DeleteShortcutFieldDataViaLink_uri(Link_uri_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_title_ := self.Args("link__title").String()
	_error := model.DeleteShortcutFieldDataViaLink_title(Link_title_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_options_ := self.Args("link__options").Bytes()
	_error := model.DeleteShortcutFieldDataViaLink_options(Link_options_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteShortcutFieldDataViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
