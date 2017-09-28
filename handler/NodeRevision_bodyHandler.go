package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetNodeRevision_bodiesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodeRevision_bodiesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node_revision__bodysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetNodeRevision_bodyCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodyCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetNodeRevision_bodyCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodyCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetNodeRevision_bodyCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodyCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetNodeRevision_bodyCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodyCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeRevision_bodyCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodyCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetNodeRevision_bodyCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodyCountViaBodyValueHandler(self *macross.Context) error {
	BodyValue_ := self.Args("body_value").String()
	_int64 := model.GetNodeRevision_bodyCountViaBodyValue(BodyValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodyCountViaBodySummaryHandler(self *macross.Context) error {
	BodySummary_ := self.Args("body_summary").String()
	_int64 := model.GetNodeRevision_bodyCountViaBodySummary(BodySummary_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodyCountViaBodyFormatHandler(self *macross.Context) error {
	BodyFormat_ := self.Args("body_format").String()
	_int64 := model.GetNodeRevision_bodyCountViaBodyFormat(BodyFormat_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__bodyCount"] = 0
	}
	m["node_revision__bodyCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_bodiesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodyValue := self.Args("body_value").String()
	if (offset > 0) && helper.IsHas(iBodyValue) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaBodyValue(offset, limit, iBodyValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodySummary := self.Args("body_summary").String()
	if (offset > 0) && helper.IsHas(iBodySummary) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaBodySummary(offset, limit, iBodySummary, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBodyFormat := self.Args("body_format").String()
	if (offset > 0) && helper.IsHas(iBodyFormat) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesViaBodyFormat(offset, limit, iBodyFormat, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeletedAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeletedAndBodyValue(offset, limit, iBundle,iDeleted,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeletedAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeletedAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeletedAndBodySummary(offset, limit, iBundle,iDeleted,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeletedAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeletedAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeletedAndBodyFormat(offset, limit, iBundle,iDeleted,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeletedAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndEntityIdAndBodyValue(offset, limit, iBundle,iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndEntityIdAndBodySummary(offset, limit, iBundle,iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndEntityIdAndBodyFormat(offset, limit, iBundle,iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyValue(offset, limit, iBundle,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndRevisionIdAndBodySummary(offset, limit, iBundle,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyFormat(offset, limit, iBundle,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndLangcodeAndBodyValue(offset, limit, iBundle,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndLangcodeAndBodySummary(offset, limit, iBundle,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndLangcodeAndBodyFormat(offset, limit, iBundle,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeltaAndBodyValue(offset, limit, iBundle,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeltaAndBodySummary(offset, limit, iBundle,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeltaAndBodyFormat(offset, limit, iBundle,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndBodyValueAndBodySummary(offset, limit, iBundle,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndBodyValueAndBodyFormat(offset, limit, iBundle,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndBodySummaryAndBodyFormat(offset, limit, iBundle,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyValue(offset, limit, iDeleted,iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndEntityIdAndBodySummary(offset, limit, iDeleted,iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyFormat(offset, limit, iDeleted,iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyValue(offset, limit, iDeleted,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodySummary(offset, limit, iDeleted,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyFormat(offset, limit, iDeleted,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyValue(offset, limit, iDeleted,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndLangcodeAndBodySummary(offset, limit, iDeleted,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyFormat(offset, limit, iDeleted,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndDeltaAndBodyValue(offset, limit, iDeleted,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndDeltaAndBodySummary(offset, limit, iDeleted,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndDeltaAndBodyFormat(offset, limit, iDeleted,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndBodyValueAndBodySummary(offset, limit, iDeleted,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndBodyValueAndBodyFormat(offset, limit, iDeleted,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndBodySummaryAndBodyFormat(offset, limit, iDeleted,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyValue(offset, limit, iEntityId,iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodySummary(offset, limit, iEntityId,iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyFormat(offset, limit, iEntityId,iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyValue(offset, limit, iEntityId,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodySummary(offset, limit, iEntityId,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyFormat(offset, limit, iEntityId,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyValue(offset, limit, iEntityId,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndDeltaAndBodySummary(offset, limit, iEntityId,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyFormat(offset, limit, iEntityId,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodySummary(offset, limit, iEntityId,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodyFormat(offset, limit, iEntityId,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndBodySummaryAndBodyFormat(offset, limit, iEntityId,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyValue(offset, limit, iRevisionId,iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodySummary(offset, limit, iRevisionId,iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyFormat(offset, limit, iRevisionId,iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyValue(offset, limit, iRevisionId,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodySummary(offset, limit, iRevisionId,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyFormat(offset, limit, iRevisionId,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodySummary(offset, limit, iRevisionId,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodyFormat(offset, limit, iRevisionId,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormat(offset, limit, iRevisionId,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyValue(offset, limit, iLangcode,iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndDeltaAndBodySummary(offset, limit, iLangcode,iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyFormat(offset, limit, iLangcode,iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodySummary(offset, limit, iLangcode,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodyFormat(offset, limit, iLangcode,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndBodySummaryAndBodyFormat(offset, limit, iLangcode,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeltaAndBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeltaAndBodyValueAndBodySummary(offset, limit, iDelta,iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeltaAndBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeltaAndBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeltaAndBodyValueAndBodyFormat(offset, limit, iDelta,iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeltaAndBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeltaAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeltaAndBodySummaryAndBodyFormat(offset, limit, iDelta,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeltaAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBodyValueAndBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodyValue) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBodyValueAndBodySummaryAndBodyFormat(offset, limit, iBodyValue,iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBodyValueAndBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndBodyValue(offset, limit, iBundle,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndBodySummary(offset, limit, iBundle,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBundleAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBundleAndBodyFormat(offset, limit, iBundle,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBundleAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndBodyValue(offset, limit, iDeleted,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndBodySummary(offset, limit, iDeleted,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeletedAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeletedAndBodyFormat(offset, limit, iDeleted,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeletedAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndBodyValue(offset, limit, iEntityId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndBodySummary(offset, limit, iEntityId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByEntityIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByEntityIdAndBodyFormat(offset, limit, iEntityId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByEntityIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndBodyValue(offset, limit, iRevisionId,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndBodySummary(offset, limit, iRevisionId,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByRevisionIdAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByRevisionIdAndBodyFormat(offset, limit, iRevisionId,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByRevisionIdAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndBodyValue(offset, limit, iLangcode,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndBodySummary(offset, limit, iLangcode,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByLangcodeAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByLangcodeAndBodyFormat(offset, limit, iLangcode,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByLangcodeAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeltaAndBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyValue := self.Args("body_value").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeltaAndBodyValue(offset, limit, iDelta,iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeltaAndBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeltaAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeltaAndBodySummary(offset, limit, iDelta,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeltaAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByDeltaAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iDelta) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByDeltaAndBodyFormat(offset, limit, iDelta,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByDeltaAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBodyValueAndBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodySummary := self.Args("body_summary").String()

	if helper.IsHas(iBodyValue) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBodyValueAndBodySummary(offset, limit, iBodyValue,iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBodyValueAndBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBodyValueAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodyValue := self.Args("body_value").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodyValue) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBodyValueAndBodyFormat(offset, limit, iBodyValue,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBodyValueAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesByBodySummaryAndBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBodySummary := self.Args("body_summary").String()
	iBodyFormat := self.Args("body_format").String()

	if helper.IsHas(iBodySummary) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodiesByBodySummaryAndBodyFormat(offset, limit, iBodySummary,iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodiesByBodySummaryAndBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodiesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodies(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodies' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyValue := self.Args("body_value").String()
	if helper.IsHas(iBodyValue) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaBodyValue(iBodyValue)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodySummary := self.Args("body_summary").String()
	if helper.IsHas(iBodySummary) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaBodySummary(iBodySummary)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyFormat := self.Args("body_format").String()
	if helper.IsHas(iBodyFormat) {
		_NodeRevision_body := model.HasNodeRevision_bodyViaBodyFormat(iBodyFormat)
		var m = map[string]interface{}{}
		m["node_revision__body"] = _NodeRevision_body
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyValue := self.Args("body_value").String()
	if helper.IsHas(iBodyValue) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaBodyValue(iBodyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodySummary := self.Args("body_summary").String()
	if helper.IsHas(iBodySummary) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaBodySummary(iBodySummary)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBodyFormat := self.Args("body_format").String()
	if helper.IsHas(iBodyFormat) {
		_NodeRevision_body, _error := model.GetNodeRevision_bodyViaBodyFormat(iBodyFormat)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the GetNodeRevision_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaBundle(Bundle_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaDeleted(Deleted_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaEntityId(EntityId_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaRevisionId(RevisionId_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaLangcode(Langcode_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaDelta(Delta_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	if helper.IsHas(BodyValue_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaBodyValue(BodyValue_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaBodyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	if helper.IsHas(BodySummary_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaBodySummary(BodySummary_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaBodySummary's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	if helper.IsHas(BodyFormat_) {
		var iNodeRevision_body model.NodeRevision_body
		self.Bind(&iNodeRevision_body)
		_NodeRevision_body, _error := model.SetNodeRevision_bodyViaBodyFormat(BodyFormat_, &iNodeRevision_body)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_body)
	}
	herr.Message = "Can't get to the SetNodeRevision_bodyViaBodyFormat's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeRevision_bodyHandler(self *macross.Context) error {
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
		_error := model.AddNodeRevision_body(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,BodyValue_,BodySummary_,BodyFormat_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNodeRevision_body's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeRevision_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_string, _error := model.PostNodeRevision_body(&iNodeRevision_body)
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

func PutNodeRevision_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_string, _error := model.PutNodeRevision_body(&iNodeRevision_body)
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

func PutNodeRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaBundle(Bundle_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaDeleted(Deleted_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaEntityId(EntityId_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaRevisionId(RevisionId_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaLangcode(Langcode_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaDelta(Delta_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaBodyValue(BodyValue_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaBodySummary(BodySummary_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	_int64, _error := model.PutNodeRevision_bodyViaBodyFormat(BodyFormat_, &iNodeRevision_body)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaBodyValue(BodyValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaBodySummary(BodySummary_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	var iNodeRevision_body model.NodeRevision_body
	self.Bind(&iNodeRevision_body)
	var iMap = helper.StructToMap(iNodeRevision_body)
	_error := model.UpdateNodeRevision_bodyViaBodyFormat(BodyFormat_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteNodeRevision_body(Bundle_)
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

func DeleteNodeRevision_bodyViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteNodeRevision_bodyViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteNodeRevision_bodyViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteNodeRevision_bodyViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteNodeRevision_bodyViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeRevision_bodyViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteNodeRevision_bodyViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyViaBodyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyValue_ := self.Args("body_value").String()
	_error := model.DeleteNodeRevision_bodyViaBodyValue(BodyValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyViaBodySummaryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodySummary_ := self.Args("body_summary").String()
	_error := model.DeleteNodeRevision_bodyViaBodySummary(BodySummary_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_bodyViaBodyFormatHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BodyFormat_ := self.Args("body_format").String()
	_error := model.DeleteNodeRevision_bodyViaBodyFormat(BodyFormat_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
