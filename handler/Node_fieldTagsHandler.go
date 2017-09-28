package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetNode_fieldTagsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNode_fieldTagsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node__field_tagssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetNode_fieldTagsCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_tagsCount"] = 0
	}
	m["node__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldTagsCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetNode_fieldTagsCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_tagsCount"] = 0
	}
	m["node__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldTagsCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetNode_fieldTagsCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_tagsCount"] = 0
	}
	m["node__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldTagsCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetNode_fieldTagsCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_tagsCount"] = 0
	}
	m["node__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldTagsCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNode_fieldTagsCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_tagsCount"] = 0
	}
	m["node__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldTagsCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetNode_fieldTagsCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_tagsCount"] = 0
	}
	m["node__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldTagsCountViaFieldTagsTargetIdHandler(self *macross.Context) error {
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	_int64 := model.GetNode_fieldTagsCountViaFieldTagsTargetId(FieldTagsTargetId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node__field_tagsCount"] = 0
	}
	m["node__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNode_fieldTagsesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()
	if (offset > 0) && helper.IsHas(iFieldTagsTargetId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesViaFieldTagsTargetId(offset, limit, iFieldTagsTargetId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesViaFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndDeletedAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndDeletedAndFieldTagsTargetId(offset, limit, iBundle,iDeleted,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndDeletedAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetId(offset, limit, iBundle,iEntityId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetId(offset, limit, iBundle,iRevisionId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetId(offset, limit, iBundle,iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndDeltaAndFieldTagsTargetId(offset, limit, iBundle,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetId(offset, limit, iDeleted,iEntityId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetId(offset, limit, iDeleted,iRevisionId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetId(offset, limit, iDeleted,iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetId(offset, limit, iDeleted,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetId(offset, limit, iEntityId,iRevisionId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetId(offset, limit, iEntityId,iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetId(offset, limit, iEntityId,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetId(offset, limit, iRevisionId,iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetId(offset, limit, iRevisionId,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetId(offset, limit, iLangcode,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByBundleAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByBundleAndFieldTagsTargetId(offset, limit, iBundle,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByBundleAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeletedAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeletedAndFieldTagsTargetId(offset, limit, iDeleted,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeletedAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByEntityIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByEntityIdAndFieldTagsTargetId(offset, limit, iEntityId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByEntityIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByRevisionIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByRevisionIdAndFieldTagsTargetId(offset, limit, iRevisionId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByRevisionIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByLangcodeAndFieldTagsTargetId(offset, limit, iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesByDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDelta) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsesByDeltaAndFieldTagsTargetId(offset, limit, iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsesByDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Node_fieldTags, _error := model.GetNode_fieldTagses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_Node_fieldTags := model.HasNode_fieldTagsViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["node__field_tags"] = _Node_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldTagsViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Node_fieldTags := model.HasNode_fieldTagsViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["node__field_tags"] = _Node_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldTagsViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_Node_fieldTags := model.HasNode_fieldTagsViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["node__field_tags"] = _Node_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldTagsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_Node_fieldTags := model.HasNode_fieldTagsViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["node__field_tags"] = _Node_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldTagsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Node_fieldTags := model.HasNode_fieldTagsViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node__field_tags"] = _Node_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldTagsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_Node_fieldTags := model.HasNode_fieldTagsViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["node__field_tags"] = _Node_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldTagsViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNode_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()
	if helper.IsHas(iFieldTagsTargetId) {
		_Node_fieldTags := model.HasNode_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId)
		var m = map[string]interface{}{}
		m["node__field_tags"] = _Node_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNode_fieldTagsViaFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNode_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()
	if helper.IsHas(iFieldTagsTargetId) {
		_Node_fieldTags, _error := model.GetNode_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the GetNode_fieldTagsViaFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iNode_fieldTags model.Node_fieldTags
		self.Bind(&iNode_fieldTags)
		_Node_fieldTags, _error := model.SetNode_fieldTagsViaBundle(Bundle_, &iNode_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the SetNode_fieldTagsViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iNode_fieldTags model.Node_fieldTags
		self.Bind(&iNode_fieldTags)
		_Node_fieldTags, _error := model.SetNode_fieldTagsViaDeleted(Deleted_, &iNode_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the SetNode_fieldTagsViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iNode_fieldTags model.Node_fieldTags
		self.Bind(&iNode_fieldTags)
		_Node_fieldTags, _error := model.SetNode_fieldTagsViaEntityId(EntityId_, &iNode_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the SetNode_fieldTagsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iNode_fieldTags model.Node_fieldTags
		self.Bind(&iNode_fieldTags)
		_Node_fieldTags, _error := model.SetNode_fieldTagsViaRevisionId(RevisionId_, &iNode_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the SetNode_fieldTagsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNode_fieldTags model.Node_fieldTags
		self.Bind(&iNode_fieldTags)
		_Node_fieldTags, _error := model.SetNode_fieldTagsViaLangcode(Langcode_, &iNode_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the SetNode_fieldTagsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iNode_fieldTags model.Node_fieldTags
		self.Bind(&iNode_fieldTags)
		_Node_fieldTags, _error := model.SetNode_fieldTagsViaDelta(Delta_, &iNode_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the SetNode_fieldTagsViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNode_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	if helper.IsHas(FieldTagsTargetId_) {
		var iNode_fieldTags model.Node_fieldTags
		self.Bind(&iNode_fieldTags)
		_Node_fieldTags, _error := model.SetNode_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_, &iNode_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node_fieldTags)
	}
	herr.Message = "Can't get to the SetNode_fieldTagsViaFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNode_fieldTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	Deleted_ := self.Args("deleted").MustInt()
	EntityId_ := self.Args("entity_id").MustInt()
	RevisionId_ := self.Args("revision_id").MustInt()
	Langcode_ := self.Args("langcode").String()
	Delta_ := self.Args("delta").MustInt()
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(Bundle_) {
		_error := model.AddNode_fieldTags(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,FieldTagsTargetId_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNode_fieldTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNode_fieldTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_string, _error := model.PostNode_fieldTags(&iNode_fieldTags)
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

func PutNode_fieldTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_string, _error := model.PutNode_fieldTags(&iNode_fieldTags)
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

func PutNode_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_int64, _error := model.PutNode_fieldTagsViaBundle(Bundle_, &iNode_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_int64, _error := model.PutNode_fieldTagsViaDeleted(Deleted_, &iNode_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_int64, _error := model.PutNode_fieldTagsViaEntityId(EntityId_, &iNode_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_int64, _error := model.PutNode_fieldTagsViaRevisionId(RevisionId_, &iNode_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_int64, _error := model.PutNode_fieldTagsViaLangcode(Langcode_, &iNode_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_int64, _error := model.PutNode_fieldTagsViaDelta(Delta_, &iNode_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNode_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	_int64, _error := model.PutNode_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_, &iNode_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	var iMap = helper.StructToMap(iNode_fieldTags)
	_error := model.UpdateNode_fieldTagsViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	var iMap = helper.StructToMap(iNode_fieldTags)
	_error := model.UpdateNode_fieldTagsViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	var iMap = helper.StructToMap(iNode_fieldTags)
	_error := model.UpdateNode_fieldTagsViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	var iMap = helper.StructToMap(iNode_fieldTags)
	_error := model.UpdateNode_fieldTagsViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	var iMap = helper.StructToMap(iNode_fieldTags)
	_error := model.UpdateNode_fieldTagsViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	var iMap = helper.StructToMap(iNode_fieldTags)
	_error := model.UpdateNode_fieldTagsViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNode_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	var iNode_fieldTags model.Node_fieldTags
	self.Bind(&iNode_fieldTags)
	var iMap = helper.StructToMap(iNode_fieldTags)
	_error := model.UpdateNode_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteNode_fieldTags(Bundle_)
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

func DeleteNode_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteNode_fieldTagsViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteNode_fieldTagsViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteNode_fieldTagsViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteNode_fieldTagsViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNode_fieldTagsViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteNode_fieldTagsViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNode_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	_error := model.DeleteNode_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
