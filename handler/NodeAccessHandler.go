package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetNodeAccessesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetNodeAccessesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["node_accesssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetNodeAccessesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessCountViaNidHandler(self *macross.Context) error {
	Nid_ := self.Args("nid").MustInt64()
	_int64 := model.GetNodeAccessCountViaNid(Nid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_accessCount"] = 0
	}
	m["node_accessCount"] = _int64
	return self.JSON(m)
}

func GetNodeAccessCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetNodeAccessCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_accessCount"] = 0
	}
	m["node_accessCount"] = _int64
	return self.JSON(m)
}

func GetNodeAccessCountViaFallbackHandler(self *macross.Context) error {
	Fallback_ := self.Args("fallback").MustInt()
	_int64 := model.GetNodeAccessCountViaFallback(Fallback_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_accessCount"] = 0
	}
	m["node_accessCount"] = _int64
	return self.JSON(m)
}

func GetNodeAccessCountViaGidHandler(self *macross.Context) error {
	Gid_ := self.Args("gid").MustInt()
	_int64 := model.GetNodeAccessCountViaGid(Gid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_accessCount"] = 0
	}
	m["node_accessCount"] = _int64
	return self.JSON(m)
}

func GetNodeAccessCountViaRealmHandler(self *macross.Context) error {
	Realm_ := self.Args("realm").String()
	_int64 := model.GetNodeAccessCountViaRealm(Realm_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_accessCount"] = 0
	}
	m["node_accessCount"] = _int64
	return self.JSON(m)
}

func GetNodeAccessCountViaGrantViewHandler(self *macross.Context) error {
	GrantView_ := self.Args("grant_view").MustInt()
	_int64 := model.GetNodeAccessCountViaGrantView(GrantView_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_accessCount"] = 0
	}
	m["node_accessCount"] = _int64
	return self.JSON(m)
}

func GetNodeAccessCountViaGrantUpdateHandler(self *macross.Context) error {
	GrantUpdate_ := self.Args("grant_update").MustInt()
	_int64 := model.GetNodeAccessCountViaGrantUpdate(GrantUpdate_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_accessCount"] = 0
	}
	m["node_accessCount"] = _int64
	return self.JSON(m)
}

func GetNodeAccessCountViaGrantDeleteHandler(self *macross.Context) error {
	GrantDelete_ := self.Args("grant_delete").MustInt()
	_int64 := model.GetNodeAccessCountViaGrantDelete(GrantDelete_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["node_accessCount"] = 0
	}
	m["node_accessCount"] = _int64
	return self.JSON(m)
}

func GetNodeAccessesViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iNid := self.Args("nid").MustInt64()
	if (offset > 0) && helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesViaNid(offset, limit, iNid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesViaFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFallback := self.Args("fallback").MustInt()
	if (offset > 0) && helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesViaFallback(offset, limit, iFallback, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesViaFallback's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesViaGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iGid := self.Args("gid").MustInt()
	if (offset > 0) && helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesViaGid(offset, limit, iGid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesViaGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesViaRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRealm := self.Args("realm").String()
	if (offset > 0) && helper.IsHas(iRealm) {
		_NodeAccess, _error := model.GetNodeAccessesViaRealm(offset, limit, iRealm, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesViaRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesViaGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iGrantView := self.Args("grant_view").MustInt()
	if (offset > 0) && helper.IsHas(iGrantView) {
		_NodeAccess, _error := model.GetNodeAccessesViaGrantView(offset, limit, iGrantView, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesViaGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesViaGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iGrantUpdate := self.Args("grant_update").MustInt()
	if (offset > 0) && helper.IsHas(iGrantUpdate) {
		_NodeAccess, _error := model.GetNodeAccessesViaGrantUpdate(offset, limit, iGrantUpdate, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesViaGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesViaGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iGrantDelete := self.Args("grant_delete").MustInt()
	if (offset > 0) && helper.IsHas(iGrantDelete) {
		_NodeAccess, _error := model.GetNodeAccessesViaGrantDelete(offset, limit, iGrantDelete, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesViaGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndLangcodeAndFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iFallback := self.Args("fallback").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndLangcodeAndFallback(offset, limit, iNid,iLangcode,iFallback)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndLangcodeAndFallback's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndLangcodeAndGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iGid := self.Args("gid").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndLangcodeAndGid(offset, limit, iNid,iLangcode,iGid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndLangcodeAndGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndLangcodeAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndLangcodeAndRealm(offset, limit, iNid,iLangcode,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndLangcodeAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndLangcodeAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndLangcodeAndGrantView(offset, limit, iNid,iLangcode,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndLangcodeAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndLangcodeAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndLangcodeAndGrantUpdate(offset, limit, iNid,iLangcode,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndLangcodeAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndLangcodeAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndLangcodeAndGrantDelete(offset, limit, iNid,iLangcode,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndLangcodeAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndFallbackAndGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iFallback := self.Args("fallback").MustInt()
	iGid := self.Args("gid").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndFallbackAndGid(offset, limit, iNid,iFallback,iGid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndFallbackAndGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndFallbackAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iFallback := self.Args("fallback").MustInt()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndFallbackAndRealm(offset, limit, iNid,iFallback,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndFallbackAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndFallbackAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iFallback := self.Args("fallback").MustInt()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndFallbackAndGrantView(offset, limit, iNid,iFallback,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndFallbackAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndFallbackAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iFallback := self.Args("fallback").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndFallbackAndGrantUpdate(offset, limit, iNid,iFallback,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndFallbackAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndFallbackAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iFallback := self.Args("fallback").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndFallbackAndGrantDelete(offset, limit, iNid,iFallback,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndFallbackAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGidAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGid := self.Args("gid").MustInt()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGidAndRealm(offset, limit, iNid,iGid,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGidAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGidAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGid := self.Args("gid").MustInt()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGidAndGrantView(offset, limit, iNid,iGid,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGidAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGidAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGid := self.Args("gid").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGidAndGrantUpdate(offset, limit, iNid,iGid,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGidAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGidAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGid := self.Args("gid").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGidAndGrantDelete(offset, limit, iNid,iGid,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGidAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndRealmAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRealm := self.Args("realm").String()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndRealmAndGrantView(offset, limit, iNid,iRealm,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndRealmAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndRealmAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRealm := self.Args("realm").String()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndRealmAndGrantUpdate(offset, limit, iNid,iRealm,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndRealmAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndRealmAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRealm := self.Args("realm").String()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndRealmAndGrantDelete(offset, limit, iNid,iRealm,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndRealmAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGrantViewAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGrantViewAndGrantUpdate(offset, limit, iNid,iGrantView,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGrantViewAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGrantViewAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGrantViewAndGrantDelete(offset, limit, iNid,iGrantView,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGrantViewAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGrantUpdateAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGrantUpdate := self.Args("grant_update").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGrantUpdateAndGrantDelete(offset, limit, iNid,iGrantUpdate,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGrantUpdateAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndFallbackAndGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFallback := self.Args("fallback").MustInt()
	iGid := self.Args("gid").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndFallbackAndGid(offset, limit, iLangcode,iFallback,iGid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndFallbackAndGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndFallbackAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFallback := self.Args("fallback").MustInt()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndFallbackAndRealm(offset, limit, iLangcode,iFallback,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndFallbackAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndFallbackAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFallback := self.Args("fallback").MustInt()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndFallbackAndGrantView(offset, limit, iLangcode,iFallback,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndFallbackAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndFallbackAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFallback := self.Args("fallback").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndFallbackAndGrantUpdate(offset, limit, iLangcode,iFallback,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndFallbackAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndFallbackAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFallback := self.Args("fallback").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndFallbackAndGrantDelete(offset, limit, iLangcode,iFallback,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndFallbackAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGidAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGid := self.Args("gid").MustInt()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGidAndRealm(offset, limit, iLangcode,iGid,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGidAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGidAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGid := self.Args("gid").MustInt()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGidAndGrantView(offset, limit, iLangcode,iGid,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGidAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGidAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGid := self.Args("gid").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGidAndGrantUpdate(offset, limit, iLangcode,iGid,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGidAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGidAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGid := self.Args("gid").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGidAndGrantDelete(offset, limit, iLangcode,iGid,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGidAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndRealmAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRealm := self.Args("realm").String()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndRealmAndGrantView(offset, limit, iLangcode,iRealm,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndRealmAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndRealmAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRealm := self.Args("realm").String()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndRealmAndGrantUpdate(offset, limit, iLangcode,iRealm,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndRealmAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndRealmAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRealm := self.Args("realm").String()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndRealmAndGrantDelete(offset, limit, iLangcode,iRealm,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndRealmAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGrantViewAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGrantViewAndGrantUpdate(offset, limit, iLangcode,iGrantView,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGrantViewAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGrantViewAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGrantViewAndGrantDelete(offset, limit, iLangcode,iGrantView,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGrantViewAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGrantUpdateAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGrantUpdate := self.Args("grant_update").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGrantUpdateAndGrantDelete(offset, limit, iLangcode,iGrantUpdate,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGrantUpdateAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGidAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGid := self.Args("gid").MustInt()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGidAndRealm(offset, limit, iFallback,iGid,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGidAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGidAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGidAndGrantView(offset, limit, iFallback,iGid,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGidAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGidAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGidAndGrantUpdate(offset, limit, iFallback,iGid,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGidAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGidAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGidAndGrantDelete(offset, limit, iFallback,iGid,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGidAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndRealmAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iRealm := self.Args("realm").String()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndRealmAndGrantView(offset, limit, iFallback,iRealm,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndRealmAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndRealmAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iRealm := self.Args("realm").String()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndRealmAndGrantUpdate(offset, limit, iFallback,iRealm,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndRealmAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndRealmAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iRealm := self.Args("realm").String()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndRealmAndGrantDelete(offset, limit, iFallback,iRealm,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndRealmAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGrantViewAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGrantViewAndGrantUpdate(offset, limit, iFallback,iGrantView,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGrantViewAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGrantViewAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGrantViewAndGrantDelete(offset, limit, iFallback,iGrantView,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGrantViewAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGrantUpdateAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGrantUpdateAndGrantDelete(offset, limit, iFallback,iGrantUpdate,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGrantUpdateAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndRealmAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iRealm := self.Args("realm").String()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndRealmAndGrantView(offset, limit, iGid,iRealm,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndRealmAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndRealmAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iRealm := self.Args("realm").String()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndRealmAndGrantUpdate(offset, limit, iGid,iRealm,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndRealmAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndRealmAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iRealm := self.Args("realm").String()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndRealmAndGrantDelete(offset, limit, iGid,iRealm,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndRealmAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndGrantViewAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndGrantViewAndGrantUpdate(offset, limit, iGid,iGrantView,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndGrantViewAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndGrantViewAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndGrantViewAndGrantDelete(offset, limit, iGid,iGrantView,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndGrantViewAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndGrantUpdateAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndGrantUpdateAndGrantDelete(offset, limit, iGid,iGrantUpdate,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndGrantUpdateAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByRealmAndGrantViewAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRealm := self.Args("realm").String()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iRealm) {
		_NodeAccess, _error := model.GetNodeAccessesByRealmAndGrantViewAndGrantUpdate(offset, limit, iRealm,iGrantView,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByRealmAndGrantViewAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByRealmAndGrantViewAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRealm := self.Args("realm").String()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iRealm) {
		_NodeAccess, _error := model.GetNodeAccessesByRealmAndGrantViewAndGrantDelete(offset, limit, iRealm,iGrantView,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByRealmAndGrantViewAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByRealmAndGrantUpdateAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRealm := self.Args("realm").String()
	iGrantUpdate := self.Args("grant_update").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iRealm) {
		_NodeAccess, _error := model.GetNodeAccessesByRealmAndGrantUpdateAndGrantDelete(offset, limit, iRealm,iGrantUpdate,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByRealmAndGrantUpdateAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGrantViewAndGrantUpdateAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iGrantView) {
		_NodeAccess, _error := model.GetNodeAccessesByGrantViewAndGrantUpdateAndGrantDelete(offset, limit, iGrantView,iGrantUpdate,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGrantViewAndGrantUpdateAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndLangcode(offset, limit, iNid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iFallback := self.Args("fallback").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndFallback(offset, limit, iNid,iFallback)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndFallback's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGid := self.Args("gid").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGid(offset, limit, iNid,iGid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndRealm(offset, limit, iNid,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGrantView(offset, limit, iNid,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGrantUpdate(offset, limit, iNid,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByNidAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessesByNidAndGrantDelete(offset, limit, iNid,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByNidAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFallback := self.Args("fallback").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndFallback(offset, limit, iLangcode,iFallback)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndFallback's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGid := self.Args("gid").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGid(offset, limit, iLangcode,iGid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndRealm(offset, limit, iLangcode,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGrantView(offset, limit, iLangcode,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGrantUpdate(offset, limit, iLangcode,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByLangcodeAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessesByLangcodeAndGrantDelete(offset, limit, iLangcode,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByLangcodeAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGid := self.Args("gid").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGid(offset, limit, iFallback,iGid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndRealm(offset, limit, iFallback,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGrantView(offset, limit, iFallback,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGrantUpdate(offset, limit, iFallback,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByFallbackAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFallback := self.Args("fallback").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessesByFallbackAndGrantDelete(offset, limit, iFallback,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByFallbackAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iRealm := self.Args("realm").String()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndRealm(offset, limit, iGid,iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndGrantView(offset, limit, iGid,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndGrantUpdate(offset, limit, iGid,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGidAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGid := self.Args("gid").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessesByGidAndGrantDelete(offset, limit, iGid,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGidAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByRealmAndGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRealm := self.Args("realm").String()
	iGrantView := self.Args("grant_view").MustInt()

	if helper.IsHas(iRealm) {
		_NodeAccess, _error := model.GetNodeAccessesByRealmAndGrantView(offset, limit, iRealm,iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByRealmAndGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByRealmAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRealm := self.Args("realm").String()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iRealm) {
		_NodeAccess, _error := model.GetNodeAccessesByRealmAndGrantUpdate(offset, limit, iRealm,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByRealmAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByRealmAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRealm := self.Args("realm").String()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iRealm) {
		_NodeAccess, _error := model.GetNodeAccessesByRealmAndGrantDelete(offset, limit, iRealm,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByRealmAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGrantViewAndGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()

	if helper.IsHas(iGrantView) {
		_NodeAccess, _error := model.GetNodeAccessesByGrantViewAndGrantUpdate(offset, limit, iGrantView,iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGrantViewAndGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGrantViewAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGrantView := self.Args("grant_view").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iGrantView) {
		_NodeAccess, _error := model.GetNodeAccessesByGrantViewAndGrantDelete(offset, limit, iGrantView,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGrantViewAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesByGrantUpdateAndGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iGrantUpdate := self.Args("grant_update").MustInt()
	iGrantDelete := self.Args("grant_delete").MustInt()

	if helper.IsHas(iGrantUpdate) {
		_NodeAccess, _error := model.GetNodeAccessesByGrantUpdateAndGrantDelete(offset, limit, iGrantUpdate,iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessesByGrantUpdateAndGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_NodeAccess, _error := model.GetNodeAccesses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccesses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeAccessViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_NodeAccess := model.HasNodeAccessViaNid(iNid)
		var m = map[string]interface{}{}
		m["node_access"] = _NodeAccess
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeAccessViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeAccessViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeAccess := model.HasNodeAccessViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["node_access"] = _NodeAccess
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeAccessViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeAccessViaFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFallback := self.Args("fallback").MustInt()
	if helper.IsHas(iFallback) {
		_NodeAccess := model.HasNodeAccessViaFallback(iFallback)
		var m = map[string]interface{}{}
		m["node_access"] = _NodeAccess
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeAccessViaFallback's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeAccessViaGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iGid := self.Args("gid").MustInt()
	if helper.IsHas(iGid) {
		_NodeAccess := model.HasNodeAccessViaGid(iGid)
		var m = map[string]interface{}{}
		m["node_access"] = _NodeAccess
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeAccessViaGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeAccessViaRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRealm := self.Args("realm").String()
	if helper.IsHas(iRealm) {
		_NodeAccess := model.HasNodeAccessViaRealm(iRealm)
		var m = map[string]interface{}{}
		m["node_access"] = _NodeAccess
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeAccessViaRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeAccessViaGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iGrantView := self.Args("grant_view").MustInt()
	if helper.IsHas(iGrantView) {
		_NodeAccess := model.HasNodeAccessViaGrantView(iGrantView)
		var m = map[string]interface{}{}
		m["node_access"] = _NodeAccess
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeAccessViaGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeAccessViaGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iGrantUpdate := self.Args("grant_update").MustInt()
	if helper.IsHas(iGrantUpdate) {
		_NodeAccess := model.HasNodeAccessViaGrantUpdate(iGrantUpdate)
		var m = map[string]interface{}{}
		m["node_access"] = _NodeAccess
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeAccessViaGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasNodeAccessViaGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iGrantDelete := self.Args("grant_delete").MustInt()
	if helper.IsHas(iGrantDelete) {
		_NodeAccess := model.HasNodeAccessViaGrantDelete(iGrantDelete)
		var m = map[string]interface{}{}
		m["node_access"] = _NodeAccess
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasNodeAccessViaGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_NodeAccess, _error := model.GetNodeAccessViaNid(iNid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_NodeAccess, _error := model.GetNodeAccessViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessViaFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFallback := self.Args("fallback").MustInt()
	if helper.IsHas(iFallback) {
		_NodeAccess, _error := model.GetNodeAccessViaFallback(iFallback)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessViaFallback's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessViaGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iGid := self.Args("gid").MustInt()
	if helper.IsHas(iGid) {
		_NodeAccess, _error := model.GetNodeAccessViaGid(iGid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessViaGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessViaRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRealm := self.Args("realm").String()
	if helper.IsHas(iRealm) {
		_NodeAccess, _error := model.GetNodeAccessViaRealm(iRealm)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessViaRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessViaGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iGrantView := self.Args("grant_view").MustInt()
	if helper.IsHas(iGrantView) {
		_NodeAccess, _error := model.GetNodeAccessViaGrantView(iGrantView)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessViaGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessViaGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iGrantUpdate := self.Args("grant_update").MustInt()
	if helper.IsHas(iGrantUpdate) {
		_NodeAccess, _error := model.GetNodeAccessViaGrantUpdate(iGrantUpdate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessViaGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetNodeAccessViaGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iGrantDelete := self.Args("grant_delete").MustInt()
	if helper.IsHas(iGrantDelete) {
		_NodeAccess, _error := model.GetNodeAccessViaGrantDelete(iGrantDelete)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the GetNodeAccessViaGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeAccessViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	if helper.IsHas(Nid_) {
		var iNodeAccess model.NodeAccess
		self.Bind(&iNodeAccess)
		_NodeAccess, _error := model.SetNodeAccessViaNid(Nid_, &iNodeAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the SetNodeAccessViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeAccessViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iNodeAccess model.NodeAccess
		self.Bind(&iNodeAccess)
		_NodeAccess, _error := model.SetNodeAccessViaLangcode(Langcode_, &iNodeAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the SetNodeAccessViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeAccessViaFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fallback_ := self.Args("fallback").MustInt()
	if helper.IsHas(Fallback_) {
		var iNodeAccess model.NodeAccess
		self.Bind(&iNodeAccess)
		_NodeAccess, _error := model.SetNodeAccessViaFallback(Fallback_, &iNodeAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the SetNodeAccessViaFallback's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeAccessViaGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Gid_ := self.Args("gid").MustInt()
	if helper.IsHas(Gid_) {
		var iNodeAccess model.NodeAccess
		self.Bind(&iNodeAccess)
		_NodeAccess, _error := model.SetNodeAccessViaGid(Gid_, &iNodeAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the SetNodeAccessViaGid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeAccessViaRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Realm_ := self.Args("realm").String()
	if helper.IsHas(Realm_) {
		var iNodeAccess model.NodeAccess
		self.Bind(&iNodeAccess)
		_NodeAccess, _error := model.SetNodeAccessViaRealm(Realm_, &iNodeAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the SetNodeAccessViaRealm's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeAccessViaGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantView_ := self.Args("grant_view").MustInt()
	if helper.IsHas(GrantView_) {
		var iNodeAccess model.NodeAccess
		self.Bind(&iNodeAccess)
		_NodeAccess, _error := model.SetNodeAccessViaGrantView(GrantView_, &iNodeAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the SetNodeAccessViaGrantView's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeAccessViaGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantUpdate_ := self.Args("grant_update").MustInt()
	if helper.IsHas(GrantUpdate_) {
		var iNodeAccess model.NodeAccess
		self.Bind(&iNodeAccess)
		_NodeAccess, _error := model.SetNodeAccessViaGrantUpdate(GrantUpdate_, &iNodeAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the SetNodeAccessViaGrantUpdate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetNodeAccessViaGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantDelete_ := self.Args("grant_delete").MustInt()
	if helper.IsHas(GrantDelete_) {
		var iNodeAccess model.NodeAccess
		self.Bind(&iNodeAccess)
		_NodeAccess, _error := model.SetNodeAccessViaGrantDelete(GrantDelete_, &iNodeAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_NodeAccess)
	}
	herr.Message = "Can't get to the SetNodeAccessViaGrantDelete's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddNodeAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	Langcode_ := self.Args("langcode").String()
	Fallback_ := self.Args("fallback").MustInt()
	Gid_ := self.Args("gid").MustInt()
	Realm_ := self.Args("realm").String()
	GrantView_ := self.Args("grant_view").MustInt()
	GrantUpdate_ := self.Args("grant_update").MustInt()
	GrantDelete_ := self.Args("grant_delete").MustInt()

	if helper.IsHas(Nid_) {
		_error := model.AddNodeAccess(Nid_,Langcode_,Fallback_,Gid_,Realm_,GrantView_,GrantUpdate_,GrantDelete_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddNodeAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostNodeAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PostNodeAccess(&iNodeAccess)
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

func PutNodeAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccess(&iNodeAccess)
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

func PutNodeAccessViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccessViaNid(Nid_, &iNodeAccess)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeAccessViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccessViaLangcode(Langcode_, &iNodeAccess)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeAccessViaFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fallback_ := self.Args("fallback").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccessViaFallback(Fallback_, &iNodeAccess)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeAccessViaGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Gid_ := self.Args("gid").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccessViaGid(Gid_, &iNodeAccess)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeAccessViaRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Realm_ := self.Args("realm").String()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccessViaRealm(Realm_, &iNodeAccess)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeAccessViaGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantView_ := self.Args("grant_view").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccessViaGrantView(GrantView_, &iNodeAccess)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeAccessViaGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantUpdate_ := self.Args("grant_update").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccessViaGrantUpdate(GrantUpdate_, &iNodeAccess)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutNodeAccessViaGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantDelete_ := self.Args("grant_delete").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	_int64, _error := model.PutNodeAccessViaGrantDelete(GrantDelete_, &iNodeAccess)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeAccessViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	var iMap = helper.StructToMap(iNodeAccess)
	_error := model.UpdateNodeAccessViaNid(Nid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeAccessViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	var iMap = helper.StructToMap(iNodeAccess)
	_error := model.UpdateNodeAccessViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeAccessViaFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fallback_ := self.Args("fallback").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	var iMap = helper.StructToMap(iNodeAccess)
	_error := model.UpdateNodeAccessViaFallback(Fallback_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeAccessViaGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Gid_ := self.Args("gid").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	var iMap = helper.StructToMap(iNodeAccess)
	_error := model.UpdateNodeAccessViaGid(Gid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeAccessViaRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Realm_ := self.Args("realm").String()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	var iMap = helper.StructToMap(iNodeAccess)
	_error := model.UpdateNodeAccessViaRealm(Realm_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeAccessViaGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantView_ := self.Args("grant_view").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	var iMap = helper.StructToMap(iNodeAccess)
	_error := model.UpdateNodeAccessViaGrantView(GrantView_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeAccessViaGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantUpdate_ := self.Args("grant_update").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	var iMap = helper.StructToMap(iNodeAccess)
	_error := model.UpdateNodeAccessViaGrantUpdate(GrantUpdate_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateNodeAccessViaGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantDelete_ := self.Args("grant_delete").MustInt()
	var iNodeAccess model.NodeAccess
	self.Bind(&iNodeAccess)
	var iMap = helper.StructToMap(iNodeAccess)
	_error := model.UpdateNodeAccessViaGrantDelete(GrantDelete_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_int64, _error := model.DeleteNodeAccess(Nid_)
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

func DeleteNodeAccessViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_error := model.DeleteNodeAccessViaNid(Nid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeAccessViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteNodeAccessViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeAccessViaFallbackHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fallback_ := self.Args("fallback").MustInt()
	_error := model.DeleteNodeAccessViaFallback(Fallback_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeAccessViaGidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Gid_ := self.Args("gid").MustInt()
	_error := model.DeleteNodeAccessViaGid(Gid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeAccessViaRealmHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Realm_ := self.Args("realm").String()
	_error := model.DeleteNodeAccessViaRealm(Realm_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeAccessViaGrantViewHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantView_ := self.Args("grant_view").MustInt()
	_error := model.DeleteNodeAccessViaGrantView(GrantView_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeAccessViaGrantUpdateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantUpdate_ := self.Args("grant_update").MustInt()
	_error := model.DeleteNodeAccessViaGrantUpdate(GrantUpdate_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteNodeAccessViaGrantDeleteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	GrantDelete_ := self.Args("grant_delete").MustInt()
	_error := model.DeleteNodeAccessViaGrantDelete(GrantDelete_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
