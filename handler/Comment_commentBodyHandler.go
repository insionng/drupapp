package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetComment_commentBodiesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetComment_commentBodiesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["comment__comment_bodysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetComment_commentBodyCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment__comment_bodyCount"] = 0
	}
	m["comment__comment_bodyCount"] = _int64
	return self.JSON(m)
}

func GetComment_commentBodyCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetComment_commentBodyCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment__comment_bodyCount"] = 0
	}
	m["comment__comment_bodyCount"] = _int64
	return self.JSON(m)
}

func GetComment_commentBodyCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetComment_commentBodyCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment__comment_bodyCount"] = 0
	}
	m["comment__comment_bodyCount"] = _int64
	return self.JSON(m)
}

func GetComment_commentBodyCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetComment_commentBodyCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment__comment_bodyCount"] = 0
	}
	m["comment__comment_bodyCount"] = _int64
	return self.JSON(m)
}

func GetComment_commentBodyCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetComment_commentBodyCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment__comment_bodyCount"] = 0
	}
	m["comment__comment_bodyCount"] = _int64
	return self.JSON(m)
}

func GetComment_commentBodyCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetComment_commentBodyCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment__comment_bodyCount"] = 0
	}
	m["comment__comment_bodyCount"] = _int64
	return self.JSON(m)
}

func GetComment_commentBodyCountViaCommentBodyValueHandler(self *macross.Context) error {
	CommentBodyValue_ := self.Args("comment_body_value").String()
	_int64 := model.GetComment_commentBodyCountViaCommentBodyValue(CommentBodyValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment__comment_bodyCount"] = 0
	}
	m["comment__comment_bodyCount"] = _int64
	return self.JSON(m)
}

func GetComment_commentBodyCountViaCommentBodyFormatHandler(self *macross.Context) error {
	CommentBodyFormat_ := self.Args("comment_body_format").String()
	_int64 := model.GetComment_commentBodyCountViaCommentBodyFormat(CommentBodyFormat_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment__comment_bodyCount"] = 0
	}
	m["comment__comment_bodyCount"] = _int64
	return self.JSON(m)
}

