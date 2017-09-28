package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetNodeRevision_commentsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodeRevision_commentsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node_revision__commentsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetNodeRevision_commentCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__commentCount"] = 0
	}
	m["node_revision__commentCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_commentCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetNodeRevision_commentCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__commentCount"] = 0
	}
	m["node_revision__commentCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_commentCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetNodeRevision_commentCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__commentCount"] = 0
	}
	m["node_revision__commentCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_commentCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetNodeRevision_commentCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__commentCount"] = 0
	}
	m["node_revision__commentCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_commentCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeRevision_commentCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__commentCount"] = 0
	}
	m["node_revision__commentCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_commentCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetNodeRevision_commentCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__commentCount"] = 0
	}
	m["node_revision__commentCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_commentCountViaCommentStatusHandler(self *macross.Context) error {
	CommentStatus_ := self.Args("comment_status").MustInt()
	_int64 := model.GetNodeRevision_commentCountViaCommentStatus(CommentStatus_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__commentCount"] = 0
	}
	m["node_revision__commentCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_commentsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsViaCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentStatus := self.Args("comment_status").MustInt()
	if (offset > 0) && helper.IsHas(iCommentStatus) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsViaCommentStatus(offset, limit, iCommentStatus, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsViaCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndDeletedAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndDeletedAndCommentStatus(offset, limit, iBundle,iDeleted,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndDeletedAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndEntityIdAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndEntityIdAndCommentStatus(offset, limit, iBundle,iEntityId,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndEntityIdAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndRevisionIdAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndRevisionIdAndCommentStatus(offset, limit, iBundle,iRevisionId,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndRevisionIdAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndLangcodeAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndLangcodeAndCommentStatus(offset, limit, iBundle,iLangcode,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndLangcodeAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndDeltaAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndDeltaAndCommentStatus(offset, limit, iBundle,iDelta,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndDeltaAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndEntityIdAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndEntityIdAndCommentStatus(offset, limit, iDeleted,iEntityId,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndEntityIdAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndRevisionIdAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndRevisionIdAndCommentStatus(offset, limit, iDeleted,iRevisionId,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndRevisionIdAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndLangcodeAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndLangcodeAndCommentStatus(offset, limit, iDeleted,iLangcode,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndLangcodeAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndDeltaAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndDeltaAndCommentStatus(offset, limit, iDeleted,iDelta,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndDeltaAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndRevisionIdAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndRevisionIdAndCommentStatus(offset, limit, iEntityId,iRevisionId,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndRevisionIdAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndLangcodeAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndLangcodeAndCommentStatus(offset, limit, iEntityId,iLangcode,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndLangcodeAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndDeltaAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndDeltaAndCommentStatus(offset, limit, iEntityId,iDelta,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndDeltaAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByRevisionIdAndLangcodeAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByRevisionIdAndLangcodeAndCommentStatus(offset, limit, iRevisionId,iLangcode,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByRevisionIdAndLangcodeAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByRevisionIdAndDeltaAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByRevisionIdAndDeltaAndCommentStatus(offset, limit, iRevisionId,iDelta,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByRevisionIdAndDeltaAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByLangcodeAndDeltaAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByLangcodeAndDeltaAndCommentStatus(offset, limit, iLangcode,iDelta,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByLangcodeAndDeltaAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByBundleAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByBundleAndCommentStatus(offset, limit, iBundle,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByBundleAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeletedAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeletedAndCommentStatus(offset, limit, iDeleted,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeletedAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByEntityIdAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByEntityIdAndCommentStatus(offset, limit, iEntityId,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByEntityIdAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByRevisionIdAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByRevisionIdAndCommentStatus(offset, limit, iRevisionId,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByRevisionIdAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByLangcodeAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByLangcodeAndCommentStatus(offset, limit, iLangcode,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByLangcodeAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsByDeltaAndCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iCommentStatus := self.Args("comment_status").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentsByDeltaAndCommentStatus(offset, limit, iDelta,iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentsByDeltaAndCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_NodeRevision_comment, _error := model.GetNodeRevision_comments(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_comments' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_commentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_NodeRevision_comment := model.HasNodeRevision_commentViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["node_revision__comment"] = _NodeRevision_comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_commentViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_commentViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_NodeRevision_comment := model.HasNodeRevision_commentViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["node_revision__comment"] = _NodeRevision_comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_commentViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_commentViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_NodeRevision_comment := model.HasNodeRevision_commentViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["node_revision__comment"] = _NodeRevision_comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_commentViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_commentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_NodeRevision_comment := model.HasNodeRevision_commentViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["node_revision__comment"] = _NodeRevision_comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_commentViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_commentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision_comment := model.HasNodeRevision_commentViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node_revision__comment"] = _NodeRevision_comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_commentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_commentViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_NodeRevision_comment := model.HasNodeRevision_commentViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["node_revision__comment"] = _NodeRevision_comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_commentViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_commentViaCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentStatus := self.Args("comment_status").MustInt()
	if helper.IsHas(iCommentStatus) {
		_NodeRevision_comment := model.HasNodeRevision_commentViaCommentStatus(iCommentStatus)
		var m = map[string]interface{}{}
		m["node_revision__comment"] = _NodeRevision_comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_commentViaCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_commentViaCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentStatus := self.Args("comment_status").MustInt()
	if helper.IsHas(iCommentStatus) {
		_NodeRevision_comment, _error := model.GetNodeRevision_commentViaCommentStatus(iCommentStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the GetNodeRevision_commentViaCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_commentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iNodeRevision_comment model.NodeRevision_comment
		self.Bind(&iNodeRevision_comment)
		_NodeRevision_comment, _error := model.SetNodeRevision_commentViaBundle(Bundle_, &iNodeRevision_comment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the SetNodeRevision_commentViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_commentViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iNodeRevision_comment model.NodeRevision_comment
		self.Bind(&iNodeRevision_comment)
		_NodeRevision_comment, _error := model.SetNodeRevision_commentViaDeleted(Deleted_, &iNodeRevision_comment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the SetNodeRevision_commentViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_commentViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iNodeRevision_comment model.NodeRevision_comment
		self.Bind(&iNodeRevision_comment)
		_NodeRevision_comment, _error := model.SetNodeRevision_commentViaEntityId(EntityId_, &iNodeRevision_comment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the SetNodeRevision_commentViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_commentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iNodeRevision_comment model.NodeRevision_comment
		self.Bind(&iNodeRevision_comment)
		_NodeRevision_comment, _error := model.SetNodeRevision_commentViaRevisionId(RevisionId_, &iNodeRevision_comment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the SetNodeRevision_commentViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_commentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNodeRevision_comment model.NodeRevision_comment
		self.Bind(&iNodeRevision_comment)
		_NodeRevision_comment, _error := model.SetNodeRevision_commentViaLangcode(Langcode_, &iNodeRevision_comment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the SetNodeRevision_commentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_commentViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iNodeRevision_comment model.NodeRevision_comment
		self.Bind(&iNodeRevision_comment)
		_NodeRevision_comment, _error := model.SetNodeRevision_commentViaDelta(Delta_, &iNodeRevision_comment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the SetNodeRevision_commentViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_commentViaCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentStatus_ := self.Args("comment_status").MustInt()
	if helper.IsHas(CommentStatus_) {
		var iNodeRevision_comment model.NodeRevision_comment
		self.Bind(&iNodeRevision_comment)
		_NodeRevision_comment, _error := model.SetNodeRevision_commentViaCommentStatus(CommentStatus_, &iNodeRevision_comment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_comment)
	}
	herr.Message = "Can't get to the SetNodeRevision_commentViaCommentStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeRevision_commentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	Deleted_ := self.Args("deleted").MustInt()
	EntityId_ := self.Args("entity_id").MustInt()
	RevisionId_ := self.Args("revision_id").MustInt()
	Langcode_ := self.Args("langcode").String()
	Delta_ := self.Args("delta").MustInt()
	CommentStatus_ := self.Args("comment_status").MustInt()

	if helper.IsHas(Bundle_) {
		_error := model.AddNodeRevision_comment(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,CommentStatus_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNodeRevision_comment's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeRevision_commentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_string, _error := model.PostNodeRevision_comment(&iNodeRevision_comment)
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

func PutNodeRevision_commentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_string, _error := model.PutNodeRevision_comment(&iNodeRevision_comment)
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

func PutNodeRevision_commentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_int64, _error := model.PutNodeRevision_commentViaBundle(Bundle_, &iNodeRevision_comment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_commentViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_int64, _error := model.PutNodeRevision_commentViaDeleted(Deleted_, &iNodeRevision_comment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_commentViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_int64, _error := model.PutNodeRevision_commentViaEntityId(EntityId_, &iNodeRevision_comment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_commentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_int64, _error := model.PutNodeRevision_commentViaRevisionId(RevisionId_, &iNodeRevision_comment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_commentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_int64, _error := model.PutNodeRevision_commentViaLangcode(Langcode_, &iNodeRevision_comment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_commentViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_int64, _error := model.PutNodeRevision_commentViaDelta(Delta_, &iNodeRevision_comment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_commentViaCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentStatus_ := self.Args("comment_status").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	_int64, _error := model.PutNodeRevision_commentViaCommentStatus(CommentStatus_, &iNodeRevision_comment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_commentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	var iMap = helper.StructToMap(iNodeRevision_comment)
	_error := model.UpdateNodeRevision_commentViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_commentViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	var iMap = helper.StructToMap(iNodeRevision_comment)
	_error := model.UpdateNodeRevision_commentViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_commentViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	var iMap = helper.StructToMap(iNodeRevision_comment)
	_error := model.UpdateNodeRevision_commentViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_commentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	var iMap = helper.StructToMap(iNodeRevision_comment)
	_error := model.UpdateNodeRevision_commentViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_commentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	var iMap = helper.StructToMap(iNodeRevision_comment)
	_error := model.UpdateNodeRevision_commentViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_commentViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	var iMap = helper.StructToMap(iNodeRevision_comment)
	_error := model.UpdateNodeRevision_commentViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_commentViaCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentStatus_ := self.Args("comment_status").MustInt()
	var iNodeRevision_comment model.NodeRevision_comment
	self.Bind(&iNodeRevision_comment)
	var iMap = helper.StructToMap(iNodeRevision_comment)
	_error := model.UpdateNodeRevision_commentViaCommentStatus(CommentStatus_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_commentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteNodeRevision_comment(Bundle_)
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

func DeleteNodeRevision_commentViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteNodeRevision_commentViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_commentViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteNodeRevision_commentViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_commentViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteNodeRevision_commentViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_commentViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteNodeRevision_commentViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_commentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeRevision_commentViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_commentViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteNodeRevision_commentViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_commentViaCommentStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentStatus_ := self.Args("comment_status").MustInt()
	_error := model.DeleteNodeRevision_commentViaCommentStatus(CommentStatus_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
