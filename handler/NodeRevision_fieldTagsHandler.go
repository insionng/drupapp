package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetNodeRevision_fieldTagsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodeRevision_fieldTagsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node_revision__field_tagssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsCountViaBundleHandler(self *macross.Context) error {
	Bundle_ := self.Args("bundle").String()
	_int64 := model.GetNodeRevision_fieldTagsCountViaBundle(Bundle_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_tagsCount"] = 0
	}
	m["node_revision__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldTagsCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetNodeRevision_fieldTagsCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_tagsCount"] = 0
	}
	m["node_revision__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldTagsCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetNodeRevision_fieldTagsCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_tagsCount"] = 0
	}
	m["node_revision__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldTagsCountViaRevisionIdHandler(self *macross.Context) error {
	RevisionId_ := self.Args("revision_id").MustInt()
	_int64 := model.GetNodeRevision_fieldTagsCountViaRevisionId(RevisionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_tagsCount"] = 0
	}
	m["node_revision__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldTagsCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeRevision_fieldTagsCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_tagsCount"] = 0
	}
	m["node_revision__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldTagsCountViaDeltaHandler(self *macross.Context) error {
	Delta_ := self.Args("delta").MustInt()
	_int64 := model.GetNodeRevision_fieldTagsCountViaDelta(Delta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_tagsCount"] = 0
	}
	m["node_revision__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldTagsCountViaFieldTagsTargetIdHandler(self *macross.Context) error {
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	_int64 := model.GetNodeRevision_fieldTagsCountViaFieldTagsTargetId(FieldTagsTargetId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revision__field_tagsCount"] = 0
	}
	m["node_revision__field_tagsCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevision_fieldTagsesViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBundle := self.Args("bundle").String()
	if (offset > 0) && helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesViaBundle(offset, limit, iBundle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionId := self.Args("revision_id").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesViaRevisionId(offset, limit, iRevisionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDelta := self.Args("delta").MustInt()
	if (offset > 0) && helper.IsHas(iDelta) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesViaDelta(offset, limit, iDelta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()
	if (offset > 0) && helper.IsHas(iFieldTagsTargetId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesViaFieldTagsTargetId(offset, limit, iFieldTagsTargetId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesViaFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndDeletedAndEntityId(offset, limit, iBundle,iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndDeletedAndRevisionId(offset, limit, iBundle,iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndDeletedAndLangcode(offset, limit, iBundle,iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndDeletedAndDelta(offset, limit, iBundle,iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndDeletedAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndDeletedAndFieldTagsTargetId(offset, limit, iBundle,iDeleted,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndDeletedAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndEntityIdAndRevisionId(offset, limit, iBundle,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndEntityIdAndLangcode(offset, limit, iBundle,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndEntityIdAndDelta(offset, limit, iBundle,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetId(offset, limit, iBundle,iEntityId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndLangcode(offset, limit, iBundle,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndDelta(offset, limit, iBundle,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetId(offset, limit, iBundle,iRevisionId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndLangcodeAndDelta(offset, limit, iBundle,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetId(offset, limit, iBundle,iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndDeltaAndFieldTagsTargetId(offset, limit, iBundle,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndRevisionId(offset, limit, iDeleted,iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndLangcode(offset, limit, iDeleted,iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndDelta(offset, limit, iDeleted,iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetId(offset, limit, iDeleted,iEntityId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndLangcode(offset, limit, iDeleted,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndDelta(offset, limit, iDeleted,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetId(offset, limit, iDeleted,iRevisionId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndDelta(offset, limit, iDeleted,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetId(offset, limit, iDeleted,iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetId(offset, limit, iDeleted,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndLangcode(offset, limit, iEntityId,iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndDelta(offset, limit, iEntityId,iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetId(offset, limit, iEntityId,iRevisionId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndDelta(offset, limit, iEntityId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetId(offset, limit, iEntityId,iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetId(offset, limit, iEntityId,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndDelta(offset, limit, iRevisionId,iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetId(offset, limit, iRevisionId,iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetId(offset, limit, iRevisionId,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetId(offset, limit, iLangcode,iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndDeleted(offset, limit, iBundle,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndEntityId(offset, limit, iBundle,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndRevisionId(offset, limit, iBundle,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndLangcode(offset, limit, iBundle,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndDelta(offset, limit, iBundle,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByBundleAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBundle := self.Args("bundle").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByBundleAndFieldTagsTargetId(offset, limit, iBundle,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByBundleAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndEntityId(offset, limit, iDeleted,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndRevisionId(offset, limit, iDeleted,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndLangcode(offset, limit, iDeleted,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndDelta(offset, limit, iDeleted,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeletedAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeletedAndFieldTagsTargetId(offset, limit, iDeleted,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeletedAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndRevisionId(offset, limit, iEntityId,iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndLangcode(offset, limit, iEntityId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndDelta(offset, limit, iEntityId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByEntityIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByEntityIdAndFieldTagsTargetId(offset, limit, iEntityId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByEntityIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByRevisionIdAndLangcode(offset, limit, iRevisionId,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByRevisionIdAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByRevisionIdAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByRevisionIdAndDelta(offset, limit, iRevisionId,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByRevisionIdAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByRevisionIdAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionId := self.Args("revision_id").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByRevisionIdAndFieldTagsTargetId(offset, limit, iRevisionId,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByRevisionIdAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByLangcodeAndDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDelta := self.Args("delta").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByLangcodeAndDelta(offset, limit, iLangcode,iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByLangcodeAndDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByLangcodeAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByLangcodeAndFieldTagsTargetId(offset, limit, iLangcode,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByLangcodeAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesByDeltaAndFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDelta := self.Args("delta").MustInt()
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()

	if helper.IsHas(iDelta) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsesByDeltaAndFieldTagsTargetId(offset, limit, iDelta,iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsesByDeltaAndFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags := model.HasNodeRevision_fieldTagsViaBundle(iBundle)
		var m = map[string]interface{}{}
		m["node_revision__field_tags"] = _NodeRevision_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldTagsViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags := model.HasNodeRevision_fieldTagsViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["node_revision__field_tags"] = _NodeRevision_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldTagsViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags := model.HasNodeRevision_fieldTagsViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["node_revision__field_tags"] = _NodeRevision_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldTagsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags := model.HasNodeRevision_fieldTagsViaRevisionId(iRevisionId)
		var m = map[string]interface{}{}
		m["node_revision__field_tags"] = _NodeRevision_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldTagsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldTags := model.HasNodeRevision_fieldTagsViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node_revision__field_tags"] = _NodeRevision_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldTagsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_NodeRevision_fieldTags := model.HasNodeRevision_fieldTagsViaDelta(iDelta)
		var m = map[string]interface{}{}
		m["node_revision__field_tags"] = _NodeRevision_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldTagsViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevision_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()
	if helper.IsHas(iFieldTagsTargetId) {
		_NodeRevision_fieldTags := model.HasNodeRevision_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId)
		var m = map[string]interface{}{}
		m["node_revision__field_tags"] = _NodeRevision_fieldTags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevision_fieldTagsViaFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBundle := self.Args("bundle").String()
	if helper.IsHas(iBundle) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsViaBundle(iBundle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionId := self.Args("revision_id").MustInt()
	if helper.IsHas(iRevisionId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsViaRevisionId(iRevisionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDelta := self.Args("delta").MustInt()
	if helper.IsHas(iDelta) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsViaDelta(iDelta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevision_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldTagsTargetId := self.Args("field_tags_target_id").MustInt()
	if helper.IsHas(iFieldTagsTargetId) {
		_NodeRevision_fieldTags, _error := model.GetNodeRevision_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the GetNodeRevision_fieldTagsViaFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	if helper.IsHas(Bundle_) {
		var iNodeRevision_fieldTags model.NodeRevision_fieldTags
		self.Bind(&iNodeRevision_fieldTags)
		_NodeRevision_fieldTags, _error := model.SetNodeRevision_fieldTagsViaBundle(Bundle_, &iNodeRevision_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldTagsViaBundle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iNodeRevision_fieldTags model.NodeRevision_fieldTags
		self.Bind(&iNodeRevision_fieldTags)
		_NodeRevision_fieldTags, _error := model.SetNodeRevision_fieldTagsViaDeleted(Deleted_, &iNodeRevision_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldTagsViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iNodeRevision_fieldTags model.NodeRevision_fieldTags
		self.Bind(&iNodeRevision_fieldTags)
		_NodeRevision_fieldTags, _error := model.SetNodeRevision_fieldTagsViaEntityId(EntityId_, &iNodeRevision_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldTagsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	if helper.IsHas(RevisionId_) {
		var iNodeRevision_fieldTags model.NodeRevision_fieldTags
		self.Bind(&iNodeRevision_fieldTags)
		_NodeRevision_fieldTags, _error := model.SetNodeRevision_fieldTagsViaRevisionId(RevisionId_, &iNodeRevision_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldTagsViaRevisionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNodeRevision_fieldTags model.NodeRevision_fieldTags
		self.Bind(&iNodeRevision_fieldTags)
		_NodeRevision_fieldTags, _error := model.SetNodeRevision_fieldTagsViaLangcode(Langcode_, &iNodeRevision_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldTagsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	if helper.IsHas(Delta_) {
		var iNodeRevision_fieldTags model.NodeRevision_fieldTags
		self.Bind(&iNodeRevision_fieldTags)
		_NodeRevision_fieldTags, _error := model.SetNodeRevision_fieldTagsViaDelta(Delta_, &iNodeRevision_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldTagsViaDelta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevision_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	if helper.IsHas(FieldTagsTargetId_) {
		var iNodeRevision_fieldTags model.NodeRevision_fieldTags
		self.Bind(&iNodeRevision_fieldTags)
		_NodeRevision_fieldTags, _error := model.SetNodeRevision_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_, &iNodeRevision_fieldTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision_fieldTags)
	}
	herr.Message = "Can't get to the SetNodeRevision_fieldTagsViaFieldTagsTargetId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeRevision_fieldTagsHandler(self *macross.Context) error {
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
		_error := model.AddNodeRevision_fieldTags(Bundle_,Deleted_,EntityId_,RevisionId_,Langcode_,Delta_,FieldTagsTargetId_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNodeRevision_fieldTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeRevision_fieldTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_string, _error := model.PostNodeRevision_fieldTags(&iNodeRevision_fieldTags)
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

func PutNodeRevision_fieldTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_string, _error := model.PutNodeRevision_fieldTags(&iNodeRevision_fieldTags)
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

func PutNodeRevision_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_int64, _error := model.PutNodeRevision_fieldTagsViaBundle(Bundle_, &iNodeRevision_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_int64, _error := model.PutNodeRevision_fieldTagsViaDeleted(Deleted_, &iNodeRevision_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_int64, _error := model.PutNodeRevision_fieldTagsViaEntityId(EntityId_, &iNodeRevision_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_int64, _error := model.PutNodeRevision_fieldTagsViaRevisionId(RevisionId_, &iNodeRevision_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_int64, _error := model.PutNodeRevision_fieldTagsViaLangcode(Langcode_, &iNodeRevision_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_int64, _error := model.PutNodeRevision_fieldTagsViaDelta(Delta_, &iNodeRevision_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevision_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	_int64, _error := model.PutNodeRevision_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_, &iNodeRevision_fieldTags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	var iMap = helper.StructToMap(iNodeRevision_fieldTags)
	_error := model.UpdateNodeRevision_fieldTagsViaBundle(Bundle_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	var iMap = helper.StructToMap(iNodeRevision_fieldTags)
	_error := model.UpdateNodeRevision_fieldTagsViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	var iMap = helper.StructToMap(iNodeRevision_fieldTags)
	_error := model.UpdateNodeRevision_fieldTagsViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	var iMap = helper.StructToMap(iNodeRevision_fieldTags)
	_error := model.UpdateNodeRevision_fieldTagsViaRevisionId(RevisionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	var iMap = helper.StructToMap(iNodeRevision_fieldTags)
	_error := model.UpdateNodeRevision_fieldTagsViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	var iMap = helper.StructToMap(iNodeRevision_fieldTags)
	_error := model.UpdateNodeRevision_fieldTagsViaDelta(Delta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevision_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	var iNodeRevision_fieldTags model.NodeRevision_fieldTags
	self.Bind(&iNodeRevision_fieldTags)
	var iMap = helper.StructToMap(iNodeRevision_fieldTags)
	_error := model.UpdateNodeRevision_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_int64, _error := model.DeleteNodeRevision_fieldTags(Bundle_)
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

func DeleteNodeRevision_fieldTagsViaBundleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bundle_ := self.Args("bundle").String()
	_error := model.DeleteNodeRevision_fieldTagsViaBundle(Bundle_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldTagsViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteNodeRevision_fieldTagsViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldTagsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteNodeRevision_fieldTagsViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldTagsViaRevisionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionId_ := self.Args("revision_id").MustInt()
	_error := model.DeleteNodeRevision_fieldTagsViaRevisionId(RevisionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldTagsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeRevision_fieldTagsViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldTagsViaDeltaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Delta_ := self.Args("delta").MustInt()
	_error := model.DeleteNodeRevision_fieldTagsViaDelta(Delta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevision_fieldTagsViaFieldTagsTargetIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldTagsTargetId_ := self.Args("field_tags_target_id").MustInt()
	_error := model.DeleteNodeRevision_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
