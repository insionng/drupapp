package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetBlockContentsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBlockContentsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["block_contentsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBlockContentsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetBlockContentCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_contentCount"] = 0
	}
	m["block_contentCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetBlockContentCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_contentCount"] = 0
	}
	m["block_contentCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentCountViaTypeHandler(self *macross.Context) error {
	Type_ := self.Args("type").String()
	_int64 := model.GetBlockContentCountViaType(Type_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_contentCount"] = 0
	}
	m["block_contentCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentCountViaUuidHandler(self *macross.Context) error {
	Uuid_ := self.Args("uuid").String()
	_int64 := model.GetBlockContentCountViaUuid(Uuid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_contentCount"] = 0
	}
	m["block_contentCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetBlockContentCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_contentCount"] = 0
	}
	m["block_contentCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentsViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_BlockContent, _error := model.GetBlockContentsViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iType := self.Args("type").String()
	if (offset > 0) && helper.IsHas(iType) {
		_BlockContent, _error := model.GetBlockContentsViaType(offset, limit, iType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUuid := self.Args("uuid").String()
	if (offset > 0) && helper.IsHas(iUuid) {
		_BlockContent, _error := model.GetBlockContentsViaUuid(offset, limit, iUuid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_BlockContent, _error := model.GetBlockContentsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndRevisionIdAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndRevisionIdAndType(offset, limit, iId,iRevisionId,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndRevisionIdAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndRevisionIdAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndRevisionIdAndUuid(offset, limit, iId,iRevisionId,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndRevisionIdAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndRevisionIdAndLangcode(offset, limit, iId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndTypeAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndTypeAndUuid(offset, limit, iId,iType,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndTypeAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndTypeAndLangcode(offset, limit, iId,iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndUuidAndLangcode(offset, limit, iId,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByRevisionIdAndTypeAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent, _error := model.GetBlockContentsByRevisionIdAndTypeAndUuid(offset, limit, iRevisionId,iType,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByRevisionIdAndTypeAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByRevisionIdAndTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent, _error := model.GetBlockContentsByRevisionIdAndTypeAndLangcode(offset, limit, iRevisionId,iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByRevisionIdAndTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByRevisionIdAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent, _error := model.GetBlockContentsByRevisionIdAndUuidAndLangcode(offset, limit, iRevisionId,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByRevisionIdAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByTypeAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iType) {
		_BlockContent, _error := model.GetBlockContentsByTypeAndUuidAndLangcode(offset, limit, iType,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByTypeAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndRevisionId(offset, limit, iId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iType := self.Args("type").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndType(offset, limit, iId,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndUuid(offset, limit, iId,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentsByIdAndLangcode(offset, limit, iId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByRevisionIdAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent, _error := model.GetBlockContentsByRevisionIdAndType(offset, limit, iRevisionId,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByRevisionIdAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByRevisionIdAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent, _error := model.GetBlockContentsByRevisionIdAndUuid(offset, limit, iRevisionId,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByRevisionIdAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent, _error := model.GetBlockContentsByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByTypeAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iType) {
		_BlockContent, _error := model.GetBlockContentsByTypeAndUuid(offset, limit, iType,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByTypeAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iType) {
		_BlockContent, _error := model.GetBlockContentsByTypeAndLangcode(offset, limit, iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsByUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUuid) {
		_BlockContent, _error := model.GetBlockContentsByUuidAndLangcode(offset, limit, iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentsByUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_BlockContent, _error := model.GetBlockContents(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContents' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_BlockContent := model.HasBlockContentViaId(iId)
		var m = map[string]interface{}{}
		m["block_content"] = _BlockContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContent := model.HasBlockContentViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["block_content"] = _BlockContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_BlockContent := model.HasBlockContentViaType(iType)
		var m = map[string]interface{}{}
		m["block_content"] = _BlockContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_BlockContent := model.HasBlockContentViaUuid(iUuid)
		var m = map[string]interface{}{}
		m["block_content"] = _BlockContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContent := model.HasBlockContentViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["block_content"] = _BlockContent
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_BlockContent, _error := model.GetBlockContentViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContent, _error := model.GetBlockContentViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_BlockContent, _error := model.GetBlockContentViaType(iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_BlockContent, _error := model.GetBlockContentViaUuid(iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContent, _error := model.GetBlockContentViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the GetBlockContentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iBlockContent model.BlockContent
		self.Bind(&iBlockContent)
		_BlockContent, _error := model.SetBlockContentViaId(Id_, &iBlockContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the SetBlockContentViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iBlockContent model.BlockContent
		self.Bind(&iBlockContent)
		_BlockContent, _error := model.SetBlockContentViaRevisionId(RevisionId_, &iBlockContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the SetBlockContentViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	if helper.IsHas(Type_) {
		var iBlockContent model.BlockContent
		self.Bind(&iBlockContent)
		_BlockContent, _error := model.SetBlockContentViaType(Type_, &iBlockContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the SetBlockContentViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	if helper.IsHas(Uuid_) {
		var iBlockContent model.BlockContent
		self.Bind(&iBlockContent)
		_BlockContent, _error := model.SetBlockContentViaUuid(Uuid_, &iBlockContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the SetBlockContentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iBlockContent model.BlockContent
		self.Bind(&iBlockContent)
		_BlockContent, _error := model.SetBlockContentViaLangcode(Langcode_, &iBlockContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent)
	}
	herr.Message = "Can't get to the SetBlockContentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBlockContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	RevisionId_ := self.Args("revision_id").MustInt()
	Type_ := self.Args("type").String()
	Uuid_ := self.Args("uuid").String()
	Langcode_ := self.Args("langcode").String()

	if helper.IsHas(Id_) {
		_error := model.AddBlockContent(Id_,RevisionId_,Type_,Uuid_,Langcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBlockContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBlockContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	_int64, _error := model.PostBlockContent(&iBlockContent)
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

func PutBlockContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	_int64, _error := model.PutBlockContent(&iBlockContent)
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

func PutBlockContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	_int64, _error := model.PutBlockContentViaId(Id_, &iBlockContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	_int64, _error := model.PutBlockContentViaRevisionId(RevisionId_, &iBlockContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	_int64, _error := model.PutBlockContentViaType(Type_, &iBlockContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	_int64, _error := model.PutBlockContentViaUuid(Uuid_, &iBlockContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	_int64, _error := model.PutBlockContentViaLangcode(Langcode_, &iBlockContent)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	var iMap = helper.StructToMap(iBlockContent)
	_error := model.UpdateBlockContentViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	var iMap = helper.StructToMap(iBlockContent)
	_error := model.UpdateBlockContentViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	var iMap = helper.StructToMap(iBlockContent)
	_error := model.UpdateBlockContentViaType(Type_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	var iMap = helper.StructToMap(iBlockContent)
	_error := model.UpdateBlockContentViaUuid(Uuid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContent model.BlockContent
	self.Bind(&iBlockContent)
	var iMap = helper.StructToMap(iBlockContent)
	_error := model.UpdateBlockContentViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteBlockContent(Id_)
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

func DeleteBlockContentViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteBlockContentViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteBlockContentViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	_error := model.DeleteBlockContentViaType(Type_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	_error := model.DeleteBlockContentViaUuid(Uuid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteBlockContentViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
