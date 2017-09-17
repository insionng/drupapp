package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetNode_bodiesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNode_bodiesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node__bodysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNode_bodiesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetNode_bodyCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodyCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetNode_bodyCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodyCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetNode_bodyCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodyCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetNode_bodyCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodyCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNode_bodyCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodyCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetNode_bodyCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodyCountViaBodyValueHandler(self *macross.Context) error {
	BodyValue_ := self.Args("body_value").String()
	_int64 := model.GetNode_bodyCountViaBodyValue(BodyValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodyCountViaBodySummaryHandler(self *macross.Context) error {
	BodySummary_ := self.Args("body_summary").String()
	_int64 := model.GetNode_bodyCountViaBodySummary(BodySummary_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodyCountViaBodyFormatHandler(self *macross.Context) error {
	BodyFormat_ := self.Args("body_format").String()
	_int64 := model.GetNode_bodyCountViaBodyFormat(BodyFormat_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__bodyCount"] = 0
	}
	m["node__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNode_bodiesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_Node_body, _error := model.GetNode_bodiesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodyValue := self.Args("body_value").String()
	if (offset > 0) && helper.IsHas(iBodyValue) {
		_Node_body, _error := model.GetNode_bodiesViaBodyValue(offset, limit, iBodyValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodySummary := self.Args("body_summary").String()
	if (offset > 0) && helper.IsHas(iBodySummary) {
		_Node_body, _error := model.GetNode_bodiesViaBodySummary(offset, limit, iBodySummary, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodyFormat := self.Args("body_format").String()
	if (offset > 0) && helper.IsHas(iBodyFormat) {
		_Node_body, _error := model.GetNode_bodiesViaBodyFormat(offset, limit, iBodyFormat, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeletedAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeletedAndBodyValue(offset, limit, iBundle,iDeleted,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeletedAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeletedAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeletedAndBodySummary(offset, limit, iBundle,iDeleted,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeletedAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeletedAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeletedAndBodyFormat(offset, limit, iBundle,iDeleted,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeletedAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndEntityIdAndBodyValue(offset, limit, iBundle,iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndEntityIdAndBodySummary(offset, limit, iBundle,iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndEntityIdAndBodyFormat(offset, limit, iBundle,iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndRevisionIdAndBodyValue(offset, limit, iBundle,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndRevisionIdAndBodySummary(offset, limit, iBundle,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndRevisionIdAndBodyFormat(offset, limit, iBundle,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndLangcodeAndBodyValue(offset, limit, iBundle,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndLangcodeAndBodySummary(offset, limit, iBundle,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndLangcodeAndBodyFormat(offset, limit, iBundle,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeltaAndBodyValue(offset, limit, iBundle,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeltaAndBodySummary(offset, limit, iBundle,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeltaAndBodyFormat(offset, limit, iBundle,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndBodyValueAndBodySummary(offset, limit, iBundle,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndBodyValueAndBodyFormat(offset, limit, iBundle,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndBodySummaryAndBodyFormat(offset, limit, iBundle,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndEntityIdAndBodyValue(offset, limit, iDeleted,iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndEntityIdAndBodySummary(offset, limit, iDeleted,iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndEntityIdAndBodyFormat(offset, limit, iDeleted,iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndRevisionIdAndBodyValue(offset, limit, iDeleted,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndRevisionIdAndBodySummary(offset, limit, iDeleted,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndRevisionIdAndBodyFormat(offset, limit, iDeleted,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndLangcodeAndBodyValue(offset, limit, iDeleted,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndLangcodeAndBodySummary(offset, limit, iDeleted,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndLangcodeAndBodyFormat(offset, limit, iDeleted,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndDeltaAndBodyValue(offset, limit, iDeleted,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndDeltaAndBodySummary(offset, limit, iDeleted,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndDeltaAndBodyFormat(offset, limit, iDeleted,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndBodyValueAndBodySummary(offset, limit, iDeleted,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndBodyValueAndBodyFormat(offset, limit, iDeleted,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndBodySummaryAndBodyFormat(offset, limit, iDeleted,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndRevisionIdAndBodyValue(offset, limit, iEntityId,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndRevisionIdAndBodySummary(offset, limit, iEntityId,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndRevisionIdAndBodyFormat(offset, limit, iEntityId,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndLangcodeAndBodyValue(offset, limit, iEntityId,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndLangcodeAndBodySummary(offset, limit, iEntityId,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndLangcodeAndBodyFormat(offset, limit, iEntityId,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndDeltaAndBodyValue(offset, limit, iEntityId,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndDeltaAndBodySummary(offset, limit, iEntityId,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndDeltaAndBodyFormat(offset, limit, iEntityId,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndBodyValueAndBodySummary(offset, limit, iEntityId,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndBodyValueAndBodyFormat(offset, limit, iEntityId,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndBodySummaryAndBodyFormat(offset, limit, iEntityId,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndLangcodeAndBodyValue(offset, limit, iRevisionId,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndLangcodeAndBodySummary(offset, limit, iRevisionId,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndLangcodeAndBodyFormat(offset, limit, iRevisionId,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndDeltaAndBodyValue(offset, limit, iRevisionId,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndDeltaAndBodySummary(offset, limit, iRevisionId,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndDeltaAndBodyFormat(offset, limit, iRevisionId,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndBodyValueAndBodySummary(offset, limit, iRevisionId,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndBodyValueAndBodyFormat(offset, limit, iRevisionId,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndBodySummaryAndBodyFormat(offset, limit, iRevisionId,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndDeltaAndBodyValue(offset, limit, iLangcode,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndDeltaAndBodySummary(offset, limit, iLangcode,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndDeltaAndBodyFormat(offset, limit, iLangcode,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndBodyValueAndBodySummary(offset, limit, iLangcode,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndBodyValueAndBodyFormat(offset, limit, iLangcode,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndBodySummaryAndBodyFormat(offset, limit, iLangcode,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeltaAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDelta) {
		_Node_body, _error := model.GetNode_bodiesByDeltaAndBodyValueAndBodySummary(offset, limit, iDelta,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeltaAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeltaAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_Node_body, _error := model.GetNode_bodiesByDeltaAndBodyValueAndBodyFormat(offset, limit, iDelta,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeltaAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeltaAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_Node_body, _error := model.GetNode_bodiesByDeltaAndBodySummaryAndBodyFormat(offset, limit, iDelta,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeltaAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBodyValueAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodyValue) {
		_Node_body, _error := model.GetNode_bodiesByBodyValueAndBodySummaryAndBodyFormat(offset, limit, iBodyValue,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBodyValueAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndBodyValue(offset, limit, iBundle,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndBodySummary(offset, limit, iBundle,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBundleAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodiesByBundleAndBodyFormat(offset, limit, iBundle,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBundleAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndBodyValue(offset, limit, iDeleted,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndBodySummary(offset, limit, iDeleted,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeletedAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodiesByDeletedAndBodyFormat(offset, limit, iDeleted,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeletedAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndBodyValue(offset, limit, iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndBodySummary(offset, limit, iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodiesByEntityIdAndBodyFormat(offset, limit, iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndBodyValue(offset, limit, iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndBodySummary(offset, limit, iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodiesByRevisionIdAndBodyFormat(offset, limit, iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndBodyValue(offset, limit, iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndBodySummary(offset, limit, iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodiesByLangcodeAndBodyFormat(offset, limit, iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDelta) {
		_Node_body, _error := model.GetNode_bodiesByDeltaAndBodyValue(offset, limit, iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDelta) {
		_Node_body, _error := model.GetNode_bodiesByDeltaAndBodySummary(offset, limit, iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_Node_body, _error := model.GetNode_bodiesByDeltaAndBodyFormat(offset, limit, iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBodyValue) {
		_Node_body, _error := model.GetNode_bodiesByBodyValueAndBodySummary(offset, limit, iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodyValue) {
		_Node_body, _error := model.GetNode_bodiesByBodyValueAndBodyFormat(offset, limit, iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesByBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodySummary) {
		_Node_body, _error := model.GetNode_bodiesByBodySummaryAndBodyFormat(offset, limit, iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodiesByBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodiesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Node_body, _error := model.GetNode_bodies(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodies' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_Node_body := model.HasNode_bodyViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Node_body := model.HasNode_bodyViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_Node_body := model.HasNode_bodyViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_Node_body := model.HasNode_bodyViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Node_body := model.HasNode_bodyViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_Node_body := model.HasNode_bodyViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyValue := self.Args("body_value").String()
	if helper.IsHas(iBodyValue) {
		_Node_body := model.HasNode_bodyViaBodyValue(iBodyValue)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodySummary := self.Args("body_summary").String()
	if helper.IsHas(iBodySummary) {
		_Node_body := model.HasNode_bodyViaBodySummary(iBodySummary)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyFormat := self.Args("body_format").String()
	if helper.IsHas(iBodyFormat) {
		_Node_body := model.HasNode_bodyViaBodyFormat(iBodyFormat)
		var m = map[string]interface{}{}
		m["node__body"] = _Node_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_Node_body, _error := model.GetNode_bodyViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Node_body, _error := model.GetNode_bodyViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_Node_body, _error := model.GetNode_bodyViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_Node_body, _error := model.GetNode_bodyViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Node_body, _error := model.GetNode_bodyViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_Node_body, _error := model.GetNode_bodyViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyValue := self.Args("body_value").String()
	if helper.IsHas(iBodyValue) {
		_Node_body, _error := model.GetNode_bodyViaBodyValue(iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodySummary := self.Args("body_summary").String()
	if helper.IsHas(iBodySummary) {
		_Node_body, _error := model.GetNode_bodyViaBodySummary(iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyFormat := self.Args("body_format").String()
	if helper.IsHas(iBodyFormat) {
		_Node_body, _error := model.GetNode_bodyViaBodyFormat(iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the GetNode_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaBundle(Bundle_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaDeleted(Deleted_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaEntityId(EntityId_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaRevisionId(RevisionId_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaLangcode(Langcode_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaDelta(Delta_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	if helper.IsHas(BodyValue_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaBodyValue(BodyValue_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	if helper.IsHas(BodySummary_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaBodySummary(BodySummary_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	if helper.IsHas(BodyFormat_) {
		var iNode_body model.Node_body
		self.Bind(&iNode_body)
		_Node_body, _error := model.SetNode_bodyViaBodyFormat(BodyFormat_, &iNode_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_body)
	}
	herr.Message = "Can't get to the SetNode_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNode_bodyHandler(self *macross.Context) error {
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
		_error := model.AddNode_body(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,BodyValue_,BodySummary_,BodyFormat_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNode_body's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNode_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_string, _error := model.PostNode_body(&iNode_body)
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

func PutNode_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_string, _error := model.PutNode_body(&iNode_body)
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

func PutNode_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaBundle(Bundle_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaDeleted(Deleted_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaEntityId(EntityId_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaRevisionId(RevisionId_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaLangcode(Langcode_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaDelta(Delta_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaBodyValue(BodyValue_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaBodySummary(BodySummary_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	_int64, _error := model.PutNode_bodyViaBodyFormat(BodyFormat_, &iNode_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaBodyValue(BodyValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaBodySummary(BodySummary_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	var iNode_body model.Node_body
	self.Bind(&iNode_body)
	var iMap = helper.StructToMap(iNode_body)
	_error := model.UpdateNode_bodyViaBodyFormat(BodyFormat_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteNode_body(Bundle_)
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

func DeleteNode_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteNode_bodyViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteNode_bodyViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteNode_bodyViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteNode_bodyViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNode_bodyViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteNode_bodyViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	_error := model.DeleteNode_bodyViaBodyValue(BodyValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	_error := model.DeleteNode_bodyViaBodySummary(BodySummary_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	_error := model.DeleteNode_bodyViaBodyFormat(BodyFormat_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
