package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetUser_userPicturesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetUser_userPicturesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["user__user_picturesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetUser_userPicturesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetUser_userPictureCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetUser_userPictureCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetUser_userPictureCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetUser_userPictureCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetUser_userPictureCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetUser_userPictureCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaUserPictureTargetIdHandler(self *macross.Context) error {
	UserPictureTargetId_ := self.Args("user_picture_target_id").MustInt()
	_int64 := model.GetUser_userPictureCountViaUserPictureTargetId(UserPictureTargetId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaUserPictureAltHandler(self *macross.Context) error {
	UserPictureAlt_ := self.Args("user_picture_alt").String()
	_int64 := model.GetUser_userPictureCountViaUserPictureAlt(UserPictureAlt_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaUserPictureTitleHandler(self *macross.Context) error {
	UserPictureTitle_ := self.Args("user_picture_title").String()
	_int64 := model.GetUser_userPictureCountViaUserPictureTitle(UserPictureTitle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaUserPictureWidthHandler(self *macross.Context) error {
	UserPictureWidth_ := self.Args("user_picture_width").MustInt()
	_int64 := model.GetUser_userPictureCountViaUserPictureWidth(UserPictureWidth_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPictureCountViaUserPictureHeightHandler(self *macross.Context) error {
	UserPictureHeight_ := self.Args("user_picture_height").MustInt()
	_int64 := model.GetUser_userPictureCountViaUserPictureHeight(UserPictureHeight_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["user__user_pictureCount"] = 0
	}
	m["user__user_pictureCount"] = _int64
	return self.JSON(m)
}

func GetUser_userPicturesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	if (offset > 0) && helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesViaUserPictureTargetId(offset, limit, iUserPictureTargetId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	if (offset > 0) && helper.IsHas(iUserPictureAlt) {
		_User_userPicture, _error := model.GetUser_userPicturesViaUserPictureAlt(offset, limit, iUserPictureAlt, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserPictureTitle := self.Args("user_picture_title").String()
	if (offset > 0) && helper.IsHas(iUserPictureTitle) {
		_User_userPicture, _error := model.GetUser_userPicturesViaUserPictureTitle(offset, limit, iUserPictureTitle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	if (offset > 0) && helper.IsHas(iUserPictureWidth) {
		_User_userPicture, _error := model.GetUser_userPicturesViaUserPictureWidth(offset, limit, iUserPictureWidth, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesViaUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()
	if (offset > 0) && helper.IsHas(iUserPictureHeight) {
		_User_userPicture, _error := model.GetUser_userPicturesViaUserPictureHeight(offset, limit, iUserPictureHeight, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesViaUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndUserPictureTargetId(offset, limit, iBundle,iDeleted,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndUserPictureAlt(offset, limit, iBundle,iDeleted,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndUserPictureTitle(offset, limit, iBundle,iDeleted,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndUserPictureWidth(offset, limit, iBundle,iDeleted,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeletedAndUserPictureHeight(offset, limit, iBundle,iDeleted,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeletedAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityIdAndUserPictureTargetId(offset, limit, iBundle,iEntityId,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityIdAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityIdAndUserPictureAlt(offset, limit, iBundle,iEntityId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityIdAndUserPictureTitle(offset, limit, iBundle,iEntityId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityIdAndUserPictureWidth(offset, limit, iBundle,iEntityId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityIdAndUserPictureHeight(offset, limit, iBundle,iEntityId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTargetId(offset, limit, iBundle,iRevisionId,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndRevisionIdAndUserPictureAlt(offset, limit, iBundle,iRevisionId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndRevisionIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTitle(offset, limit, iBundle,iRevisionId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndRevisionIdAndUserPictureWidth(offset, limit, iBundle,iRevisionId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndRevisionIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndRevisionIdAndUserPictureHeight(offset, limit, iBundle,iRevisionId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndRevisionIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndLangcodeAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndLangcodeAndUserPictureTargetId(offset, limit, iBundle,iLangcode,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndLangcodeAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndLangcodeAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndLangcodeAndUserPictureAlt(offset, limit, iBundle,iLangcode,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndLangcodeAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndLangcodeAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndLangcodeAndUserPictureTitle(offset, limit, iBundle,iLangcode,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndLangcodeAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndLangcodeAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndLangcodeAndUserPictureWidth(offset, limit, iBundle,iLangcode,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndLangcodeAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndLangcodeAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndLangcodeAndUserPictureHeight(offset, limit, iBundle,iLangcode,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndLangcodeAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeltaAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeltaAndUserPictureTargetId(offset, limit, iBundle,iDelta,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeltaAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeltaAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeltaAndUserPictureAlt(offset, limit, iBundle,iDelta,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeltaAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeltaAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeltaAndUserPictureTitle(offset, limit, iBundle,iDelta,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeltaAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeltaAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeltaAndUserPictureWidth(offset, limit, iBundle,iDelta,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeltaAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeltaAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeltaAndUserPictureHeight(offset, limit, iBundle,iDelta,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeltaAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureAlt(offset, limit, iBundle,iUserPictureTargetId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureTitle(offset, limit, iBundle,iUserPictureTargetId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureWidth(offset, limit, iBundle,iUserPictureTargetId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureHeight(offset, limit, iBundle,iUserPictureTargetId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureTitle(offset, limit, iBundle,iUserPictureAlt,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureWidth(offset, limit, iBundle,iUserPictureAlt,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureHeight(offset, limit, iBundle,iUserPictureAlt,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureWidth(offset, limit, iBundle,iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureHeight(offset, limit, iBundle,iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureWidthAndUserPictureHeight(offset, limit, iBundle,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTargetId(offset, limit, iDeleted,iEntityId,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityIdAndUserPictureAlt(offset, limit, iDeleted,iEntityId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTitle(offset, limit, iDeleted,iEntityId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityIdAndUserPictureWidth(offset, limit, iDeleted,iEntityId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityIdAndUserPictureHeight(offset, limit, iDeleted,iEntityId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTargetId(offset, limit, iDeleted,iRevisionId,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureAlt(offset, limit, iDeleted,iRevisionId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTitle(offset, limit, iDeleted,iRevisionId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureWidth(offset, limit, iDeleted,iRevisionId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureHeight(offset, limit, iDeleted,iRevisionId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTargetId(offset, limit, iDeleted,iLangcode,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndLangcodeAndUserPictureAlt(offset, limit, iDeleted,iLangcode,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndLangcodeAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTitle(offset, limit, iDeleted,iLangcode,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndLangcodeAndUserPictureWidth(offset, limit, iDeleted,iLangcode,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndLangcodeAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndLangcodeAndUserPictureHeight(offset, limit, iDeleted,iLangcode,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndLangcodeAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndDeltaAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndDeltaAndUserPictureTargetId(offset, limit, iDeleted,iDelta,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndDeltaAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndDeltaAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndDeltaAndUserPictureAlt(offset, limit, iDeleted,iDelta,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndDeltaAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndDeltaAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndDeltaAndUserPictureTitle(offset, limit, iDeleted,iDelta,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndDeltaAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndDeltaAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndDeltaAndUserPictureWidth(offset, limit, iDeleted,iDelta,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndDeltaAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndDeltaAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndDeltaAndUserPictureHeight(offset, limit, iDeleted,iDelta,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndDeltaAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureAlt(offset, limit, iDeleted,iUserPictureTargetId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureTitle(offset, limit, iDeleted,iUserPictureTargetId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureWidth(offset, limit, iDeleted,iUserPictureTargetId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureHeight(offset, limit, iDeleted,iUserPictureTargetId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureTitle(offset, limit, iDeleted,iUserPictureAlt,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureWidth(offset, limit, iDeleted,iUserPictureAlt,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureHeight(offset, limit, iDeleted,iUserPictureAlt,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureWidth(offset, limit, iDeleted,iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureHeight(offset, limit, iDeleted,iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureWidthAndUserPictureHeight(offset, limit, iDeleted,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTargetId(offset, limit, iEntityId,iRevisionId,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureAlt(offset, limit, iEntityId,iRevisionId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTitle(offset, limit, iEntityId,iRevisionId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureWidth(offset, limit, iEntityId,iRevisionId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureHeight(offset, limit, iEntityId,iRevisionId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTargetId(offset, limit, iEntityId,iLangcode,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureAlt(offset, limit, iEntityId,iLangcode,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTitle(offset, limit, iEntityId,iLangcode,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureWidth(offset, limit, iEntityId,iLangcode,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureHeight(offset, limit, iEntityId,iLangcode,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTargetId(offset, limit, iEntityId,iDelta,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndDeltaAndUserPictureAlt(offset, limit, iEntityId,iDelta,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndDeltaAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTitle(offset, limit, iEntityId,iDelta,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndDeltaAndUserPictureWidth(offset, limit, iEntityId,iDelta,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndDeltaAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndDeltaAndUserPictureHeight(offset, limit, iEntityId,iDelta,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndDeltaAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureAlt(offset, limit, iEntityId,iUserPictureTargetId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureTitle(offset, limit, iEntityId,iUserPictureTargetId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureWidth(offset, limit, iEntityId,iUserPictureTargetId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureHeight(offset, limit, iEntityId,iUserPictureTargetId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureTitle(offset, limit, iEntityId,iUserPictureAlt,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureWidth(offset, limit, iEntityId,iUserPictureAlt,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureHeight(offset, limit, iEntityId,iUserPictureAlt,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureWidth(offset, limit, iEntityId,iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureHeight(offset, limit, iEntityId,iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureWidthAndUserPictureHeight(offset, limit, iEntityId,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTargetId(offset, limit, iRevisionId,iLangcode,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureAlt(offset, limit, iRevisionId,iLangcode,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTitle(offset, limit, iRevisionId,iLangcode,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureWidth(offset, limit, iRevisionId,iLangcode,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureHeight(offset, limit, iRevisionId,iLangcode,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTargetId(offset, limit, iRevisionId,iDelta,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureAlt(offset, limit, iRevisionId,iDelta,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTitle(offset, limit, iRevisionId,iDelta,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureWidth(offset, limit, iRevisionId,iDelta,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureHeight(offset, limit, iRevisionId,iDelta,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureAlt(offset, limit, iRevisionId,iUserPictureTargetId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureTitle(offset, limit, iRevisionId,iUserPictureTargetId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureWidth(offset, limit, iRevisionId,iUserPictureTargetId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureHeight(offset, limit, iRevisionId,iUserPictureTargetId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureTitle(offset, limit, iRevisionId,iUserPictureAlt,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureWidth(offset, limit, iRevisionId,iUserPictureAlt,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureHeight(offset, limit, iRevisionId,iUserPictureAlt,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureWidth(offset, limit, iRevisionId,iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureHeight(offset, limit, iRevisionId,iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureWidthAndUserPictureHeight(offset, limit, iRevisionId,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTargetId(offset, limit, iLangcode,iDelta,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndDeltaAndUserPictureAlt(offset, limit, iLangcode,iDelta,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndDeltaAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTitle(offset, limit, iLangcode,iDelta,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndDeltaAndUserPictureWidth(offset, limit, iLangcode,iDelta,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndDeltaAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndDeltaAndUserPictureHeight(offset, limit, iLangcode,iDelta,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndDeltaAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureAlt(offset, limit, iLangcode,iUserPictureTargetId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureTitle(offset, limit, iLangcode,iUserPictureTargetId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureWidth(offset, limit, iLangcode,iUserPictureTargetId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureHeight(offset, limit, iLangcode,iUserPictureTargetId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureTitle(offset, limit, iLangcode,iUserPictureAlt,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureWidth(offset, limit, iLangcode,iUserPictureAlt,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureHeight(offset, limit, iLangcode,iUserPictureAlt,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureWidth(offset, limit, iLangcode,iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureHeight(offset, limit, iLangcode,iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureWidthAndUserPictureHeight(offset, limit, iLangcode,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureAlt(offset, limit, iDelta,iUserPictureTargetId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureTitle(offset, limit, iDelta,iUserPictureTargetId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureWidth(offset, limit, iDelta,iUserPictureTargetId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureHeight(offset, limit, iDelta,iUserPictureTargetId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureTitle(offset, limit, iDelta,iUserPictureAlt,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureWidth(offset, limit, iDelta,iUserPictureAlt,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureHeight(offset, limit, iDelta,iUserPictureAlt,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureWidth(offset, limit, iDelta,iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureHeight(offset, limit, iDelta,iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureWidthAndUserPictureHeight(offset, limit, iDelta,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureTitle(offset, limit, iUserPictureTargetId,iUserPictureAlt,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureWidth(offset, limit, iUserPictureTargetId,iUserPictureAlt,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureHeight(offset, limit, iUserPictureTargetId,iUserPictureAlt,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureWidth(offset, limit, iUserPictureTargetId,iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureHeight(offset, limit, iUserPictureTargetId,iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidthAndUserPictureHeight(offset, limit, iUserPictureTargetId,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iUserPictureAlt) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureWidth(offset, limit, iUserPictureAlt,iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureAlt) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureHeight(offset, limit, iUserPictureAlt,iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureAltAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureAlt) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureAltAndUserPictureWidthAndUserPictureHeight(offset, limit, iUserPictureAlt,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureAltAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTitleAndUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureTitle) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTitleAndUserPictureWidthAndUserPictureHeight(offset, limit, iUserPictureTitle,iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTitleAndUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureTargetId(offset, limit, iBundle,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureAlt(offset, limit, iBundle,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureTitle(offset, limit, iBundle,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureWidth(offset, limit, iBundle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByBundleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPicturesByBundleAndUserPictureHeight(offset, limit, iBundle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByBundleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureTargetId(offset, limit, iDeleted,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureAlt(offset, limit, iDeleted,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureTitle(offset, limit, iDeleted,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureWidth(offset, limit, iDeleted,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeletedAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeletedAndUserPictureHeight(offset, limit, iDeleted,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeletedAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureTargetId(offset, limit, iEntityId,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureAlt(offset, limit, iEntityId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureTitle(offset, limit, iEntityId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureWidth(offset, limit, iEntityId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByEntityIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPicturesByEntityIdAndUserPictureHeight(offset, limit, iEntityId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByEntityIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureTargetId(offset, limit, iRevisionId,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureAlt(offset, limit, iRevisionId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureTitle(offset, limit, iRevisionId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureWidth(offset, limit, iRevisionId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByRevisionIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPicturesByRevisionIdAndUserPictureHeight(offset, limit, iRevisionId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByRevisionIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureTargetId(offset, limit, iLangcode,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureAlt(offset, limit, iLangcode,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureTitle(offset, limit, iLangcode,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureWidth(offset, limit, iLangcode,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByLangcodeAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPicturesByLangcodeAndUserPictureHeight(offset, limit, iLangcode,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByLangcodeAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureTargetId(offset, limit, iDelta,iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureAlt(offset, limit, iDelta,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureTitle(offset, limit, iDelta,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureWidth(offset, limit, iDelta,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByDeltaAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPicturesByDeltaAndUserPictureHeight(offset, limit, iDelta,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByDeltaAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureAlt(offset, limit, iUserPictureTargetId,iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitle(offset, limit, iUserPictureTargetId,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidth(offset, limit, iUserPictureTargetId,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTargetIdAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTargetIdAndUserPictureHeight(offset, limit, iUserPictureTargetId,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTargetIdAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureAltAndUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureTitle := self.Args("user_picture_title").String()

	if helper.IsHas(iUserPictureAlt) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureAltAndUserPictureTitle(offset, limit, iUserPictureAlt,iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureAltAndUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureAltAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iUserPictureAlt) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureAltAndUserPictureWidth(offset, limit, iUserPictureAlt,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureAltAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureAltAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureAlt := self.Args("user_picture_alt").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureAlt) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureAltAndUserPictureHeight(offset, limit, iUserPictureAlt,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureAltAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTitleAndUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()

	if helper.IsHas(iUserPictureTitle) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTitleAndUserPictureWidth(offset, limit, iUserPictureTitle,iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTitleAndUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureTitleAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureTitle := self.Args("user_picture_title").String()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureTitle) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureTitleAndUserPictureHeight(offset, limit, iUserPictureTitle,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureTitleAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesByUserPictureWidthAndUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	iUserPictureHeight := self.Args("user_picture_height").MustInt()

	if helper.IsHas(iUserPictureWidth) {
		_User_userPicture, _error := model.GetUser_userPicturesByUserPictureWidthAndUserPictureHeight(offset, limit, iUserPictureWidth,iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPicturesByUserPictureWidthAndUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPicturesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_User_userPicture, _error := model.GetUser_userPictures(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictures' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_User_userPicture := model.HasUser_userPictureViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_User_userPicture := model.HasUser_userPictureViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_User_userPicture := model.HasUser_userPictureViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_User_userPicture := model.HasUser_userPictureViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_User_userPicture := model.HasUser_userPictureViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_User_userPicture := model.HasUser_userPictureViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture := model.HasUser_userPictureViaUserPictureTargetId(iUserPictureTargetId)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureAlt := self.Args("user_picture_alt").String()
	if helper.IsHas(iUserPictureAlt) {
		_User_userPicture := model.HasUser_userPictureViaUserPictureAlt(iUserPictureAlt)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureTitle := self.Args("user_picture_title").String()
	if helper.IsHas(iUserPictureTitle) {
		_User_userPicture := model.HasUser_userPictureViaUserPictureTitle(iUserPictureTitle)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	if helper.IsHas(iUserPictureWidth) {
		_User_userPicture := model.HasUser_userPictureViaUserPictureWidth(iUserPictureWidth)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUser_userPictureViaUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureHeight := self.Args("user_picture_height").MustInt()
	if helper.IsHas(iUserPictureHeight) {
		_User_userPicture := model.HasUser_userPictureViaUserPictureHeight(iUserPictureHeight)
		var m = map[string]interface{}{}
		m["user__user_picture"] = _User_userPicture
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUser_userPictureViaUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_User_userPicture, _error := model.GetUser_userPictureViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_User_userPicture, _error := model.GetUser_userPictureViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_User_userPicture, _error := model.GetUser_userPictureViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_User_userPicture, _error := model.GetUser_userPictureViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_User_userPicture, _error := model.GetUser_userPictureViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_User_userPicture, _error := model.GetUser_userPictureViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureTargetId := self.Args("user_picture_target_id").MustInt()
	if helper.IsHas(iUserPictureTargetId) {
		_User_userPicture, _error := model.GetUser_userPictureViaUserPictureTargetId(iUserPictureTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureAlt := self.Args("user_picture_alt").String()
	if helper.IsHas(iUserPictureAlt) {
		_User_userPicture, _error := model.GetUser_userPictureViaUserPictureAlt(iUserPictureAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureTitle := self.Args("user_picture_title").String()
	if helper.IsHas(iUserPictureTitle) {
		_User_userPicture, _error := model.GetUser_userPictureViaUserPictureTitle(iUserPictureTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureWidth := self.Args("user_picture_width").MustInt()
	if helper.IsHas(iUserPictureWidth) {
		_User_userPicture, _error := model.GetUser_userPictureViaUserPictureWidth(iUserPictureWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUser_userPictureViaUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPictureHeight := self.Args("user_picture_height").MustInt()
	if helper.IsHas(iUserPictureHeight) {
		_User_userPicture, _error := model.GetUser_userPictureViaUserPictureHeight(iUserPictureHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the GetUser_userPictureViaUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaBundle(Bundle_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaDeleted(Deleted_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaEntityId(EntityId_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaRevisionId(RevisionId_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaLangcode(Langcode_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaDelta(Delta_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureTargetId_ := self.Args("user_picture_target_id").MustInt()
	if helper.IsHas(UserPictureTargetId_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaUserPictureTargetId(UserPictureTargetId_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaUserPictureTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureAlt_ := self.Args("user_picture_alt").String()
	if helper.IsHas(UserPictureAlt_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaUserPictureAlt(UserPictureAlt_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaUserPictureAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureTitle_ := self.Args("user_picture_title").String()
	if helper.IsHas(UserPictureTitle_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaUserPictureTitle(UserPictureTitle_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaUserPictureTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureWidth_ := self.Args("user_picture_width").MustInt()
	if helper.IsHas(UserPictureWidth_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaUserPictureWidth(UserPictureWidth_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaUserPictureWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUser_userPictureViaUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureHeight_ := self.Args("user_picture_height").MustInt()
	if helper.IsHas(UserPictureHeight_) {
		var iUser_userPicture model.User_userPicture
		self.Bind(&iUser_userPicture)
		_User_userPicture, _error := model.SetUser_userPictureViaUserPictureHeight(UserPictureHeight_, &iUser_userPicture)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_User_userPicture)
	}
	herr.Message = "Can't get to the SetUser_userPictureViaUserPictureHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddUser_userPictureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	Deleted_ := self.Args("deleted").MustInt()
	EntityId_ := self.Args("entity_id").MustInt()
	RevisionId_ := self.Args("revision_id").MustInt()
	Langcode_ := self.Args("langcode").String()
	Delta_ := self.Args("delta").MustInt()
	UserPictureTargetId_ := self.Args("user_picture_target_id").MustInt()
	UserPictureAlt_ := self.Args("user_picture_alt").String()
	UserPictureTitle_ := self.Args("user_picture_title").String()
	UserPictureWidth_ := self.Args("user_picture_width").MustInt()
	UserPictureHeight_ := self.Args("user_picture_height").MustInt()

	if helper.IsHas(Bundle_) {
		_error := model.AddUser_userPicture(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,UserPictureTargetId_,UserPictureAlt_,UserPictureTitle_,UserPictureWidth_,UserPictureHeight_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddUser_userPicture's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostUser_userPictureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_string, _error := model.PostUser_userPicture(&iUser_userPicture)
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

func PutUser_userPictureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_string, _error := model.PutUser_userPicture(&iUser_userPicture)
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

func PutUser_userPictureViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaBundle(Bundle_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaDeleted(Deleted_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaEntityId(EntityId_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaRevisionId(RevisionId_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaLangcode(Langcode_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaDelta(Delta_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureTargetId_ := self.Args("user_picture_target_id").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaUserPictureTargetId(UserPictureTargetId_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureAlt_ := self.Args("user_picture_alt").String()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaUserPictureAlt(UserPictureAlt_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureTitle_ := self.Args("user_picture_title").String()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaUserPictureTitle(UserPictureTitle_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureWidth_ := self.Args("user_picture_width").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaUserPictureWidth(UserPictureWidth_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUser_userPictureViaUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureHeight_ := self.Args("user_picture_height").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	_int64, _error := model.PutUser_userPictureViaUserPictureHeight(UserPictureHeight_, &iUser_userPicture)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureTargetId_ := self.Args("user_picture_target_id").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaUserPictureTargetId(UserPictureTargetId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureAlt_ := self.Args("user_picture_alt").String()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaUserPictureAlt(UserPictureAlt_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureTitle_ := self.Args("user_picture_title").String()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaUserPictureTitle(UserPictureTitle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureWidth_ := self.Args("user_picture_width").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaUserPictureWidth(UserPictureWidth_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUser_userPictureViaUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureHeight_ := self.Args("user_picture_height").MustInt()
	var iUser_userPicture model.User_userPicture
	self.Bind(&iUser_userPicture)
	var iMap = helper.StructToMap(iUser_userPicture)
	_error := model.UpdateUser_userPictureViaUserPictureHeight(UserPictureHeight_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteUser_userPicture(Bundle_)
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

func DeleteUser_userPictureViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteUser_userPictureViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteUser_userPictureViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteUser_userPictureViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteUser_userPictureViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteUser_userPictureViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteUser_userPictureViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaUserPictureTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureTargetId_ := self.Args("user_picture_target_id").MustInt()
	_error := model.DeleteUser_userPictureViaUserPictureTargetId(UserPictureTargetId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaUserPictureAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureAlt_ := self.Args("user_picture_alt").String()
	_error := model.DeleteUser_userPictureViaUserPictureAlt(UserPictureAlt_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaUserPictureTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureTitle_ := self.Args("user_picture_title").String()
	_error := model.DeleteUser_userPictureViaUserPictureTitle(UserPictureTitle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaUserPictureWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureWidth_ := self.Args("user_picture_width").MustInt()
	_error := model.DeleteUser_userPictureViaUserPictureWidth(UserPictureWidth_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUser_userPictureViaUserPictureHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPictureHeight_ := self.Args("user_picture_height").MustInt()
	_error := model.DeleteUser_userPictureViaUserPictureHeight(UserPictureHeight_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
