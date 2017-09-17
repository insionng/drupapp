package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetBlockContentFieldRevisionsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBlockContentFieldRevisionsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["block_content_field_revisionsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetBlockContentFieldRevisionCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetBlockContentFieldRevisionCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetBlockContentFieldRevisionCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionCountViaInfoHandler(self *macross.Context) error {
	Info_ := self.Args("info").String()
	_int64 := model.GetBlockContentFieldRevisionCountViaInfo(Info_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetBlockContentFieldRevisionCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionCountViaRevisionCreatedHandler(self *macross.Context) error {
	RevisionCreated_ := self.Args("revision_created").MustInt()
	_int64 := model.GetBlockContentFieldRevisionCountViaRevisionCreated(RevisionCreated_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionCountViaRevisionUserHandler(self *macross.Context) error {
	RevisionUser_ := self.Args("revision_user").MustInt()
	_int64 := model.GetBlockContentFieldRevisionCountViaRevisionUser(RevisionUser_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionCountViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	_int64 := model.GetBlockContentFieldRevisionCountViaRevisionTranslationAffected(RevisionTranslationAffected_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetBlockContentFieldRevisionCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_field_revisionCount"] = 0
	}
	m["block_content_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentFieldRevisionsViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iInfo := self.Args("info").String()
	if (offset > 0) && helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaInfo(offset, limit, iInfo, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaRevisionCreated(offset, limit, iRevisionCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionUser) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaRevisionUser(offset, limit, iRevisionUser, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionTranslationAffected) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaRevisionTranslationAffected(offset, limit, iRevisionTranslationAffected, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionIdAndLangcode(offset, limit, iId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionIdAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionIdAndInfo(offset, limit, iId,iRevisionId,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionIdAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionIdAndChanged(offset, limit, iId,iRevisionId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionCreated(offset, limit, iId,iRevisionId,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionUser(offset, limit, iId,iRevisionId,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionTranslationAffected(offset, limit, iId,iRevisionId,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionIdAndDefaultLangcode(offset, limit, iId,iRevisionId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndLangcodeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndLangcodeAndInfo(offset, limit, iId,iLangcode,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndLangcodeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndLangcodeAndChanged(offset, limit, iId,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionCreated(offset, limit, iId,iLangcode,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionUser(offset, limit, iId,iLangcode,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionTranslationAffected(offset, limit, iId,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndLangcodeAndDefaultLangcode(offset, limit, iId,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndInfoAndChanged(offset, limit, iId,iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndInfoAndRevisionCreated(offset, limit, iId,iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndInfoAndRevisionUser(offset, limit, iId,iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndInfoAndRevisionTranslationAffected(offset, limit, iId,iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndInfoAndDefaultLangcode(offset, limit, iId,iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndChangedAndRevisionCreated(offset, limit, iId,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndChangedAndRevisionUser(offset, limit, iId,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndChangedAndRevisionTranslationAffected(offset, limit, iId,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndChangedAndDefaultLangcode(offset, limit, iId,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionUser(offset, limit, iId,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iId,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndDefaultLangcode(offset, limit, iId,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionUserAndRevisionTranslationAffected(offset, limit, iId,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionUserAndDefaultLangcode(offset, limit, iId,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iId,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndInfo(offset, limit, iRevisionId,iLangcode,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndChanged(offset, limit, iRevisionId,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionCreated(offset, limit, iRevisionId,iLangcode,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionUser(offset, limit, iRevisionId,iLangcode,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionTranslationAffected(offset, limit, iRevisionId,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndDefaultLangcode(offset, limit, iRevisionId,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndInfoAndChanged(offset, limit, iRevisionId,iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionCreated(offset, limit, iRevisionId,iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionUser(offset, limit, iRevisionId,iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionTranslationAffected(offset, limit, iRevisionId,iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndInfoAndDefaultLangcode(offset, limit, iRevisionId,iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionCreated(offset, limit, iRevisionId,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionUser(offset, limit, iRevisionId,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionTranslationAffected(offset, limit, iRevisionId,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndChangedAndDefaultLangcode(offset, limit, iRevisionId,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionUser(offset, limit, iRevisionId,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iRevisionId,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndDefaultLangcode(offset, limit, iRevisionId,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndRevisionTranslationAffected(offset, limit, iRevisionId,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndDefaultLangcode(offset, limit, iRevisionId,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionId,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndInfoAndChanged(offset, limit, iLangcode,iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionCreated(offset, limit, iLangcode,iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionUser(offset, limit, iLangcode,iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionTranslationAffected(offset, limit, iLangcode,iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndInfoAndDefaultLangcode(offset, limit, iLangcode,iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionCreated(offset, limit, iLangcode,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionUser(offset, limit, iLangcode,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffected(offset, limit, iLangcode,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndChangedAndDefaultLangcode(offset, limit, iLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionUser(offset, limit, iLangcode,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iLangcode,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndDefaultLangcode(offset, limit, iLangcode,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndRevisionTranslationAffected(offset, limit, iLangcode,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndDefaultLangcode(offset, limit, iLangcode,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iLangcode,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionCreated(offset, limit, iInfo,iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionUser(offset, limit, iInfo,iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionTranslationAffected(offset, limit, iInfo,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndChangedAndDefaultLangcode(offset, limit, iInfo,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionUser(offset, limit, iInfo,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iInfo,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndDefaultLangcode(offset, limit, iInfo,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionUserAndRevisionTranslationAffected(offset, limit, iInfo,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionUserAndDefaultLangcode(offset, limit, iInfo,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iInfo,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionUser(offset, limit, iChanged,iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionTranslationAffected(offset, limit, iChanged,iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndDefaultLangcode(offset, limit, iChanged,iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionUserAndRevisionTranslationAffected(offset, limit, iChanged,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionUserAndDefaultLangcode(offset, limit, iChanged,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iChanged,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndRevisionTranslationAffected(offset, limit, iRevisionCreated,iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndDefaultLangcode(offset, limit, iRevisionCreated,iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionCreated,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionUser,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionId(offset, limit, iId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndLangcode(offset, limit, iId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iInfo := self.Args("info").String()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndInfo(offset, limit, iId,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndChanged(offset, limit, iId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionCreated(offset, limit, iId,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionUser(offset, limit, iId,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffected(offset, limit, iId,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByIdAndDefaultLangcode(offset, limit, iId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iInfo := self.Args("info").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndInfo(offset, limit, iRevisionId,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndChanged(offset, limit, iRevisionId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreated(offset, limit, iRevisionId,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionUser(offset, limit, iRevisionId,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffected(offset, limit, iRevisionId,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionIdAndDefaultLangcode(offset, limit, iRevisionId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInfo := self.Args("info").String()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndInfo(offset, limit, iLangcode,iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionCreated(offset, limit, iLangcode,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionUser(offset, limit, iLangcode,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffected(offset, limit, iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndChanged(offset, limit, iInfo,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionCreated(offset, limit, iInfo,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionUser(offset, limit, iInfo,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffected(offset, limit, iInfo,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByInfoAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInfo := self.Args("info").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByInfoAndDefaultLangcode(offset, limit, iInfo,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByInfoAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionCreated(offset, limit, iChanged,iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionUser(offset, limit, iChanged,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffected(offset, limit, iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByChangedAndDefaultLangcode(offset, limit, iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUser(offset, limit, iRevisionCreated,iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffected(offset, limit, iRevisionCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionCreated := self.Args("revision_created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionCreatedAndDefaultLangcode(offset, limit, iRevisionCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffected(offset, limit, iRevisionUser,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionUserAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionUser := self.Args("revision_user").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionUserAndDefaultLangcode(offset, limit, iRevisionUser,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionUserAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionTranslationAffected) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisions(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisions' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaId(iId)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iInfo := self.Args("info").String()
	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaInfo(iInfo)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionCreated := self.Args("revision_created").MustInt()
	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaRevisionCreated(iRevisionCreated)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionUser := self.Args("revision_user").MustInt()
	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaRevisionUser(iRevisionUser)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(iRevisionTranslationAffected) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_BlockContentFieldRevision := model.HasBlockContentFieldRevisionViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["block_content_field_revision"] = _BlockContentFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentFieldRevisionViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iInfo := self.Args("info").String()
	if helper.IsHas(iInfo) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaInfo(iInfo)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionCreated := self.Args("revision_created").MustInt()
	if helper.IsHas(iRevisionCreated) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaRevisionCreated(iRevisionCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionUser := self.Args("revision_user").MustInt()
	if helper.IsHas(iRevisionUser) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaRevisionUser(iRevisionUser)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(iRevisionTranslationAffected) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_BlockContentFieldRevision, _error := model.GetBlockContentFieldRevisionViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the GetBlockContentFieldRevisionViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaId(Id_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaRevisionId(RevisionId_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaLangcode(Langcode_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Info_ := self.Args("info").String()
	if helper.IsHas(Info_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaInfo(Info_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaInfo's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaChanged(Changed_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionCreated_ := self.Args("revision_created").MustInt()
	if helper.IsHas(RevisionCreated_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaRevisionCreated(RevisionCreated_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaRevisionCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUser_ := self.Args("revision_user").MustInt()
	if helper.IsHas(RevisionUser_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaRevisionUser(RevisionUser_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaRevisionUser's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(RevisionTranslationAffected_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iBlockContentFieldRevision model.BlockContentFieldRevision
		self.Bind(&iBlockContentFieldRevision)
		_BlockContentFieldRevision, _error := model.SetBlockContentFieldRevisionViaDefaultLangcode(DefaultLangcode_, &iBlockContentFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentFieldRevision)
	}
	herr.Message = "Can't get to the SetBlockContentFieldRevisionViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBlockContentFieldRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	RevisionId_ := self.Args("revision_id").MustInt()
	Langcode_ := self.Args("langcode").String()
	Info_ := self.Args("info").String()
	Changed_ := self.Args("changed").MustInt()
	RevisionCreated_ := self.Args("revision_created").MustInt()
	RevisionUser_ := self.Args("revision_user").MustInt()
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Id_) {
		_error := model.AddBlockContentFieldRevision(Id_,RevisionId_,Langcode_,Info_,Changed_,RevisionCreated_,RevisionUser_,RevisionTranslationAffected_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBlockContentFieldRevision's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBlockContentFieldRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PostBlockContentFieldRevision(&iBlockContentFieldRevision)
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

func PutBlockContentFieldRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevision(&iBlockContentFieldRevision)
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

func PutBlockContentFieldRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaId(Id_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaRevisionId(RevisionId_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaLangcode(Langcode_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldRevisionViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Info_ := self.Args("info").String()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaInfo(Info_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaChanged(Changed_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldRevisionViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionCreated_ := self.Args("revision_created").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaRevisionCreated(RevisionCreated_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldRevisionViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUser_ := self.Args("revision_user").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaRevisionUser(RevisionUser_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	_int64, _error := model.PutBlockContentFieldRevisionViaDefaultLangcode(DefaultLangcode_, &iBlockContentFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Info_ := self.Args("info").String()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaInfo(Info_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionCreated_ := self.Args("revision_created").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaRevisionCreated(RevisionCreated_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUser_ := self.Args("revision_user").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaRevisionUser(RevisionUser_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iBlockContentFieldRevision model.BlockContentFieldRevision
	self.Bind(&iBlockContentFieldRevision)
	var iMap = helper.StructToMap(iBlockContentFieldRevision)
	_error := model.UpdateBlockContentFieldRevisionViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteBlockContentFieldRevision(Id_)
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

func DeleteBlockContentFieldRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteBlockContentFieldRevisionViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteBlockContentFieldRevisionViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteBlockContentFieldRevisionViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionViaInfoHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Info_ := self.Args("info").String()
	_error := model.DeleteBlockContentFieldRevisionViaInfo(Info_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteBlockContentFieldRevisionViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionViaRevisionCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionCreated_ := self.Args("revision_created").MustInt()
	_error := model.DeleteBlockContentFieldRevisionViaRevisionCreated(RevisionCreated_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionViaRevisionUserHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUser_ := self.Args("revision_user").MustInt()
	_error := model.DeleteBlockContentFieldRevisionViaRevisionUser(RevisionUser_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	_error := model.DeleteBlockContentFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteBlockContentFieldRevisionViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
