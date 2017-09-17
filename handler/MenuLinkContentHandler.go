package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetMenuLinkContentsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetMenuLinkContentsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["menu_link_contentsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetMenuLinkContentCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_contentCount"] = 0
	}
	m["menu_link_contentCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetMenuLinkContentCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_contentCount"] = 0
	}
	m["menu_link_contentCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentCountViaUuidHandler(self *macross.Context) error {
	Uuid_ := self.Args("uuid").String()
	_int64 := model.GetMenuLinkContentCountViaUuid(Uuid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_contentCount"] = 0
	}
	m["menu_link_contentCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetMenuLinkContentCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_contentCount"] = 0
	}
	m["menu_link_contentCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentsViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUuid := self.Args("uuid").String()
	if (offset > 0) && helper.IsHas(iUuid) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsViaUuid(offset, limit, iUuid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByIdAndBundleAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iId) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByIdAndBundleAndUuid(offset, limit, iId,iBundle,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByIdAndBundleAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByIdAndBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByIdAndBundleAndLangcode(offset, limit, iId,iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByIdAndBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByIdAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByIdAndUuidAndLangcode(offset, limit, iId,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByIdAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByBundleAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByBundleAndUuidAndLangcode(offset, limit, iBundle,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByBundleAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByIdAndBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()

	if helper.IsHas(iId) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByIdAndBundle(offset, limit, iId,iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByIdAndBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByIdAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iId) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByIdAndUuid(offset, limit, iId,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByIdAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByIdAndLangcode(offset, limit, iId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByBundleAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByBundleAndUuid(offset, limit, iBundle,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByBundleAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsByUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUuid) {
		_MenuLinkContent, _error := model.GetMenuLinkContentsByUuidAndLangcode(offset, limit, iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentsByUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_MenuLinkContent, _error := model.GetMenuLinkContents(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContents' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_MenuLinkContent := model.HasMenuLinkContentViaId(iId)
		var m = map[string]interface{}{}
		m["menu_link_content"] = _MenuLinkContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_MenuLinkContent := model.HasMenuLinkContentViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["menu_link_content"] = _MenuLinkContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_MenuLinkContent := model.HasMenuLinkContentViaUuid(iUuid)
		var m = map[string]interface{}{}
		m["menu_link_content"] = _MenuLinkContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_MenuLinkContent := model.HasMenuLinkContentViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["menu_link_content"] = _MenuLinkContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_MenuLinkContent, _error := model.GetMenuLinkContentViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_MenuLinkContent, _error := model.GetMenuLinkContentViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_MenuLinkContent, _error := model.GetMenuLinkContentViaUuid(iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_MenuLinkContent, _error := model.GetMenuLinkContentViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the GetMenuLinkContentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iMenuLinkContent model.MenuLinkContent
		self.Bind(&iMenuLinkContent)
		_MenuLinkContent, _error := model.SetMenuLinkContentViaId(Id_, &iMenuLinkContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the SetMenuLinkContentViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iMenuLinkContent model.MenuLinkContent
		self.Bind(&iMenuLinkContent)
		_MenuLinkContent, _error := model.SetMenuLinkContentViaBundle(Bundle_, &iMenuLinkContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the SetMenuLinkContentViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	if helper.IsHas(Uuid_) {
		var iMenuLinkContent model.MenuLinkContent
		self.Bind(&iMenuLinkContent)
		_MenuLinkContent, _error := model.SetMenuLinkContentViaUuid(Uuid_, &iMenuLinkContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the SetMenuLinkContentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iMenuLinkContent model.MenuLinkContent
		self.Bind(&iMenuLinkContent)
		_MenuLinkContent, _error := model.SetMenuLinkContentViaLangcode(Langcode_, &iMenuLinkContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContent)
	}
	herr.Message = "Can't get to the SetMenuLinkContentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddMenuLinkContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	Bundle_ := self.Args("bundle").String()
	Uuid_ := self.Args("uuid").String()
	Langcode_ := self.Args("langcode").String()

	if helper.IsHas(Id_) {
		_error := model.AddMenuLinkContent(Id_,Bundle_,Uuid_,Langcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddMenuLinkContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostMenuLinkContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	_int64, _error := model.PostMenuLinkContent(&iMenuLinkContent)
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

func PutMenuLinkContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	_int64, _error := model.PutMenuLinkContent(&iMenuLinkContent)
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

func PutMenuLinkContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	_int64, _error := model.PutMenuLinkContentViaId(Id_, &iMenuLinkContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	_int64, _error := model.PutMenuLinkContentViaBundle(Bundle_, &iMenuLinkContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	_int64, _error := model.PutMenuLinkContentViaUuid(Uuid_, &iMenuLinkContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	_int64, _error := model.PutMenuLinkContentViaLangcode(Langcode_, &iMenuLinkContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	var iMap = helper.StructToMap(iMenuLinkContent)
	_error := model.UpdateMenuLinkContentViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	var iMap = helper.StructToMap(iMenuLinkContent)
	_error := model.UpdateMenuLinkContentViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	var iMap = helper.StructToMap(iMenuLinkContent)
	_error := model.UpdateMenuLinkContentViaUuid(Uuid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iMenuLinkContent model.MenuLinkContent
	self.Bind(&iMenuLinkContent)
	var iMap = helper.StructToMap(iMenuLinkContent)
	_error := model.UpdateMenuLinkContentViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteMenuLinkContent(Id_)
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

func DeleteMenuLinkContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteMenuLinkContentViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteMenuLinkContentViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	_error := model.DeleteMenuLinkContentViaUuid(Uuid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteMenuLinkContentViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
