package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetNode_fieldImagesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNode_fieldImagesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node__field_imagesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetNode_fieldImageCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetNode_fieldImageCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetNode_fieldImageCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetNode_fieldImageCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNode_fieldImageCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetNode_fieldImageCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaFieldImageTargetIdHandler(self *macross.Context) error {
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	_int64 := model.GetNode_fieldImageCountViaFieldImageTargetId(FieldImageTargetId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaFieldImageAltHandler(self *macross.Context) error {
	FieldImageAlt_ := self.Args("field_image_alt").String()
	_int64 := model.GetNode_fieldImageCountViaFieldImageAlt(FieldImageAlt_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaFieldImageTitleHandler(self *macross.Context) error {
	FieldImageTitle_ := self.Args("field_image_title").String()
	_int64 := model.GetNode_fieldImageCountViaFieldImageTitle(FieldImageTitle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaFieldImageWidthHandler(self *macross.Context) error {
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	_int64 := model.GetNode_fieldImageCountViaFieldImageWidth(FieldImageWidth_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImageCountViaFieldImageHeightHandler(self *macross.Context) error {
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	_int64 := model.GetNode_fieldImageCountViaFieldImageHeight(FieldImageHeight_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_imageCount"] = 0
	}
	m["node__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldImagesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	if (offset > 0) && helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaFieldImageTargetId(offset, limit, iFieldImageTargetId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	if (offset > 0) && helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaFieldImageAlt(offset, limit, iFieldImageAlt, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	if (offset > 0) && helper.IsHas(iFieldImageTitle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaFieldImageTitle(offset, limit, iFieldImageTitle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	if (offset > 0) && helper.IsHas(iFieldImageWidth) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaFieldImageWidth(offset, limit, iFieldImageWidth, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()
	if (offset > 0) && helper.IsHas(iFieldImageHeight) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesViaFieldImageHeight(offset, limit, iFieldImageHeight, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesViaFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndFieldImageTargetId(offset, limit, iBundle,iDeleted,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndFieldImageAlt(offset, limit, iBundle,iDeleted,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndFieldImageTitle(offset, limit, iBundle,iDeleted,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndFieldImageWidth(offset, limit, iBundle,iDeleted,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeletedAndFieldImageHeight(offset, limit, iBundle,iDeleted,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeletedAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTargetId(offset, limit, iBundle,iEntityId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityIdAndFieldImageAlt(offset, limit, iBundle,iEntityId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTitle(offset, limit, iBundle,iEntityId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityIdAndFieldImageWidth(offset, limit, iBundle,iEntityId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityIdAndFieldImageHeight(offset, limit, iBundle,iEntityId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTargetId(offset, limit, iBundle,iRevisionId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageAlt(offset, limit, iBundle,iRevisionId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTitle(offset, limit, iBundle,iRevisionId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageWidth(offset, limit, iBundle,iRevisionId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageHeight(offset, limit, iBundle,iRevisionId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTargetId(offset, limit, iBundle,iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndLangcodeAndFieldImageAlt(offset, limit, iBundle,iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTitle(offset, limit, iBundle,iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndLangcodeAndFieldImageWidth(offset, limit, iBundle,iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndLangcodeAndFieldImageHeight(offset, limit, iBundle,iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeltaAndFieldImageTargetId(offset, limit, iBundle,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeltaAndFieldImageAlt(offset, limit, iBundle,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeltaAndFieldImageTitle(offset, limit, iBundle,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeltaAndFieldImageWidth(offset, limit, iBundle,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeltaAndFieldImageHeight(offset, limit, iBundle,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iBundle,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iBundle,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iBundle,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iBundle,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageTitle(offset, limit, iBundle,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageWidth(offset, limit, iBundle,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageHeight(offset, limit, iBundle,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidth(offset, limit, iBundle,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeight(offset, limit, iBundle,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeight(offset, limit, iBundle,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTargetId(offset, limit, iDeleted,iEntityId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageAlt(offset, limit, iDeleted,iEntityId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTitle(offset, limit, iDeleted,iEntityId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageWidth(offset, limit, iDeleted,iEntityId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageHeight(offset, limit, iDeleted,iEntityId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetId(offset, limit, iDeleted,iRevisionId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageAlt(offset, limit, iDeleted,iRevisionId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTitle(offset, limit, iDeleted,iRevisionId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageWidth(offset, limit, iDeleted,iRevisionId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageHeight(offset, limit, iDeleted,iRevisionId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTargetId(offset, limit, iDeleted,iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageAlt(offset, limit, iDeleted,iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTitle(offset, limit, iDeleted,iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageWidth(offset, limit, iDeleted,iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageHeight(offset, limit, iDeleted,iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTargetId(offset, limit, iDeleted,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndDeltaAndFieldImageAlt(offset, limit, iDeleted,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTitle(offset, limit, iDeleted,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndDeltaAndFieldImageWidth(offset, limit, iDeleted,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndDeltaAndFieldImageHeight(offset, limit, iDeleted,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iDeleted,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iDeleted,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iDeleted,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iDeleted,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitle(offset, limit, iDeleted,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidth(offset, limit, iDeleted,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeight(offset, limit, iDeleted,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidth(offset, limit, iDeleted,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeight(offset, limit, iDeleted,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeight(offset, limit, iDeleted,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetId(offset, limit, iEntityId,iRevisionId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageAlt(offset, limit, iEntityId,iRevisionId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitle(offset, limit, iEntityId,iRevisionId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidth(offset, limit, iEntityId,iRevisionId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeight(offset, limit, iEntityId,iRevisionId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetId(offset, limit, iEntityId,iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageAlt(offset, limit, iEntityId,iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTitle(offset, limit, iEntityId,iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageWidth(offset, limit, iEntityId,iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageHeight(offset, limit, iEntityId,iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTargetId(offset, limit, iEntityId,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageAlt(offset, limit, iEntityId,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTitle(offset, limit, iEntityId,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageWidth(offset, limit, iEntityId,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageHeight(offset, limit, iEntityId,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iEntityId,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iEntityId,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iEntityId,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iEntityId,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitle(offset, limit, iEntityId,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidth(offset, limit, iEntityId,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeight(offset, limit, iEntityId,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidth(offset, limit, iEntityId,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeight(offset, limit, iEntityId,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeight(offset, limit, iEntityId,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetId(offset, limit, iRevisionId,iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageAlt(offset, limit, iRevisionId,iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitle(offset, limit, iRevisionId,iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidth(offset, limit, iRevisionId,iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeight(offset, limit, iRevisionId,iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetId(offset, limit, iRevisionId,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageAlt(offset, limit, iRevisionId,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTitle(offset, limit, iRevisionId,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageWidth(offset, limit, iRevisionId,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageHeight(offset, limit, iRevisionId,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iRevisionId,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iRevisionId,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iRevisionId,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitle(offset, limit, iRevisionId,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidth(offset, limit, iRevisionId,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidth(offset, limit, iRevisionId,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTargetId(offset, limit, iLangcode,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageAlt(offset, limit, iLangcode,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTitle(offset, limit, iLangcode,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageWidth(offset, limit, iLangcode,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageHeight(offset, limit, iLangcode,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iLangcode,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iLangcode,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iLangcode,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iLangcode,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitle(offset, limit, iLangcode,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidth(offset, limit, iLangcode,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeight(offset, limit, iLangcode,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidth(offset, limit, iLangcode,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeight(offset, limit, iLangcode,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeight(offset, limit, iLangcode,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iDelta,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iDelta,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iDelta,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iDelta,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitle(offset, limit, iDelta,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidth(offset, limit, iDelta,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeight(offset, limit, iDelta,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidth(offset, limit, iDelta,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeight(offset, limit, iDelta,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeight(offset, limit, iDelta,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitle(offset, limit, iFieldImageTargetId,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidth(offset, limit, iFieldImageTargetId,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeight(offset, limit, iFieldImageTargetId,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidth(offset, limit, iFieldImageTargetId,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeight(offset, limit, iFieldImageTargetId,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeight(offset, limit, iFieldImageTargetId,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidth(offset, limit, iFieldImageAlt,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeight(offset, limit, iFieldImageAlt,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeight(offset, limit, iFieldImageAlt,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTitle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeight(offset, limit, iFieldImageTitle,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageTargetId(offset, limit, iBundle,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageAlt(offset, limit, iBundle,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageTitle(offset, limit, iBundle,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageWidth(offset, limit, iBundle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByBundleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByBundleAndFieldImageHeight(offset, limit, iBundle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByBundleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageTargetId(offset, limit, iDeleted,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageAlt(offset, limit, iDeleted,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageTitle(offset, limit, iDeleted,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageWidth(offset, limit, iDeleted,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeletedAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeletedAndFieldImageHeight(offset, limit, iDeleted,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeletedAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageTargetId(offset, limit, iEntityId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageAlt(offset, limit, iEntityId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageTitle(offset, limit, iEntityId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageWidth(offset, limit, iEntityId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByEntityIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByEntityIdAndFieldImageHeight(offset, limit, iEntityId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByEntityIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageTargetId(offset, limit, iRevisionId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageAlt(offset, limit, iRevisionId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageTitle(offset, limit, iRevisionId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageWidth(offset, limit, iRevisionId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByRevisionIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByRevisionIdAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByRevisionIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageTargetId(offset, limit, iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageAlt(offset, limit, iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageTitle(offset, limit, iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageWidth(offset, limit, iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByLangcodeAndFieldImageHeight(offset, limit, iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageTargetId(offset, limit, iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageAlt(offset, limit, iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageTitle(offset, limit, iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageWidth(offset, limit, iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByDeltaAndFieldImageHeight(offset, limit, iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAlt(offset, limit, iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitle(offset, limit, iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidth(offset, limit, iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTargetIdAndFieldImageHeight(offset, limit, iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageAltAndFieldImageTitle(offset, limit, iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageAltAndFieldImageWidth(offset, limit, iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageAltAndFieldImageHeight(offset, limit, iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageTitle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTitleAndFieldImageWidth(offset, limit, iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTitle) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageTitleAndFieldImageHeight(offset, limit, iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesByFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageWidth) {
		_Node_fieldImage, _error := model.GetNode_fieldImagesByFieldImageWidthAndFieldImageHeight(offset, limit, iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImagesByFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImagesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Node_fieldImage, _error := model.GetNode_fieldImages(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImages' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_Node_fieldImage := model.HasNode_fieldImageViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Node_fieldImage := model.HasNode_fieldImageViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_Node_fieldImage := model.HasNode_fieldImageViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_Node_fieldImage := model.HasNode_fieldImageViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Node_fieldImage := model.HasNode_fieldImageViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_Node_fieldImage := model.HasNode_fieldImageViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage := model.HasNode_fieldImageViaFieldImageTargetId(iFieldImageTargetId)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageAlt := self.Args("field_image_alt").String()
	if helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage := model.HasNode_fieldImageViaFieldImageAlt(iFieldImageAlt)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageTitle := self.Args("field_image_title").String()
	if helper.IsHas(iFieldImageTitle) {
		_Node_fieldImage := model.HasNode_fieldImageViaFieldImageTitle(iFieldImageTitle)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	if helper.IsHas(iFieldImageWidth) {
		_Node_fieldImage := model.HasNode_fieldImageViaFieldImageWidth(iFieldImageWidth)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageHeight := self.Args("field_image_height").MustInt()
	if helper.IsHas(iFieldImageHeight) {
		_Node_fieldImage := model.HasNode_fieldImageViaFieldImageHeight(iFieldImageHeight)
		var m = map[string]interface{}{}
		m["node__field_image"] = _Node_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldImageViaFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	if helper.IsHas(iFieldImageTargetId) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaFieldImageTargetId(iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageAlt := self.Args("field_image_alt").String()
	if helper.IsHas(iFieldImageAlt) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaFieldImageAlt(iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageTitle := self.Args("field_image_title").String()
	if helper.IsHas(iFieldImageTitle) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaFieldImageTitle(iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	if helper.IsHas(iFieldImageWidth) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaFieldImageWidth(iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageHeight := self.Args("field_image_height").MustInt()
	if helper.IsHas(iFieldImageHeight) {
		_Node_fieldImage, _error := model.GetNode_fieldImageViaFieldImageHeight(iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the GetNode_fieldImageViaFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaBundle(Bundle_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaDeleted(Deleted_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaEntityId(EntityId_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaRevisionId(RevisionId_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaLangcode(Langcode_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaDelta(Delta_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	if helper.IsHas(FieldImageTargetId_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaFieldImageTargetId(FieldImageTargetId_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageAlt_ := self.Args("field_image_alt").String()
	if helper.IsHas(FieldImageAlt_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaFieldImageAlt(FieldImageAlt_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTitle_ := self.Args("field_image_title").String()
	if helper.IsHas(FieldImageTitle_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaFieldImageTitle(FieldImageTitle_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	if helper.IsHas(FieldImageWidth_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaFieldImageWidth(FieldImageWidth_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	if helper.IsHas(FieldImageHeight_) {
		var iNode_fieldImage model.Node_fieldImage
		self.Bind(&iNode_fieldImage)
		_Node_fieldImage, _error := model.SetNode_fieldImageViaFieldImageHeight(FieldImageHeight_, &iNode_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldImage)
	}
	herr.Message = "Can't get to the SetNode_fieldImageViaFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNode_fieldImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	Deleted_ := self.Args("deleted").MustInt()
	EntityId_ := self.Args("entity_id").MustInt()
	RevisionId_ := self.Args("revision_id").MustInt()
	Langcode_ := self.Args("langcode").String()
	Delta_ := self.Args("delta").MustInt()
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	FieldImageAlt_ := self.Args("field_image_alt").String()
	FieldImageTitle_ := self.Args("field_image_title").String()
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	FieldImageHeight_ := self.Args("field_image_height").MustInt()

	if helper.IsHas(Bundle_) {
		_error := model.AddNode_fieldImage(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,FieldImageTargetId_,FieldImageAlt_,FieldImageTitle_,FieldImageWidth_,FieldImageHeight_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNode_fieldImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNode_fieldImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_string, _error := model.PostNode_fieldImage(&iNode_fieldImage)
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

func PutNode_fieldImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_string, _error := model.PutNode_fieldImage(&iNode_fieldImage)
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

func PutNode_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaBundle(Bundle_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaDeleted(Deleted_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaEntityId(EntityId_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaRevisionId(RevisionId_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaLangcode(Langcode_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaDelta(Delta_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaFieldImageTargetId(FieldImageTargetId_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageAlt_ := self.Args("field_image_alt").String()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaFieldImageAlt(FieldImageAlt_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTitle_ := self.Args("field_image_title").String()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaFieldImageTitle(FieldImageTitle_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaFieldImageWidth(FieldImageWidth_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	_int64, _error := model.PutNode_fieldImageViaFieldImageHeight(FieldImageHeight_, &iNode_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaFieldImageTargetId(FieldImageTargetId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageAlt_ := self.Args("field_image_alt").String()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaFieldImageAlt(FieldImageAlt_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTitle_ := self.Args("field_image_title").String()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaFieldImageTitle(FieldImageTitle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaFieldImageWidth(FieldImageWidth_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	var iNode_fieldImage model.Node_fieldImage
	self.Bind(&iNode_fieldImage)
	var iMap = helper.StructToMap(iNode_fieldImage)
	_error := model.UpdateNode_fieldImageViaFieldImageHeight(FieldImageHeight_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteNode_fieldImage(Bundle_)
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

func DeleteNode_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteNode_fieldImageViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteNode_fieldImageViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteNode_fieldImageViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteNode_fieldImageViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNode_fieldImageViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteNode_fieldImageViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	_error := model.DeleteNode_fieldImageViaFieldImageTargetId(FieldImageTargetId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageAlt_ := self.Args("field_image_alt").String()
	_error := model.DeleteNode_fieldImageViaFieldImageAlt(FieldImageAlt_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTitle_ := self.Args("field_image_title").String()
	_error := model.DeleteNode_fieldImageViaFieldImageTitle(FieldImageTitle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	_error := model.DeleteNode_fieldImageViaFieldImageWidth(FieldImageWidth_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	_error := model.DeleteNode_fieldImageViaFieldImageHeight(FieldImageHeight_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
