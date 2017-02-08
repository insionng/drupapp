package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetBlockContentRevisionsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBlockContentRevisionsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["block_content_revisionsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetBlockContentRevisionCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revisionCount"] = 0
	}
	m["block_content_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevisionCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetBlockContentRevisionCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revisionCount"] = 0
	}
	m["block_content_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevisionCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetBlockContentRevisionCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revisionCount"] = 0
	}
	m["block_content_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevisionCountViaRevisionLogHandler(self *macross.Context) error {
	RevisionLog_ := self.Args("revision_log").String()
	_int64 := model.GetBlockContentRevisionCountViaRevisionLog(RevisionLog_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revisionCount"] = 0
	}
	m["block_content_revisionCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevisionsViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionLog := self.Args("revision_log").String()
	if (offset > 0) && helper.IsHas(iRevisionLog) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsViaRevisionLog(offset, limit, iRevisionLog, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsViaRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByIdAndRevisionIdAndLangcode(offset, limit, iId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByIdAndRevisionIdAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByIdAndRevisionIdAndRevisionLog(offset, limit, iId,iRevisionId,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByIdAndRevisionIdAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByIdAndLangcodeAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByIdAndLangcodeAndRevisionLog(offset, limit, iId,iLangcode,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByIdAndLangcodeAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByRevisionIdAndLangcodeAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByRevisionIdAndLangcodeAndRevisionLog(offset, limit, iRevisionId,iLangcode,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByRevisionIdAndLangcodeAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByIdAndRevisionId(offset, limit, iId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByIdAndLangcode(offset, limit, iId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByIdAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByIdAndRevisionLog(offset, limit, iId,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByIdAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByRevisionIdAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByRevisionIdAndRevisionLog(offset, limit, iRevisionId,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByRevisionIdAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsByLangcodeAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionsByLangcodeAndRevisionLog(offset, limit, iLangcode,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionsByLangcodeAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_BlockContentRevision, _error := model.GetBlockContentRevisions(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisions' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_BlockContentRevision := model.HasBlockContentRevisionViaId(iId)
		var m = map[string]interface{}{}
		m["block_content_revision"] = _BlockContentRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevisionViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContentRevision := model.HasBlockContentRevisionViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["block_content_revision"] = _BlockContentRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevisionViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContentRevision := model.HasBlockContentRevisionViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["block_content_revision"] = _BlockContentRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionLog := self.Args("revision_log").String()
	if helper.IsHas(iRevisionLog) {
		_BlockContentRevision := model.HasBlockContentRevisionViaRevisionLog(iRevisionLog)
		var m = map[string]interface{}{}
		m["block_content_revision"] = _BlockContentRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevisionViaRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionLog := self.Args("revision_log").String()
	if helper.IsHas(iRevisionLog) {
		_BlockContentRevision, _error := model.GetBlockContentRevisionViaRevisionLog(iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the GetBlockContentRevisionViaRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iBlockContentRevision model.BlockContentRevision
		self.Bind(&iBlockContentRevision)
		_BlockContentRevision, _error := model.SetBlockContentRevisionViaId(Id_, &iBlockContentRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the SetBlockContentRevisionViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iBlockContentRevision model.BlockContentRevision
		self.Bind(&iBlockContentRevision)
		_BlockContentRevision, _error := model.SetBlockContentRevisionViaRevisionId(RevisionId_, &iBlockContentRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the SetBlockContentRevisionViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iBlockContentRevision model.BlockContentRevision
		self.Bind(&iBlockContentRevision)
		_BlockContentRevision, _error := model.SetBlockContentRevisionViaLangcode(Langcode_, &iBlockContentRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the SetBlockContentRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionLog_ := self.Args("revision_log").String()
	if helper.IsHas(RevisionLog_) {
		var iBlockContentRevision model.BlockContentRevision
		self.Bind(&iBlockContentRevision)
		_BlockContentRevision, _error := model.SetBlockContentRevisionViaRevisionLog(RevisionLog_, &iBlockContentRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision)
	}
	herr.Message = "Can't get to the SetBlockContentRevisionViaRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBlockContentRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	RevisionId_ := self.Args("revision_id").MustInt()
	Langcode_ := self.Args("langcode").String()
	RevisionLog_ := self.Args("revision_log").String()

	if helper.IsHas(Id_) {
		_error := model.AddBlockContentRevision(Id_,RevisionId_,Langcode_,RevisionLog_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBlockContentRevision's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBlockContentRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	_int64, _error := model.PostBlockContentRevision(&iBlockContentRevision)
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

func PutBlockContentRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	_int64, _error := model.PutBlockContentRevision(&iBlockContentRevision)
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

func PutBlockContentRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	_int64, _error := model.PutBlockContentRevisionViaId(Id_, &iBlockContentRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	_int64, _error := model.PutBlockContentRevisionViaRevisionId(RevisionId_, &iBlockContentRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	_int64, _error := model.PutBlockContentRevisionViaLangcode(Langcode_, &iBlockContentRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionLog_ := self.Args("revision_log").String()
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	_int64, _error := model.PutBlockContentRevisionViaRevisionLog(RevisionLog_, &iBlockContentRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	var iMap = helper.StructToMap(iBlockContentRevision)
	_error := model.UpdateBlockContentRevisionViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	var iMap = helper.StructToMap(iBlockContentRevision)
	_error := model.UpdateBlockContentRevisionViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	var iMap = helper.StructToMap(iBlockContentRevision)
	_error := model.UpdateBlockContentRevisionViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionLog_ := self.Args("revision_log").String()
	var iBlockContentRevision model.BlockContentRevision
	self.Bind(&iBlockContentRevision)
	var iMap = helper.StructToMap(iBlockContentRevision)
	_error := model.UpdateBlockContentRevisionViaRevisionLog(RevisionLog_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteBlockContentRevision(Id_)
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

func DeleteBlockContentRevisionViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteBlockContentRevisionViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevisionViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteBlockContentRevisionViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteBlockContentRevisionViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionLog_ := self.Args("revision_log").String()
	_error := model.DeleteBlockContentRevisionViaRevisionLog(RevisionLog_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