func GetComment_commentBodiesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesViaCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentBodyValue := self.Args("comment_body_value").String()
	if (offset > 0) && helper.IsHas(iCommentBodyValue) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesViaCommentBodyValue(offset, limit, iCommentBodyValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesViaCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesViaCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()
	if (offset > 0) && helper.IsHas(iCommentBodyFormat) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesViaCommentBodyFormat(offset, limit, iCommentBodyFormat, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesViaCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeletedAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeletedAndCommentBodyValue(offset, limit, iBundle,iDeleted,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeletedAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeletedAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeletedAndCommentBodyFormat(offset, limit, iBundle,iDeleted,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeletedAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyValue(offset, limit, iBundle,iEntityId,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyFormat(offset, limit, iBundle,iEntityId,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyValue(offset, limit, iBundle,iRevisionId,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyFormat(offset, limit, iBundle,iRevisionId,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyValue(offset, limit, iBundle,iLangcode,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyFormat(offset, limit, iBundle,iLangcode,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeltaAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeltaAndCommentBodyValue(offset, limit, iBundle,iDelta,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeltaAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeltaAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeltaAndCommentBodyFormat(offset, limit, iBundle,iDelta,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeltaAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndCommentBodyValueAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iCommentBodyValue := self.Args("comment_body_value").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndCommentBodyValueAndCommentBodyFormat(offset, limit, iBundle,iCommentBodyValue,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndCommentBodyValueAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyValue(offset, limit, iDeleted,iEntityId,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyFormat(offset, limit, iDeleted,iEntityId,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyValue(offset, limit, iDeleted,iRevisionId,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyFormat(offset, limit, iDeleted,iRevisionId,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyValue(offset, limit, iDeleted,iLangcode,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyFormat(offset, limit, iDeleted,iLangcode,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyValue(offset, limit, iDeleted,iDelta,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyFormat(offset, limit, iDeleted,iDelta,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndCommentBodyValueAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndCommentBodyValueAndCommentBodyFormat(offset, limit, iDeleted,iCommentBodyValue,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndCommentBodyValueAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyValue(offset, limit, iEntityId,iRevisionId,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyFormat(offset, limit, iEntityId,iRevisionId,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyValue(offset, limit, iEntityId,iLangcode,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyFormat(offset, limit, iEntityId,iLangcode,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyValue(offset, limit, iEntityId,iDelta,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyFormat(offset, limit, iEntityId,iDelta,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndCommentBodyValueAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndCommentBodyValueAndCommentBodyFormat(offset, limit, iEntityId,iCommentBodyValue,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndCommentBodyValueAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyValue(offset, limit, iRevisionId,iLangcode,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyFormat(offset, limit, iRevisionId,iLangcode,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyValue(offset, limit, iRevisionId,iDelta,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyFormat(offset, limit, iRevisionId,iDelta,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndCommentBodyValueAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndCommentBodyValueAndCommentBodyFormat(offset, limit, iRevisionId,iCommentBodyValue,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndCommentBodyValueAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iLangcode) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyValue(offset, limit, iLangcode,iDelta,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iLangcode) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyFormat(offset, limit, iLangcode,iDelta,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByLangcodeAndCommentBodyValueAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyValue := self.Args("comment_body_value").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iLangcode) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByLangcodeAndCommentBodyValueAndCommentBodyFormat(offset, limit, iLangcode,iCommentBodyValue,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByLangcodeAndCommentBodyValueAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeltaAndCommentBodyValueAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iDelta) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeltaAndCommentBodyValueAndCommentBodyFormat(offset, limit, iDelta,iCommentBodyValue,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeltaAndCommentBodyValueAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndCommentBodyValue(offset, limit, iBundle,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByBundleAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByBundleAndCommentBodyFormat(offset, limit, iBundle,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByBundleAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndCommentBodyValue(offset, limit, iDeleted,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeletedAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeletedAndCommentBodyFormat(offset, limit, iDeleted,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeletedAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndCommentBodyValue(offset, limit, iEntityId,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByEntityIdAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByEntityIdAndCommentBodyFormat(offset, limit, iEntityId,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByEntityIdAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndCommentBodyValue(offset, limit, iRevisionId,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByRevisionIdAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByRevisionIdAndCommentBodyFormat(offset, limit, iRevisionId,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByRevisionIdAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByLangcodeAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iLangcode) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByLangcodeAndCommentBodyValue(offset, limit, iLangcode,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByLangcodeAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByLangcodeAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iLangcode) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByLangcodeAndCommentBodyFormat(offset, limit, iLangcode,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByLangcodeAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeltaAndCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()

	if helper.IsHas(iDelta) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeltaAndCommentBodyValue(offset, limit, iDelta,iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeltaAndCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByDeltaAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iDelta) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByDeltaAndCommentBodyFormat(offset, limit, iDelta,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByDeltaAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesByCommentBodyValueAndCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentBodyValue := self.Args("comment_body_value").String()
	iCommentBodyFormat := self.Args("comment_body_format").String()

	if helper.IsHas(iCommentBodyValue) {
		_Comment_commentBody, _error := model.GetComment_commentBodiesByCommentBodyValueAndCommentBodyFormat(offset, limit, iCommentBodyValue,iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodiesByCommentBodyValueAndCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodiesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Comment_commentBody, _error := model.GetComment_commentBodies(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodies' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasComment_commentBodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_Comment_commentBody := model.HasComment_commentBodyViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["comment__comment_body"] = _Comment_commentBody
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasComment_commentBodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasComment_commentBodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Comment_commentBody := model.HasComment_commentBodyViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["comment__comment_body"] = _Comment_commentBody
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasComment_commentBodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasComment_commentBodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_Comment_commentBody := model.HasComment_commentBodyViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["comment__comment_body"] = _Comment_commentBody
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasComment_commentBodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasComment_commentBodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_Comment_commentBody := model.HasComment_commentBodyViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["comment__comment_body"] = _Comment_commentBody
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasComment_commentBodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasComment_commentBodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Comment_commentBody := model.HasComment_commentBodyViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["comment__comment_body"] = _Comment_commentBody
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasComment_commentBodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasComment_commentBodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_Comment_commentBody := model.HasComment_commentBodyViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["comment__comment_body"] = _Comment_commentBody
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasComment_commentBodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasComment_commentBodyViaCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentBodyValue := self.Args("comment_body_value").String()
	if helper.IsHas(iCommentBodyValue) {
		_Comment_commentBody := model.HasComment_commentBodyViaCommentBodyValue(iCommentBodyValue)
		var m = map[string]interface{}{}
		m["comment__comment_body"] = _Comment_commentBody
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasComment_commentBodyViaCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasComment_commentBodyViaCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentBodyFormat := self.Args("comment_body_format").String()
	if helper.IsHas(iCommentBodyFormat) {
		_Comment_commentBody := model.HasComment_commentBodyViaCommentBodyFormat(iCommentBodyFormat)
		var m = map[string]interface{}{}
		m["comment__comment_body"] = _Comment_commentBody
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasComment_commentBodyViaCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_Comment_commentBody, _error := model.GetComment_commentBodyViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Comment_commentBody, _error := model.GetComment_commentBodyViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_Comment_commentBody, _error := model.GetComment_commentBodyViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_Comment_commentBody, _error := model.GetComment_commentBodyViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Comment_commentBody, _error := model.GetComment_commentBodyViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_Comment_commentBody, _error := model.GetComment_commentBodyViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyViaCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentBodyValue := self.Args("comment_body_value").String()
	if helper.IsHas(iCommentBodyValue) {
		_Comment_commentBody, _error := model.GetComment_commentBodyViaCommentBodyValue(iCommentBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodyViaCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetComment_commentBodyViaCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentBodyFormat := self.Args("comment_body_format").String()
	if helper.IsHas(iCommentBodyFormat) {
		_Comment_commentBody, _error := model.GetComment_commentBodyViaCommentBodyFormat(iCommentBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the GetComment_commentBodyViaCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetComment_commentBodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iComment_commentBody model.Comment_commentBody
		self.Bind(&iComment_commentBody)
		_Comment_commentBody, _error := model.SetComment_commentBodyViaBundle(Bundle_, &iComment_commentBody)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the SetComment_commentBodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetComment_commentBodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iComment_commentBody model.Comment_commentBody
		self.Bind(&iComment_commentBody)
		_Comment_commentBody, _error := model.SetComment_commentBodyViaDeleted(Deleted_, &iComment_commentBody)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the SetComment_commentBodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetComment_commentBodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iComment_commentBody model.Comment_commentBody
		self.Bind(&iComment_commentBody)
		_Comment_commentBody, _error := model.SetComment_commentBodyViaEntityId(EntityId_, &iComment_commentBody)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the SetComment_commentBodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetComment_commentBodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iComment_commentBody model.Comment_commentBody
		self.Bind(&iComment_commentBody)
		_Comment_commentBody, _error := model.SetComment_commentBodyViaRevisionId(RevisionId_, &iComment_commentBody)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the SetComment_commentBodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetComment_commentBodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iComment_commentBody model.Comment_commentBody
		self.Bind(&iComment_commentBody)
		_Comment_commentBody, _error := model.SetComment_commentBodyViaLangcode(Langcode_, &iComment_commentBody)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the SetComment_commentBodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetComment_commentBodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iComment_commentBody model.Comment_commentBody
		self.Bind(&iComment_commentBody)
		_Comment_commentBody, _error := model.SetComment_commentBodyViaDelta(Delta_, &iComment_commentBody)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the SetComment_commentBodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetComment_commentBodyViaCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentBodyValue_ := self.Args("comment_body_value").String()
	if helper.IsHas(CommentBodyValue_) {
		var iComment_commentBody model.Comment_commentBody
		self.Bind(&iComment_commentBody)
		_Comment_commentBody, _error := model.SetComment_commentBodyViaCommentBodyValue(CommentBodyValue_, &iComment_commentBody)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the SetComment_commentBodyViaCommentBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetComment_commentBodyViaCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentBodyFormat_ := self.Args("comment_body_format").String()
	if helper.IsHas(CommentBodyFormat_) {
		var iComment_commentBody model.Comment_commentBody
		self.Bind(&iComment_commentBody)
		_Comment_commentBody, _error := model.SetComment_commentBodyViaCommentBodyFormat(CommentBodyFormat_, &iComment_commentBody)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment_commentBody)
	}
	herr.Message = "Can't get to the SetComment_commentBodyViaCommentBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddComment_commentBodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	Deleted_ := self.Args("deleted").MustInt()
	EntityId_ := self.Args("entity_id").MustInt()
	RevisionId_ := self.Args("revision_id").MustInt()
	Langcode_ := self.Args("langcode").String()
	Delta_ := self.Args("delta").MustInt()
	CommentBodyValue_ := self.Args("comment_body_value").String()
	CommentBodyFormat_ := self.Args("comment_body_format").String()

	if helper.IsHas(Bundle_) {
		_error := model.AddComment_commentBody(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,CommentBodyValue_,CommentBodyFormat_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddComment_commentBody's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostComment_commentBodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_string, _error := model.PostComment_commentBody(&iComment_commentBody)
	if (helper.IsHas(_string)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["created"] = _string
		return self.JSON(m, macross.StatusCreated)
	}
	return self.JSON(herr)
}

func PutComment_commentBodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_string, _error := model.PutComment_commentBody(&iComment_commentBody)
	if (helper.IsHas(_string)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["updated"] = _string
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func PutComment_commentBodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_int64, _error := model.PutComment_commentBodyViaBundle(Bundle_, &iComment_commentBody)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutComment_commentBodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_int64, _error := model.PutComment_commentBodyViaDeleted(Deleted_, &iComment_commentBody)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutComment_commentBodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_int64, _error := model.PutComment_commentBodyViaEntityId(EntityId_, &iComment_commentBody)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutComment_commentBodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_int64, _error := model.PutComment_commentBodyViaRevisionId(RevisionId_, &iComment_commentBody)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutComment_commentBodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_int64, _error := model.PutComment_commentBodyViaLangcode(Langcode_, &iComment_commentBody)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutComment_commentBodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_int64, _error := model.PutComment_commentBodyViaDelta(Delta_, &iComment_commentBody)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutComment_commentBodyViaCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentBodyValue_ := self.Args("comment_body_value").String()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_int64, _error := model.PutComment_commentBodyViaCommentBodyValue(CommentBodyValue_, &iComment_commentBody)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutComment_commentBodyViaCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentBodyFormat_ := self.Args("comment_body_format").String()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	_int64, _error := model.PutComment_commentBodyViaCommentBodyFormat(CommentBodyFormat_, &iComment_commentBody)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateComment_commentBodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	var iMap = helper.StructToMap(iComment_commentBody)
	_error := model.UpdateComment_commentBodyViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateComment_commentBodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	var iMap = helper.StructToMap(iComment_commentBody)
	_error := model.UpdateComment_commentBodyViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateComment_commentBodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	var iMap = helper.StructToMap(iComment_commentBody)
	_error := model.UpdateComment_commentBodyViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateComment_commentBodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	var iMap = helper.StructToMap(iComment_commentBody)
	_error := model.UpdateComment_commentBodyViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateComment_commentBodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	var iMap = helper.StructToMap(iComment_commentBody)
	_error := model.UpdateComment_commentBodyViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateComment_commentBodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	var iMap = helper.StructToMap(iComment_commentBody)
	_error := model.UpdateComment_commentBodyViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateComment_commentBodyViaCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentBodyValue_ := self.Args("comment_body_value").String()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	var iMap = helper.StructToMap(iComment_commentBody)
	_error := model.UpdateComment_commentBodyViaCommentBodyValue(CommentBodyValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateComment_commentBodyViaCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentBodyFormat_ := self.Args("comment_body_format").String()
	var iComment_commentBody model.Comment_commentBody
	self.Bind(&iComment_commentBody)
	var iMap = helper.StructToMap(iComment_commentBody)
	_error := model.UpdateComment_commentBodyViaCommentBodyFormat(CommentBodyFormat_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteComment_commentBodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteComment_commentBody(Bundle_)
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

func DeleteComment_commentBodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteComment_commentBodyViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteComment_commentBodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteComment_commentBodyViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteComment_commentBodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteComment_commentBodyViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteComment_commentBodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteComment_commentBodyViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteComment_commentBodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteComment_commentBodyViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteComment_commentBodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteComment_commentBodyViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteComment_commentBodyViaCommentBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentBodyValue_ := self.Args("comment_body_value").String()
	_error := model.DeleteComment_commentBodyViaCommentBodyValue(CommentBodyValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteComment_commentBodyViaCommentBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentBodyFormat_ := self.Args("comment_body_format").String()
	_error := model.DeleteComment_commentBodyViaCommentBodyFormat(CommentBodyFormat_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
