package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetNodeRevisionsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodeRevisionsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node_revisionsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodeRevisionsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionCountViaNidHandler(self *macross.Context) error {
	Nid_ := self.Args("nid").MustInt64()
	_int64 := model.GetNodeRevisionCountViaNid(Nid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revisionCount"] = 0
	}
	m["node_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevisionCountViaVidHandler(self *macross.Context) error {
	Vid_ := self.Args("vid").MustInt()
	_int64 := model.GetNodeRevisionCountViaVid(Vid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revisionCount"] = 0
	}
	m["node_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevisionCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeRevisionCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revisionCount"] = 0
	}
	m["node_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevisionCountViaRevisionTimestampHandler(self *macross.Context) error {
	RevisionTimestamp_ := self.Args("revision_timestamp").MustInt()
	_int64 := model.GetNodeRevisionCountViaRevisionTimestamp(RevisionTimestamp_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revisionCount"] = 0
	}
	m["node_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevisionCountViaRevisionUidHandler(self *macross.Context) error {
	RevisionUid_ := self.Args("revision_uid").MustInt()
	_int64 := model.GetNodeRevisionCountViaRevisionUid(RevisionUid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revisionCount"] = 0
	}
	m["node_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevisionCountViaRevisionLogHandler(self *macross.Context) error {
	RevisionLog_ := self.Args("revision_log").String()
	_int64 := model.GetNodeRevisionCountViaRevisionLog(RevisionLog_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_revisionCount"] = 0
	}
	m["node_revisionCount"] = _int64
	return self.JSON(m)
}

func GetNodeRevisionsViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iNid := self.Args("nid").MustInt64()
	if (offset > 0) && helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsViaNid(offset, limit, iNid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iVid := self.Args("vid").MustInt()
	if (offset > 0) && helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsViaVid(offset, limit, iVid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_NodeRevision, _error := model.GetNodeRevisionsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsViaRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionTimestamp) {
		_NodeRevision, _error := model.GetNodeRevisionsViaRevisionTimestamp(offset, limit, iRevisionTimestamp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsViaRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsViaRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionUid := self.Args("revision_uid").MustInt()
	if (offset > 0) && helper.IsHas(iRevisionUid) {
		_NodeRevision, _error := model.GetNodeRevisionsViaRevisionUid(offset, limit, iRevisionUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsViaRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRevisionLog := self.Args("revision_log").String()
	if (offset > 0) && helper.IsHas(iRevisionLog) {
		_NodeRevision, _error := model.GetNodeRevisionsViaRevisionLog(offset, limit, iRevisionLog, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsViaRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndVidAndLangcode(offset, limit, iNid,iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndVidAndRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndVidAndRevisionTimestamp(offset, limit, iNid,iVid,iRevisionTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndVidAndRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndVidAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndVidAndRevisionUid(offset, limit, iNid,iVid,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndVidAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndVidAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndVidAndRevisionLog(offset, limit, iNid,iVid,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndVidAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndLangcodeAndRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndLangcodeAndRevisionTimestamp(offset, limit, iNid,iLangcode,iRevisionTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndLangcodeAndRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndLangcodeAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndLangcodeAndRevisionUid(offset, limit, iNid,iLangcode,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndLangcodeAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndLangcodeAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndLangcodeAndRevisionLog(offset, limit, iNid,iLangcode,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndLangcodeAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndRevisionTimestampAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndRevisionTimestampAndRevisionUid(offset, limit, iNid,iRevisionTimestamp,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndRevisionTimestampAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndRevisionTimestampAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndRevisionTimestampAndRevisionLog(offset, limit, iNid,iRevisionTimestamp,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndRevisionTimestampAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndRevisionUidAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionUid := self.Args("revision_uid").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndRevisionUidAndRevisionLog(offset, limit, iNid,iRevisionUid,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndRevisionUidAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndLangcodeAndRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndLangcodeAndRevisionTimestamp(offset, limit, iVid,iLangcode,iRevisionTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndLangcodeAndRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndLangcodeAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndLangcodeAndRevisionUid(offset, limit, iVid,iLangcode,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndLangcodeAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndLangcodeAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndLangcodeAndRevisionLog(offset, limit, iVid,iLangcode,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndLangcodeAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndRevisionTimestampAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndRevisionTimestampAndRevisionUid(offset, limit, iVid,iRevisionTimestamp,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndRevisionTimestampAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndRevisionTimestampAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndRevisionTimestampAndRevisionLog(offset, limit, iVid,iRevisionTimestamp,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndRevisionTimestampAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndRevisionUidAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndRevisionUidAndRevisionLog(offset, limit, iVid,iRevisionUid,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndRevisionUidAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision, _error := model.GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionUid(offset, limit, iLangcode,iRevisionTimestamp,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision, _error := model.GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionLog(offset, limit, iLangcode,iRevisionTimestamp,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByLangcodeAndRevisionUidAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUid := self.Args("revision_uid").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision, _error := model.GetNodeRevisionsByLangcodeAndRevisionUidAndRevisionLog(offset, limit, iLangcode,iRevisionUid,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByLangcodeAndRevisionUidAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByRevisionTimestampAndRevisionUidAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iRevisionTimestamp) {
		_NodeRevision, _error := model.GetNodeRevisionsByRevisionTimestampAndRevisionUidAndRevisionLog(offset, limit, iRevisionTimestamp,iRevisionUid,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByRevisionTimestampAndRevisionUidAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndVid(offset, limit, iNid,iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndLangcode(offset, limit, iNid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndRevisionTimestamp(offset, limit, iNid,iRevisionTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndRevisionUid(offset, limit, iNid,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByNidAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionsByNidAndRevisionLog(offset, limit, iNid,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByNidAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndLangcode(offset, limit, iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndRevisionTimestamp(offset, limit, iVid,iRevisionTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndRevisionUid(offset, limit, iVid,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByVidAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionsByVidAndRevisionLog(offset, limit, iVid,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByVidAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByLangcodeAndRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision, _error := model.GetNodeRevisionsByLangcodeAndRevisionTimestamp(offset, limit, iLangcode,iRevisionTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByLangcodeAndRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByLangcodeAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeRevision, _error := model.GetNodeRevisionsByLangcodeAndRevisionUid(offset, limit, iLangcode,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByLangcodeAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByLangcodeAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iLangcode) {
		_NodeRevision, _error := model.GetNodeRevisionsByLangcodeAndRevisionLog(offset, limit, iLangcode,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByLangcodeAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByRevisionTimestampAndRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()

	if helper.IsHas(iRevisionTimestamp) {
		_NodeRevision, _error := model.GetNodeRevisionsByRevisionTimestampAndRevisionUid(offset, limit, iRevisionTimestamp,iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByRevisionTimestampAndRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByRevisionTimestampAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iRevisionTimestamp) {
		_NodeRevision, _error := model.GetNodeRevisionsByRevisionTimestampAndRevisionLog(offset, limit, iRevisionTimestamp,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByRevisionTimestampAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsByRevisionUidAndRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRevisionUid := self.Args("revision_uid").MustInt()
	iRevisionLog := self.Args("revision_log").String()

	if helper.IsHas(iRevisionUid) {
		_NodeRevision, _error := model.GetNodeRevisionsByRevisionUidAndRevisionLog(offset, limit, iRevisionUid,iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionsByRevisionUidAndRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_NodeRevision, _error := model.GetNodeRevisions(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisions' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_NodeRevision := model.HasNodeRevisionViaNid(iNid)
		var m = map[string]interface{}{}
		m["node_revision"] = _NodeRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevisionViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").MustInt()
	if helper.IsHas(iVid) {
		_NodeRevision := model.HasNodeRevisionViaVid(iVid)
		var m = map[string]interface{}{}
		m["node_revision"] = _NodeRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevisionViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision := model.HasNodeRevisionViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node_revision"] = _NodeRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevisionViaRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	if helper.IsHas(iRevisionTimestamp) {
		_NodeRevision := model.HasNodeRevisionViaRevisionTimestamp(iRevisionTimestamp)
		var m = map[string]interface{}{}
		m["node_revision"] = _NodeRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevisionViaRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevisionViaRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionUid := self.Args("revision_uid").MustInt()
	if helper.IsHas(iRevisionUid) {
		_NodeRevision := model.HasNodeRevisionViaRevisionUid(iRevisionUid)
		var m = map[string]interface{}{}
		m["node_revision"] = _NodeRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevisionViaRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionLog := self.Args("revision_log").String()
	if helper.IsHas(iRevisionLog) {
		_NodeRevision := model.HasNodeRevisionViaRevisionLog(iRevisionLog)
		var m = map[string]interface{}{}
		m["node_revision"] = _NodeRevision
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeRevisionViaRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_NodeRevision, _error := model.GetNodeRevisionViaNid(iNid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").MustInt()
	if helper.IsHas(iVid) {
		_NodeRevision, _error := model.GetNodeRevisionViaVid(iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeRevision, _error := model.GetNodeRevisionViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionViaRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionTimestamp := self.Args("revision_timestamp").MustInt()
	if helper.IsHas(iRevisionTimestamp) {
		_NodeRevision, _error := model.GetNodeRevisionViaRevisionTimestamp(iRevisionTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionViaRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionViaRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionUid := self.Args("revision_uid").MustInt()
	if helper.IsHas(iRevisionUid) {
		_NodeRevision, _error := model.GetNodeRevisionViaRevisionUid(iRevisionUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionViaRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRevisionLog := self.Args("revision_log").String()
	if helper.IsHas(iRevisionLog) {
		_NodeRevision, _error := model.GetNodeRevisionViaRevisionLog(iRevisionLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the GetNodeRevisionViaRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	if helper.IsHas(Nid_) {
		var iNodeRevision model.NodeRevision
		self.Bind(&iNodeRevision)
		_NodeRevision, _error := model.SetNodeRevisionViaNid(Nid_, &iNodeRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the SetNodeRevisionViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	if helper.IsHas(Vid_) {
		var iNodeRevision model.NodeRevision
		self.Bind(&iNodeRevision)
		_NodeRevision, _error := model.SetNodeRevisionViaVid(Vid_, &iNodeRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the SetNodeRevisionViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNodeRevision model.NodeRevision
		self.Bind(&iNodeRevision)
		_NodeRevision, _error := model.SetNodeRevisionViaLangcode(Langcode_, &iNodeRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the SetNodeRevisionViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevisionViaRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTimestamp_ := self.Args("revision_timestamp").MustInt()
	if helper.IsHas(RevisionTimestamp_) {
		var iNodeRevision model.NodeRevision
		self.Bind(&iNodeRevision)
		_NodeRevision, _error := model.SetNodeRevisionViaRevisionTimestamp(RevisionTimestamp_, &iNodeRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the SetNodeRevisionViaRevisionTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevisionViaRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUid_ := self.Args("revision_uid").MustInt()
	if helper.IsHas(RevisionUid_) {
		var iNodeRevision model.NodeRevision
		self.Bind(&iNodeRevision)
		_NodeRevision, _error := model.SetNodeRevisionViaRevisionUid(RevisionUid_, &iNodeRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the SetNodeRevisionViaRevisionUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionLog_ := self.Args("revision_log").String()
	if helper.IsHas(RevisionLog_) {
		var iNodeRevision model.NodeRevision
		self.Bind(&iNodeRevision)
		_NodeRevision, _error := model.SetNodeRevisionViaRevisionLog(RevisionLog_, &iNodeRevision)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeRevision)
	}
	herr.Message = "Can't get to the SetNodeRevisionViaRevisionLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	Vid_ := self.Args("vid").MustInt()
	Langcode_ := self.Args("langcode").String()
	RevisionTimestamp_ := self.Args("revision_timestamp").MustInt()
	RevisionUid_ := self.Args("revision_uid").MustInt()
	RevisionLog_ := self.Args("revision_log").String()

	if helper.IsHas(Nid_) {
		_error := model.AddNodeRevision(Nid_,Vid_,Langcode_,RevisionTimestamp_,RevisionUid_,RevisionLog_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNodeRevision's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	_int64, _error := model.PostNodeRevision(&iNodeRevision)
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

func PutNodeRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	_int64, _error := model.PutNodeRevision(&iNodeRevision)
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

func PutNodeRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	_int64, _error := model.PutNodeRevisionViaNid(Nid_, &iNodeRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	_int64, _error := model.PutNodeRevisionViaVid(Vid_, &iNodeRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	_int64, _error := model.PutNodeRevisionViaLangcode(Langcode_, &iNodeRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevisionViaRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTimestamp_ := self.Args("revision_timestamp").MustInt()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	_int64, _error := model.PutNodeRevisionViaRevisionTimestamp(RevisionTimestamp_, &iNodeRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevisionViaRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUid_ := self.Args("revision_uid").MustInt()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	_int64, _error := model.PutNodeRevisionViaRevisionUid(RevisionUid_, &iNodeRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionLog_ := self.Args("revision_log").String()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	_int64, _error := model.PutNodeRevisionViaRevisionLog(RevisionLog_, &iNodeRevision)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	var iMap = helper.StructToMap(iNodeRevision)
	_error := model.UpdateNodeRevisionViaNid(Nid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	var iMap = helper.StructToMap(iNodeRevision)
	_error := model.UpdateNodeRevisionViaVid(Vid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	var iMap = helper.StructToMap(iNodeRevision)
	_error := model.UpdateNodeRevisionViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevisionViaRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTimestamp_ := self.Args("revision_timestamp").MustInt()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	var iMap = helper.StructToMap(iNodeRevision)
	_error := model.UpdateNodeRevisionViaRevisionTimestamp(RevisionTimestamp_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevisionViaRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUid_ := self.Args("revision_uid").MustInt()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	var iMap = helper.StructToMap(iNodeRevision)
	_error := model.UpdateNodeRevisionViaRevisionUid(RevisionUid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionLog_ := self.Args("revision_log").String()
	var iNodeRevision model.NodeRevision
	self.Bind(&iNodeRevision)
	var iMap = helper.StructToMap(iNodeRevision)
	_error := model.UpdateNodeRevisionViaRevisionLog(RevisionLog_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevisionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_int64, _error := model.DeleteNodeRevision(Nid_)
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

func DeleteNodeRevisionViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_error := model.DeleteNodeRevisionViaNid(Nid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevisionViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	_error := model.DeleteNodeRevisionViaVid(Vid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevisionViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeRevisionViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevisionViaRevisionTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionTimestamp_ := self.Args("revision_timestamp").MustInt()
	_error := model.DeleteNodeRevisionViaRevisionTimestamp(RevisionTimestamp_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevisionViaRevisionUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionUid_ := self.Args("revision_uid").MustInt()
	_error := model.DeleteNodeRevisionViaRevisionUid(RevisionUid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeRevisionViaRevisionLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	RevisionLog_ := self.Args("revision_log").String()
	_error := model.DeleteNodeRevisionViaRevisionLog(RevisionLog_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
