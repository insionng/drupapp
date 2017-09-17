package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetBlockContent_bodiesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBlockContent_bodiesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["block_content__bodysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetBlockContent_bodyCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodyCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetBlockContent_bodyCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodyCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetBlockContent_bodyCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodyCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetBlockContent_bodyCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodyCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetBlockContent_bodyCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodyCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetBlockContent_bodyCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodyCountViaBodyValueHandler(self *macross.Context) error {
	BodyValue_ := self.Args("body_value").String()
	_int64 := model.GetBlockContent_bodyCountViaBodyValue(BodyValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodyCountViaBodySummaryHandler(self *macross.Context) error {
	BodySummary_ := self.Args("body_summary").String()
	_int64 := model.GetBlockContent_bodyCountViaBodySummary(BodySummary_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodyCountViaBodyFormatHandler(self *macross.Context) error {
	BodyFormat_ := self.Args("body_format").String()
	_int64 := model.GetBlockContent_bodyCountViaBodyFormat(BodyFormat_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["block_content__bodyCount"] = 0
	}
	m["block_content__bodyCount"] = _int64
	return self.JSON(m)
}

func GetBlockContent_bodiesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodyValue := self.Args("body_value").String()
	if (offset > 0) && helper.IsHas(iBodyValue) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaBodyValue(offset, limit, iBodyValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodySummary := self.Args("body_summary").String()
	if (offset > 0) && helper.IsHas(iBodySummary) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaBodySummary(offset, limit, iBodySummary, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodyFormat := self.Args("body_format").String()
	if (offset > 0) && helper.IsHas(iBodyFormat) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesViaBodyFormat(offset, limit, iBodyFormat, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeletedAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeletedAndBodyValue(offset, limit, iBundle,iDeleted,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeletedAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeletedAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeletedAndBodySummary(offset, limit, iBundle,iDeleted,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeletedAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeletedAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeletedAndBodyFormat(offset, limit, iBundle,iDeleted,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeletedAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndEntityIdAndBodyValue(offset, limit, iBundle,iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndEntityIdAndBodySummary(offset, limit, iBundle,iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndEntityIdAndBodyFormat(offset, limit, iBundle,iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndRevisionIdAndBodyValue(offset, limit, iBundle,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndRevisionIdAndBodySummary(offset, limit, iBundle,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndRevisionIdAndBodyFormat(offset, limit, iBundle,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndLangcodeAndBodyValue(offset, limit, iBundle,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndLangcodeAndBodySummary(offset, limit, iBundle,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndLangcodeAndBodyFormat(offset, limit, iBundle,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeltaAndBodyValue(offset, limit, iBundle,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeltaAndBodySummary(offset, limit, iBundle,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeltaAndBodyFormat(offset, limit, iBundle,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndBodyValueAndBodySummary(offset, limit, iBundle,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndBodyValueAndBodyFormat(offset, limit, iBundle,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndBodySummaryAndBodyFormat(offset, limit, iBundle,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndEntityIdAndBodyValue(offset, limit, iDeleted,iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndEntityIdAndBodySummary(offset, limit, iDeleted,iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndEntityIdAndBodyFormat(offset, limit, iDeleted,iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyValue(offset, limit, iDeleted,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndRevisionIdAndBodySummary(offset, limit, iDeleted,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyFormat(offset, limit, iDeleted,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndLangcodeAndBodyValue(offset, limit, iDeleted,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndLangcodeAndBodySummary(offset, limit, iDeleted,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndLangcodeAndBodyFormat(offset, limit, iDeleted,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndDeltaAndBodyValue(offset, limit, iDeleted,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndDeltaAndBodySummary(offset, limit, iDeleted,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndDeltaAndBodyFormat(offset, limit, iDeleted,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndBodyValueAndBodySummary(offset, limit, iDeleted,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndBodyValueAndBodyFormat(offset, limit, iDeleted,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndBodySummaryAndBodyFormat(offset, limit, iDeleted,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyValue(offset, limit, iEntityId,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodySummary(offset, limit, iEntityId,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyFormat(offset, limit, iEntityId,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyValue(offset, limit, iEntityId,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndLangcodeAndBodySummary(offset, limit, iEntityId,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyFormat(offset, limit, iEntityId,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndDeltaAndBodyValue(offset, limit, iEntityId,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndDeltaAndBodySummary(offset, limit, iEntityId,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndDeltaAndBodyFormat(offset, limit, iEntityId,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndBodyValueAndBodySummary(offset, limit, iEntityId,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndBodyValueAndBodyFormat(offset, limit, iEntityId,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndBodySummaryAndBodyFormat(offset, limit, iEntityId,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyValue(offset, limit, iRevisionId,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodySummary(offset, limit, iRevisionId,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyFormat(offset, limit, iRevisionId,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyValue(offset, limit, iRevisionId,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndDeltaAndBodySummary(offset, limit, iRevisionId,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyFormat(offset, limit, iRevisionId,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodySummary(offset, limit, iRevisionId,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodyFormat(offset, limit, iRevisionId,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndBodySummaryAndBodyFormat(offset, limit, iRevisionId,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndDeltaAndBodyValue(offset, limit, iLangcode,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndDeltaAndBodySummary(offset, limit, iLangcode,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndDeltaAndBodyFormat(offset, limit, iLangcode,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndBodyValueAndBodySummary(offset, limit, iLangcode,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndBodyValueAndBodyFormat(offset, limit, iLangcode,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndBodySummaryAndBodyFormat(offset, limit, iLangcode,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeltaAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDelta) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeltaAndBodyValueAndBodySummary(offset, limit, iDelta,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeltaAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeltaAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeltaAndBodyValueAndBodyFormat(offset, limit, iDelta,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeltaAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeltaAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeltaAndBodySummaryAndBodyFormat(offset, limit, iDelta,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeltaAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBodyValueAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodyValue) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBodyValueAndBodySummaryAndBodyFormat(offset, limit, iBodyValue,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBodyValueAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndBodyValue(offset, limit, iBundle,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndBodySummary(offset, limit, iBundle,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBundleAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBundleAndBodyFormat(offset, limit, iBundle,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBundleAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndBodyValue(offset, limit, iDeleted,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndBodySummary(offset, limit, iDeleted,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeletedAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeletedAndBodyFormat(offset, limit, iDeleted,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeletedAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndBodyValue(offset, limit, iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndBodySummary(offset, limit, iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByEntityIdAndBodyFormat(offset, limit, iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndBodyValue(offset, limit, iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndBodySummary(offset, limit, iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByRevisionIdAndBodyFormat(offset, limit, iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndBodyValue(offset, limit, iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndBodySummary(offset, limit, iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByLangcodeAndBodyFormat(offset, limit, iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDelta) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeltaAndBodyValue(offset, limit, iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDelta) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeltaAndBodySummary(offset, limit, iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByDeltaAndBodyFormat(offset, limit, iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBodyValue) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBodyValueAndBodySummary(offset, limit, iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodyValue) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBodyValueAndBodyFormat(offset, limit, iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesByBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodySummary) {
		_BlockContent_body, _error := model.GetBlockContent_bodiesByBodySummaryAndBodyFormat(offset, limit, iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodiesByBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodiesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_BlockContent_body, _error := model.GetBlockContent_bodies(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodies' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_BlockContent_body := model.HasBlockContent_bodyViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_BlockContent_body := model.HasBlockContent_bodyViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_BlockContent_body := model.HasBlockContent_bodyViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContent_body := model.HasBlockContent_bodyViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContent_body := model.HasBlockContent_bodyViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_BlockContent_body := model.HasBlockContent_bodyViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyValue := self.Args("body_value").String()
	if helper.IsHas(iBodyValue) {
		_BlockContent_body := model.HasBlockContent_bodyViaBodyValue(iBodyValue)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodySummary := self.Args("body_summary").String()
	if helper.IsHas(iBodySummary) {
		_BlockContent_body := model.HasBlockContent_bodyViaBodySummary(iBodySummary)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlockContent_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyFormat := self.Args("body_format").String()
	if helper.IsHas(iBodyFormat) {
		_BlockContent_body := model.HasBlockContent_bodyViaBodyFormat(iBodyFormat)
		var m = map[string]interface{}{}
		m["block_content__body"] = _BlockContent_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlockContent_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyValue := self.Args("body_value").String()
	if helper.IsHas(iBodyValue) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaBodyValue(iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodySummary := self.Args("body_summary").String()
	if helper.IsHas(iBodySummary) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaBodySummary(iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlockContent_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyFormat := self.Args("body_format").String()
	if helper.IsHas(iBodyFormat) {
		_BlockContent_body, _error := model.GetBlockContent_bodyViaBodyFormat(iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the GetBlockContent_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaBundle(Bundle_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaDeleted(Deleted_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaEntityId(EntityId_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaRevisionId(RevisionId_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaLangcode(Langcode_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaDelta(Delta_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	if helper.IsHas(BodyValue_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaBodyValue(BodyValue_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	if helper.IsHas(BodySummary_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaBodySummary(BodySummary_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlockContent_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	if helper.IsHas(BodyFormat_) {
		var iBlockContent_body model.BlockContent_body
		self.Bind(&iBlockContent_body)
		_BlockContent_body, _error := model.SetBlockContent_bodyViaBodyFormat(BodyFormat_, &iBlockContent_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlockContent_body)
	}
	herr.Message = "Can't get to the SetBlockContent_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBlockContent_bodyHandler(self *macross.Context) error {
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
		_error := model.AddBlockContent_body(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,BodyValue_,BodySummary_,BodyFormat_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBlockContent_body's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBlockContent_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_string, _error := model.PostBlockContent_body(&iBlockContent_body)
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

func PutBlockContent_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_string, _error := model.PutBlockContent_body(&iBlockContent_body)
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

func PutBlockContent_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaBundle(Bundle_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContent_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaDeleted(Deleted_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContent_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaEntityId(EntityId_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContent_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaRevisionId(RevisionId_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContent_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaLangcode(Langcode_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContent_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaDelta(Delta_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContent_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaBodyValue(BodyValue_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContent_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaBodySummary(BodySummary_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlockContent_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	_int64, _error := model.PutBlockContent_bodyViaBodyFormat(BodyFormat_, &iBlockContent_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaBodyValue(BodyValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaBodySummary(BodySummary_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlockContent_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	var iBlockContent_body model.BlockContent_body
	self.Bind(&iBlockContent_body)
	var iMap = helper.StructToMap(iBlockContent_body)
	_error := model.UpdateBlockContent_bodyViaBodyFormat(BodyFormat_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteBlockContent_body(Bundle_)
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

func DeleteBlockContent_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteBlockContent_bodyViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteBlockContent_bodyViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteBlockContent_bodyViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteBlockContent_bodyViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteBlockContent_bodyViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteBlockContent_bodyViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	_error := model.DeleteBlockContent_bodyViaBodyValue(BodyValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	_error := model.DeleteBlockContent_bodyViaBodySummary(BodySummary_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlockContent_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	_error := model.DeleteBlockContent_bodyViaBodyFormat(BodyFormat_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
