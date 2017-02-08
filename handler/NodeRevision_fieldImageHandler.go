package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetNodeRevision_fieldImagesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodeRevision_fieldImagesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node_revision__field_imagesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetNodeRevision_fieldImageCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetNodeRevision_fieldImageCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetNodeRevision_fieldImageCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetNodeRevision_fieldImageCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeRevision_fieldImageCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetNodeRevision_fieldImageCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaFieldImageTargetIdHandler(self *macross.Context) error {
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	_int64 := model.GetNodeRevision_fieldImageCountViaFieldImageTargetId(FieldImageTargetId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaFieldImageAltHandler(self *macross.Context) error {
	FieldImageAlt_ := self.Args("field_image_alt").String()
	_int64 := model.GetNodeRevision_fieldImageCountViaFieldImageAlt(FieldImageAlt_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaFieldImageTitleHandler(self *macross.Context) error {
	FieldImageTitle_ := self.Args("field_image_title").String()
	_int64 := model.GetNodeRevision_fieldImageCountViaFieldImageTitle(FieldImageTitle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaFieldImageWidthHandler(self *macross.Context) error {
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	_int64 := model.GetNodeRevision_fieldImageCountViaFieldImageWidth(FieldImageWidth_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImageCountViaFieldImageHeightHandler(self *macross.Context) error {
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	_int64 := model.GetNodeRevision_fieldImageCountViaFieldImageHeight(FieldImageHeight_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_imageCount"] = 0
	}
	m["node_revision__field_imageCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldImagesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	if (offset > 0) && helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaFieldImageTargetId(offset, limit, iFieldImageTargetId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	if (offset > 0) && helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaFieldImageAlt(offset, limit, iFieldImageAlt, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	if (offset > 0) && helper.IsHas(iFieldImageTitle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaFieldImageTitle(offset, limit, iFieldImageTitle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	if (offset > 0) && helper.IsHas(iFieldImageWidth) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaFieldImageWidth(offset, limit, iFieldImageWidth, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()
	if (offset > 0) && helper.IsHas(iFieldImageHeight) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesViaFieldImageHeight(offset, limit, iFieldImageHeight, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesViaFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTargetId(offset, limit, iBundle,iDeleted,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageAlt(offset, limit, iBundle,iDeleted,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTitle(offset, limit, iBundle,iDeleted,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageWidth(offset, limit, iBundle,iDeleted,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageHeight(offset, limit, iBundle,iDeleted,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTargetId(offset, limit, iBundle,iEntityId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageAlt(offset, limit, iBundle,iEntityId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTitle(offset, limit, iBundle,iEntityId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageWidth(offset, limit, iBundle,iEntityId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageHeight(offset, limit, iBundle,iEntityId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTargetId(offset, limit, iBundle,iRevisionId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageAlt(offset, limit, iBundle,iRevisionId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTitle(offset, limit, iBundle,iRevisionId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageWidth(offset, limit, iBundle,iRevisionId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageHeight(offset, limit, iBundle,iRevisionId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTargetId(offset, limit, iBundle,iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageAlt(offset, limit, iBundle,iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTitle(offset, limit, iBundle,iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageWidth(offset, limit, iBundle,iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageHeight(offset, limit, iBundle,iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTargetId(offset, limit, iBundle,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageAlt(offset, limit, iBundle,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTitle(offset, limit, iBundle,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageWidth(offset, limit, iBundle,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageHeight(offset, limit, iBundle,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iBundle,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iBundle,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iBundle,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iBundle,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageTitle(offset, limit, iBundle,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageWidth(offset, limit, iBundle,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageHeight(offset, limit, iBundle,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidth(offset, limit, iBundle,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeight(offset, limit, iBundle,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeight(offset, limit, iBundle,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTargetId(offset, limit, iDeleted,iEntityId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageAlt(offset, limit, iDeleted,iEntityId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTitle(offset, limit, iDeleted,iEntityId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageWidth(offset, limit, iDeleted,iEntityId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageHeight(offset, limit, iDeleted,iEntityId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetId(offset, limit, iDeleted,iRevisionId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageAlt(offset, limit, iDeleted,iRevisionId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTitle(offset, limit, iDeleted,iRevisionId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageWidth(offset, limit, iDeleted,iRevisionId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageHeight(offset, limit, iDeleted,iRevisionId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTargetId(offset, limit, iDeleted,iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageAlt(offset, limit, iDeleted,iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTitle(offset, limit, iDeleted,iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageWidth(offset, limit, iDeleted,iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageHeight(offset, limit, iDeleted,iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTargetId(offset, limit, iDeleted,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageAlt(offset, limit, iDeleted,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTitle(offset, limit, iDeleted,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageWidth(offset, limit, iDeleted,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageHeight(offset, limit, iDeleted,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iDeleted,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iDeleted,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iDeleted,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iDeleted,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitle(offset, limit, iDeleted,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidth(offset, limit, iDeleted,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeight(offset, limit, iDeleted,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidth(offset, limit, iDeleted,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeight(offset, limit, iDeleted,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeight(offset, limit, iDeleted,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetId(offset, limit, iEntityId,iRevisionId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageAlt(offset, limit, iEntityId,iRevisionId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitle(offset, limit, iEntityId,iRevisionId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidth(offset, limit, iEntityId,iRevisionId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeight(offset, limit, iEntityId,iRevisionId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetId(offset, limit, iEntityId,iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageAlt(offset, limit, iEntityId,iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTitle(offset, limit, iEntityId,iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageWidth(offset, limit, iEntityId,iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageHeight(offset, limit, iEntityId,iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTargetId(offset, limit, iEntityId,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageAlt(offset, limit, iEntityId,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTitle(offset, limit, iEntityId,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageWidth(offset, limit, iEntityId,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageHeight(offset, limit, iEntityId,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iEntityId,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iEntityId,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iEntityId,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iEntityId,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitle(offset, limit, iEntityId,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidth(offset, limit, iEntityId,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeight(offset, limit, iEntityId,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidth(offset, limit, iEntityId,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeight(offset, limit, iEntityId,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeight(offset, limit, iEntityId,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetId(offset, limit, iRevisionId,iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageAlt(offset, limit, iRevisionId,iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitle(offset, limit, iRevisionId,iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidth(offset, limit, iRevisionId,iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeight(offset, limit, iRevisionId,iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetId(offset, limit, iRevisionId,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageAlt(offset, limit, iRevisionId,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTitle(offset, limit, iRevisionId,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageWidth(offset, limit, iRevisionId,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageHeight(offset, limit, iRevisionId,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iRevisionId,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iRevisionId,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iRevisionId,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitle(offset, limit, iRevisionId,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidth(offset, limit, iRevisionId,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidth(offset, limit, iRevisionId,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTargetId(offset, limit, iLangcode,iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageAlt(offset, limit, iLangcode,iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTitle(offset, limit, iLangcode,iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageWidth(offset, limit, iLangcode,iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageHeight(offset, limit, iLangcode,iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iLangcode,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iLangcode,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iLangcode,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iLangcode,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitle(offset, limit, iLangcode,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidth(offset, limit, iLangcode,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeight(offset, limit, iLangcode,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidth(offset, limit, iLangcode,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeight(offset, limit, iLangcode,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeight(offset, limit, iLangcode,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAlt(offset, limit, iDelta,iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitle(offset, limit, iDelta,iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidth(offset, limit, iDelta,iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeight(offset, limit, iDelta,iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitle(offset, limit, iDelta,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidth(offset, limit, iDelta,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeight(offset, limit, iDelta,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidth(offset, limit, iDelta,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeight(offset, limit, iDelta,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeight(offset, limit, iDelta,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitle(offset, limit, iFieldImageTargetId,iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidth(offset, limit, iFieldImageTargetId,iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeight(offset, limit, iFieldImageTargetId,iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidth(offset, limit, iFieldImageTargetId,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeight(offset, limit, iFieldImageTargetId,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeight(offset, limit, iFieldImageTargetId,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidth(offset, limit, iFieldImageAlt,iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeight(offset, limit, iFieldImageAlt,iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeight(offset, limit, iFieldImageAlt,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTitle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeight(offset, limit, iFieldImageTitle,iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageTargetId(offset, limit, iBundle,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageAlt(offset, limit, iBundle,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageTitle(offset, limit, iBundle,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageWidth(offset, limit, iBundle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByBundleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByBundleAndFieldImageHeight(offset, limit, iBundle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByBundleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetId(offset, limit, iDeleted,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageAlt(offset, limit, iDeleted,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageTitle(offset, limit, iDeleted,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageWidth(offset, limit, iDeleted,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeletedAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeletedAndFieldImageHeight(offset, limit, iDeleted,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeletedAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetId(offset, limit, iEntityId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageAlt(offset, limit, iEntityId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitle(offset, limit, iEntityId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidth(offset, limit, iEntityId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByEntityIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByEntityIdAndFieldImageHeight(offset, limit, iEntityId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByEntityIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetId(offset, limit, iRevisionId,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAlt(offset, limit, iRevisionId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitle(offset, limit, iRevisionId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidth(offset, limit, iRevisionId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByRevisionIdAndFieldImageHeight(offset, limit, iRevisionId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByRevisionIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetId(offset, limit, iLangcode,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageAlt(offset, limit, iLangcode,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitle(offset, limit, iLangcode,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidth(offset, limit, iLangcode,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByLangcodeAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByLangcodeAndFieldImageHeight(offset, limit, iLangcode,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByLangcodeAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetId(offset, limit, iDelta,iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageAlt(offset, limit, iDelta,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageTitle(offset, limit, iDelta,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageWidth(offset, limit, iDelta,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByDeltaAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByDeltaAndFieldImageHeight(offset, limit, iDelta,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByDeltaAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAlt(offset, limit, iFieldImageTargetId,iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitle(offset, limit, iFieldImageTargetId,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidth(offset, limit, iFieldImageTargetId,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageHeight(offset, limit, iFieldImageTargetId,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageTitle := self.Args("field_image_title").String()

	if helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitle(offset, limit, iFieldImageAlt,iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidth(offset, limit, iFieldImageAlt,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageAlt := self.Args("field_image_alt").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageHeight(offset, limit, iFieldImageAlt,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageWidth := self.Args("field_image_width").MustInt()

	if helper.IsHas(iFieldImageTitle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidth(offset, limit, iFieldImageTitle,iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageTitle := self.Args("field_image_title").String()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageTitle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageHeight(offset, limit, iFieldImageTitle,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesByFieldImageWidthAndFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	iFieldImageHeight := self.Args("field_image_height").MustInt()

	if helper.IsHas(iFieldImageWidth) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImagesByFieldImageWidthAndFieldImageHeight(offset, limit, iFieldImageWidth,iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImagesByFieldImageWidthAndFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImagesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImages(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImages' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaFieldImageTargetId(iFieldImageTargetId)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageAlt := self.Args("field_image_alt").String()
	if helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaFieldImageAlt(iFieldImageAlt)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageTitle := self.Args("field_image_title").String()
	if helper.IsHas(iFieldImageTitle) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaFieldImageTitle(iFieldImageTitle)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	if helper.IsHas(iFieldImageWidth) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaFieldImageWidth(iFieldImageWidth)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageHeight := self.Args("field_image_height").MustInt()
	if helper.IsHas(iFieldImageHeight) {
		_NodeRevision_fieldImage := model.HasNodeRevision_fieldImageViaFieldImageHeight(iFieldImageHeight)
		var m = map[string]interface{}{}
		m["node_revision__field_image"] = _NodeRevision_fieldImage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldImageViaFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageTargetId := self.Args("field_image_target_id").MustInt()
	if helper.IsHas(iFieldImageTargetId) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaFieldImageTargetId(iFieldImageTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageAlt := self.Args("field_image_alt").String()
	if helper.IsHas(iFieldImageAlt) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaFieldImageAlt(iFieldImageAlt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageTitle := self.Args("field_image_title").String()
	if helper.IsHas(iFieldImageTitle) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaFieldImageTitle(iFieldImageTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageWidth := self.Args("field_image_width").MustInt()
	if helper.IsHas(iFieldImageWidth) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaFieldImageWidth(iFieldImageWidth)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldImageHeight := self.Args("field_image_height").MustInt()
	if helper.IsHas(iFieldImageHeight) {
		_NodeRevision_fieldImage, _error := model.GetNodeRevision_fieldImageViaFieldImageHeight(iFieldImageHeight)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldImageViaFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaBundle(Bundle_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaDeleted(Deleted_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaEntityId(EntityId_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaRevisionId(RevisionId_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaLangcode(Langcode_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaDelta(Delta_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	if helper.IsHas(FieldImageTargetId_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaFieldImageTargetId(FieldImageTargetId_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaFieldImageTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageAlt_ := self.Args("field_image_alt").String()
	if helper.IsHas(FieldImageAlt_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaFieldImageAlt(FieldImageAlt_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaFieldImageAlt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTitle_ := self.Args("field_image_title").String()
	if helper.IsHas(FieldImageTitle_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaFieldImageTitle(FieldImageTitle_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaFieldImageTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	if helper.IsHas(FieldImageWidth_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaFieldImageWidth(FieldImageWidth_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaFieldImageWidth's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	if helper.IsHas(FieldImageHeight_) {
		var iNodeRevision_fieldImage model.NodeRevision_fieldImage
		self.Bind(&iNodeRevision_fieldImage)
		_NodeRevision_fieldImage, _error := model.SetNodeRevision_fieldImageViaFieldImageHeight(FieldImageHeight_, &iNodeRevision_fieldImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldImage)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldImageViaFieldImageHeight's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeRevision_fieldImageHandler(self *macross.Context) error {
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
		_error := model.AddNodeRevision_fieldImage(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,FieldImageTargetId_,FieldImageAlt_,FieldImageTitle_,FieldImageWidth_,FieldImageHeight_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNodeRevision_fieldImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeRevision_fieldImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_string, _error := model.PostNodeRevision_fieldImage(&iNodeRevision_fieldImage)
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

func PutNodeRevision_fieldImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_string, _error := model.PutNodeRevision_fieldImage(&iNodeRevision_fieldImage)
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

func PutNodeRevision_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaBundle(Bundle_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaDeleted(Deleted_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaEntityId(EntityId_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaRevisionId(RevisionId_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaLangcode(Langcode_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaDelta(Delta_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaFieldImageTargetId(FieldImageTargetId_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageAlt_ := self.Args("field_image_alt").String()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaFieldImageAlt(FieldImageAlt_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTitle_ := self.Args("field_image_title").String()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaFieldImageTitle(FieldImageTitle_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaFieldImageWidth(FieldImageWidth_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	_int64, _error := model.PutNodeRevision_fieldImageViaFieldImageHeight(FieldImageHeight_, &iNodeRevision_fieldImage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaFieldImageTargetId(FieldImageTargetId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageAlt_ := self.Args("field_image_alt").String()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaFieldImageAlt(FieldImageAlt_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTitle_ := self.Args("field_image_title").String()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaFieldImageTitle(FieldImageTitle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaFieldImageWidth(FieldImageWidth_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	var iNodeRevision_fieldImage model.NodeRevision_fieldImage
	self.Bind(&iNodeRevision_fieldImage)
	var iMap = helper.StructToMap(iNodeRevision_fieldImage)
	_error := model.UpdateNodeRevision_fieldImageViaFieldImageHeight(FieldImageHeight_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteNodeRevision_fieldImage(Bundle_)
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

func DeleteNodeRevision_fieldImageViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteNodeRevision_fieldImageViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteNodeRevision_fieldImageViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteNodeRevision_fieldImageViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteNodeRevision_fieldImageViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeRevision_fieldImageViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteNodeRevision_fieldImageViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaFieldImageTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTargetId_ := self.Args("field_image_target_id").MustInt()
	_error := model.DeleteNodeRevision_fieldImageViaFieldImageTargetId(FieldImageTargetId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaFieldImageAltHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageAlt_ := self.Args("field_image_alt").String()
	_error := model.DeleteNodeRevision_fieldImageViaFieldImageAlt(FieldImageAlt_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaFieldImageTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageTitle_ := self.Args("field_image_title").String()
	_error := model.DeleteNodeRevision_fieldImageViaFieldImageTitle(FieldImageTitle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaFieldImageWidthHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageWidth_ := self.Args("field_image_width").MustInt()
	_error := model.DeleteNodeRevision_fieldImageViaFieldImageWidth(FieldImageWidth_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldImageViaFieldImageHeightHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldImageHeight_ := self.Args("field_image_height").MustInt()
	_error := model.DeleteNodeRevision_fieldImageViaFieldImageHeight(FieldImageHeight_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
