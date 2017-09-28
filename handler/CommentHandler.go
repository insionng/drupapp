package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetCommentsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCommentsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["commentsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCommentsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentCountViaCidHandler(self *macross.Context) error {
	Cid_ := self.Args("cid").MustInt64()
	_int64 := model.GetCommentCountViaCid(Cid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentCount"] = 0
	}
	m["commentCount"] = _int64
	return self.JSON(m)
}

func GetCommentCountViaCommentTypeHandler(self *macross.Context) error {
	CommentType_ := self.Args("comment_type").String()
	_int64 := model.GetCommentCountViaCommentType(CommentType_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentCount"] = 0
	}
	m["commentCount"] = _int64
	return self.JSON(m)
}

func GetCommentCountViaUuidHandler(self *macross.Context) error {
	Uuid_ := self.Args("uuid").String()
	_int64 := model.GetCommentCountViaUuid(Uuid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentCount"] = 0
	}
	m["commentCount"] = _int64
	return self.JSON(m)
}

func GetCommentCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetCommentCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentCount"] = 0
	}
	m["commentCount"] = _int64
	return self.JSON(m)
}

func GetCommentsViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCid := self.Args("cid").MustInt64()
	if (offset > 0) && helper.IsHas(iCid) {
		_Comment, _error := model.GetCommentsViaCid(offset, limit, iCid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentType := self.Args("comment_type").String()
	if (offset > 0) && helper.IsHas(iCommentType) {
		_Comment, _error := model.GetCommentsViaCommentType(offset, limit, iCommentType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsViaCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUuid := self.Args("uuid").String()
	if (offset > 0) && helper.IsHas(iUuid) {
		_Comment, _error := model.GetCommentsViaUuid(offset, limit, iUuid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_Comment, _error := model.GetCommentsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCidAndCommentTypeAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iCid) {
		_Comment, _error := model.GetCommentsByCidAndCommentTypeAndUuid(offset, limit, iCid,iCommentType,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCidAndCommentTypeAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCidAndCommentTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iCid) {
		_Comment, _error := model.GetCommentsByCidAndCommentTypeAndLangcode(offset, limit, iCid,iCommentType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCidAndCommentTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCidAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iCid) {
		_Comment, _error := model.GetCommentsByCidAndUuidAndLangcode(offset, limit, iCid,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCidAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentTypeAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iCommentType) {
		_Comment, _error := model.GetCommentsByCommentTypeAndUuidAndLangcode(offset, limit, iCommentType,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCommentTypeAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCidAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCid) {
		_Comment, _error := model.GetCommentsByCidAndCommentType(offset, limit, iCid,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCidAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iCid) {
		_Comment, _error := model.GetCommentsByCidAndUuid(offset, limit, iCid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iCid) {
		_Comment, _error := model.GetCommentsByCidAndLangcode(offset, limit, iCid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentTypeAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iCommentType) {
		_Comment, _error := model.GetCommentsByCommentTypeAndUuid(offset, limit, iCommentType,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCommentTypeAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iCommentType) {
		_Comment, _error := model.GetCommentsByCommentTypeAndLangcode(offset, limit, iCommentType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByCommentTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUuid) {
		_Comment, _error := model.GetCommentsByUuidAndLangcode(offset, limit, iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentsByUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Comment, _error := model.GetComments(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetComments' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").MustInt64()
	if helper.IsHas(iCid) {
		_Comment := model.HasCommentViaCid(iCid)
		var m = map[string]interface{}{}
		m["comment"] = _Comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentType := self.Args("comment_type").String()
	if helper.IsHas(iCommentType) {
		_Comment := model.HasCommentViaCommentType(iCommentType)
		var m = map[string]interface{}{}
		m["comment"] = _Comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentViaCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_Comment := model.HasCommentViaUuid(iUuid)
		var m = map[string]interface{}{}
		m["comment"] = _Comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Comment := model.HasCommentViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["comment"] = _Comment
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").MustInt64()
	if helper.IsHas(iCid) {
		_Comment, _error := model.GetCommentViaCid(iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentType := self.Args("comment_type").String()
	if helper.IsHas(iCommentType) {
		_Comment, _error := model.GetCommentViaCommentType(iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentViaCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_Comment, _error := model.GetCommentViaUuid(iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Comment, _error := model.GetCommentViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the GetCommentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	if helper.IsHas(Cid_) {
		var iComment model.Comment
		self.Bind(&iComment)
		_Comment, _error := model.SetCommentViaCid(Cid_, &iComment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the SetCommentViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	if helper.IsHas(CommentType_) {
		var iComment model.Comment
		self.Bind(&iComment)
		_Comment, _error := model.SetCommentViaCommentType(CommentType_, &iComment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the SetCommentViaCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	if helper.IsHas(Uuid_) {
		var iComment model.Comment
		self.Bind(&iComment)
		_Comment, _error := model.SetCommentViaUuid(Uuid_, &iComment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the SetCommentViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iComment model.Comment
		self.Bind(&iComment)
		_Comment, _error := model.SetCommentViaLangcode(Langcode_, &iComment)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comment)
	}
	herr.Message = "Can't get to the SetCommentViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCommentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	CommentType_ := self.Args("comment_type").String()
	Uuid_ := self.Args("uuid").String()
	Langcode_ := self.Args("langcode").String()

	if helper.IsHas(Cid_) {
		_error := model.AddComment(Cid_,CommentType_,Uuid_,Langcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddComment's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCommentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iComment model.Comment
	self.Bind(&iComment)
	_int64, _error := model.PostComment(&iComment)
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

func PutCommentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iComment model.Comment
	self.Bind(&iComment)
	_int64, _error := model.PutComment(&iComment)
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

func PutCommentViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	var iComment model.Comment
	self.Bind(&iComment)
	_int64, _error := model.PutCommentViaCid(Cid_, &iComment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	var iComment model.Comment
	self.Bind(&iComment)
	_int64, _error := model.PutCommentViaCommentType(CommentType_, &iComment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iComment model.Comment
	self.Bind(&iComment)
	_int64, _error := model.PutCommentViaUuid(Uuid_, &iComment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iComment model.Comment
	self.Bind(&iComment)
	_int64, _error := model.PutCommentViaLangcode(Langcode_, &iComment)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	var iComment model.Comment
	self.Bind(&iComment)
	var iMap = helper.StructToMap(iComment)
	_error := model.UpdateCommentViaCid(Cid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	var iComment model.Comment
	self.Bind(&iComment)
	var iMap = helper.StructToMap(iComment)
	_error := model.UpdateCommentViaCommentType(CommentType_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iComment model.Comment
	self.Bind(&iComment)
	var iMap = helper.StructToMap(iComment)
	_error := model.UpdateCommentViaUuid(Uuid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iComment model.Comment
	self.Bind(&iComment)
	var iMap = helper.StructToMap(iComment)
	_error := model.UpdateCommentViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	_int64, _error := model.DeleteComment(Cid_)
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

func DeleteCommentViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	_error := model.DeleteCommentViaCid(Cid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	_error := model.DeleteCommentViaCommentType(CommentType_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	_error := model.DeleteCommentViaUuid(Uuid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteCommentViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
