package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetBlockContentFieldDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBlockContentFieldDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["block_content_field_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetBlockContentFieldDataCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetBlockContentFieldDataCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaTypeHandler(self *macross.Context) error {
	Type_ := self.Args("type").String()
	_int64 := model.GetBlockContentFieldDataCountViaType(Type_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetBlockContentFieldDataCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaInfoHandler(self *macross.Context) error {
	Info_ := self.Args("info").String()
	_int64 := model.GetBlockContentFieldDataCountViaInfo(Info_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetBlockContentFieldDataCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaRevisionCreatedHandler(self *macross.Context) error {
	RevisionCreated_ := self.Args("revision_created").MustInt()
	_int64 := model.GetBlockContentFieldDataCountViaRevisionCreated(RevisionCreated_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaRevisionUserHandler(self *macross.Context) error {
	RevisionUser_ := self.Args("revision_user").MustInt()
	_int64 := model.GetBlockContentFieldDataCountViaRevisionUser(RevisionUser_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	_int64 := model.GetBlockContentFieldDataCountViaRevisionTranslationAffected(RevisionTranslationAffected_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDataCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetBlockContentFieldDataCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_dataCount"] = 0
	}
	m["block_content_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldDatasViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iType := self.Args("type").String()
	if (offset > 0) && helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaType(offset, limit, iType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iInfo := self.Args("info").String()
	if (offset > 0) && helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaInfo(offset, limit, iInfo, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaRevisionCreated(offset, limit, iRevisionCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionUser) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaRevisionUser(offset, limit, iRevisionUser, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionTranslationAffected) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaRevisionTranslationAffected(offset, limit, iRevisionTranslationAffected, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionIdAndType(offset, limit, iId,iRevisionId,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionIdAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionIdAndLangcode(offset, limit, iId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionIdAndInfo(offset, limit, iId,iRevisionId,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionIdAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionIdAndChanged(offset, limit, iId,iRevisionId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionCreated(offset, limit, iId,iRevisionId,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionUser(offset, limit, iId,iRevisionId,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionTranslationAffected(offset, limit, iId,iRevisionId,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionIdAndDefaultLangcode(offset, limit, iId,iRevisionId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndTypeAndLangcode(offset, limit, iId,iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndTypeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndTypeAndInfo(offset, limit, iId,iType,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndTypeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndTypeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndTypeAndChanged(offset, limit, iId,iType,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndTypeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndTypeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndTypeAndRevisionCreated(offset, limit, iId,iType,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndTypeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndTypeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndTypeAndRevisionUser(offset, limit, iId,iType,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndTypeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndTypeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndTypeAndRevisionTranslationAffected(offset, limit, iId,iType,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndTypeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndTypeAndDefaultLangcode(offset, limit, iId,iType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndLangcodeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndLangcodeAndInfo(offset, limit, iId,iLangcode,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndLangcodeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndLangcodeAndChanged(offset, limit, iId,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndLangcodeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndLangcodeAndRevisionCreated(offset, limit, iId,iLangcode,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndLangcodeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndLangcodeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndLangcodeAndRevisionUser(offset, limit, iId,iLangcode,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndLangcodeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndLangcodeAndRevisionTranslationAffected(offset, limit, iId,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndLangcodeAndDefaultLangcode(offset, limit, iId,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndInfoAndChanged(offset, limit, iId,iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndInfoAndRevisionCreated(offset, limit, iId,iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndInfoAndRevisionUser(offset, limit, iId,iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndInfoAndRevisionTranslationAffected(offset, limit, iId,iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndInfoAndDefaultLangcode(offset, limit, iId,iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndChangedAndRevisionCreated(offset, limit, iId,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndChangedAndRevisionUser(offset, limit, iId,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndChangedAndRevisionTranslationAffected(offset, limit, iId,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndChangedAndDefaultLangcode(offset, limit, iId,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionUser(offset, limit, iId,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iId,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionCreatedAndDefaultLangcode(offset, limit, iId,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionUserAndRevisionTranslationAffected(offset, limit, iId,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionUserAndDefaultLangcode(offset, limit, iId,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iId,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndTypeAndLangcode(offset, limit, iRevisionId,iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndTypeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndTypeAndInfo(offset, limit, iRevisionId,iType,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndTypeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndTypeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndTypeAndChanged(offset, limit, iRevisionId,iType,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndTypeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionCreated(offset, limit, iRevisionId,iType,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionUser(offset, limit, iRevisionId,iType,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionTranslationAffected(offset, limit, iRevisionId,iType,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndTypeAndDefaultLangcode(offset, limit, iRevisionId,iType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndLangcodeAndInfo(offset, limit, iRevisionId,iLangcode,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndLangcodeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndLangcodeAndChanged(offset, limit, iRevisionId,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionCreated(offset, limit, iRevisionId,iLangcode,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionUser(offset, limit, iRevisionId,iLangcode,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionTranslationAffected(offset, limit, iRevisionId,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndLangcodeAndDefaultLangcode(offset, limit, iRevisionId,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndInfoAndChanged(offset, limit, iRevisionId,iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionCreated(offset, limit, iRevisionId,iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionUser(offset, limit, iRevisionId,iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionTranslationAffected(offset, limit, iRevisionId,iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndInfoAndDefaultLangcode(offset, limit, iRevisionId,iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionCreated(offset, limit, iRevisionId,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionUser(offset, limit, iRevisionId,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionTranslationAffected(offset, limit, iRevisionId,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndChangedAndDefaultLangcode(offset, limit, iRevisionId,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionUser(offset, limit, iRevisionId,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iRevisionId,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndDefaultLangcode(offset, limit, iRevisionId,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndRevisionTranslationAffected(offset, limit, iRevisionId,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndDefaultLangcode(offset, limit, iRevisionId,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionId,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndLangcodeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndLangcodeAndInfo(offset, limit, iType,iLangcode,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndLangcodeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndLangcodeAndChanged(offset, limit, iType,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionCreated(offset, limit, iType,iLangcode,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionUser(offset, limit, iType,iLangcode,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionTranslationAffected(offset, limit, iType,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndLangcodeAndDefaultLangcode(offset, limit, iType,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndInfoAndChanged(offset, limit, iType,iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndInfoAndRevisionCreated(offset, limit, iType,iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndInfoAndRevisionUser(offset, limit, iType,iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndInfoAndRevisionTranslationAffected(offset, limit, iType,iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndInfoAndDefaultLangcode(offset, limit, iType,iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndChangedAndRevisionCreated(offset, limit, iType,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndChangedAndRevisionUser(offset, limit, iType,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndChangedAndRevisionTranslationAffected(offset, limit, iType,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndChangedAndDefaultLangcode(offset, limit, iType,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionUser(offset, limit, iType,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iType,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionCreatedAndDefaultLangcode(offset, limit, iType,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionUserAndRevisionTranslationAffected(offset, limit, iType,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionUserAndDefaultLangcode(offset, limit, iType,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iType,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndInfoAndChanged(offset, limit, iLangcode,iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionCreated(offset, limit, iLangcode,iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionUser(offset, limit, iLangcode,iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionTranslationAffected(offset, limit, iLangcode,iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndInfoAndDefaultLangcode(offset, limit, iLangcode,iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionCreated(offset, limit, iLangcode,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionUser(offset, limit, iLangcode,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionTranslationAffected(offset, limit, iLangcode,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset, limit, iLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionUser(offset, limit, iLangcode,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iLangcode,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndDefaultLangcode(offset, limit, iLangcode,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionUserAndRevisionTranslationAffected(offset, limit, iLangcode,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionUserAndDefaultLangcode(offset, limit, iLangcode,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iLangcode,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndChangedAndRevisionCreated(offset, limit, iInfo,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndChangedAndRevisionUser(offset, limit, iInfo,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndChangedAndRevisionTranslationAffected(offset, limit, iInfo,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndChangedAndDefaultLangcode(offset, limit, iInfo,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionUser(offset, limit, iInfo,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iInfo,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionCreatedAndDefaultLangcode(offset, limit, iInfo,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionUserAndRevisionTranslationAffected(offset, limit, iInfo,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionUserAndDefaultLangcode(offset, limit, iInfo,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iInfo,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionUser(offset, limit, iChanged,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iChanged,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionCreatedAndDefaultLangcode(offset, limit, iChanged,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionUserAndRevisionTranslationAffected(offset, limit, iChanged,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionUserAndDefaultLangcode(offset, limit, iChanged,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iChanged,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndRevisionTranslationAffected(offset, limit, iRevisionCreated,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndDefaultLangcode(offset, limit, iRevisionCreated,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionCreated,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionUser,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionId(offset, limit, iId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndType(offset, limit, iId,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndLangcode(offset, limit, iId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndInfo(offset, limit, iId,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndChanged(offset, limit, iId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionCreated(offset, limit, iId,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionUser(offset, limit, iId,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndRevisionTranslationAffected(offset, limit, iId,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByIdAndDefaultLangcode(offset, limit, iId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndType(offset, limit, iRevisionId,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndInfo(offset, limit, iRevisionId,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndChanged(offset, limit, iRevisionId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionCreated(offset, limit, iRevisionId,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionUser(offset, limit, iRevisionId,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffected(offset, limit, iRevisionId,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionIdAndDefaultLangcode(offset, limit, iRevisionId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndLangcode(offset, limit, iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndInfo(offset, limit, iType,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndChanged(offset, limit, iType,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionCreated(offset, limit, iType,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionUser(offset, limit, iType,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndRevisionTranslationAffected(offset, limit, iType,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByTypeAndDefaultLangcode(offset, limit, iType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndInfo(offset, limit, iLangcode,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionCreated(offset, limit, iLangcode,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionUser(offset, limit, iLangcode,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffected(offset, limit, iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndChanged(offset, limit, iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionCreated(offset, limit, iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionUser(offset, limit, iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndRevisionTranslationAffected(offset, limit, iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByInfoAndDefaultLangcode(offset, limit, iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionCreated(offset, limit, iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionUser(offset, limit, iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndRevisionTranslationAffected(offset, limit, iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByChangedAndDefaultLangcode(offset, limit, iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionCreatedAndRevisionUser(offset, limit, iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffected(offset, limit, iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionCreatedAndDefaultLangcode(offset, limit, iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffected(offset, limit, iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionUserAndDefaultLangcode(offset, limit, iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasByRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionTranslationAffected) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatasByRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatasByRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaId(iId)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaType(iType)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iInfo := self.Args("info").String()
	if helper.IsHas(iInfo) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaInfo(iInfo)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionCreated := self.Args("revision_created").MustInt()
	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaRevisionCreated(iRevisionCreated)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionUser := self.Args("revision_user").MustInt()
	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaRevisionUser(iRevisionUser)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(iRevisionTranslationAffected) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_BlockContentFieldData := model.HasBlockContentFieldDataViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["block_content_field_data"] = _BlockContentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaType(iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iInfo := self.Args("info").String()
	if helper.IsHas(iInfo) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaInfo(iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionCreated := self.Args("revision_created").MustInt()
	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaRevisionCreated(iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionUser := self.Args("revision_user").MustInt()
	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaRevisionUser(iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(iRevisionTranslationAffected) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_BlockContentFieldData, _error := model.GetBlockContentFieldDataViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the GetBlockContentFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaId(Id_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaRevisionId(RevisionId_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	if helper.IsHas(Type_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaType(Type_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaLangcode(Langcode_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Info_ := self.Args("info").String()
	if helper.IsHas(Info_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaInfo(Info_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaChanged(Changed_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionCreated_ := self.Args("revision_created").MustInt()
	if helper.IsHas(RevisionCreated_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaRevisionCreated(RevisionCreated_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUser_ := self.Args("revision_user").MustInt()
	if helper.IsHas(RevisionUser_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaRevisionUser(RevisionUser_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(RevisionTranslationAffected_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iBlockContentFieldData model.BlockContentFieldData
		self.Bind(&iBlockContentFieldData)
		_BlockContentFieldData, _error := model.SetBlockContentFieldDataViaDefaultLangcode(DefaultLangcode_, &iBlockContentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldData)
	}
	herr.Message = "Can't get to the SetBlockContentFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBlockContentFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	RevisionId_ := self.Args("revision_id").MustInt()
	Type_ := self.Args("type").String()
	Langcode_ := self.Args("langcode").String()
	Info_ := self.Args("info").String()
	Changed_ := self.Args("changed").MustInt()
	RevisionCreated_ := self.Args("revision_created").MustInt()
	RevisionUser_ := self.Args("revision_user").MustInt()
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Id_) {
		_error := model.AddBlockContentFieldData(Id_,RevisionId_,Type_,Langcode_,Info_,Changed_,RevisionCreated_,RevisionUser_,RevisionTranslationAffected_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBlockContentFieldData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBlockContentFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PostBlockContentFieldData(&iBlockContentFieldData)
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

func PutBlockContentFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldData(&iBlockContentFieldData)
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

func PutBlockContentFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaId(Id_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaRevisionId(RevisionId_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaType(Type_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaLangcode(Langcode_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Info_ := self.Args("info").String()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaInfo(Info_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaChanged(Changed_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionCreated_ := self.Args("revision_created").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaRevisionCreated(RevisionCreated_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUser_ := self.Args("revision_user").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaRevisionUser(RevisionUser_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	_int64, _error := model.PutBlockContentFieldDataViaDefaultLangcode(DefaultLangcode_, &iBlockContentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaType(Type_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Info_ := self.Args("info").String()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaInfo(Info_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionCreated_ := self.Args("revision_created").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaRevisionCreated(RevisionCreated_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUser_ := self.Args("revision_user").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaRevisionUser(RevisionUser_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iBlockContentFieldData model.BlockContentFieldData
	self.Bind(&iBlockContentFieldData)
	var iMap = helper.StructToMap(iBlockContentFieldData)
	_error := model.UpdateBlockContentFieldDataViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteBlockContentFieldData(Id_)
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

func DeleteBlockContentFieldDataViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteBlockContentFieldDataViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteBlockContentFieldDataViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	_error := model.DeleteBlockContentFieldDataViaType(Type_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteBlockContentFieldDataViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Info_ := self.Args("info").String()
	_error := model.DeleteBlockContentFieldDataViaInfo(Info_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteBlockContentFieldDataViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionCreated_ := self.Args("revision_created").MustInt()
	_error := model.DeleteBlockContentFieldDataViaRevisionCreated(RevisionCreated_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUser_ := self.Args("revision_user").MustInt()
	_error := model.DeleteBlockContentFieldDataViaRevisionUser(RevisionUser_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	_error := model.DeleteBlockContentFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteBlockContentFieldDataViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
