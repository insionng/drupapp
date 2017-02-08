package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetNodeFieldDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodeFieldDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node_field_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataCountViaNidHandler(self *macross.Context) error {
	Nid_ := self.Args("nid").MustInt64()
	_int64 := model.GetNodeFieldDataCountViaNid(Nid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaVidHandler(self *macross.Context) error {
	Vid_ := self.Args("vid").MustInt()
	_int64 := model.GetNodeFieldDataCountViaVid(Vid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaTypeHandler(self *macross.Context) error {
	Type_ := self.Args("type").String()
	_int64 := model.GetNodeFieldDataCountViaType(Type_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeFieldDataCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaTitleHandler(self *macross.Context) error {
	Title_ := self.Args("title").String()
	_int64 := model.GetNodeFieldDataCountViaTitle(Title_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt()
	_int64 := model.GetNodeFieldDataCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaStatusHandler(self *macross.Context) error {
	Status_ := self.Args("status").MustInt()
	_int64 := model.GetNodeFieldDataCountViaStatus(Status_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").MustInt()
	_int64 := model.GetNodeFieldDataCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetNodeFieldDataCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaPromoteHandler(self *macross.Context) error {
	Promote_ := self.Args("promote").MustInt()
	_int64 := model.GetNodeFieldDataCountViaPromote(Promote_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaStickyHandler(self *macross.Context) error {
	Sticky_ := self.Args("sticky").MustInt()
	_int64 := model.GetNodeFieldDataCountViaSticky(Sticky_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	_int64 := model.GetNodeFieldDataCountViaRevisionTranslationAffected(RevisionTranslationAffected_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDataCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetNodeFieldDataCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_dataCount"] = 0
	}
	m["node_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldDatasViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iNid := self.Args("nid").MustInt64()
	if (offset > 0) && helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaNid(offset, limit, iNid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iVid := self.Args("vid").MustInt()
	if (offset > 0) && helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaVid(offset, limit, iVid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iType := self.Args("type").String()
	if (offset > 0) && helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaType(offset, limit, iType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTitle := self.Args("title").String()
	if (offset > 0) && helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaTitle(offset, limit, iTitle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt()
	if (offset > 0) && helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iStatus := self.Args("status").MustInt()
	if (offset > 0) && helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaStatus(offset, limit, iStatus, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").MustInt()
	if (offset > 0) && helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPromote := self.Args("promote").MustInt()
	if (offset > 0) && helper.IsHas(iPromote) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaPromote(offset, limit, iPromote, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSticky := self.Args("sticky").MustInt()
	if (offset > 0) && helper.IsHas(iSticky) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaSticky(offset, limit, iSticky, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionTranslationAffected) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaRevisionTranslationAffected(offset, limit, iRevisionTranslationAffected, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndType(offset, limit, iNid,iVid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndLangcode(offset, limit, iNid,iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndTitle(offset, limit, iNid,iVid,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndUid(offset, limit, iNid,iVid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndStatus(offset, limit, iNid,iVid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndCreated(offset, limit, iNid,iVid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndChanged(offset, limit, iNid,iVid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndPromote(offset, limit, iNid,iVid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndSticky(offset, limit, iNid,iVid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndRevisionTranslationAffected(offset, limit, iNid,iVid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVidAndDefaultLangcode(offset, limit, iNid,iVid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndLangcode(offset, limit, iNid,iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndTitle(offset, limit, iNid,iType,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndUid(offset, limit, iNid,iType,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndStatus(offset, limit, iNid,iType,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndCreated(offset, limit, iNid,iType,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndChanged(offset, limit, iNid,iType,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndPromote(offset, limit, iNid,iType,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndSticky(offset, limit, iNid,iType,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndRevisionTranslationAffected(offset, limit, iNid,iType,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTypeAndDefaultLangcode(offset, limit, iNid,iType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndTitle(offset, limit, iNid,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndUid(offset, limit, iNid,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndStatus(offset, limit, iNid,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndCreated(offset, limit, iNid,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndChanged(offset, limit, iNid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndPromote(offset, limit, iNid,iLangcode,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndSticky(offset, limit, iNid,iLangcode,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndRevisionTranslationAffected(offset, limit, iNid,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcodeAndDefaultLangcode(offset, limit, iNid,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitleAndUid(offset, limit, iNid,iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitleAndStatus(offset, limit, iNid,iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitleAndCreated(offset, limit, iNid,iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitleAndChanged(offset, limit, iNid,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitleAndPromote(offset, limit, iNid,iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitleAndSticky(offset, limit, iNid,iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitleAndRevisionTranslationAffected(offset, limit, iNid,iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitleAndDefaultLangcode(offset, limit, iNid,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndUidAndStatus(offset, limit, iNid,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndUidAndCreated(offset, limit, iNid,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndUidAndChanged(offset, limit, iNid,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndUidAndPromote(offset, limit, iNid,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndUidAndSticky(offset, limit, iNid,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndUidAndRevisionTranslationAffected(offset, limit, iNid,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndUidAndDefaultLangcode(offset, limit, iNid,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStatusAndCreated(offset, limit, iNid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStatusAndChanged(offset, limit, iNid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStatusAndPromote(offset, limit, iNid,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStatusAndSticky(offset, limit, iNid,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStatusAndRevisionTranslationAffected(offset, limit, iNid,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStatusAndDefaultLangcode(offset, limit, iNid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndCreatedAndChanged(offset, limit, iNid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndCreatedAndPromote(offset, limit, iNid,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndCreatedAndSticky(offset, limit, iNid,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndCreatedAndRevisionTranslationAffected(offset, limit, iNid,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndCreatedAndDefaultLangcode(offset, limit, iNid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndChangedAndPromote(offset, limit, iNid,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndChangedAndSticky(offset, limit, iNid,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndChangedAndRevisionTranslationAffected(offset, limit, iNid,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndChangedAndDefaultLangcode(offset, limit, iNid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndPromoteAndSticky(offset, limit, iNid,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndPromoteAndRevisionTranslationAffected(offset, limit, iNid,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndPromoteAndDefaultLangcode(offset, limit, iNid,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStickyAndRevisionTranslationAffected(offset, limit, iNid,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStickyAndDefaultLangcode(offset, limit, iNid,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iNid,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndLangcode(offset, limit, iVid,iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndTitle(offset, limit, iVid,iType,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndUid(offset, limit, iVid,iType,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndStatus(offset, limit, iVid,iType,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndCreated(offset, limit, iVid,iType,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndChanged(offset, limit, iVid,iType,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndPromote(offset, limit, iVid,iType,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndSticky(offset, limit, iVid,iType,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndRevisionTranslationAffected(offset, limit, iVid,iType,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTypeAndDefaultLangcode(offset, limit, iVid,iType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndTitle(offset, limit, iVid,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndUid(offset, limit, iVid,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndStatus(offset, limit, iVid,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndCreated(offset, limit, iVid,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndChanged(offset, limit, iVid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndPromote(offset, limit, iVid,iLangcode,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndSticky(offset, limit, iVid,iLangcode,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndRevisionTranslationAffected(offset, limit, iVid,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcodeAndDefaultLangcode(offset, limit, iVid,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitleAndUid(offset, limit, iVid,iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitleAndStatus(offset, limit, iVid,iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitleAndCreated(offset, limit, iVid,iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitleAndChanged(offset, limit, iVid,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitleAndPromote(offset, limit, iVid,iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitleAndSticky(offset, limit, iVid,iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitleAndRevisionTranslationAffected(offset, limit, iVid,iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitleAndDefaultLangcode(offset, limit, iVid,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndUidAndStatus(offset, limit, iVid,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndUidAndCreated(offset, limit, iVid,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndUidAndChanged(offset, limit, iVid,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndUidAndPromote(offset, limit, iVid,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndUidAndSticky(offset, limit, iVid,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndUidAndRevisionTranslationAffected(offset, limit, iVid,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndUidAndDefaultLangcode(offset, limit, iVid,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStatusAndCreated(offset, limit, iVid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStatusAndChanged(offset, limit, iVid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStatusAndPromote(offset, limit, iVid,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStatusAndSticky(offset, limit, iVid,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStatusAndRevisionTranslationAffected(offset, limit, iVid,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStatusAndDefaultLangcode(offset, limit, iVid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndCreatedAndChanged(offset, limit, iVid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndCreatedAndPromote(offset, limit, iVid,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndCreatedAndSticky(offset, limit, iVid,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndCreatedAndRevisionTranslationAffected(offset, limit, iVid,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndCreatedAndDefaultLangcode(offset, limit, iVid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndChangedAndPromote(offset, limit, iVid,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndChangedAndSticky(offset, limit, iVid,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndChangedAndRevisionTranslationAffected(offset, limit, iVid,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndChangedAndDefaultLangcode(offset, limit, iVid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndPromoteAndSticky(offset, limit, iVid,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndPromoteAndRevisionTranslationAffected(offset, limit, iVid,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndPromoteAndDefaultLangcode(offset, limit, iVid,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStickyAndRevisionTranslationAffected(offset, limit, iVid,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStickyAndDefaultLangcode(offset, limit, iVid,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iVid,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndTitle(offset, limit, iType,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndUid(offset, limit, iType,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndStatus(offset, limit, iType,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndCreated(offset, limit, iType,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndChanged(offset, limit, iType,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndPromote(offset, limit, iType,iLangcode,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndSticky(offset, limit, iType,iLangcode,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndRevisionTranslationAffected(offset, limit, iType,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcodeAndDefaultLangcode(offset, limit, iType,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitleAndUid(offset, limit, iType,iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitleAndStatus(offset, limit, iType,iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitleAndCreated(offset, limit, iType,iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitleAndChanged(offset, limit, iType,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitleAndPromote(offset, limit, iType,iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitleAndSticky(offset, limit, iType,iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitleAndRevisionTranslationAffected(offset, limit, iType,iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitleAndDefaultLangcode(offset, limit, iType,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndUidAndStatus(offset, limit, iType,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndUidAndCreated(offset, limit, iType,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndUidAndChanged(offset, limit, iType,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndUidAndPromote(offset, limit, iType,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndUidAndSticky(offset, limit, iType,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndUidAndRevisionTranslationAffected(offset, limit, iType,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndUidAndDefaultLangcode(offset, limit, iType,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStatusAndCreated(offset, limit, iType,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStatusAndChanged(offset, limit, iType,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStatusAndPromote(offset, limit, iType,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStatusAndSticky(offset, limit, iType,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStatusAndRevisionTranslationAffected(offset, limit, iType,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStatusAndDefaultLangcode(offset, limit, iType,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndCreatedAndChanged(offset, limit, iType,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndCreatedAndPromote(offset, limit, iType,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndCreatedAndSticky(offset, limit, iType,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndCreatedAndRevisionTranslationAffected(offset, limit, iType,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndCreatedAndDefaultLangcode(offset, limit, iType,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndChangedAndPromote(offset, limit, iType,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndChangedAndSticky(offset, limit, iType,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndChangedAndRevisionTranslationAffected(offset, limit, iType,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndChangedAndDefaultLangcode(offset, limit, iType,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndPromoteAndSticky(offset, limit, iType,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndPromoteAndRevisionTranslationAffected(offset, limit, iType,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndPromoteAndDefaultLangcode(offset, limit, iType,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStickyAndRevisionTranslationAffected(offset, limit, iType,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStickyAndDefaultLangcode(offset, limit, iType,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iType,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitleAndUid(offset, limit, iLangcode,iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitleAndStatus(offset, limit, iLangcode,iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitleAndCreated(offset, limit, iLangcode,iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitleAndChanged(offset, limit, iLangcode,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitleAndPromote(offset, limit, iLangcode,iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitleAndSticky(offset, limit, iLangcode,iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitleAndRevisionTranslationAffected(offset, limit, iLangcode,iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitleAndDefaultLangcode(offset, limit, iLangcode,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndUidAndStatus(offset, limit, iLangcode,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndUidAndCreated(offset, limit, iLangcode,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndUidAndChanged(offset, limit, iLangcode,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndUidAndPromote(offset, limit, iLangcode,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndUidAndSticky(offset, limit, iLangcode,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndUidAndRevisionTranslationAffected(offset, limit, iLangcode,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndUidAndDefaultLangcode(offset, limit, iLangcode,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStatusAndCreated(offset, limit, iLangcode,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStatusAndChanged(offset, limit, iLangcode,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStatusAndPromote(offset, limit, iLangcode,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStatusAndSticky(offset, limit, iLangcode,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStatusAndRevisionTranslationAffected(offset, limit, iLangcode,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStatusAndDefaultLangcode(offset, limit, iLangcode,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndCreatedAndChanged(offset, limit, iLangcode,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndCreatedAndPromote(offset, limit, iLangcode,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndCreatedAndSticky(offset, limit, iLangcode,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndCreatedAndRevisionTranslationAffected(offset, limit, iLangcode,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndCreatedAndDefaultLangcode(offset, limit, iLangcode,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndChangedAndPromote(offset, limit, iLangcode,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndChangedAndSticky(offset, limit, iLangcode,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndChangedAndRevisionTranslationAffected(offset, limit, iLangcode,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset, limit, iLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndPromoteAndSticky(offset, limit, iLangcode,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndPromoteAndRevisionTranslationAffected(offset, limit, iLangcode,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndPromoteAndDefaultLangcode(offset, limit, iLangcode,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStickyAndRevisionTranslationAffected(offset, limit, iLangcode,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStickyAndDefaultLangcode(offset, limit, iLangcode,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iLangcode,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndUidAndStatus(offset, limit, iTitle,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndUidAndCreated(offset, limit, iTitle,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndUidAndChanged(offset, limit, iTitle,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndUidAndPromote(offset, limit, iTitle,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndUidAndSticky(offset, limit, iTitle,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndUidAndRevisionTranslationAffected(offset, limit, iTitle,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndUidAndDefaultLangcode(offset, limit, iTitle,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStatusAndCreated(offset, limit, iTitle,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStatusAndChanged(offset, limit, iTitle,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStatusAndPromote(offset, limit, iTitle,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStatusAndSticky(offset, limit, iTitle,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStatusAndRevisionTranslationAffected(offset, limit, iTitle,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStatusAndDefaultLangcode(offset, limit, iTitle,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndCreatedAndChanged(offset, limit, iTitle,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndCreatedAndPromote(offset, limit, iTitle,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndCreatedAndSticky(offset, limit, iTitle,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndCreatedAndRevisionTranslationAffected(offset, limit, iTitle,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndCreatedAndDefaultLangcode(offset, limit, iTitle,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndChangedAndPromote(offset, limit, iTitle,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndChangedAndSticky(offset, limit, iTitle,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndChangedAndRevisionTranslationAffected(offset, limit, iTitle,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndChangedAndDefaultLangcode(offset, limit, iTitle,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndPromoteAndSticky(offset, limit, iTitle,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndPromoteAndRevisionTranslationAffected(offset, limit, iTitle,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndPromoteAndDefaultLangcode(offset, limit, iTitle,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStickyAndRevisionTranslationAffected(offset, limit, iTitle,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStickyAndDefaultLangcode(offset, limit, iTitle,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iTitle,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStatusAndCreated(offset, limit, iUid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStatusAndChanged(offset, limit, iUid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStatusAndPromote(offset, limit, iUid,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStatusAndSticky(offset, limit, iUid,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStatusAndRevisionTranslationAffected(offset, limit, iUid,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStatusAndDefaultLangcode(offset, limit, iUid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndCreatedAndChanged(offset, limit, iUid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndCreatedAndPromote(offset, limit, iUid,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndCreatedAndSticky(offset, limit, iUid,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndCreatedAndRevisionTranslationAffected(offset, limit, iUid,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndCreatedAndDefaultLangcode(offset, limit, iUid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndChangedAndPromote(offset, limit, iUid,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndChangedAndSticky(offset, limit, iUid,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndChangedAndRevisionTranslationAffected(offset, limit, iUid,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndChangedAndDefaultLangcode(offset, limit, iUid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndPromoteAndSticky(offset, limit, iUid,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndPromoteAndRevisionTranslationAffected(offset, limit, iUid,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndPromoteAndDefaultLangcode(offset, limit, iUid,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStickyAndRevisionTranslationAffected(offset, limit, iUid,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStickyAndDefaultLangcode(offset, limit, iUid,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iUid,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndCreatedAndChanged(offset, limit, iStatus,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndCreatedAndPromote(offset, limit, iStatus,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndCreatedAndSticky(offset, limit, iStatus,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndCreatedAndRevisionTranslationAffected(offset, limit, iStatus,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndCreatedAndDefaultLangcode(offset, limit, iStatus,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndChangedAndPromote(offset, limit, iStatus,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndChangedAndSticky(offset, limit, iStatus,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndChangedAndRevisionTranslationAffected(offset, limit, iStatus,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndChangedAndDefaultLangcode(offset, limit, iStatus,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndPromoteAndSticky(offset, limit, iStatus,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndPromoteAndRevisionTranslationAffected(offset, limit, iStatus,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndPromoteAndDefaultLangcode(offset, limit, iStatus,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndStickyAndRevisionTranslationAffected(offset, limit, iStatus,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndStickyAndDefaultLangcode(offset, limit, iStatus,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iStatus,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndChangedAndPromote(offset, limit, iCreated,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndChangedAndSticky(offset, limit, iCreated,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndChangedAndRevisionTranslationAffected(offset, limit, iCreated,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndChangedAndDefaultLangcode(offset, limit, iCreated,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndPromoteAndSticky(offset, limit, iCreated,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndPromoteAndRevisionTranslationAffected(offset, limit, iCreated,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndPromoteAndDefaultLangcode(offset, limit, iCreated,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndStickyAndRevisionTranslationAffected(offset, limit, iCreated,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndStickyAndDefaultLangcode(offset, limit, iCreated,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iCreated,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndPromoteAndSticky(offset, limit, iChanged,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndPromoteAndRevisionTranslationAffected(offset, limit, iChanged,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndPromoteAndDefaultLangcode(offset, limit, iChanged,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndStickyAndRevisionTranslationAffected(offset, limit, iChanged,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndStickyAndDefaultLangcode(offset, limit, iChanged,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iChanged,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByPromoteAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByPromoteAndStickyAndRevisionTranslationAffected(offset, limit, iPromote,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByPromoteAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByPromoteAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByPromoteAndStickyAndDefaultLangcode(offset, limit, iPromote,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByPromoteAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByPromoteAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByPromoteAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iPromote,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByPromoteAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStickyAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSticky) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStickyAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iSticky,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStickyAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndVid(offset, limit, iNid,iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndType(offset, limit, iNid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndLangcode(offset, limit, iNid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndTitle(offset, limit, iNid,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndUid(offset, limit, iNid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndStatus(offset, limit, iNid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndCreated(offset, limit, iNid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndChanged(offset, limit, iNid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndPromote(offset, limit, iNid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndSticky(offset, limit, iNid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndRevisionTranslationAffected(offset, limit, iNid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByNidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByNidAndDefaultLangcode(offset, limit, iNid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByNidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndType(offset, limit, iVid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndLangcode(offset, limit, iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndTitle(offset, limit, iVid,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndUid(offset, limit, iVid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndStatus(offset, limit, iVid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndCreated(offset, limit, iVid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndChanged(offset, limit, iVid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndPromote(offset, limit, iVid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndSticky(offset, limit, iVid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndRevisionTranslationAffected(offset, limit, iVid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByVidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByVidAndDefaultLangcode(offset, limit, iVid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByVidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndLangcode(offset, limit, iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndTitle(offset, limit, iType,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndUid(offset, limit, iType,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndStatus(offset, limit, iType,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndCreated(offset, limit, iType,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndChanged(offset, limit, iType,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndPromote(offset, limit, iType,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndSticky(offset, limit, iType,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndRevisionTranslationAffected(offset, limit, iType,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTypeAndDefaultLangcode(offset, limit, iType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndTitle(offset, limit, iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndUid(offset, limit, iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndStatus(offset, limit, iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndCreated(offset, limit, iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndPromote(offset, limit, iLangcode,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndSticky(offset, limit, iLangcode,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndRevisionTranslationAffected(offset, limit, iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndUid(offset, limit, iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndStatus(offset, limit, iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndCreated(offset, limit, iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndChanged(offset, limit, iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndPromote(offset, limit, iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndSticky(offset, limit, iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndRevisionTranslationAffected(offset, limit, iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByTitleAndDefaultLangcode(offset, limit, iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndStatus(offset, limit, iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndCreated(offset, limit, iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndChanged(offset, limit, iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndPromote(offset, limit, iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndSticky(offset, limit, iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndRevisionTranslationAffected(offset, limit, iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByUidAndDefaultLangcode(offset, limit, iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndCreated(offset, limit, iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndChanged(offset, limit, iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndPromote(offset, limit, iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndSticky(offset, limit, iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndRevisionTranslationAffected(offset, limit, iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStatusAndDefaultLangcode(offset, limit, iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndChanged(offset, limit, iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndPromote(offset, limit, iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndSticky(offset, limit, iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndRevisionTranslationAffected(offset, limit, iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByCreatedAndDefaultLangcode(offset, limit, iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndPromote(offset, limit, iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndSticky(offset, limit, iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndRevisionTranslationAffected(offset, limit, iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByChangedAndDefaultLangcode(offset, limit, iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByPromoteAndSticky(offset, limit, iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByPromoteAndRevisionTranslationAffected(offset, limit, iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByPromoteAndDefaultLangcode(offset, limit, iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iSticky) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStickyAndRevisionTranslationAffected(offset, limit, iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSticky) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByStickyAndDefaultLangcode(offset, limit, iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasByRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionTranslationAffected) {
		_NodeFieldData, _error := model.GetNodeFieldDatasByRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatasByRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_NodeFieldData, _error := model.GetNodeFieldDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_NodeFieldData := model.HasNodeFieldDataViaNid(iNid)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").MustInt()
	if helper.IsHas(iVid) {
		_NodeFieldData := model.HasNodeFieldDataViaVid(iVid)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_NodeFieldData := model.HasNodeFieldDataViaType(iType)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeFieldData := model.HasNodeFieldDataViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_NodeFieldData := model.HasNodeFieldDataViaTitle(iTitle)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_NodeFieldData := model.HasNodeFieldDataViaUid(iUid)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_NodeFieldData := model.HasNodeFieldDataViaStatus(iStatus)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_NodeFieldData := model.HasNodeFieldDataViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_NodeFieldData := model.HasNodeFieldDataViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPromote := self.Args("promote").MustInt()
	if helper.IsHas(iPromote) {
		_NodeFieldData := model.HasNodeFieldDataViaPromote(iPromote)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSticky := self.Args("sticky").MustInt()
	if helper.IsHas(iSticky) {
		_NodeFieldData := model.HasNodeFieldDataViaSticky(iSticky)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(iRevisionTranslationAffected) {
		_NodeFieldData := model.HasNodeFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_NodeFieldData := model.HasNodeFieldDataViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["node_field_data"] = _NodeFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaNid(iNid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").MustInt()
	if helper.IsHas(iVid) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaVid(iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaType(iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaTitle(iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaStatus(iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPromote := self.Args("promote").MustInt()
	if helper.IsHas(iPromote) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaPromote(iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSticky := self.Args("sticky").MustInt()
	if helper.IsHas(iSticky) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaSticky(iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(iRevisionTranslationAffected) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_NodeFieldData, _error := model.GetNodeFieldDataViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the GetNodeFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	if helper.IsHas(Nid_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaNid(Nid_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	if helper.IsHas(Vid_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaVid(Vid_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	if helper.IsHas(Type_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaType(Type_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaLangcode(Langcode_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	if helper.IsHas(Title_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaTitle(Title_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	if helper.IsHas(Uid_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaUid(Uid_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	if helper.IsHas(Status_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaStatus(Status_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	if helper.IsHas(Created_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaCreated(Created_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaChanged(Changed_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Promote_ := self.Args("promote").MustInt()
	if helper.IsHas(Promote_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaPromote(Promote_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	if helper.IsHas(Sticky_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaSticky(Sticky_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(RevisionTranslationAffected_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iNodeFieldData model.NodeFieldData
		self.Bind(&iNodeFieldData)
		_NodeFieldData, _error := model.SetNodeFieldDataViaDefaultLangcode(DefaultLangcode_, &iNodeFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldData)
	}
	herr.Message = "Can't get to the SetNodeFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	Vid_ := self.Args("vid").MustInt()
	Type_ := self.Args("type").String()
	Langcode_ := self.Args("langcode").String()
	Title_ := self.Args("title").String()
	Uid_ := self.Args("uid").MustInt()
	Status_ := self.Args("status").MustInt()
	Created_ := self.Args("created").MustInt()
	Changed_ := self.Args("changed").MustInt()
	Promote_ := self.Args("promote").MustInt()
	Sticky_ := self.Args("sticky").MustInt()
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Nid_) {
		_error := model.AddNodeFieldData(Nid_,Vid_,Type_,Langcode_,Title_,Uid_,Status_,Created_,Changed_,Promote_,Sticky_,RevisionTranslationAffected_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNodeFieldData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PostNodeFieldData(&iNodeFieldData)
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

func PutNodeFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldData(&iNodeFieldData)
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

func PutNodeFieldDataViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaNid(Nid_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaVid(Vid_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaType(Type_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaLangcode(Langcode_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaTitle(Title_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaUid(Uid_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaStatus(Status_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaCreated(Created_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaChanged(Changed_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Promote_ := self.Args("promote").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaPromote(Promote_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaSticky(Sticky_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	_int64, _error := model.PutNodeFieldDataViaDefaultLangcode(DefaultLangcode_, &iNodeFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaNid(Nid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaVid(Vid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaType(Type_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaTitle(Title_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaStatus(Status_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Promote_ := self.Args("promote").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaPromote(Promote_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaSticky(Sticky_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iNodeFieldData model.NodeFieldData
	self.Bind(&iNodeFieldData)
	var iMap = helper.StructToMap(iNodeFieldData)
	_error := model.UpdateNodeFieldDataViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_int64, _error := model.DeleteNodeFieldData(Nid_)
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

func DeleteNodeFieldDataViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_error := model.DeleteNodeFieldDataViaNid(Nid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	_error := model.DeleteNodeFieldDataViaVid(Vid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	_error := model.DeleteNodeFieldDataViaType(Type_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeFieldDataViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	_error := model.DeleteNodeFieldDataViaTitle(Title_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	_error := model.DeleteNodeFieldDataViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	_error := model.DeleteNodeFieldDataViaStatus(Status_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	_error := model.DeleteNodeFieldDataViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteNodeFieldDataViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Promote_ := self.Args("promote").MustInt()
	_error := model.DeleteNodeFieldDataViaPromote(Promote_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	_error := model.DeleteNodeFieldDataViaSticky(Sticky_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	_error := model.DeleteNodeFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteNodeFieldDataViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
