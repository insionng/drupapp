package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetShortcutsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetShortcutsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["shortcutsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetShortcutsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetShortcutCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcutCount"] = 0
	}
	m["shortcutCount"] = _int64
	return self.JSON(m)
}

func GetShortcutCountViaShortcutSetHandler(self *macross.Context) error {
	ShortcutSet_ := self.Args("shortcut_set").String()
	_int64 := model.GetShortcutCountViaShortcutSet(ShortcutSet_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcutCount"] = 0
	}
	m["shortcutCount"] = _int64
	return self.JSON(m)
}

func GetShortcutCountViaUuidHandler(self *macross.Context) error {
	Uuid_ := self.Args("uuid").String()
	_int64 := model.GetShortcutCountViaUuid(Uuid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcutCount"] = 0
	}
	m["shortcutCount"] = _int64
	return self.JSON(m)
}

func GetShortcutCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetShortcutCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcutCount"] = 0
	}
	m["shortcutCount"] = _int64
	return self.JSON(m)
}

func GetShortcutsViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_Shortcut, _error := model.GetShortcutsViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iShortcutSet := self.Args("shortcut_set").String()
	if (offset > 0) && helper.IsHas(iShortcutSet) {
		_Shortcut, _error := model.GetShortcutsViaShortcutSet(offset, limit, iShortcutSet, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsViaShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUuid := self.Args("uuid").String()
	if (offset > 0) && helper.IsHas(iUuid) {
		_Shortcut, _error := model.GetShortcutsViaUuid(offset, limit, iUuid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_Shortcut, _error := model.GetShortcutsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByIdAndShortcutSetAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iId) {
		_Shortcut, _error := model.GetShortcutsByIdAndShortcutSetAndUuid(offset, limit, iId,iShortcutSet,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByIdAndShortcutSetAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByIdAndShortcutSetAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_Shortcut, _error := model.GetShortcutsByIdAndShortcutSetAndLangcode(offset, limit, iId,iShortcutSet,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByIdAndShortcutSetAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByIdAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_Shortcut, _error := model.GetShortcutsByIdAndUuidAndLangcode(offset, limit, iId,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByIdAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByShortcutSetAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iShortcutSet) {
		_Shortcut, _error := model.GetShortcutsByShortcutSetAndUuidAndLangcode(offset, limit, iShortcutSet,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByShortcutSetAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByIdAndShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iShortcutSet := self.Args("shortcut_set").String()

	if helper.IsHas(iId) {
		_Shortcut, _error := model.GetShortcutsByIdAndShortcutSet(offset, limit, iId,iShortcutSet)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByIdAndShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByIdAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iId) {
		_Shortcut, _error := model.GetShortcutsByIdAndUuid(offset, limit, iId,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByIdAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_Shortcut, _error := model.GetShortcutsByIdAndLangcode(offset, limit, iId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByShortcutSetAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iShortcutSet) {
		_Shortcut, _error := model.GetShortcutsByShortcutSetAndUuid(offset, limit, iShortcutSet,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByShortcutSetAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByShortcutSetAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iShortcutSet := self.Args("shortcut_set").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iShortcutSet) {
		_Shortcut, _error := model.GetShortcutsByShortcutSetAndLangcode(offset, limit, iShortcutSet,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByShortcutSetAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsByUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUuid) {
		_Shortcut, _error := model.GetShortcutsByUuidAndLangcode(offset, limit, iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutsByUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Shortcut, _error := model.GetShortcuts(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcuts' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_Shortcut := model.HasShortcutViaId(iId)
		var m = map[string]interface{}{}
		m["shortcut"] = _Shortcut
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iShortcutSet := self.Args("shortcut_set").String()
	if helper.IsHas(iShortcutSet) {
		_Shortcut := model.HasShortcutViaShortcutSet(iShortcutSet)
		var m = map[string]interface{}{}
		m["shortcut"] = _Shortcut
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutViaShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_Shortcut := model.HasShortcutViaUuid(iUuid)
		var m = map[string]interface{}{}
		m["shortcut"] = _Shortcut
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Shortcut := model.HasShortcutViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["shortcut"] = _Shortcut
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_Shortcut, _error := model.GetShortcutViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iShortcutSet := self.Args("shortcut_set").String()
	if helper.IsHas(iShortcutSet) {
		_Shortcut, _error := model.GetShortcutViaShortcutSet(iShortcutSet)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutViaShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_Shortcut, _error := model.GetShortcutViaUuid(iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Shortcut, _error := model.GetShortcutViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the GetShortcutViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iShortcut model.Shortcut
		self.Bind(&iShortcut)
		_Shortcut, _error := model.SetShortcutViaId(Id_, &iShortcut)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the SetShortcutViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ShortcutSet_ := self.Args("shortcut_set").String()
	if helper.IsHas(ShortcutSet_) {
		var iShortcut model.Shortcut
		self.Bind(&iShortcut)
		_Shortcut, _error := model.SetShortcutViaShortcutSet(ShortcutSet_, &iShortcut)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the SetShortcutViaShortcutSet's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	if helper.IsHas(Uuid_) {
		var iShortcut model.Shortcut
		self.Bind(&iShortcut)
		_Shortcut, _error := model.SetShortcutViaUuid(Uuid_, &iShortcut)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the SetShortcutViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iShortcut model.Shortcut
		self.Bind(&iShortcut)
		_Shortcut, _error := model.SetShortcutViaLangcode(Langcode_, &iShortcut)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Shortcut)
	}
	herr.Message = "Can't get to the SetShortcutViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddShortcutHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	ShortcutSet_ := self.Args("shortcut_set").String()
	Uuid_ := self.Args("uuid").String()
	Langcode_ := self.Args("langcode").String()

	if helper.IsHas(Id_) {
		_error := model.AddShortcut(Id_,ShortcutSet_,Uuid_,Langcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddShortcut's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostShortcutHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	_int64, _error := model.PostShortcut(&iShortcut)
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

func PutShortcutHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	_int64, _error := model.PutShortcut(&iShortcut)
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

func PutShortcutViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	_int64, _error := model.PutShortcutViaId(Id_, &iShortcut)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ShortcutSet_ := self.Args("shortcut_set").String()
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	_int64, _error := model.PutShortcutViaShortcutSet(ShortcutSet_, &iShortcut)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	_int64, _error := model.PutShortcutViaUuid(Uuid_, &iShortcut)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	_int64, _error := model.PutShortcutViaLangcode(Langcode_, &iShortcut)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	var iMap = helper.StructToMap(iShortcut)
	_error := model.UpdateShortcutViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ShortcutSet_ := self.Args("shortcut_set").String()
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	var iMap = helper.StructToMap(iShortcut)
	_error := model.UpdateShortcutViaShortcutSet(ShortcutSet_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	var iMap = helper.StructToMap(iShortcut)
	_error := model.UpdateShortcutViaUuid(Uuid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iShortcut model.Shortcut
	self.Bind(&iShortcut)
	var iMap = helper.StructToMap(iShortcut)
	_error := model.UpdateShortcutViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteShortcut(Id_)
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

func DeleteShortcutViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteShortcutViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutViaShortcutSetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ShortcutSet_ := self.Args("shortcut_set").String()
	_error := model.DeleteShortcutViaShortcutSet(ShortcutSet_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	_error := model.DeleteShortcutViaUuid(Uuid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteShortcutViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
