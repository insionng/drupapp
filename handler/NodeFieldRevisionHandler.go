package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetNodeFieldRevisionsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodeFieldRevisionsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node_field_revisionsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionCountViaNidHandler(self *macross.Context) error {
	Nid_ := self.Args("nid").MustInt64()
	_int64 := model.GetNodeFieldRevisionCountViaNid(Nid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaVidHandler(self *macross.Context) error {
	Vid_ := self.Args("vid").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaVid(Vid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeFieldRevisionCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaTitleHandler(self *macross.Context) error {
	Title_ := self.Args("title").String()
	_int64 := model.GetNodeFieldRevisionCountViaTitle(Title_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaStatusHandler(self *macross.Context) error {
	Status_ := self.Args("status").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaStatus(Status_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaPromoteHandler(self *macross.Context) error {
	Promote_ := self.Args("promote").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaPromote(Promote_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaStickyHandler(self *macross.Context) error {
	Sticky_ := self.Args("sticky").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaSticky(Sticky_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaRevisionTranslationAffected(RevisionTranslationAffected_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetNodeFieldRevisionCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_field_revisionCount"] = 0
	}
	m["node_field_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeFieldRevisionsViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iNid := self.Args("nid").MustInt64()
	if (offset > 0) && helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaNid(offset, limit, iNid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iVid := self.Args("vid").MustInt()
	if (offset > 0) && helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaVid(offset, limit, iVid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTitle := self.Args("title").String()
	if (offset > 0) && helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaTitle(offset, limit, iTitle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt()
	if (offset > 0) && helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iStatus := self.Args("status").MustInt()
	if (offset > 0) && helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaStatus(offset, limit, iStatus, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").MustInt()
	if (offset > 0) && helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPromote := self.Args("promote").MustInt()
	if (offset > 0) && helper.IsHas(iPromote) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaPromote(offset, limit, iPromote, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSticky := self.Args("sticky").MustInt()
	if (offset > 0) && helper.IsHas(iSticky) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaSticky(offset, limit, iSticky, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionTranslationAffected) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaRevisionTranslationAffected(offset, limit, iRevisionTranslationAffected, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndLangcode(offset, limit, iNid,iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndTitle(offset, limit, iNid,iVid,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndUid(offset, limit, iNid,iVid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndStatus(offset, limit, iNid,iVid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndCreated(offset, limit, iNid,iVid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndChanged(offset, limit, iNid,iVid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndPromote(offset, limit, iNid,iVid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndSticky(offset, limit, iNid,iVid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndRevisionTranslationAffected(offset, limit, iNid,iVid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVidAndDefaultLangcode(offset, limit, iNid,iVid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndTitle(offset, limit, iNid,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndUid(offset, limit, iNid,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndStatus(offset, limit, iNid,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndCreated(offset, limit, iNid,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndChanged(offset, limit, iNid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndPromote(offset, limit, iNid,iLangcode,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndSticky(offset, limit, iNid,iLangcode,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndRevisionTranslationAffected(offset, limit, iNid,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcodeAndDefaultLangcode(offset, limit, iNid,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitleAndUid(offset, limit, iNid,iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitleAndStatus(offset, limit, iNid,iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitleAndCreated(offset, limit, iNid,iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitleAndChanged(offset, limit, iNid,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitleAndPromote(offset, limit, iNid,iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitleAndSticky(offset, limit, iNid,iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitleAndRevisionTranslationAffected(offset, limit, iNid,iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitleAndDefaultLangcode(offset, limit, iNid,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndUidAndStatus(offset, limit, iNid,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndUidAndCreated(offset, limit, iNid,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndUidAndChanged(offset, limit, iNid,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndUidAndPromote(offset, limit, iNid,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndUidAndSticky(offset, limit, iNid,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndUidAndRevisionTranslationAffected(offset, limit, iNid,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndUidAndDefaultLangcode(offset, limit, iNid,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStatusAndCreated(offset, limit, iNid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStatusAndChanged(offset, limit, iNid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStatusAndPromote(offset, limit, iNid,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStatusAndSticky(offset, limit, iNid,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStatusAndRevisionTranslationAffected(offset, limit, iNid,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStatusAndDefaultLangcode(offset, limit, iNid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndCreatedAndChanged(offset, limit, iNid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndCreatedAndPromote(offset, limit, iNid,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndCreatedAndSticky(offset, limit, iNid,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndCreatedAndRevisionTranslationAffected(offset, limit, iNid,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndCreatedAndDefaultLangcode(offset, limit, iNid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndChangedAndPromote(offset, limit, iNid,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndChangedAndSticky(offset, limit, iNid,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndChangedAndRevisionTranslationAffected(offset, limit, iNid,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndChangedAndDefaultLangcode(offset, limit, iNid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndPromoteAndSticky(offset, limit, iNid,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndPromoteAndRevisionTranslationAffected(offset, limit, iNid,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndPromoteAndDefaultLangcode(offset, limit, iNid,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStickyAndRevisionTranslationAffected(offset, limit, iNid,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStickyAndDefaultLangcode(offset, limit, iNid,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iNid,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndTitle(offset, limit, iVid,iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndUid(offset, limit, iVid,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndStatus(offset, limit, iVid,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndCreated(offset, limit, iVid,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndChanged(offset, limit, iVid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndPromote(offset, limit, iVid,iLangcode,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndSticky(offset, limit, iVid,iLangcode,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndRevisionTranslationAffected(offset, limit, iVid,iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcodeAndDefaultLangcode(offset, limit, iVid,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitleAndUid(offset, limit, iVid,iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitleAndStatus(offset, limit, iVid,iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitleAndCreated(offset, limit, iVid,iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitleAndChanged(offset, limit, iVid,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitleAndPromote(offset, limit, iVid,iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitleAndSticky(offset, limit, iVid,iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitleAndRevisionTranslationAffected(offset, limit, iVid,iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitleAndDefaultLangcode(offset, limit, iVid,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndUidAndStatus(offset, limit, iVid,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndUidAndCreated(offset, limit, iVid,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndUidAndChanged(offset, limit, iVid,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndUidAndPromote(offset, limit, iVid,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndUidAndSticky(offset, limit, iVid,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndUidAndRevisionTranslationAffected(offset, limit, iVid,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndUidAndDefaultLangcode(offset, limit, iVid,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStatusAndCreated(offset, limit, iVid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStatusAndChanged(offset, limit, iVid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStatusAndPromote(offset, limit, iVid,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStatusAndSticky(offset, limit, iVid,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStatusAndRevisionTranslationAffected(offset, limit, iVid,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStatusAndDefaultLangcode(offset, limit, iVid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndCreatedAndChanged(offset, limit, iVid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndCreatedAndPromote(offset, limit, iVid,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndCreatedAndSticky(offset, limit, iVid,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndCreatedAndRevisionTranslationAffected(offset, limit, iVid,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndCreatedAndDefaultLangcode(offset, limit, iVid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndChangedAndPromote(offset, limit, iVid,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndChangedAndSticky(offset, limit, iVid,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndChangedAndRevisionTranslationAffected(offset, limit, iVid,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndChangedAndDefaultLangcode(offset, limit, iVid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndPromoteAndSticky(offset, limit, iVid,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndPromoteAndRevisionTranslationAffected(offset, limit, iVid,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndPromoteAndDefaultLangcode(offset, limit, iVid,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStickyAndRevisionTranslationAffected(offset, limit, iVid,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStickyAndDefaultLangcode(offset, limit, iVid,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iVid,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitleAndUid(offset, limit, iLangcode,iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitleAndStatus(offset, limit, iLangcode,iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitleAndCreated(offset, limit, iLangcode,iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitleAndChanged(offset, limit, iLangcode,iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitleAndPromote(offset, limit, iLangcode,iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitleAndSticky(offset, limit, iLangcode,iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitleAndRevisionTranslationAffected(offset, limit, iLangcode,iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitleAndDefaultLangcode(offset, limit, iLangcode,iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndUidAndStatus(offset, limit, iLangcode,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndUidAndCreated(offset, limit, iLangcode,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndUidAndChanged(offset, limit, iLangcode,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndUidAndPromote(offset, limit, iLangcode,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndUidAndSticky(offset, limit, iLangcode,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndUidAndRevisionTranslationAffected(offset, limit, iLangcode,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndUidAndDefaultLangcode(offset, limit, iLangcode,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStatusAndCreated(offset, limit, iLangcode,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStatusAndChanged(offset, limit, iLangcode,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStatusAndPromote(offset, limit, iLangcode,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStatusAndSticky(offset, limit, iLangcode,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStatusAndRevisionTranslationAffected(offset, limit, iLangcode,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStatusAndDefaultLangcode(offset, limit, iLangcode,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndCreatedAndChanged(offset, limit, iLangcode,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndCreatedAndPromote(offset, limit, iLangcode,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndCreatedAndSticky(offset, limit, iLangcode,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndCreatedAndRevisionTranslationAffected(offset, limit, iLangcode,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndCreatedAndDefaultLangcode(offset, limit, iLangcode,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndChangedAndPromote(offset, limit, iLangcode,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndChangedAndSticky(offset, limit, iLangcode,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffected(offset, limit, iLangcode,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndChangedAndDefaultLangcode(offset, limit, iLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndPromoteAndSticky(offset, limit, iLangcode,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndPromoteAndRevisionTranslationAffected(offset, limit, iLangcode,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndPromoteAndDefaultLangcode(offset, limit, iLangcode,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStickyAndRevisionTranslationAffected(offset, limit, iLangcode,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStickyAndDefaultLangcode(offset, limit, iLangcode,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iLangcode,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndUidAndStatus(offset, limit, iTitle,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndUidAndCreated(offset, limit, iTitle,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndUidAndChanged(offset, limit, iTitle,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndUidAndPromote(offset, limit, iTitle,iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndUidAndSticky(offset, limit, iTitle,iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndUidAndRevisionTranslationAffected(offset, limit, iTitle,iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndUidAndDefaultLangcode(offset, limit, iTitle,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStatusAndCreated(offset, limit, iTitle,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStatusAndChanged(offset, limit, iTitle,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStatusAndPromote(offset, limit, iTitle,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStatusAndSticky(offset, limit, iTitle,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStatusAndRevisionTranslationAffected(offset, limit, iTitle,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStatusAndDefaultLangcode(offset, limit, iTitle,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndCreatedAndChanged(offset, limit, iTitle,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndCreatedAndPromote(offset, limit, iTitle,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndCreatedAndSticky(offset, limit, iTitle,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndCreatedAndRevisionTranslationAffected(offset, limit, iTitle,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndCreatedAndDefaultLangcode(offset, limit, iTitle,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndChangedAndPromote(offset, limit, iTitle,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndChangedAndSticky(offset, limit, iTitle,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndChangedAndRevisionTranslationAffected(offset, limit, iTitle,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndChangedAndDefaultLangcode(offset, limit, iTitle,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndPromoteAndSticky(offset, limit, iTitle,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndPromoteAndRevisionTranslationAffected(offset, limit, iTitle,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndPromoteAndDefaultLangcode(offset, limit, iTitle,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStickyAndRevisionTranslationAffected(offset, limit, iTitle,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStickyAndDefaultLangcode(offset, limit, iTitle,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iTitle,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStatusAndCreated(offset, limit, iUid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStatusAndChanged(offset, limit, iUid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStatusAndPromote(offset, limit, iUid,iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStatusAndSticky(offset, limit, iUid,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStatusAndRevisionTranslationAffected(offset, limit, iUid,iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStatusAndDefaultLangcode(offset, limit, iUid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndCreatedAndChanged(offset, limit, iUid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndCreatedAndPromote(offset, limit, iUid,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndCreatedAndSticky(offset, limit, iUid,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndCreatedAndRevisionTranslationAffected(offset, limit, iUid,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndCreatedAndDefaultLangcode(offset, limit, iUid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndChangedAndPromote(offset, limit, iUid,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndChangedAndSticky(offset, limit, iUid,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndChangedAndRevisionTranslationAffected(offset, limit, iUid,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndChangedAndDefaultLangcode(offset, limit, iUid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndPromoteAndSticky(offset, limit, iUid,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndPromoteAndRevisionTranslationAffected(offset, limit, iUid,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndPromoteAndDefaultLangcode(offset, limit, iUid,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStickyAndRevisionTranslationAffected(offset, limit, iUid,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStickyAndDefaultLangcode(offset, limit, iUid,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iUid,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndCreatedAndChanged(offset, limit, iStatus,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndCreatedAndPromote(offset, limit, iStatus,iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndCreatedAndSticky(offset, limit, iStatus,iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndCreatedAndRevisionTranslationAffected(offset, limit, iStatus,iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndCreatedAndDefaultLangcode(offset, limit, iStatus,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndChangedAndPromote(offset, limit, iStatus,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndChangedAndSticky(offset, limit, iStatus,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndChangedAndRevisionTranslationAffected(offset, limit, iStatus,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndChangedAndDefaultLangcode(offset, limit, iStatus,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndPromoteAndSticky(offset, limit, iStatus,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndPromoteAndRevisionTranslationAffected(offset, limit, iStatus,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndPromoteAndDefaultLangcode(offset, limit, iStatus,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndStickyAndRevisionTranslationAffected(offset, limit, iStatus,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndStickyAndDefaultLangcode(offset, limit, iStatus,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iStatus,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndChangedAndPromote(offset, limit, iCreated,iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndChangedAndSticky(offset, limit, iCreated,iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndChangedAndRevisionTranslationAffected(offset, limit, iCreated,iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndChangedAndDefaultLangcode(offset, limit, iCreated,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndPromoteAndSticky(offset, limit, iCreated,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndPromoteAndRevisionTranslationAffected(offset, limit, iCreated,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndPromoteAndDefaultLangcode(offset, limit, iCreated,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndStickyAndRevisionTranslationAffected(offset, limit, iCreated,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndStickyAndDefaultLangcode(offset, limit, iCreated,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iCreated,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndPromoteAndSticky(offset, limit, iChanged,iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndPromoteAndRevisionTranslationAffected(offset, limit, iChanged,iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndPromoteAndDefaultLangcode(offset, limit, iChanged,iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndStickyAndRevisionTranslationAffected(offset, limit, iChanged,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndStickyAndDefaultLangcode(offset, limit, iChanged,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iChanged,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByPromoteAndStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByPromoteAndStickyAndRevisionTranslationAffected(offset, limit, iPromote,iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByPromoteAndStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByPromoteAndStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByPromoteAndStickyAndDefaultLangcode(offset, limit, iPromote,iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByPromoteAndStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iPromote,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStickyAndRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSticky) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStickyAndRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iSticky,iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStickyAndRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndVid(offset, limit, iNid,iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndLangcode(offset, limit, iNid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTitle := self.Args("title").String()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndTitle(offset, limit, iNid,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndUid(offset, limit, iNid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndStatus(offset, limit, iNid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndCreated(offset, limit, iNid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndChanged(offset, limit, iNid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndPromote(offset, limit, iNid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndSticky(offset, limit, iNid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndRevisionTranslationAffected(offset, limit, iNid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByNidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByNidAndDefaultLangcode(offset, limit, iNid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByNidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndLangcode(offset, limit, iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iTitle := self.Args("title").String()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndTitle(offset, limit, iVid,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndUid(offset, limit, iVid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndStatus(offset, limit, iVid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndCreated(offset, limit, iVid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndChanged(offset, limit, iVid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndPromote(offset, limit, iVid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndSticky(offset, limit, iVid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndRevisionTranslationAffected(offset, limit, iVid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByVidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByVidAndDefaultLangcode(offset, limit, iVid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByVidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndTitle(offset, limit, iLangcode,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndUid(offset, limit, iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndStatus(offset, limit, iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndCreated(offset, limit, iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndPromote(offset, limit, iLangcode,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndSticky(offset, limit, iLangcode,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffected(offset, limit, iLangcode,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndUid(offset, limit, iTitle,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndStatus(offset, limit, iTitle,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndCreated(offset, limit, iTitle,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndChanged(offset, limit, iTitle,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndPromote(offset, limit, iTitle,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndSticky(offset, limit, iTitle,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndRevisionTranslationAffected(offset, limit, iTitle,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByTitleAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByTitleAndDefaultLangcode(offset, limit, iTitle,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByTitleAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndStatus(offset, limit, iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndCreated(offset, limit, iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndChanged(offset, limit, iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndPromote(offset, limit, iUid,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndSticky(offset, limit, iUid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndRevisionTranslationAffected(offset, limit, iUid,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByUidAndDefaultLangcode(offset, limit, iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndCreated(offset, limit, iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndChanged(offset, limit, iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndPromote(offset, limit, iStatus,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndSticky(offset, limit, iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndRevisionTranslationAffected(offset, limit, iStatus,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStatusAndDefaultLangcode(offset, limit, iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndChanged(offset, limit, iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndPromote(offset, limit, iCreated,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndSticky(offset, limit, iCreated,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffected(offset, limit, iCreated,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByCreatedAndDefaultLangcode(offset, limit, iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iPromote := self.Args("promote").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndPromote(offset, limit, iChanged,iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndSticky(offset, limit, iChanged,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndRevisionTranslationAffected(offset, limit, iChanged,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByChangedAndDefaultLangcode(offset, limit, iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByPromoteAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByPromoteAndSticky(offset, limit, iPromote,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByPromoteAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffected(offset, limit, iPromote,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByPromoteAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPromote := self.Args("promote").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPromote) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByPromoteAndDefaultLangcode(offset, limit, iPromote,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByPromoteAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStickyAndRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()

	if helper.IsHas(iSticky) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStickyAndRevisionTranslationAffected(offset, limit, iSticky,iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStickyAndRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByStickyAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSticky) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByStickyAndDefaultLangcode(offset, limit, iSticky,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByStickyAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iRevisionTranslationAffected) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcode(offset, limit, iRevisionTranslationAffected,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisions(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisions' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaNid(iNid)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").MustInt()
	if helper.IsHas(iVid) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaVid(iVid)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaTitle(iTitle)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaUid(iUid)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaStatus(iStatus)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPromote := self.Args("promote").MustInt()
	if helper.IsHas(iPromote) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaPromote(iPromote)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSticky := self.Args("sticky").MustInt()
	if helper.IsHas(iSticky) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaSticky(iSticky)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(iRevisionTranslationAffected) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_NodeFieldRevision := model.HasNodeFieldRevisionViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["node_field_revision"] = _NodeFieldRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeFieldRevisionViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaNid(iNid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").MustInt()
	if helper.IsHas(iVid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaVid(iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaTitle(iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaStatus(iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPromote := self.Args("promote").MustInt()
	if helper.IsHas(iPromote) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaPromote(iPromote)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSticky := self.Args("sticky").MustInt()
	if helper.IsHas(iSticky) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaSticky(iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTranslationAffected := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(iRevisionTranslationAffected) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_NodeFieldRevision, _error := model.GetNodeFieldRevisionViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the GetNodeFieldRevisionViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	if helper.IsHas(Nid_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaNid(Nid_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	if helper.IsHas(Vid_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaVid(Vid_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaLangcode(Langcode_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	if helper.IsHas(Title_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaTitle(Title_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	if helper.IsHas(Uid_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaUid(Uid_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	if helper.IsHas(Status_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaStatus(Status_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	if helper.IsHas(Created_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaCreated(Created_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaChanged(Changed_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Promote_ := self.Args("promote").MustInt()
	if helper.IsHas(Promote_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaPromote(Promote_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaPromote's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	if helper.IsHas(Sticky_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaSticky(Sticky_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	if helper.IsHas(RevisionTranslationAffected_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaRevisionTranslationAffected's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iNodeFieldRevision model.NodeFieldRevision
		self.Bind(&iNodeFieldRevision)
		_NodeFieldRevision, _error := model.SetNodeFieldRevisionViaDefaultLangcode(DefaultLangcode_, &iNodeFieldRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeFieldRevision)
	}
	herr.Message = "Can't get to the SetNodeFieldRevisionViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeFieldRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	Vid_ := self.Args("vid").MustInt()
	Langcode_ := self.Args("langcode").String()
	Title_ := self.Args("title").String()
	Uid_ := self.Args("uid").MustInt()
	Status_ := self.Args("status").MustInt()
	Created_ := self.Args("created").MustInt()
	Changed_ := self.Args("changed").MustInt()
	Promote_ := self.Args("promote").MustInt()
	Sticky_ := self.Args("sticky").MustInt()
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Nid_) {
		_error := model.AddNodeFieldRevision(Nid_,Vid_,Langcode_,Title_,Uid_,Status_,Created_,Changed_,Promote_,Sticky_,RevisionTranslationAffected_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNodeFieldRevision's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeFieldRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PostNodeFieldRevision(&iNodeFieldRevision)
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

func PutNodeFieldRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevision(&iNodeFieldRevision)
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

func PutNodeFieldRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaNid(Nid_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaVid(Vid_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaLangcode(Langcode_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaTitle(Title_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaUid(Uid_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaStatus(Status_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaCreated(Created_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaChanged(Changed_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Promote_ := self.Args("promote").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaPromote(Promote_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaSticky(Sticky_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	_int64, _error := model.PutNodeFieldRevisionViaDefaultLangcode(DefaultLangcode_, &iNodeFieldRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaNid(Nid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaVid(Vid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaTitle(Title_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaStatus(Status_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Promote_ := self.Args("promote").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaPromote(Promote_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaSticky(Sticky_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iNodeFieldRevision model.NodeFieldRevision
	self.Bind(&iNodeFieldRevision)
	var iMap = helper.StructToMap(iNodeFieldRevision)
	_error := model.UpdateNodeFieldRevisionViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_int64, _error := model.DeleteNodeFieldRevision(Nid_)
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

func DeleteNodeFieldRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_error := model.DeleteNodeFieldRevisionViaNid(Nid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	_error := model.DeleteNodeFieldRevisionViaVid(Vid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeFieldRevisionViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	_error := model.DeleteNodeFieldRevisionViaTitle(Title_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	_error := model.DeleteNodeFieldRevisionViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	_error := model.DeleteNodeFieldRevisionViaStatus(Status_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	_error := model.DeleteNodeFieldRevisionViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteNodeFieldRevisionViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaPromoteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Promote_ := self.Args("promote").MustInt()
	_error := model.DeleteNodeFieldRevisionViaPromote(Promote_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	_error := model.DeleteNodeFieldRevisionViaSticky(Sticky_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaRevisionTranslationAffectedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTranslationAffected_ := self.Args("revision_translation_affected").MustInt()
	_error := model.DeleteNodeFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeFieldRevisionViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteNodeFieldRevisionViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
