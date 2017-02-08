package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetBlockContentRevision_bodiesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBlockContentRevision_bodiesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["block_content_revision__bodysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetBlockContentRevision_bodyCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodyCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetBlockContentRevision_bodyCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodyCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetBlockContentRevision_bodyCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodyCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetBlockContentRevision_bodyCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodyCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetBlockContentRevision_bodyCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodyCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetBlockContentRevision_bodyCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodyCountViaBodyValueHandler(self *macross.Context) error {
	BodyValue_ := self.Args("body_value").String()
	_int64 := model.GetBlockContentRevision_bodyCountViaBodyValue(BodyValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodyCountViaBodySummaryHandler(self *macross.Context) error {
	BodySummary_ := self.Args("body_summary").String()
	_int64 := model.GetBlockContentRevision_bodyCountViaBodySummary(BodySummary_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodyCountViaBodyFormatHandler(self *macross.Context) error {
	BodyFormat_ := self.Args("body_format").String()
	_int64 := model.GetBlockContentRevision_bodyCountViaBodyFormat(BodyFormat_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content_revision__bodyCount"] = 0
	}
	m["block_content_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContentRevision_bodiesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodyValue := self.Args("body_value").String()
	if (offset > 0) && helper.IsHas(iBodyValue) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaBodyValue(offset, limit, iBodyValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodySummary := self.Args("body_summary").String()
	if (offset > 0) && helper.IsHas(iBodySummary) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaBodySummary(offset, limit, iBodySummary, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodyFormat := self.Args("body_format").String()
	if (offset > 0) && helper.IsHas(iBodyFormat) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesViaBodyFormat(offset, limit, iBodyFormat, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyValue(offset, limit, iBundle,iDeleted,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeletedAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeletedAndBodySummary(offset, limit, iBundle,iDeleted,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeletedAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyFormat(offset, limit, iBundle,iDeleted,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyValue(offset, limit, iBundle,iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodySummary(offset, limit, iBundle,iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyFormat(offset, limit, iBundle,iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyValue(offset, limit, iBundle,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodySummary(offset, limit, iBundle,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyFormat(offset, limit, iBundle,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyValue(offset, limit, iBundle,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodySummary(offset, limit, iBundle,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyFormat(offset, limit, iBundle,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyValue(offset, limit, iBundle,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeltaAndBodySummary(offset, limit, iBundle,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyFormat(offset, limit, iBundle,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodySummary(offset, limit, iBundle,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodyFormat(offset, limit, iBundle,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndBodySummaryAndBodyFormat(offset, limit, iBundle,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyValue(offset, limit, iDeleted,iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodySummary(offset, limit, iDeleted,iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyFormat(offset, limit, iDeleted,iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyValue(offset, limit, iDeleted,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodySummary(offset, limit, iDeleted,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyFormat(offset, limit, iDeleted,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyValue(offset, limit, iDeleted,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodySummary(offset, limit, iDeleted,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyFormat(offset, limit, iDeleted,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyValue(offset, limit, iDeleted,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodySummary(offset, limit, iDeleted,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyFormat(offset, limit, iDeleted,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodySummary(offset, limit, iDeleted,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodyFormat(offset, limit, iDeleted,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndBodySummaryAndBodyFormat(offset, limit, iDeleted,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyValue(offset, limit, iEntityId,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodySummary(offset, limit, iEntityId,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyFormat(offset, limit, iEntityId,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyValue(offset, limit, iEntityId,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodySummary(offset, limit, iEntityId,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyFormat(offset, limit, iEntityId,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyValue(offset, limit, iEntityId,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodySummary(offset, limit, iEntityId,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyFormat(offset, limit, iEntityId,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodySummary(offset, limit, iEntityId,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodyFormat(offset, limit, iEntityId,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndBodySummaryAndBodyFormat(offset, limit, iEntityId,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyValue(offset, limit, iRevisionId,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodySummary(offset, limit, iRevisionId,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyFormat(offset, limit, iRevisionId,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyValue(offset, limit, iRevisionId,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodySummary(offset, limit, iRevisionId,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyFormat(offset, limit, iRevisionId,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodySummary(offset, limit, iRevisionId,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodyFormat(offset, limit, iRevisionId,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormat(offset, limit, iRevisionId,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyValue(offset, limit, iLangcode,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodySummary(offset, limit, iLangcode,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyFormat(offset, limit, iLangcode,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodySummary(offset, limit, iLangcode,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodyFormat(offset, limit, iLangcode,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndBodySummaryAndBodyFormat(offset, limit, iLangcode,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDelta) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodySummary(offset, limit, iDelta,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodyFormat(offset, limit, iDelta,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeltaAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeltaAndBodySummaryAndBodyFormat(offset, limit, iDelta,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeltaAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBodyValueAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodyValue) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBodyValueAndBodySummaryAndBodyFormat(offset, limit, iBodyValue,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBodyValueAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndBodyValue(offset, limit, iBundle,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndBodySummary(offset, limit, iBundle,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBundleAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBundleAndBodyFormat(offset, limit, iBundle,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBundleAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndBodyValue(offset, limit, iDeleted,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndBodySummary(offset, limit, iDeleted,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeletedAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeletedAndBodyFormat(offset, limit, iDeleted,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeletedAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndBodyValue(offset, limit, iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndBodySummary(offset, limit, iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByEntityIdAndBodyFormat(offset, limit, iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndBodyValue(offset, limit, iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndBodySummary(offset, limit, iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByRevisionIdAndBodyFormat(offset, limit, iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndBodyValue(offset, limit, iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndBodySummary(offset, limit, iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByLangcodeAndBodyFormat(offset, limit, iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDelta) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeltaAndBodyValue(offset, limit, iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDelta) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeltaAndBodySummary(offset, limit, iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByDeltaAndBodyFormat(offset, limit, iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBodyValue) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBodyValueAndBodySummary(offset, limit, iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodyValue) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBodyValueAndBodyFormat(offset, limit, iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesByBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodySummary) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodiesByBodySummaryAndBodyFormat(offset, limit, iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodiesByBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodiesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodies(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodies' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyValue := self.Args("body_value").String()
	if helper.IsHas(iBodyValue) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaBodyValue(iBodyValue)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodySummary := self.Args("body_summary").String()
	if helper.IsHas(iBodySummary) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaBodySummary(iBodySummary)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContentRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyFormat := self.Args("body_format").String()
	if helper.IsHas(iBodyFormat) {
		_BlockContentRevision_body := model.HasBlockContentRevision_bodyViaBodyFormat(iBodyFormat)
		var m = map[string]interface{}{}
		m["block_content_revision__body"] = _BlockContentRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContentRevision_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyValue := self.Args("body_value").String()
	if helper.IsHas(iBodyValue) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaBodyValue(iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodySummary := self.Args("body_summary").String()
	if helper.IsHas(iBodySummary) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaBodySummary(iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContentRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyFormat := self.Args("body_format").String()
	if helper.IsHas(iBodyFormat) {
		_BlockContentRevision_body, _error := model.GetBlockContentRevision_bodyViaBodyFormat(iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the GetBlockContentRevision_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaBundle(Bundle_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaDeleted(Deleted_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaEntityId(EntityId_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaRevisionId(RevisionId_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaLangcode(Langcode_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaDelta(Delta_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	if helper.IsHas(BodyValue_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaBodyValue(BodyValue_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	if helper.IsHas(BodySummary_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaBodySummary(BodySummary_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContentRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	if helper.IsHas(BodyFormat_) {
		var iBlockContentRevision_body model.BlockContentRevision_body
		self.Bind(&iBlockContentRevision_body)
		_BlockContentRevision_body, _error := model.SetBlockContentRevision_bodyViaBodyFormat(BodyFormat_, &iBlockContentRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContentRevision_body)
	}
	herr.Message = "Can't get to the SetBlockContentRevision_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBlockContentRevision_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	Deleted_ := self.Args("deleted").MustInt()
	EntityId_ := self.Args("entity_id").MustInt()
	RevisionId_ := self.Args("revision_id").MustInt()
	Langcode_ := self.Args("langcode").String()
	Delta_ := self.Args("delta").MustInt()
	BodyValue_ := self.Args("body_value").String()
	BodySummary_ := self.Args("body_summary").String()
	BodyFormat_ := self.Args("body_format").String()

	if helper.IsHas(Bundle_) {
		_error := model.AddBlockContentRevision_body(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,BodyValue_,BodySummary_,BodyFormat_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBlockContentRevision_body's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBlockContentRevision_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_string, _error := model.PostBlockContentRevision_body(&iBlockContentRevision_body)
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

func PutBlockContentRevision_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_string, _error := model.PutBlockContentRevision_body(&iBlockContentRevision_body)
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

func PutBlockContentRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaBundle(Bundle_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaDeleted(Deleted_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaEntityId(EntityId_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaRevisionId(RevisionId_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaLangcode(Langcode_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaDelta(Delta_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaBodyValue(BodyValue_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaBodySummary(BodySummary_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContentRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	_int64, _error := model.PutBlockContentRevision_bodyViaBodyFormat(BodyFormat_, &iBlockContentRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaBodyValue(BodyValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaBodySummary(BodySummary_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContentRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	var iBlockContentRevision_body model.BlockContentRevision_body
	self.Bind(&iBlockContentRevision_body)
	var iMap = helper.StructToMap(iBlockContentRevision_body)
	_error := model.UpdateBlockContentRevision_bodyViaBodyFormat(BodyFormat_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteBlockContentRevision_body(Bundle_)
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

func DeleteBlockContentRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteBlockContentRevision_bodyViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteBlockContentRevision_bodyViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteBlockContentRevision_bodyViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteBlockContentRevision_bodyViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteBlockContentRevision_bodyViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteBlockContentRevision_bodyViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	_error := model.DeleteBlockContentRevision_bodyViaBodyValue(BodyValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	_error := model.DeleteBlockContentRevision_bodyViaBodySummary(BodySummary_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContentRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	_error := model.DeleteBlockContentRevision_bodyViaBodyFormat(BodyFormat_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
