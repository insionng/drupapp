package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetMenuLinkContentDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetMenuLinkContentDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["menu_link_content_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetMenuLinkContentDataCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetMenuLinkContentDataCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetMenuLinkContentDataCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaTitleHandler(self *macross.Context) error {
	Title_ := self.Args("title").String()
	_int64 := model.GetMenuLinkContentDataCountViaTitle(Title_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaDescriptionHandler(self *macross.Context) error {
	Description_ := self.Args("description").String()
	_int64 := model.GetMenuLinkContentDataCountViaDescription(Description_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaMenuNameHandler(self *macross.Context) error {
	MenuName_ := self.Args("menu_name").String()
	_int64 := model.GetMenuLinkContentDataCountViaMenuName(MenuName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaLink_uriHandler(self *macross.Context) error {
	Link_uri_ := self.Args("link__uri").String()
	_int64 := model.GetMenuLinkContentDataCountViaLink_uri(Link_uri_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaLink_titleHandler(self *macross.Context) error {
	Link_title_ := self.Args("link__title").String()
	_int64 := model.GetMenuLinkContentDataCountViaLink_title(Link_title_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaLink_optionsHandler(self *macross.Context) error {
	Link_options_ := self.Args("link__options").Bytes()
	_int64 := model.GetMenuLinkContentDataCountViaLink_options(Link_options_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaExternalHandler(self *macross.Context) error {
	External_ := self.Args("external").MustInt()
	_int64 := model.GetMenuLinkContentDataCountViaExternal(External_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaRediscoverHandler(self *macross.Context) error {
	Rediscover_ := self.Args("rediscover").MustInt()
	_int64 := model.GetMenuLinkContentDataCountViaRediscover(Rediscover_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaWeightHandler(self *macross.Context) error {
	Weight_ := self.Args("weight").MustInt()
	_int64 := model.GetMenuLinkContentDataCountViaWeight(Weight_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaExpandedHandler(self *macross.Context) error {
	Expanded_ := self.Args("expanded").MustInt()
	_int64 := model.GetMenuLinkContentDataCountViaExpanded(Expanded_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaEnabledHandler(self *macross.Context) error {
	Enabled_ := self.Args("enabled").MustInt()
	_int64 := model.GetMenuLinkContentDataCountViaEnabled(Enabled_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaParentHandler(self *macross.Context) error {
	Parent_ := self.Args("parent").String()
	_int64 := model.GetMenuLinkContentDataCountViaParent(Parent_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetMenuLinkContentDataCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDataCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetMenuLinkContentDataCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["menu_link_content_dataCount"] = 0
	}
	m["menu_link_content_dataCount"] = _int64
	return self.JSON(m)
}

func GetMenuLinkContentDatasViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTitle := self.Args("title").String()
	if (offset > 0) && helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaTitle(offset, limit, iTitle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDescription := self.Args("description").String()
	if (offset > 0) && helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaDescription(offset, limit, iDescription, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMenuName := self.Args("menu_name").String()
	if (offset > 0) && helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaMenuName(offset, limit, iMenuName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLink_uri := self.Args("link__uri").String()
	if (offset > 0) && helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaLink_uri(offset, limit, iLink_uri, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLink_title := self.Args("link__title").String()
	if (offset > 0) && helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaLink_title(offset, limit, iLink_title, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLink_options := self.Args("link__options").Bytes()
	if (offset > 0) && helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaLink_options(offset, limit, iLink_options, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExternal := self.Args("external").MustInt()
	if (offset > 0) && helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaExternal(offset, limit, iExternal, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRediscover := self.Args("rediscover").MustInt()
	if (offset > 0) && helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaRediscover(offset, limit, iRediscover, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iWeight := self.Args("weight").MustInt()
	if (offset > 0) && helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaWeight(offset, limit, iWeight, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpanded := self.Args("expanded").MustInt()
	if (offset > 0) && helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaExpanded(offset, limit, iExpanded, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEnabled := self.Args("enabled").MustInt()
	if (offset > 0) && helper.IsHas(iEnabled) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaEnabled(offset, limit, iEnabled, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iParent := self.Args("parent").String()
	if (offset > 0) && helper.IsHas(iParent) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaParent(offset, limit, iParent, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndLangcode(offset, limit, iId,iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndTitle(offset, limit, iId,iBundle,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndDescription(offset, limit, iId,iBundle,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndMenuName(offset, limit, iId,iBundle,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndLink_uri(offset, limit, iId,iBundle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndLink_title(offset, limit, iId,iBundle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndLink_options(offset, limit, iId,iBundle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndExternal(offset, limit, iId,iBundle,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndRediscover(offset, limit, iId,iBundle,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndWeight(offset, limit, iId,iBundle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndExpanded(offset, limit, iId,iBundle,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndEnabled(offset, limit, iId,iBundle,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndParent(offset, limit, iId,iBundle,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndChanged(offset, limit, iId,iBundle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundleAndDefaultLangcode(offset, limit, iId,iBundle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndTitle(offset, limit, iId,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndDescription(offset, limit, iId,iLangcode,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndMenuName(offset, limit, iId,iLangcode,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndLink_uri(offset, limit, iId,iLangcode,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndLink_title(offset, limit, iId,iLangcode,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndLink_options(offset, limit, iId,iLangcode,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndExternal(offset, limit, iId,iLangcode,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndRediscover(offset, limit, iId,iLangcode,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndWeight(offset, limit, iId,iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndExpanded(offset, limit, iId,iLangcode,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndEnabled(offset, limit, iId,iLangcode,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndParent(offset, limit, iId,iLangcode,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndChanged(offset, limit, iId,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcodeAndDefaultLangcode(offset, limit, iId,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndDescription(offset, limit, iId,iTitle,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndMenuName(offset, limit, iId,iTitle,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndLink_uri(offset, limit, iId,iTitle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndLink_title(offset, limit, iId,iTitle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndLink_options(offset, limit, iId,iTitle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndExternal(offset, limit, iId,iTitle,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndRediscover(offset, limit, iId,iTitle,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndWeight(offset, limit, iId,iTitle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndExpanded(offset, limit, iId,iTitle,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndEnabled(offset, limit, iId,iTitle,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndParent(offset, limit, iId,iTitle,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndChanged(offset, limit, iId,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitleAndDefaultLangcode(offset, limit, iId,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndMenuName(offset, limit, iId,iDescription,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndLink_uri(offset, limit, iId,iDescription,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndLink_title(offset, limit, iId,iDescription,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndLink_options(offset, limit, iId,iDescription,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndExternal(offset, limit, iId,iDescription,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndRediscover(offset, limit, iId,iDescription,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndWeight(offset, limit, iId,iDescription,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndExpanded(offset, limit, iId,iDescription,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndEnabled(offset, limit, iId,iDescription,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndParent(offset, limit, iId,iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndChanged(offset, limit, iId,iDescription,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescriptionAndDefaultLangcode(offset, limit, iId,iDescription,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescriptionAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndLink_uri(offset, limit, iId,iMenuName,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndLink_title(offset, limit, iId,iMenuName,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndLink_options(offset, limit, iId,iMenuName,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndExternal(offset, limit, iId,iMenuName,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndRediscover(offset, limit, iId,iMenuName,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndWeight(offset, limit, iId,iMenuName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndExpanded(offset, limit, iId,iMenuName,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndEnabled(offset, limit, iId,iMenuName,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndParent(offset, limit, iId,iMenuName,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndChanged(offset, limit, iId,iMenuName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuNameAndDefaultLangcode(offset, limit, iId,iMenuName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndLink_title(offset, limit, iId,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndLink_options(offset, limit, iId,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndExternal(offset, limit, iId,iLink_uri,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndRediscover(offset, limit, iId,iLink_uri,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndWeight(offset, limit, iId,iLink_uri,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndExpanded(offset, limit, iId,iLink_uri,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndEnabled(offset, limit, iId,iLink_uri,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndParent(offset, limit, iId,iLink_uri,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndChanged(offset, limit, iId,iLink_uri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uriAndDefaultLangcode(offset, limit, iId,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndLink_options(offset, limit, iId,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndExternal(offset, limit, iId,iLink_title,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndRediscover(offset, limit, iId,iLink_title,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndWeight(offset, limit, iId,iLink_title,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndExpanded(offset, limit, iId,iLink_title,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndEnabled(offset, limit, iId,iLink_title,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndParent(offset, limit, iId,iLink_title,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndChanged(offset, limit, iId,iLink_title,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_titleAndDefaultLangcode(offset, limit, iId,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_optionsAndExternal(offset, limit, iId,iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_optionsAndRediscover(offset, limit, iId,iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_optionsAndWeight(offset, limit, iId,iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_optionsAndExpanded(offset, limit, iId,iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_optionsAndEnabled(offset, limit, iId,iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_optionsAndParent(offset, limit, iId,iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_optionsAndChanged(offset, limit, iId,iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_optionsAndDefaultLangcode(offset, limit, iId,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExternalAndRediscover(offset, limit, iId,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExternalAndWeight(offset, limit, iId,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExternalAndExpanded(offset, limit, iId,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExternalAndEnabled(offset, limit, iId,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExternalAndParent(offset, limit, iId,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExternalAndChanged(offset, limit, iId,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExternalAndDefaultLangcode(offset, limit, iId,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndRediscoverAndWeight(offset, limit, iId,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndRediscoverAndExpanded(offset, limit, iId,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndRediscoverAndEnabled(offset, limit, iId,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndRediscoverAndParent(offset, limit, iId,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndRediscoverAndChanged(offset, limit, iId,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndRediscoverAndDefaultLangcode(offset, limit, iId,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndWeightAndExpanded(offset, limit, iId,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndWeightAndEnabled(offset, limit, iId,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndWeightAndParent(offset, limit, iId,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndWeightAndChanged(offset, limit, iId,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndWeightAndDefaultLangcode(offset, limit, iId,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExpandedAndEnabled(offset, limit, iId,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExpandedAndParent(offset, limit, iId,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExpandedAndChanged(offset, limit, iId,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExpandedAndDefaultLangcode(offset, limit, iId,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndEnabledAndParent(offset, limit, iId,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndEnabledAndChanged(offset, limit, iId,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndEnabledAndDefaultLangcode(offset, limit, iId,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndParentAndChanged(offset, limit, iId,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndParentAndDefaultLangcode(offset, limit, iId,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndChangedAndDefaultLangcode(offset, limit, iId,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndTitle(offset, limit, iBundle,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndDescription(offset, limit, iBundle,iLangcode,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndMenuName(offset, limit, iBundle,iLangcode,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndLink_uri(offset, limit, iBundle,iLangcode,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndLink_title(offset, limit, iBundle,iLangcode,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndLink_options(offset, limit, iBundle,iLangcode,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndExternal(offset, limit, iBundle,iLangcode,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndRediscover(offset, limit, iBundle,iLangcode,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndWeight(offset, limit, iBundle,iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndExpanded(offset, limit, iBundle,iLangcode,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndEnabled(offset, limit, iBundle,iLangcode,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndParent(offset, limit, iBundle,iLangcode,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndChanged(offset, limit, iBundle,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcodeAndDefaultLangcode(offset, limit, iBundle,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndDescription(offset, limit, iBundle,iTitle,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndMenuName(offset, limit, iBundle,iTitle,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndLink_uri(offset, limit, iBundle,iTitle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndLink_title(offset, limit, iBundle,iTitle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndLink_options(offset, limit, iBundle,iTitle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndExternal(offset, limit, iBundle,iTitle,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndRediscover(offset, limit, iBundle,iTitle,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndWeight(offset, limit, iBundle,iTitle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndExpanded(offset, limit, iBundle,iTitle,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndEnabled(offset, limit, iBundle,iTitle,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndParent(offset, limit, iBundle,iTitle,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndChanged(offset, limit, iBundle,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitleAndDefaultLangcode(offset, limit, iBundle,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndMenuName(offset, limit, iBundle,iDescription,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndLink_uri(offset, limit, iBundle,iDescription,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndLink_title(offset, limit, iBundle,iDescription,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndLink_options(offset, limit, iBundle,iDescription,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndExternal(offset, limit, iBundle,iDescription,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndRediscover(offset, limit, iBundle,iDescription,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndWeight(offset, limit, iBundle,iDescription,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndExpanded(offset, limit, iBundle,iDescription,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndEnabled(offset, limit, iBundle,iDescription,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndParent(offset, limit, iBundle,iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndChanged(offset, limit, iBundle,iDescription,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescriptionAndDefaultLangcode(offset, limit, iBundle,iDescription,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescriptionAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndLink_uri(offset, limit, iBundle,iMenuName,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndLink_title(offset, limit, iBundle,iMenuName,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndLink_options(offset, limit, iBundle,iMenuName,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndExternal(offset, limit, iBundle,iMenuName,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndRediscover(offset, limit, iBundle,iMenuName,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndWeight(offset, limit, iBundle,iMenuName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndExpanded(offset, limit, iBundle,iMenuName,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndEnabled(offset, limit, iBundle,iMenuName,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndParent(offset, limit, iBundle,iMenuName,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndChanged(offset, limit, iBundle,iMenuName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuNameAndDefaultLangcode(offset, limit, iBundle,iMenuName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndLink_title(offset, limit, iBundle,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndLink_options(offset, limit, iBundle,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndExternal(offset, limit, iBundle,iLink_uri,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndRediscover(offset, limit, iBundle,iLink_uri,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndWeight(offset, limit, iBundle,iLink_uri,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndExpanded(offset, limit, iBundle,iLink_uri,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndEnabled(offset, limit, iBundle,iLink_uri,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndParent(offset, limit, iBundle,iLink_uri,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndChanged(offset, limit, iBundle,iLink_uri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uriAndDefaultLangcode(offset, limit, iBundle,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndLink_options(offset, limit, iBundle,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndExternal(offset, limit, iBundle,iLink_title,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndRediscover(offset, limit, iBundle,iLink_title,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndWeight(offset, limit, iBundle,iLink_title,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndExpanded(offset, limit, iBundle,iLink_title,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndEnabled(offset, limit, iBundle,iLink_title,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndParent(offset, limit, iBundle,iLink_title,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndChanged(offset, limit, iBundle,iLink_title,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_titleAndDefaultLangcode(offset, limit, iBundle,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_optionsAndExternal(offset, limit, iBundle,iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_optionsAndRediscover(offset, limit, iBundle,iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_optionsAndWeight(offset, limit, iBundle,iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_optionsAndExpanded(offset, limit, iBundle,iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_optionsAndEnabled(offset, limit, iBundle,iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_optionsAndParent(offset, limit, iBundle,iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_optionsAndChanged(offset, limit, iBundle,iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_optionsAndDefaultLangcode(offset, limit, iBundle,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExternalAndRediscover(offset, limit, iBundle,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExternalAndWeight(offset, limit, iBundle,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExternalAndExpanded(offset, limit, iBundle,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExternalAndEnabled(offset, limit, iBundle,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExternalAndParent(offset, limit, iBundle,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExternalAndChanged(offset, limit, iBundle,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExternalAndDefaultLangcode(offset, limit, iBundle,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndRediscoverAndWeight(offset, limit, iBundle,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndRediscoverAndExpanded(offset, limit, iBundle,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndRediscoverAndEnabled(offset, limit, iBundle,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndRediscoverAndParent(offset, limit, iBundle,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndRediscoverAndChanged(offset, limit, iBundle,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndRediscoverAndDefaultLangcode(offset, limit, iBundle,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndWeightAndExpanded(offset, limit, iBundle,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndWeightAndEnabled(offset, limit, iBundle,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndWeightAndParent(offset, limit, iBundle,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndWeightAndChanged(offset, limit, iBundle,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndWeightAndDefaultLangcode(offset, limit, iBundle,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExpandedAndEnabled(offset, limit, iBundle,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExpandedAndParent(offset, limit, iBundle,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExpandedAndChanged(offset, limit, iBundle,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExpandedAndDefaultLangcode(offset, limit, iBundle,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndEnabledAndParent(offset, limit, iBundle,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndEnabledAndChanged(offset, limit, iBundle,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndEnabledAndDefaultLangcode(offset, limit, iBundle,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndParentAndChanged(offset, limit, iBundle,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndParentAndDefaultLangcode(offset, limit, iBundle,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndChangedAndDefaultLangcode(offset, limit, iBundle,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndDescription(offset, limit, iLangcode,iTitle,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndMenuName(offset, limit, iLangcode,iTitle,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndLink_uri(offset, limit, iLangcode,iTitle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndLink_title(offset, limit, iLangcode,iTitle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndLink_options(offset, limit, iLangcode,iTitle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndExternal(offset, limit, iLangcode,iTitle,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndRediscover(offset, limit, iLangcode,iTitle,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndWeight(offset, limit, iLangcode,iTitle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndExpanded(offset, limit, iLangcode,iTitle,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndEnabled(offset, limit, iLangcode,iTitle,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndParent(offset, limit, iLangcode,iTitle,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndChanged(offset, limit, iLangcode,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitleAndDefaultLangcode(offset, limit, iLangcode,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndMenuName(offset, limit, iLangcode,iDescription,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_uri(offset, limit, iLangcode,iDescription,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_title(offset, limit, iLangcode,iDescription,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_options(offset, limit, iLangcode,iDescription,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndExternal(offset, limit, iLangcode,iDescription,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndRediscover(offset, limit, iLangcode,iDescription,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndWeight(offset, limit, iLangcode,iDescription,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndExpanded(offset, limit, iLangcode,iDescription,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndEnabled(offset, limit, iLangcode,iDescription,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndParent(offset, limit, iLangcode,iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndChanged(offset, limit, iLangcode,iDescription,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescriptionAndDefaultLangcode(offset, limit, iLangcode,iDescription,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescriptionAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_uri(offset, limit, iLangcode,iMenuName,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_title(offset, limit, iLangcode,iMenuName,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_options(offset, limit, iLangcode,iMenuName,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndExternal(offset, limit, iLangcode,iMenuName,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndRediscover(offset, limit, iLangcode,iMenuName,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndWeight(offset, limit, iLangcode,iMenuName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndExpanded(offset, limit, iLangcode,iMenuName,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndEnabled(offset, limit, iLangcode,iMenuName,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndParent(offset, limit, iLangcode,iMenuName,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndChanged(offset, limit, iLangcode,iMenuName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuNameAndDefaultLangcode(offset, limit, iLangcode,iMenuName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_title(offset, limit, iLangcode,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_options(offset, limit, iLangcode,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndExternal(offset, limit, iLangcode,iLink_uri,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndRediscover(offset, limit, iLangcode,iLink_uri,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndWeight(offset, limit, iLangcode,iLink_uri,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndExpanded(offset, limit, iLangcode,iLink_uri,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndEnabled(offset, limit, iLangcode,iLink_uri,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndParent(offset, limit, iLangcode,iLink_uri,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndChanged(offset, limit, iLangcode,iLink_uri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uriAndDefaultLangcode(offset, limit, iLangcode,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndLink_options(offset, limit, iLangcode,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndExternal(offset, limit, iLangcode,iLink_title,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndRediscover(offset, limit, iLangcode,iLink_title,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndWeight(offset, limit, iLangcode,iLink_title,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndExpanded(offset, limit, iLangcode,iLink_title,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndEnabled(offset, limit, iLangcode,iLink_title,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndParent(offset, limit, iLangcode,iLink_title,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndChanged(offset, limit, iLangcode,iLink_title,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_titleAndDefaultLangcode(offset, limit, iLangcode,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExternal(offset, limit, iLangcode,iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_optionsAndRediscover(offset, limit, iLangcode,iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_optionsAndWeight(offset, limit, iLangcode,iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExpanded(offset, limit, iLangcode,iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_optionsAndEnabled(offset, limit, iLangcode,iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_optionsAndParent(offset, limit, iLangcode,iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_optionsAndChanged(offset, limit, iLangcode,iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_optionsAndDefaultLangcode(offset, limit, iLangcode,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExternalAndRediscover(offset, limit, iLangcode,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExternalAndWeight(offset, limit, iLangcode,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExternalAndExpanded(offset, limit, iLangcode,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExternalAndEnabled(offset, limit, iLangcode,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExternalAndParent(offset, limit, iLangcode,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExternalAndChanged(offset, limit, iLangcode,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExternalAndDefaultLangcode(offset, limit, iLangcode,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndRediscoverAndWeight(offset, limit, iLangcode,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndRediscoverAndExpanded(offset, limit, iLangcode,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndRediscoverAndEnabled(offset, limit, iLangcode,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndRediscoverAndParent(offset, limit, iLangcode,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndRediscoverAndChanged(offset, limit, iLangcode,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndRediscoverAndDefaultLangcode(offset, limit, iLangcode,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndWeightAndExpanded(offset, limit, iLangcode,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndWeightAndEnabled(offset, limit, iLangcode,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndWeightAndParent(offset, limit, iLangcode,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndWeightAndChanged(offset, limit, iLangcode,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndWeightAndDefaultLangcode(offset, limit, iLangcode,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExpandedAndEnabled(offset, limit, iLangcode,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExpandedAndParent(offset, limit, iLangcode,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExpandedAndChanged(offset, limit, iLangcode,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExpandedAndDefaultLangcode(offset, limit, iLangcode,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndEnabledAndParent(offset, limit, iLangcode,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndEnabledAndChanged(offset, limit, iLangcode,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndEnabledAndDefaultLangcode(offset, limit, iLangcode,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndParentAndChanged(offset, limit, iLangcode,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndParentAndDefaultLangcode(offset, limit, iLangcode,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndChangedAndDefaultLangcode(offset, limit, iLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndMenuName(offset, limit, iTitle,iDescription,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndLink_uri(offset, limit, iTitle,iDescription,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndLink_title(offset, limit, iTitle,iDescription,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndLink_options(offset, limit, iTitle,iDescription,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndExternal(offset, limit, iTitle,iDescription,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndRediscover(offset, limit, iTitle,iDescription,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndWeight(offset, limit, iTitle,iDescription,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndExpanded(offset, limit, iTitle,iDescription,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndEnabled(offset, limit, iTitle,iDescription,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndParent(offset, limit, iTitle,iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndChanged(offset, limit, iTitle,iDescription,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescriptionAndDefaultLangcode(offset, limit, iTitle,iDescription,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescriptionAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndLink_uri(offset, limit, iTitle,iMenuName,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndLink_title(offset, limit, iTitle,iMenuName,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndLink_options(offset, limit, iTitle,iMenuName,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndExternal(offset, limit, iTitle,iMenuName,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndRediscover(offset, limit, iTitle,iMenuName,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndWeight(offset, limit, iTitle,iMenuName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndExpanded(offset, limit, iTitle,iMenuName,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndEnabled(offset, limit, iTitle,iMenuName,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndParent(offset, limit, iTitle,iMenuName,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndChanged(offset, limit, iTitle,iMenuName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuNameAndDefaultLangcode(offset, limit, iTitle,iMenuName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndLink_title(offset, limit, iTitle,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndLink_options(offset, limit, iTitle,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndExternal(offset, limit, iTitle,iLink_uri,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndRediscover(offset, limit, iTitle,iLink_uri,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndWeight(offset, limit, iTitle,iLink_uri,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndExpanded(offset, limit, iTitle,iLink_uri,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndEnabled(offset, limit, iTitle,iLink_uri,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndParent(offset, limit, iTitle,iLink_uri,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndChanged(offset, limit, iTitle,iLink_uri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uriAndDefaultLangcode(offset, limit, iTitle,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndLink_options(offset, limit, iTitle,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndExternal(offset, limit, iTitle,iLink_title,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndRediscover(offset, limit, iTitle,iLink_title,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndWeight(offset, limit, iTitle,iLink_title,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndExpanded(offset, limit, iTitle,iLink_title,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndEnabled(offset, limit, iTitle,iLink_title,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndParent(offset, limit, iTitle,iLink_title,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndChanged(offset, limit, iTitle,iLink_title,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_titleAndDefaultLangcode(offset, limit, iTitle,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_optionsAndExternal(offset, limit, iTitle,iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_optionsAndRediscover(offset, limit, iTitle,iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_optionsAndWeight(offset, limit, iTitle,iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_optionsAndExpanded(offset, limit, iTitle,iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_optionsAndEnabled(offset, limit, iTitle,iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_optionsAndParent(offset, limit, iTitle,iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_optionsAndChanged(offset, limit, iTitle,iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_optionsAndDefaultLangcode(offset, limit, iTitle,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExternalAndRediscover(offset, limit, iTitle,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExternalAndWeight(offset, limit, iTitle,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExternalAndExpanded(offset, limit, iTitle,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExternalAndEnabled(offset, limit, iTitle,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExternalAndParent(offset, limit, iTitle,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExternalAndChanged(offset, limit, iTitle,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExternalAndDefaultLangcode(offset, limit, iTitle,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndRediscoverAndWeight(offset, limit, iTitle,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndRediscoverAndExpanded(offset, limit, iTitle,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndRediscoverAndEnabled(offset, limit, iTitle,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndRediscoverAndParent(offset, limit, iTitle,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndRediscoverAndChanged(offset, limit, iTitle,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndRediscoverAndDefaultLangcode(offset, limit, iTitle,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndWeightAndExpanded(offset, limit, iTitle,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndWeightAndEnabled(offset, limit, iTitle,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndWeightAndParent(offset, limit, iTitle,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndWeightAndChanged(offset, limit, iTitle,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndWeightAndDefaultLangcode(offset, limit, iTitle,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExpandedAndEnabled(offset, limit, iTitle,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExpandedAndParent(offset, limit, iTitle,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExpandedAndChanged(offset, limit, iTitle,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExpandedAndDefaultLangcode(offset, limit, iTitle,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndEnabledAndParent(offset, limit, iTitle,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndEnabledAndChanged(offset, limit, iTitle,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndEnabledAndDefaultLangcode(offset, limit, iTitle,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndParentAndChanged(offset, limit, iTitle,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndParentAndDefaultLangcode(offset, limit, iTitle,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndChangedAndDefaultLangcode(offset, limit, iTitle,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_uri(offset, limit, iDescription,iMenuName,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_title(offset, limit, iDescription,iMenuName,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_options(offset, limit, iDescription,iMenuName,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndExternal(offset, limit, iDescription,iMenuName,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndRediscover(offset, limit, iDescription,iMenuName,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndWeight(offset, limit, iDescription,iMenuName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndExpanded(offset, limit, iDescription,iMenuName,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndEnabled(offset, limit, iDescription,iMenuName,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndParent(offset, limit, iDescription,iMenuName,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndChanged(offset, limit, iDescription,iMenuName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuNameAndDefaultLangcode(offset, limit, iDescription,iMenuName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_title(offset, limit, iDescription,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_options(offset, limit, iDescription,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndExternal(offset, limit, iDescription,iLink_uri,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndRediscover(offset, limit, iDescription,iLink_uri,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndWeight(offset, limit, iDescription,iLink_uri,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndExpanded(offset, limit, iDescription,iLink_uri,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndEnabled(offset, limit, iDescription,iLink_uri,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndParent(offset, limit, iDescription,iLink_uri,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndChanged(offset, limit, iDescription,iLink_uri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uriAndDefaultLangcode(offset, limit, iDescription,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndLink_options(offset, limit, iDescription,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndExternal(offset, limit, iDescription,iLink_title,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndRediscover(offset, limit, iDescription,iLink_title,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndWeight(offset, limit, iDescription,iLink_title,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndExpanded(offset, limit, iDescription,iLink_title,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndEnabled(offset, limit, iDescription,iLink_title,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndParent(offset, limit, iDescription,iLink_title,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndChanged(offset, limit, iDescription,iLink_title,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_titleAndDefaultLangcode(offset, limit, iDescription,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExternal(offset, limit, iDescription,iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_optionsAndRediscover(offset, limit, iDescription,iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_optionsAndWeight(offset, limit, iDescription,iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExpanded(offset, limit, iDescription,iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_optionsAndEnabled(offset, limit, iDescription,iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_optionsAndParent(offset, limit, iDescription,iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_optionsAndChanged(offset, limit, iDescription,iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_optionsAndDefaultLangcode(offset, limit, iDescription,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExternalAndRediscover(offset, limit, iDescription,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExternalAndWeight(offset, limit, iDescription,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExternalAndExpanded(offset, limit, iDescription,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExternalAndEnabled(offset, limit, iDescription,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExternalAndParent(offset, limit, iDescription,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExternalAndChanged(offset, limit, iDescription,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExternalAndDefaultLangcode(offset, limit, iDescription,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndRediscoverAndWeight(offset, limit, iDescription,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndRediscoverAndExpanded(offset, limit, iDescription,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndRediscoverAndEnabled(offset, limit, iDescription,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndRediscoverAndParent(offset, limit, iDescription,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndRediscoverAndChanged(offset, limit, iDescription,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndRediscoverAndDefaultLangcode(offset, limit, iDescription,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndWeightAndExpanded(offset, limit, iDescription,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndWeightAndEnabled(offset, limit, iDescription,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndWeightAndParent(offset, limit, iDescription,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndWeightAndChanged(offset, limit, iDescription,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndWeightAndDefaultLangcode(offset, limit, iDescription,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExpandedAndEnabled(offset, limit, iDescription,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExpandedAndParent(offset, limit, iDescription,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExpandedAndChanged(offset, limit, iDescription,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExpandedAndDefaultLangcode(offset, limit, iDescription,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndEnabledAndParent(offset, limit, iDescription,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndEnabledAndChanged(offset, limit, iDescription,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndEnabledAndDefaultLangcode(offset, limit, iDescription,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndParentAndChanged(offset, limit, iDescription,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndParentAndDefaultLangcode(offset, limit, iDescription,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndChangedAndDefaultLangcode(offset, limit, iDescription,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_title(offset, limit, iMenuName,iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_options(offset, limit, iMenuName,iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndExternal(offset, limit, iMenuName,iLink_uri,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndRediscover(offset, limit, iMenuName,iLink_uri,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndWeight(offset, limit, iMenuName,iLink_uri,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndExpanded(offset, limit, iMenuName,iLink_uri,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndEnabled(offset, limit, iMenuName,iLink_uri,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndParent(offset, limit, iMenuName,iLink_uri,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndChanged(offset, limit, iMenuName,iLink_uri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uriAndDefaultLangcode(offset, limit, iMenuName,iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndLink_options(offset, limit, iMenuName,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndExternal(offset, limit, iMenuName,iLink_title,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndRediscover(offset, limit, iMenuName,iLink_title,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndWeight(offset, limit, iMenuName,iLink_title,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndExpanded(offset, limit, iMenuName,iLink_title,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndEnabled(offset, limit, iMenuName,iLink_title,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndParent(offset, limit, iMenuName,iLink_title,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndChanged(offset, limit, iMenuName,iLink_title,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_titleAndDefaultLangcode(offset, limit, iMenuName,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExternal(offset, limit, iMenuName,iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_optionsAndRediscover(offset, limit, iMenuName,iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_optionsAndWeight(offset, limit, iMenuName,iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExpanded(offset, limit, iMenuName,iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_optionsAndEnabled(offset, limit, iMenuName,iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_optionsAndParent(offset, limit, iMenuName,iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_optionsAndChanged(offset, limit, iMenuName,iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_optionsAndDefaultLangcode(offset, limit, iMenuName,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExternalAndRediscover(offset, limit, iMenuName,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExternalAndWeight(offset, limit, iMenuName,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExternalAndExpanded(offset, limit, iMenuName,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExternalAndEnabled(offset, limit, iMenuName,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExternalAndParent(offset, limit, iMenuName,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExternalAndChanged(offset, limit, iMenuName,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExternalAndDefaultLangcode(offset, limit, iMenuName,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndRediscoverAndWeight(offset, limit, iMenuName,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndRediscoverAndExpanded(offset, limit, iMenuName,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndRediscoverAndEnabled(offset, limit, iMenuName,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndRediscoverAndParent(offset, limit, iMenuName,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndRediscoverAndChanged(offset, limit, iMenuName,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndRediscoverAndDefaultLangcode(offset, limit, iMenuName,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndWeightAndExpanded(offset, limit, iMenuName,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndWeightAndEnabled(offset, limit, iMenuName,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndWeightAndParent(offset, limit, iMenuName,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndWeightAndChanged(offset, limit, iMenuName,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndWeightAndDefaultLangcode(offset, limit, iMenuName,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExpandedAndEnabled(offset, limit, iMenuName,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExpandedAndParent(offset, limit, iMenuName,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExpandedAndChanged(offset, limit, iMenuName,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExpandedAndDefaultLangcode(offset, limit, iMenuName,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndEnabledAndParent(offset, limit, iMenuName,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndEnabledAndChanged(offset, limit, iMenuName,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndEnabledAndDefaultLangcode(offset, limit, iMenuName,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndParentAndChanged(offset, limit, iMenuName,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndParentAndDefaultLangcode(offset, limit, iMenuName,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndChangedAndDefaultLangcode(offset, limit, iMenuName,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndLink_options(offset, limit, iLink_uri,iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndExternal(offset, limit, iLink_uri,iLink_title,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndRediscover(offset, limit, iLink_uri,iLink_title,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndWeight(offset, limit, iLink_uri,iLink_title,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndExpanded(offset, limit, iLink_uri,iLink_title,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndEnabled(offset, limit, iLink_uri,iLink_title,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndParent(offset, limit, iLink_uri,iLink_title,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndChanged(offset, limit, iLink_uri,iLink_title,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_titleAndDefaultLangcode(offset, limit, iLink_uri,iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExternal(offset, limit, iLink_uri,iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_optionsAndRediscover(offset, limit, iLink_uri,iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_optionsAndWeight(offset, limit, iLink_uri,iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExpanded(offset, limit, iLink_uri,iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_optionsAndEnabled(offset, limit, iLink_uri,iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_optionsAndParent(offset, limit, iLink_uri,iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_optionsAndChanged(offset, limit, iLink_uri,iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_optionsAndDefaultLangcode(offset, limit, iLink_uri,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExternalAndRediscover(offset, limit, iLink_uri,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExternalAndWeight(offset, limit, iLink_uri,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExternalAndExpanded(offset, limit, iLink_uri,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExternalAndEnabled(offset, limit, iLink_uri,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExternalAndParent(offset, limit, iLink_uri,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExternalAndChanged(offset, limit, iLink_uri,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExternalAndDefaultLangcode(offset, limit, iLink_uri,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndRediscoverAndWeight(offset, limit, iLink_uri,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndRediscoverAndExpanded(offset, limit, iLink_uri,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndRediscoverAndEnabled(offset, limit, iLink_uri,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndRediscoverAndParent(offset, limit, iLink_uri,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndRediscoverAndChanged(offset, limit, iLink_uri,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndRediscoverAndDefaultLangcode(offset, limit, iLink_uri,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndWeightAndExpanded(offset, limit, iLink_uri,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndWeightAndEnabled(offset, limit, iLink_uri,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndWeightAndParent(offset, limit, iLink_uri,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndWeightAndChanged(offset, limit, iLink_uri,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndWeightAndDefaultLangcode(offset, limit, iLink_uri,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExpandedAndEnabled(offset, limit, iLink_uri,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExpandedAndParent(offset, limit, iLink_uri,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExpandedAndChanged(offset, limit, iLink_uri,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExpandedAndDefaultLangcode(offset, limit, iLink_uri,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndEnabledAndParent(offset, limit, iLink_uri,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndEnabledAndChanged(offset, limit, iLink_uri,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndEnabledAndDefaultLangcode(offset, limit, iLink_uri,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndParentAndChanged(offset, limit, iLink_uri,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndParentAndDefaultLangcode(offset, limit, iLink_uri,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndChangedAndDefaultLangcode(offset, limit, iLink_uri,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExternal(offset, limit, iLink_title,iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_optionsAndRediscover(offset, limit, iLink_title,iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_optionsAndWeight(offset, limit, iLink_title,iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExpanded(offset, limit, iLink_title,iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_optionsAndEnabled(offset, limit, iLink_title,iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_optionsAndParent(offset, limit, iLink_title,iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_optionsAndChanged(offset, limit, iLink_title,iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_optionsAndDefaultLangcode(offset, limit, iLink_title,iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExternalAndRediscover(offset, limit, iLink_title,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExternalAndWeight(offset, limit, iLink_title,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExternalAndExpanded(offset, limit, iLink_title,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExternalAndEnabled(offset, limit, iLink_title,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExternalAndParent(offset, limit, iLink_title,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExternalAndChanged(offset, limit, iLink_title,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExternalAndDefaultLangcode(offset, limit, iLink_title,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndRediscoverAndWeight(offset, limit, iLink_title,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndRediscoverAndExpanded(offset, limit, iLink_title,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndRediscoverAndEnabled(offset, limit, iLink_title,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndRediscoverAndParent(offset, limit, iLink_title,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndRediscoverAndChanged(offset, limit, iLink_title,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndRediscoverAndDefaultLangcode(offset, limit, iLink_title,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndWeightAndExpanded(offset, limit, iLink_title,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndWeightAndEnabled(offset, limit, iLink_title,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndWeightAndParent(offset, limit, iLink_title,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndWeightAndChanged(offset, limit, iLink_title,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndWeightAndDefaultLangcode(offset, limit, iLink_title,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExpandedAndEnabled(offset, limit, iLink_title,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExpandedAndParent(offset, limit, iLink_title,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExpandedAndChanged(offset, limit, iLink_title,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExpandedAndDefaultLangcode(offset, limit, iLink_title,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndEnabledAndParent(offset, limit, iLink_title,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndEnabledAndChanged(offset, limit, iLink_title,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndEnabledAndDefaultLangcode(offset, limit, iLink_title,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndParentAndChanged(offset, limit, iLink_title,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndParentAndDefaultLangcode(offset, limit, iLink_title,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndChangedAndDefaultLangcode(offset, limit, iLink_title,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExternalAndRediscover(offset, limit, iLink_options,iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExternalAndWeight(offset, limit, iLink_options,iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExternalAndExpanded(offset, limit, iLink_options,iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExternalAndEnabled(offset, limit, iLink_options,iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExternalAndParent(offset, limit, iLink_options,iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExternalAndChanged(offset, limit, iLink_options,iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExternalAndDefaultLangcode(offset, limit, iLink_options,iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndRediscoverAndWeight(offset, limit, iLink_options,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndRediscoverAndExpanded(offset, limit, iLink_options,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndRediscoverAndEnabled(offset, limit, iLink_options,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndRediscoverAndParent(offset, limit, iLink_options,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndRediscoverAndChanged(offset, limit, iLink_options,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndRediscoverAndDefaultLangcode(offset, limit, iLink_options,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndWeightAndExpanded(offset, limit, iLink_options,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndWeightAndEnabled(offset, limit, iLink_options,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndWeightAndParent(offset, limit, iLink_options,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndWeightAndChanged(offset, limit, iLink_options,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndWeightAndDefaultLangcode(offset, limit, iLink_options,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExpandedAndEnabled(offset, limit, iLink_options,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExpandedAndParent(offset, limit, iLink_options,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExpandedAndChanged(offset, limit, iLink_options,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExpandedAndDefaultLangcode(offset, limit, iLink_options,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndEnabledAndParent(offset, limit, iLink_options,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndEnabledAndChanged(offset, limit, iLink_options,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndEnabledAndDefaultLangcode(offset, limit, iLink_options,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndParentAndChanged(offset, limit, iLink_options,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndParentAndDefaultLangcode(offset, limit, iLink_options,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndChangedAndDefaultLangcode(offset, limit, iLink_options,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndRediscoverAndWeight(offset, limit, iExternal,iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndRediscoverAndExpanded(offset, limit, iExternal,iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndRediscoverAndEnabled(offset, limit, iExternal,iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndRediscoverAndParent(offset, limit, iExternal,iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndRediscoverAndChanged(offset, limit, iExternal,iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndRediscoverAndDefaultLangcode(offset, limit, iExternal,iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndWeightAndExpanded(offset, limit, iExternal,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndWeightAndEnabled(offset, limit, iExternal,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndWeightAndParent(offset, limit, iExternal,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndWeightAndChanged(offset, limit, iExternal,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndWeightAndDefaultLangcode(offset, limit, iExternal,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndExpandedAndEnabled(offset, limit, iExternal,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndExpandedAndParent(offset, limit, iExternal,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndExpandedAndChanged(offset, limit, iExternal,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndExpandedAndDefaultLangcode(offset, limit, iExternal,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndEnabledAndParent(offset, limit, iExternal,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndEnabledAndChanged(offset, limit, iExternal,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndEnabledAndDefaultLangcode(offset, limit, iExternal,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndParentAndChanged(offset, limit, iExternal,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndParentAndDefaultLangcode(offset, limit, iExternal,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndChangedAndDefaultLangcode(offset, limit, iExternal,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndWeightAndExpanded(offset, limit, iRediscover,iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndWeightAndEnabled(offset, limit, iRediscover,iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndWeightAndParent(offset, limit, iRediscover,iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndWeightAndChanged(offset, limit, iRediscover,iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndWeightAndDefaultLangcode(offset, limit, iRediscover,iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndExpandedAndEnabled(offset, limit, iRediscover,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndExpandedAndParent(offset, limit, iRediscover,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndExpandedAndChanged(offset, limit, iRediscover,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndExpandedAndDefaultLangcode(offset, limit, iRediscover,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndEnabledAndParent(offset, limit, iRediscover,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndEnabledAndChanged(offset, limit, iRediscover,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndEnabledAndDefaultLangcode(offset, limit, iRediscover,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndParentAndChanged(offset, limit, iRediscover,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndParentAndDefaultLangcode(offset, limit, iRediscover,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndChangedAndDefaultLangcode(offset, limit, iRediscover,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndExpandedAndEnabled(offset, limit, iWeight,iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndExpandedAndParent(offset, limit, iWeight,iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndExpandedAndChanged(offset, limit, iWeight,iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndExpandedAndDefaultLangcode(offset, limit, iWeight,iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndEnabledAndParent(offset, limit, iWeight,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndEnabledAndChanged(offset, limit, iWeight,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndEnabledAndDefaultLangcode(offset, limit, iWeight,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndParentAndChanged(offset, limit, iWeight,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndParentAndDefaultLangcode(offset, limit, iWeight,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndChangedAndDefaultLangcode(offset, limit, iWeight,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndEnabledAndParent(offset, limit, iExpanded,iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndEnabledAndChanged(offset, limit, iExpanded,iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndEnabledAndDefaultLangcode(offset, limit, iExpanded,iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndParentAndChanged(offset, limit, iExpanded,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndParentAndDefaultLangcode(offset, limit, iExpanded,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndChangedAndDefaultLangcode(offset, limit, iExpanded,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByEnabledAndParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEnabled) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByEnabledAndParentAndChanged(offset, limit, iEnabled,iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByEnabledAndParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByEnabledAndParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEnabled) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByEnabledAndParentAndDefaultLangcode(offset, limit, iEnabled,iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByEnabledAndParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByEnabledAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEnabled) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByEnabledAndChangedAndDefaultLangcode(offset, limit, iEnabled,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByEnabledAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByParentAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iParent) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByParentAndChangedAndDefaultLangcode(offset, limit, iParent,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByParentAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBundle := self.Args("bundle").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndBundle(offset, limit, iId,iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLangcode(offset, limit, iId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iTitle := self.Args("title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndTitle(offset, limit, iId,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDescription := self.Args("description").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDescription(offset, limit, iId,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndMenuName(offset, limit, iId,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_uri(offset, limit, iId,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_title(offset, limit, iId,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndLink_options(offset, limit, iId,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExternal(offset, limit, iId,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndRediscover(offset, limit, iId,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndWeight(offset, limit, iId,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndExpanded(offset, limit, iId,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndEnabled(offset, limit, iId,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iParent := self.Args("parent").String()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndParent(offset, limit, iId,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndChanged(offset, limit, iId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByIdAndDefaultLangcode(offset, limit, iId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndTitle(offset, limit, iBundle,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDescription(offset, limit, iBundle,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndMenuName(offset, limit, iBundle,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_uri(offset, limit, iBundle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_title(offset, limit, iBundle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndLink_options(offset, limit, iBundle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExternal(offset, limit, iBundle,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndRediscover(offset, limit, iBundle,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndWeight(offset, limit, iBundle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndExpanded(offset, limit, iBundle,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndEnabled(offset, limit, iBundle,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndParent(offset, limit, iBundle,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndChanged(offset, limit, iBundle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByBundleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByBundleAndDefaultLangcode(offset, limit, iBundle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByBundleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndTitle(offset, limit, iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDescription(offset, limit, iLangcode,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndMenuName(offset, limit, iLangcode,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_uri(offset, limit, iLangcode,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_title(offset, limit, iLangcode,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndLink_options(offset, limit, iLangcode,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExternal(offset, limit, iLangcode,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndRediscover(offset, limit, iLangcode,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndWeight(offset, limit, iLangcode,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndExpanded(offset, limit, iLangcode,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndEnabled(offset, limit, iLangcode,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndParent(offset, limit, iLangcode,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDescription(offset, limit, iTitle,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndMenuName(offset, limit, iTitle,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_uri(offset, limit, iTitle,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_title(offset, limit, iTitle,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndLink_options(offset, limit, iTitle,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExternal(offset, limit, iTitle,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndRediscover(offset, limit, iTitle,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndWeight(offset, limit, iTitle,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndExpanded(offset, limit, iTitle,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndEnabled(offset, limit, iTitle,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndParent(offset, limit, iTitle,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndChanged(offset, limit, iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByTitleAndDefaultLangcode(offset, limit, iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iMenuName := self.Args("menu_name").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndMenuName(offset, limit, iDescription,iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_uri(offset, limit, iDescription,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_title(offset, limit, iDescription,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndLink_options(offset, limit, iDescription,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExternal(offset, limit, iDescription,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndRediscover(offset, limit, iDescription,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndWeight(offset, limit, iDescription,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndExpanded(offset, limit, iDescription,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndEnabled(offset, limit, iDescription,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndParent(offset, limit, iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndChanged(offset, limit, iDescription,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByDescriptionAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByDescriptionAndDefaultLangcode(offset, limit, iDescription,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByDescriptionAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_uri := self.Args("link_uri").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_uri(offset, limit, iMenuName,iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_title(offset, limit, iMenuName,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndLink_options(offset, limit, iMenuName,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExternal(offset, limit, iMenuName,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndRediscover(offset, limit, iMenuName,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndWeight(offset, limit, iMenuName,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndExpanded(offset, limit, iMenuName,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndEnabled(offset, limit, iMenuName,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndParent(offset, limit, iMenuName,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndChanged(offset, limit, iMenuName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByMenuNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMenuName := self.Args("menu_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByMenuNameAndDefaultLangcode(offset, limit, iMenuName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByMenuNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_title := self.Args("link_title").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_title(offset, limit, iLink_uri,iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndLink_options(offset, limit, iLink_uri,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExternal(offset, limit, iLink_uri,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndRediscover(offset, limit, iLink_uri,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndWeight(offset, limit, iLink_uri,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndExpanded(offset, limit, iLink_uri,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndEnabled(offset, limit, iLink_uri,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndParent(offset, limit, iLink_uri,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndChanged(offset, limit, iLink_uri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_uriAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_uri := self.Args("link_uri").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_uriAndDefaultLangcode(offset, limit, iLink_uri,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_uriAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iLink_options := self.Args("link_options").Bytes()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndLink_options(offset, limit, iLink_title,iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExternal(offset, limit, iLink_title,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndRediscover(offset, limit, iLink_title,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndWeight(offset, limit, iLink_title,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndExpanded(offset, limit, iLink_title,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndEnabled(offset, limit, iLink_title,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndParent(offset, limit, iLink_title,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndChanged(offset, limit, iLink_title,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_titleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_title := self.Args("link_title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_titleAndDefaultLangcode(offset, limit, iLink_title,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_titleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExternal := self.Args("external").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExternal(offset, limit, iLink_options,iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndRediscover(offset, limit, iLink_options,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndWeight(offset, limit, iLink_options,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndExpanded(offset, limit, iLink_options,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndEnabled(offset, limit, iLink_options,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iParent := self.Args("parent").String()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndParent(offset, limit, iLink_options,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndChanged(offset, limit, iLink_options,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByLink_optionsAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink_options := self.Args("link_options").Bytes()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByLink_optionsAndDefaultLangcode(offset, limit, iLink_options,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByLink_optionsAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iRediscover := self.Args("rediscover").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndRediscover(offset, limit, iExternal,iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndWeight(offset, limit, iExternal,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndExpanded(offset, limit, iExternal,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndEnabled(offset, limit, iExternal,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndParent(offset, limit, iExternal,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndChanged(offset, limit, iExternal,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExternalAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExternal := self.Args("external").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExternalAndDefaultLangcode(offset, limit, iExternal,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExternalAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iWeight := self.Args("weight").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndWeight(offset, limit, iRediscover,iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndExpanded(offset, limit, iRediscover,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndEnabled(offset, limit, iRediscover,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndParent(offset, limit, iRediscover,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndChanged(offset, limit, iRediscover,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByRediscoverAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRediscover := self.Args("rediscover").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByRediscoverAndDefaultLangcode(offset, limit, iRediscover,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByRediscoverAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iExpanded := self.Args("expanded").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndExpanded(offset, limit, iWeight,iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndEnabled(offset, limit, iWeight,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndParent(offset, limit, iWeight,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndChanged(offset, limit, iWeight,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByWeightAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWeight := self.Args("weight").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByWeightAndDefaultLangcode(offset, limit, iWeight,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByWeightAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iEnabled := self.Args("enabled").MustInt()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndEnabled(offset, limit, iExpanded,iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndParent(offset, limit, iExpanded,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndChanged(offset, limit, iExpanded,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByExpandedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpanded := self.Args("expanded").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByExpandedAndDefaultLangcode(offset, limit, iExpanded,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByExpandedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByEnabledAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iParent := self.Args("parent").String()

	if helper.IsHas(iEnabled) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByEnabledAndParent(offset, limit, iEnabled,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByEnabledAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByEnabledAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEnabled) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByEnabledAndChanged(offset, limit, iEnabled,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByEnabledAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByEnabledAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEnabled := self.Args("enabled").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEnabled) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByEnabledAndDefaultLangcode(offset, limit, iEnabled,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByEnabledAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByParentAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iParent := self.Args("parent").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iParent) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByParentAndChanged(offset, limit, iParent,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByParentAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByParentAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iParent := self.Args("parent").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iParent) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByParentAndDefaultLangcode(offset, limit, iParent,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByParentAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasByChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatasByChangedAndDefaultLangcode(offset, limit, iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatasByChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaId(iId)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaTitle(iTitle)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription := self.Args("description").String()
	if helper.IsHas(iDescription) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaDescription(iDescription)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMenuName := self.Args("menu_name").String()
	if helper.IsHas(iMenuName) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaMenuName(iMenuName)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_uri := self.Args("link__uri").String()
	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaLink_uri(iLink_uri)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_title := self.Args("link__title").String()
	if helper.IsHas(iLink_title) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaLink_title(iLink_title)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_options := self.Args("link__options").Bytes()
	if helper.IsHas(iLink_options) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaLink_options(iLink_options)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExternal := self.Args("external").MustInt()
	if helper.IsHas(iExternal) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaExternal(iExternal)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRediscover := self.Args("rediscover").MustInt()
	if helper.IsHas(iRediscover) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaRediscover(iRediscover)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWeight := self.Args("weight").MustInt()
	if helper.IsHas(iWeight) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaWeight(iWeight)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpanded := self.Args("expanded").MustInt()
	if helper.IsHas(iExpanded) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaExpanded(iExpanded)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEnabled := self.Args("enabled").MustInt()
	if helper.IsHas(iEnabled) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaEnabled(iEnabled)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iParent := self.Args("parent").String()
	if helper.IsHas(iParent) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaParent(iParent)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasMenuLinkContentDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_MenuLinkContentData := model.HasMenuLinkContentDataViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["menu_link_content_data"] = _MenuLinkContentData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasMenuLinkContentDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaTitle(iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription := self.Args("description").String()
	if helper.IsHas(iDescription) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaDescription(iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMenuName := self.Args("menu_name").String()
	if helper.IsHas(iMenuName) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaMenuName(iMenuName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_uri := self.Args("link__uri").String()
	if helper.IsHas(iLink_uri) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaLink_uri(iLink_uri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_title := self.Args("link__title").String()
	if helper.IsHas(iLink_title) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaLink_title(iLink_title)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink_options := self.Args("link__options").Bytes()
	if helper.IsHas(iLink_options) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaLink_options(iLink_options)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExternal := self.Args("external").MustInt()
	if helper.IsHas(iExternal) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaExternal(iExternal)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRediscover := self.Args("rediscover").MustInt()
	if helper.IsHas(iRediscover) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaRediscover(iRediscover)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWeight := self.Args("weight").MustInt()
	if helper.IsHas(iWeight) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaWeight(iWeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpanded := self.Args("expanded").MustInt()
	if helper.IsHas(iExpanded) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaExpanded(iExpanded)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEnabled := self.Args("enabled").MustInt()
	if helper.IsHas(iEnabled) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaEnabled(iEnabled)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iParent := self.Args("parent").String()
	if helper.IsHas(iParent) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaParent(iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetMenuLinkContentDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_MenuLinkContentData, _error := model.GetMenuLinkContentDataViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the GetMenuLinkContentDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaId(Id_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaBundle(Bundle_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaLangcode(Langcode_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	if helper.IsHas(Title_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaTitle(Title_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	if helper.IsHas(Description_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaDescription(Description_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MenuName_ := self.Args("menu_name").String()
	if helper.IsHas(MenuName_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaMenuName(MenuName_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaMenuName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_uri_ := self.Args("link__uri").String()
	if helper.IsHas(Link_uri_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaLink_uri(Link_uri_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaLink_uri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_title_ := self.Args("link__title").String()
	if helper.IsHas(Link_title_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaLink_title(Link_title_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaLink_title's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_options_ := self.Args("link__options").Bytes()
	if helper.IsHas(Link_options_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaLink_options(Link_options_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaLink_options's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	External_ := self.Args("external").MustInt()
	if helper.IsHas(External_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaExternal(External_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaExternal's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Rediscover_ := self.Args("rediscover").MustInt()
	if helper.IsHas(Rediscover_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaRediscover(Rediscover_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaRediscover's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	if helper.IsHas(Weight_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaWeight(Weight_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaWeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expanded_ := self.Args("expanded").MustInt()
	if helper.IsHas(Expanded_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaExpanded(Expanded_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaExpanded's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Enabled_ := self.Args("enabled").MustInt()
	if helper.IsHas(Enabled_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaEnabled(Enabled_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaEnabled's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").String()
	if helper.IsHas(Parent_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaParent(Parent_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaChanged(Changed_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetMenuLinkContentDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iMenuLinkContentData model.MenuLinkContentData
		self.Bind(&iMenuLinkContentData)
		_MenuLinkContentData, _error := model.SetMenuLinkContentDataViaDefaultLangcode(DefaultLangcode_, &iMenuLinkContentData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_MenuLinkContentData)
	}
	herr.Message = "Can't get to the SetMenuLinkContentDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddMenuLinkContentDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	Bundle_ := self.Args("bundle").String()
	Langcode_ := self.Args("langcode").String()
	Title_ := self.Args("title").String()
	Description_ := self.Args("description").String()
	MenuName_ := self.Args("menu_name").String()
	Link_uri_ := self.Args("link__uri").String()
	Link_title_ := self.Args("link__title").String()
	Link_options_ := self.Args("link__options").Bytes()
	External_ := self.Args("external").MustInt()
	Rediscover_ := self.Args("rediscover").MustInt()
	Weight_ := self.Args("weight").MustInt()
	Expanded_ := self.Args("expanded").MustInt()
	Enabled_ := self.Args("enabled").MustInt()
	Parent_ := self.Args("parent").String()
	Changed_ := self.Args("changed").MustInt()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Id_) {
		_error := model.AddMenuLinkContentData(Id_,Bundle_,Langcode_,Title_,Description_,MenuName_,Link_uri_,Link_title_,Link_options_,External_,Rediscover_,Weight_,Expanded_,Enabled_,Parent_,Changed_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddMenuLinkContentData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostMenuLinkContentDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PostMenuLinkContentData(&iMenuLinkContentData)
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

func PutMenuLinkContentDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentData(&iMenuLinkContentData)
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

func PutMenuLinkContentDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaId(Id_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaBundle(Bundle_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaLangcode(Langcode_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaTitle(Title_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaDescription(Description_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MenuName_ := self.Args("menu_name").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaMenuName(MenuName_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_uri_ := self.Args("link__uri").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaLink_uri(Link_uri_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_title_ := self.Args("link__title").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaLink_title(Link_title_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_options_ := self.Args("link__options").Bytes()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaLink_options(Link_options_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	External_ := self.Args("external").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaExternal(External_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Rediscover_ := self.Args("rediscover").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaRediscover(Rediscover_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaWeight(Weight_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expanded_ := self.Args("expanded").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaExpanded(Expanded_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Enabled_ := self.Args("enabled").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaEnabled(Enabled_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaParent(Parent_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaChanged(Changed_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutMenuLinkContentDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	_int64, _error := model.PutMenuLinkContentDataViaDefaultLangcode(DefaultLangcode_, &iMenuLinkContentData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaTitle(Title_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaDescription(Description_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MenuName_ := self.Args("menu_name").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaMenuName(MenuName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_uri_ := self.Args("link__uri").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaLink_uri(Link_uri_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_title_ := self.Args("link__title").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaLink_title(Link_title_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_options_ := self.Args("link__options").Bytes()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaLink_options(Link_options_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	External_ := self.Args("external").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaExternal(External_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Rediscover_ := self.Args("rediscover").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaRediscover(Rediscover_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaWeight(Weight_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expanded_ := self.Args("expanded").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaExpanded(Expanded_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Enabled_ := self.Args("enabled").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaEnabled(Enabled_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").String()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaParent(Parent_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateMenuLinkContentDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iMenuLinkContentData model.MenuLinkContentData
	self.Bind(&iMenuLinkContentData)
	var iMap = helper.StructToMap(iMenuLinkContentData)
	_error := model.UpdateMenuLinkContentDataViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteMenuLinkContentData(Id_)
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

func DeleteMenuLinkContentDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteMenuLinkContentDataViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteMenuLinkContentDataViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteMenuLinkContentDataViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	_error := model.DeleteMenuLinkContentDataViaTitle(Title_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	_error := model.DeleteMenuLinkContentDataViaDescription(Description_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaMenuNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MenuName_ := self.Args("menu_name").String()
	_error := model.DeleteMenuLinkContentDataViaMenuName(MenuName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaLink_uriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_uri_ := self.Args("link__uri").String()
	_error := model.DeleteMenuLinkContentDataViaLink_uri(Link_uri_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaLink_titleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_title_ := self.Args("link__title").String()
	_error := model.DeleteMenuLinkContentDataViaLink_title(Link_title_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaLink_optionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_options_ := self.Args("link__options").Bytes()
	_error := model.DeleteMenuLinkContentDataViaLink_options(Link_options_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaExternalHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	External_ := self.Args("external").MustInt()
	_error := model.DeleteMenuLinkContentDataViaExternal(External_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaRediscoverHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Rediscover_ := self.Args("rediscover").MustInt()
	_error := model.DeleteMenuLinkContentDataViaRediscover(Rediscover_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaWeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Weight_ := self.Args("weight").MustInt()
	_error := model.DeleteMenuLinkContentDataViaWeight(Weight_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaExpandedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expanded_ := self.Args("expanded").MustInt()
	_error := model.DeleteMenuLinkContentDataViaExpanded(Expanded_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaEnabledHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Enabled_ := self.Args("enabled").MustInt()
	_error := model.DeleteMenuLinkContentDataViaEnabled(Enabled_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").String()
	_error := model.DeleteMenuLinkContentDataViaParent(Parent_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteMenuLinkContentDataViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteMenuLinkContentDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteMenuLinkContentDataViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
