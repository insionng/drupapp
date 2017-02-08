package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetNodesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["nodesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeCountViaNidHandler(self *macross.Context) error {
	Nid_ := self.Args("nid").MustInt64()
	_int64 := model.GetNodeCountViaNid(Nid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["nodeCount"] = 0
	}
	m["nodeCount"] = _int64
	return self.JSON(m)
}

func GetNodeCountViaVidHandler(self *macross.Context) error {
	Vid_ := self.Args("vid").MustInt()
	_int64 := model.GetNodeCountViaVid(Vid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["nodeCount"] = 0
	}
	m["nodeCount"] = _int64
	return self.JSON(m)
}

func GetNodeCountViaTypeHandler(self *macross.Context) error {
	Type_ := self.Args("type").String()
	_int64 := model.GetNodeCountViaType(Type_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["nodeCount"] = 0
	}
	m["nodeCount"] = _int64
	return self.JSON(m)
}

func GetNodeCountViaUuidHandler(self *macross.Context) error {
	Uuid_ := self.Args("uuid").String()
	_int64 := model.GetNodeCountViaUuid(Uuid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["nodeCount"] = 0
	}
	m["nodeCount"] = _int64
	return self.JSON(m)
}

func GetNodeCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["nodeCount"] = 0
	}
	m["nodeCount"] = _int64
	return self.JSON(m)
}

func GetNodesViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iNid := self.Args("nid").MustInt64()
	if (offset > 0) && helper.IsHas(iNid) {
		_Node, _error := model.GetNodesViaNid(offset, limit, iNid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iVid := self.Args("vid").MustInt()
	if (offset > 0) && helper.IsHas(iVid) {
		_Node, _error := model.GetNodesViaVid(offset, limit, iVid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iType := self.Args("type").String()
	if (offset > 0) && helper.IsHas(iType) {
		_Node, _error := model.GetNodesViaType(offset, limit, iType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUuid := self.Args("uuid").String()
	if (offset > 0) && helper.IsHas(iUuid) {
		_Node, _error := model.GetNodesViaUuid(offset, limit, iUuid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_Node, _error := model.GetNodesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndVidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndVidAndType(offset, limit, iNid,iVid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndVidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndVidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndVidAndUuid(offset, limit, iNid,iVid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndVidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndVidAndLangcode(offset, limit, iNid,iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndTypeAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndTypeAndUuid(offset, limit, iNid,iType,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndTypeAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndTypeAndLangcode(offset, limit, iNid,iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndUuidAndLangcode(offset, limit, iNid,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByVidAndTypeAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iVid) {
		_Node, _error := model.GetNodesByVidAndTypeAndUuid(offset, limit, iVid,iType,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByVidAndTypeAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByVidAndTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_Node, _error := model.GetNodesByVidAndTypeAndLangcode(offset, limit, iVid,iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByVidAndTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByVidAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_Node, _error := model.GetNodesByVidAndUuidAndLangcode(offset, limit, iVid,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByVidAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByTypeAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iType) {
		_Node, _error := model.GetNodesByTypeAndUuidAndLangcode(offset, limit, iType,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByTypeAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iVid := self.Args("vid").MustInt()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndVid(offset, limit, iNid,iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iType := self.Args("type").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndType(offset, limit, iNid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndUuid(offset, limit, iNid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByNidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodesByNidAndLangcode(offset, limit, iNid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByNidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByVidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iVid) {
		_Node, _error := model.GetNodesByVidAndType(offset, limit, iVid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByVidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByVidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iVid) {
		_Node, _error := model.GetNodesByVidAndUuid(offset, limit, iVid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByVidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByVidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVid := self.Args("vid").MustInt()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iVid) {
		_Node, _error := model.GetNodesByVidAndLangcode(offset, limit, iVid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByVidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByTypeAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iType) {
		_Node, _error := model.GetNodesByTypeAndUuid(offset, limit, iType,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByTypeAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iType) {
		_Node, _error := model.GetNodesByTypeAndLangcode(offset, limit, iType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesByUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUuid) {
		_Node, _error := model.GetNodesByUuidAndLangcode(offset, limit, iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodesByUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Node, _error := model.GetNodes(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodes' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_Node := model.HasNodeViaNid(iNid)
		var m = map[string]interface{}{}
		m["node"] = _Node
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").MustInt()
	if helper.IsHas(iVid) {
		_Node := model.HasNodeViaVid(iVid)
		var m = map[string]interface{}{}
		m["node"] = _Node
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_Node := model.HasNodeViaType(iType)
		var m = map[string]interface{}{}
		m["node"] = _Node
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_Node := model.HasNodeViaUuid(iUuid)
		var m = map[string]interface{}{}
		m["node"] = _Node
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Node := model.HasNodeViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node"] = _Node
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_Node, _error := model.GetNodeViaNid(iNid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodeViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVid := self.Args("vid").MustInt()
	if helper.IsHas(iVid) {
		_Node, _error := model.GetNodeViaVid(iVid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodeViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_Node, _error := model.GetNodeViaType(iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodeViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_Node, _error := model.GetNodeViaUuid(iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodeViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Node, _error := model.GetNodeViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the GetNodeViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	if helper.IsHas(Nid_) {
		var iNode model.Node
		self.Bind(&iNode)
		_Node, _error := model.SetNodeViaNid(Nid_, &iNode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the SetNodeViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	if helper.IsHas(Vid_) {
		var iNode model.Node
		self.Bind(&iNode)
		_Node, _error := model.SetNodeViaVid(Vid_, &iNode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the SetNodeViaVid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	if helper.IsHas(Type_) {
		var iNode model.Node
		self.Bind(&iNode)
		_Node, _error := model.SetNodeViaType(Type_, &iNode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the SetNodeViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	if helper.IsHas(Uuid_) {
		var iNode model.Node
		self.Bind(&iNode)
		_Node, _error := model.SetNodeViaUuid(Uuid_, &iNode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the SetNodeViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNode model.Node
		self.Bind(&iNode)
		_Node, _error := model.SetNodeViaLangcode(Langcode_, &iNode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Node)
	}
	herr.Message = "Can't get to the SetNodeViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	Vid_ := self.Args("vid").MustInt()
	Type_ := self.Args("type").String()
	Uuid_ := self.Args("uuid").String()
	Langcode_ := self.Args("langcode").String()

	if helper.IsHas(Nid_) {
		_error := model.AddNode(Nid_,Vid_,Type_,Uuid_,Langcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNode model.Node
	self.Bind(&iNode)
	_int64, _error := model.PostNode(&iNode)
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

func PutNodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNode model.Node
	self.Bind(&iNode)
	_int64, _error := model.PutNode(&iNode)
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

func PutNodeViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNode model.Node
	self.Bind(&iNode)
	_int64, _error := model.PutNodeViaNid(Nid_, &iNode)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	var iNode model.Node
	self.Bind(&iNode)
	_int64, _error := model.PutNodeViaVid(Vid_, &iNode)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iNode model.Node
	self.Bind(&iNode)
	_int64, _error := model.PutNodeViaType(Type_, &iNode)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iNode model.Node
	self.Bind(&iNode)
	_int64, _error := model.PutNodeViaUuid(Uuid_, &iNode)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNode model.Node
	self.Bind(&iNode)
	_int64, _error := model.PutNodeViaLangcode(Langcode_, &iNode)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNode model.Node
	self.Bind(&iNode)
	var iMap = helper.StructToMap(iNode)
	_error := model.UpdateNodeViaNid(Nid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	var iNode model.Node
	self.Bind(&iNode)
	var iMap = helper.StructToMap(iNode)
	_error := model.UpdateNodeViaVid(Vid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iNode model.Node
	self.Bind(&iNode)
	var iMap = helper.StructToMap(iNode)
	_error := model.UpdateNodeViaType(Type_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iNode model.Node
	self.Bind(&iNode)
	var iMap = helper.StructToMap(iNode)
	_error := model.UpdateNodeViaUuid(Uuid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNode model.Node
	self.Bind(&iNode)
	var iMap = helper.StructToMap(iNode)
	_error := model.UpdateNodeViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_int64, _error := model.DeleteNode(Nid_)
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

func DeleteNodeViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_error := model.DeleteNodeViaNid(Nid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeViaVidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Vid_ := self.Args("vid").MustInt()
	_error := model.DeleteNodeViaVid(Vid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	_error := model.DeleteNodeViaType(Type_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	_error := model.DeleteNodeViaUuid(Uuid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
