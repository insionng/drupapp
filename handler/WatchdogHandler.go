package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetWatchdogsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetWatchdogsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["watchdogsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetWatchdogsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogCountViaWidHandler(self *macross.Context) error {
	Wid_ := self.Args("wid").MustInt64()
	_int64 := model.GetWatchdogCountViaWid(Wid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt()
	_int64 := model.GetWatchdogCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaTypeHandler(self *macross.Context) error {
	Type_ := self.Args("type").String()
	_int64 := model.GetWatchdogCountViaType(Type_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaMessageHandler(self *macross.Context) error {
	Message_ := self.Args("message").String()
	_int64 := model.GetWatchdogCountViaMessage(Message_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaVariablesHandler(self *macross.Context) error {
	Variables_ := self.Args("variables").Bytes()
	_int64 := model.GetWatchdogCountViaVariables(Variables_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaSeverityHandler(self *macross.Context) error {
	Severity_ := self.Args("severity").MustInt()
	_int64 := model.GetWatchdogCountViaSeverity(Severity_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaLinkHandler(self *macross.Context) error {
	Link_ := self.Args("link").String()
	_int64 := model.GetWatchdogCountViaLink(Link_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaLocationHandler(self *macross.Context) error {
	Location_ := self.Args("location").String()
	_int64 := model.GetWatchdogCountViaLocation(Location_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaRefererHandler(self *macross.Context) error {
	Referer_ := self.Args("referer").String()
	_int64 := model.GetWatchdogCountViaReferer(Referer_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaHostnameHandler(self *macross.Context) error {
	Hostname_ := self.Args("hostname").String()
	_int64 := model.GetWatchdogCountViaHostname(Hostname_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogCountViaTimestampHandler(self *macross.Context) error {
	Timestamp_ := self.Args("timestamp").MustInt()
	_int64 := model.GetWatchdogCountViaTimestamp(Timestamp_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["watchdogCount"] = 0
	}
	m["watchdogCount"] = _int64
	return self.JSON(m)
}

func GetWatchdogsViaWidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iWid := self.Args("wid").MustInt64()
	if (offset > 0) && helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsViaWid(offset, limit, iWid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaWid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt()
	if (offset > 0) && helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iType := self.Args("type").String()
	if (offset > 0) && helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsViaType(offset, limit, iType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMessage := self.Args("message").String()
	if (offset > 0) && helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsViaMessage(offset, limit, iMessage, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iVariables := self.Args("variables").Bytes()
	if (offset > 0) && helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsViaVariables(offset, limit, iVariables, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSeverity := self.Args("severity").MustInt()
	if (offset > 0) && helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsViaSeverity(offset, limit, iSeverity, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLink := self.Args("link").String()
	if (offset > 0) && helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsViaLink(offset, limit, iLink, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLocation := self.Args("location").String()
	if (offset > 0) && helper.IsHas(iLocation) {
		_Watchdog, _error := model.GetWatchdogsViaLocation(offset, limit, iLocation, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iReferer := self.Args("referer").String()
	if (offset > 0) && helper.IsHas(iReferer) {
		_Watchdog, _error := model.GetWatchdogsViaReferer(offset, limit, iReferer, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iHostname := self.Args("hostname").String()
	if (offset > 0) && helper.IsHas(iHostname) {
		_Watchdog, _error := model.GetWatchdogsViaHostname(offset, limit, iHostname, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTimestamp := self.Args("timestamp").MustInt()
	if (offset > 0) && helper.IsHas(iTimestamp) {
		_Watchdog, _error := model.GetWatchdogsViaTimestamp(offset, limit, iTimestamp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndType(offset, limit, iWid,iUid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndMessage(offset, limit, iWid,iUid,iMessage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndVariables(offset, limit, iWid,iUid,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndSeverity(offset, limit, iWid,iUid,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iLink := self.Args("link").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndLink(offset, limit, iWid,iUid,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iLocation := self.Args("location").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndLocation(offset, limit, iWid,iUid,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndReferer(offset, limit, iWid,iUid,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndHostname(offset, limit, iWid,iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUidAndTimestamp(offset, limit, iWid,iUid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeAndMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTypeAndMessage(offset, limit, iWid,iType,iMessage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTypeAndMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTypeAndVariables(offset, limit, iWid,iType,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTypeAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTypeAndSeverity(offset, limit, iWid,iType,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTypeAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()
	iLink := self.Args("link").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTypeAndLink(offset, limit, iWid,iType,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTypeAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTypeAndLocation(offset, limit, iWid,iType,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTypeAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTypeAndReferer(offset, limit, iWid,iType,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTypeAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTypeAndHostname(offset, limit, iWid,iType,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTypeAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTypeAndTimestamp(offset, limit, iWid,iType,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTypeAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndMessageAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndMessageAndVariables(offset, limit, iWid,iMessage,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndMessageAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndMessageAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndMessageAndSeverity(offset, limit, iWid,iMessage,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndMessageAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndMessageAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iMessage := self.Args("message").String()
	iLink := self.Args("link").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndMessageAndLink(offset, limit, iWid,iMessage,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndMessageAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndMessageAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iMessage := self.Args("message").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndMessageAndLocation(offset, limit, iWid,iMessage,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndMessageAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndMessageAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iMessage := self.Args("message").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndMessageAndReferer(offset, limit, iWid,iMessage,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndMessageAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndMessageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iMessage := self.Args("message").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndMessageAndHostname(offset, limit, iWid,iMessage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndMessageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndMessageAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iMessage := self.Args("message").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndMessageAndTimestamp(offset, limit, iWid,iMessage,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndMessageAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndVariablesAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndVariablesAndSeverity(offset, limit, iWid,iVariables,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndVariablesAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndVariablesAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndVariablesAndLink(offset, limit, iWid,iVariables,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndVariablesAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndVariablesAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iVariables := self.Args("variables").Bytes()
	iLocation := self.Args("location").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndVariablesAndLocation(offset, limit, iWid,iVariables,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndVariablesAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndVariablesAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iVariables := self.Args("variables").Bytes()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndVariablesAndReferer(offset, limit, iWid,iVariables,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndVariablesAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndVariablesAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iVariables := self.Args("variables").Bytes()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndVariablesAndHostname(offset, limit, iWid,iVariables,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndVariablesAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndVariablesAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iVariables := self.Args("variables").Bytes()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndVariablesAndTimestamp(offset, limit, iWid,iVariables,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndVariablesAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndSeverityAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndSeverityAndLink(offset, limit, iWid,iSeverity,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndSeverityAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndSeverityAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndSeverityAndLocation(offset, limit, iWid,iSeverity,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndSeverityAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndSeverityAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iSeverity := self.Args("severity").MustInt()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndSeverityAndReferer(offset, limit, iWid,iSeverity,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndSeverityAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndSeverityAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iSeverity := self.Args("severity").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndSeverityAndHostname(offset, limit, iWid,iSeverity,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndSeverityAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndSeverityAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iSeverity := self.Args("severity").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndSeverityAndTimestamp(offset, limit, iWid,iSeverity,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndSeverityAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLinkAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLinkAndLocation(offset, limit, iWid,iLink,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLinkAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLinkAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLinkAndReferer(offset, limit, iWid,iLink,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLinkAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLinkAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLink := self.Args("link").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLinkAndHostname(offset, limit, iWid,iLink,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLinkAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLinkAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLink := self.Args("link").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLinkAndTimestamp(offset, limit, iWid,iLink,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLinkAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLocationAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLocationAndReferer(offset, limit, iWid,iLocation,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLocationAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLocationAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLocationAndHostname(offset, limit, iWid,iLocation,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLocationAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLocationAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLocation := self.Args("location").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLocationAndTimestamp(offset, limit, iWid,iLocation,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLocationAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndRefererAndHostname(offset, limit, iWid,iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndRefererAndTimestamp(offset, limit, iWid,iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndHostnameAndTimestamp(offset, limit, iWid,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeAndMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTypeAndMessage(offset, limit, iUid,iType,iMessage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTypeAndMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTypeAndVariables(offset, limit, iUid,iType,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTypeAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTypeAndSeverity(offset, limit, iUid,iType,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTypeAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()
	iLink := self.Args("link").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTypeAndLink(offset, limit, iUid,iType,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTypeAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTypeAndLocation(offset, limit, iUid,iType,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTypeAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTypeAndReferer(offset, limit, iUid,iType,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTypeAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTypeAndHostname(offset, limit, iUid,iType,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTypeAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTypeAndTimestamp(offset, limit, iUid,iType,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTypeAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndMessageAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndMessageAndVariables(offset, limit, iUid,iMessage,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndMessageAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndMessageAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndMessageAndSeverity(offset, limit, iUid,iMessage,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndMessageAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndMessageAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()
	iLink := self.Args("link").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndMessageAndLink(offset, limit, iUid,iMessage,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndMessageAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndMessageAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndMessageAndLocation(offset, limit, iUid,iMessage,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndMessageAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndMessageAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndMessageAndReferer(offset, limit, iUid,iMessage,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndMessageAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndMessageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndMessageAndHostname(offset, limit, iUid,iMessage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndMessageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndMessageAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndMessageAndTimestamp(offset, limit, iUid,iMessage,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndMessageAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndVariablesAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndVariablesAndSeverity(offset, limit, iUid,iVariables,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndVariablesAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndVariablesAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndVariablesAndLink(offset, limit, iUid,iVariables,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndVariablesAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndVariablesAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLocation := self.Args("location").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndVariablesAndLocation(offset, limit, iUid,iVariables,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndVariablesAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndVariablesAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iVariables := self.Args("variables").Bytes()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndVariablesAndReferer(offset, limit, iUid,iVariables,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndVariablesAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndVariablesAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iVariables := self.Args("variables").Bytes()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndVariablesAndHostname(offset, limit, iUid,iVariables,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndVariablesAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndVariablesAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iVariables := self.Args("variables").Bytes()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndVariablesAndTimestamp(offset, limit, iUid,iVariables,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndVariablesAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndSeverityAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndSeverityAndLink(offset, limit, iUid,iSeverity,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndSeverityAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndSeverityAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndSeverityAndLocation(offset, limit, iUid,iSeverity,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndSeverityAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndSeverityAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndSeverityAndReferer(offset, limit, iUid,iSeverity,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndSeverityAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndSeverityAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndSeverityAndHostname(offset, limit, iUid,iSeverity,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndSeverityAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndSeverityAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndSeverityAndTimestamp(offset, limit, iUid,iSeverity,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndSeverityAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLinkAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLinkAndLocation(offset, limit, iUid,iLink,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLinkAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLinkAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLinkAndReferer(offset, limit, iUid,iLink,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLinkAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLinkAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLink := self.Args("link").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLinkAndHostname(offset, limit, iUid,iLink,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLinkAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLinkAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLink := self.Args("link").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLinkAndTimestamp(offset, limit, iUid,iLink,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLinkAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLocationAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLocationAndReferer(offset, limit, iUid,iLocation,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLocationAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLocationAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLocationAndHostname(offset, limit, iUid,iLocation,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLocationAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLocationAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLocation := self.Args("location").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLocationAndTimestamp(offset, limit, iUid,iLocation,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLocationAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndRefererAndHostname(offset, limit, iUid,iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndRefererAndTimestamp(offset, limit, iUid,iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndHostnameAndTimestamp(offset, limit, iUid,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndMessageAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndMessageAndVariables(offset, limit, iType,iMessage,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndMessageAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndMessageAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndMessageAndSeverity(offset, limit, iType,iMessage,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndMessageAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndMessageAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()
	iLink := self.Args("link").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndMessageAndLink(offset, limit, iType,iMessage,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndMessageAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndMessageAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndMessageAndLocation(offset, limit, iType,iMessage,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndMessageAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndMessageAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndMessageAndReferer(offset, limit, iType,iMessage,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndMessageAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndMessageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndMessageAndHostname(offset, limit, iType,iMessage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndMessageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndMessageAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndMessageAndTimestamp(offset, limit, iType,iMessage,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndMessageAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndVariablesAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndVariablesAndSeverity(offset, limit, iType,iVariables,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndVariablesAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndVariablesAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndVariablesAndLink(offset, limit, iType,iVariables,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndVariablesAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndVariablesAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()
	iLocation := self.Args("location").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndVariablesAndLocation(offset, limit, iType,iVariables,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndVariablesAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndVariablesAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndVariablesAndReferer(offset, limit, iType,iVariables,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndVariablesAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndVariablesAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndVariablesAndHostname(offset, limit, iType,iVariables,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndVariablesAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndVariablesAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndVariablesAndTimestamp(offset, limit, iType,iVariables,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndVariablesAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndSeverityAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndSeverityAndLink(offset, limit, iType,iSeverity,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndSeverityAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndSeverityAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndSeverityAndLocation(offset, limit, iType,iSeverity,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndSeverityAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndSeverityAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSeverity := self.Args("severity").MustInt()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndSeverityAndReferer(offset, limit, iType,iSeverity,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndSeverityAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndSeverityAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSeverity := self.Args("severity").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndSeverityAndHostname(offset, limit, iType,iSeverity,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndSeverityAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndSeverityAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSeverity := self.Args("severity").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndSeverityAndTimestamp(offset, limit, iType,iSeverity,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndSeverityAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLinkAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLinkAndLocation(offset, limit, iType,iLink,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLinkAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLinkAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLinkAndReferer(offset, limit, iType,iLink,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLinkAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLinkAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLink := self.Args("link").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLinkAndHostname(offset, limit, iType,iLink,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLinkAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLinkAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLink := self.Args("link").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLinkAndTimestamp(offset, limit, iType,iLink,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLinkAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLocationAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLocationAndReferer(offset, limit, iType,iLocation,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLocationAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLocationAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLocationAndHostname(offset, limit, iType,iLocation,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLocationAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLocationAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLocation := self.Args("location").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLocationAndTimestamp(offset, limit, iType,iLocation,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLocationAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndRefererAndHostname(offset, limit, iType,iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndRefererAndTimestamp(offset, limit, iType,iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndHostnameAndTimestamp(offset, limit, iType,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndVariablesAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndVariablesAndSeverity(offset, limit, iMessage,iVariables,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndVariablesAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndVariablesAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndVariablesAndLink(offset, limit, iMessage,iVariables,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndVariablesAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndVariablesAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()
	iLocation := self.Args("location").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndVariablesAndLocation(offset, limit, iMessage,iVariables,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndVariablesAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndVariablesAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndVariablesAndReferer(offset, limit, iMessage,iVariables,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndVariablesAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndVariablesAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndVariablesAndHostname(offset, limit, iMessage,iVariables,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndVariablesAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndVariablesAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndVariablesAndTimestamp(offset, limit, iMessage,iVariables,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndVariablesAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndSeverityAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndSeverityAndLink(offset, limit, iMessage,iSeverity,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndSeverityAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndSeverityAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndSeverityAndLocation(offset, limit, iMessage,iSeverity,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndSeverityAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndSeverityAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndSeverityAndReferer(offset, limit, iMessage,iSeverity,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndSeverityAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndSeverityAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndSeverityAndHostname(offset, limit, iMessage,iSeverity,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndSeverityAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndSeverityAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndSeverityAndTimestamp(offset, limit, iMessage,iSeverity,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndSeverityAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLinkAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLinkAndLocation(offset, limit, iMessage,iLink,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLinkAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLinkAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLinkAndReferer(offset, limit, iMessage,iLink,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLinkAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLinkAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLink := self.Args("link").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLinkAndHostname(offset, limit, iMessage,iLink,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLinkAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLinkAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLink := self.Args("link").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLinkAndTimestamp(offset, limit, iMessage,iLink,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLinkAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLocationAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLocationAndReferer(offset, limit, iMessage,iLocation,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLocationAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLocationAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLocationAndHostname(offset, limit, iMessage,iLocation,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLocationAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLocationAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLocation := self.Args("location").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLocationAndTimestamp(offset, limit, iMessage,iLocation,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLocationAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndRefererAndHostname(offset, limit, iMessage,iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndRefererAndTimestamp(offset, limit, iMessage,iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndHostnameAndTimestamp(offset, limit, iMessage,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndSeverityAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndSeverityAndLink(offset, limit, iVariables,iSeverity,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndSeverityAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndSeverityAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndSeverityAndLocation(offset, limit, iVariables,iSeverity,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndSeverityAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndSeverityAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndSeverityAndReferer(offset, limit, iVariables,iSeverity,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndSeverityAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndSeverityAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndSeverityAndHostname(offset, limit, iVariables,iSeverity,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndSeverityAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndSeverityAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndSeverityAndTimestamp(offset, limit, iVariables,iSeverity,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndSeverityAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLinkAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLinkAndLocation(offset, limit, iVariables,iLink,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLinkAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLinkAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLinkAndReferer(offset, limit, iVariables,iLink,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLinkAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLinkAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLinkAndHostname(offset, limit, iVariables,iLink,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLinkAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLinkAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLinkAndTimestamp(offset, limit, iVariables,iLink,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLinkAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLocationAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLocationAndReferer(offset, limit, iVariables,iLocation,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLocationAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLocationAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLocationAndHostname(offset, limit, iVariables,iLocation,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLocationAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLocationAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLocation := self.Args("location").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLocationAndTimestamp(offset, limit, iVariables,iLocation,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLocationAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndRefererAndHostname(offset, limit, iVariables,iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndRefererAndTimestamp(offset, limit, iVariables,iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndHostnameAndTimestamp(offset, limit, iVariables,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLinkAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLinkAndLocation(offset, limit, iSeverity,iLink,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLinkAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLinkAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLinkAndReferer(offset, limit, iSeverity,iLink,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLinkAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLinkAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLinkAndHostname(offset, limit, iSeverity,iLink,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLinkAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLinkAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLinkAndTimestamp(offset, limit, iSeverity,iLink,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLinkAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLocationAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLocationAndReferer(offset, limit, iSeverity,iLocation,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLocationAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLocationAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLocationAndHostname(offset, limit, iSeverity,iLocation,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLocationAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLocationAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLocationAndTimestamp(offset, limit, iSeverity,iLocation,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLocationAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndRefererAndHostname(offset, limit, iSeverity,iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndRefererAndTimestamp(offset, limit, iSeverity,iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndHostnameAndTimestamp(offset, limit, iSeverity,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndLocationAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndLocationAndReferer(offset, limit, iLink,iLocation,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndLocationAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndLocationAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndLocationAndHostname(offset, limit, iLink,iLocation,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndLocationAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndLocationAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndLocationAndTimestamp(offset, limit, iLink,iLocation,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndLocationAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndRefererAndHostname(offset, limit, iLink,iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndRefererAndTimestamp(offset, limit, iLink,iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndHostnameAndTimestamp(offset, limit, iLink,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLocationAndRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLocation) {
		_Watchdog, _error := model.GetWatchdogsByLocationAndRefererAndHostname(offset, limit, iLocation,iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLocationAndRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLocationAndRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iLocation) {
		_Watchdog, _error := model.GetWatchdogsByLocationAndRefererAndTimestamp(offset, limit, iLocation,iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLocationAndRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLocationAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iLocation) {
		_Watchdog, _error := model.GetWatchdogsByLocationAndHostnameAndTimestamp(offset, limit, iLocation,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLocationAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByRefererAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iReferer) {
		_Watchdog, _error := model.GetWatchdogsByRefererAndHostnameAndTimestamp(offset, limit, iReferer,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByRefererAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndUid(offset, limit, iWid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iType := self.Args("type").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndType(offset, limit, iWid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iMessage := self.Args("message").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndMessage(offset, limit, iWid,iMessage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndVariables(offset, limit, iWid,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndSeverity(offset, limit, iWid,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLink := self.Args("link").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLink(offset, limit, iWid,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iLocation := self.Args("location").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndLocation(offset, limit, iWid,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndReferer(offset, limit, iWid,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndHostname(offset, limit, iWid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByWidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iWid := self.Args("wid").MustInt64()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogsByWidAndTimestamp(offset, limit, iWid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByWidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iType := self.Args("type").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndType(offset, limit, iUid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMessage := self.Args("message").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndMessage(offset, limit, iUid,iMessage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndVariables(offset, limit, iUid,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndSeverity(offset, limit, iUid,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLink := self.Args("link").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLink(offset, limit, iUid,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iLocation := self.Args("location").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndLocation(offset, limit, iUid,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndReferer(offset, limit, iUid,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndHostname(offset, limit, iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByUidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogsByUidAndTimestamp(offset, limit, iUid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByUidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iMessage := self.Args("message").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndMessage(offset, limit, iType,iMessage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndVariables(offset, limit, iType,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndSeverity(offset, limit, iType,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLink := self.Args("link").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLink(offset, limit, iType,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndLocation(offset, limit, iType,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndReferer(offset, limit, iType,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndHostname(offset, limit, iType,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByTypeAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogsByTypeAndTimestamp(offset, limit, iType,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByTypeAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iVariables := self.Args("variables").Bytes()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndVariables(offset, limit, iMessage,iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndSeverity(offset, limit, iMessage,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLink := self.Args("link").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLink(offset, limit, iMessage,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndLocation(offset, limit, iMessage,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndReferer(offset, limit, iMessage,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndHostname(offset, limit, iMessage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByMessageAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMessage := self.Args("message").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogsByMessageAndTimestamp(offset, limit, iMessage,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByMessageAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iSeverity := self.Args("severity").MustInt()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndSeverity(offset, limit, iVariables,iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLink := self.Args("link").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLink(offset, limit, iVariables,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iLocation := self.Args("location").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndLocation(offset, limit, iVariables,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndReferer(offset, limit, iVariables,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndHostname(offset, limit, iVariables,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByVariablesAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iVariables := self.Args("variables").Bytes()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogsByVariablesAndTimestamp(offset, limit, iVariables,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByVariablesAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLink := self.Args("link").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLink(offset, limit, iSeverity,iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iLocation := self.Args("location").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndLocation(offset, limit, iSeverity,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndReferer(offset, limit, iSeverity,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndHostname(offset, limit, iSeverity,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsBySeverityAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSeverity := self.Args("severity").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogsBySeverityAndTimestamp(offset, limit, iSeverity,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsBySeverityAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iLocation := self.Args("location").String()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndLocation(offset, limit, iLink,iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndReferer(offset, limit, iLink,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndHostname(offset, limit, iLink,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLinkAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLink := self.Args("link").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogsByLinkAndTimestamp(offset, limit, iLink,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLinkAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLocationAndRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLocation := self.Args("location").String()
	iReferer := self.Args("referer").String()

	if helper.IsHas(iLocation) {
		_Watchdog, _error := model.GetWatchdogsByLocationAndReferer(offset, limit, iLocation,iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLocationAndReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLocationAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLocation := self.Args("location").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLocation) {
		_Watchdog, _error := model.GetWatchdogsByLocationAndHostname(offset, limit, iLocation,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLocationAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByLocationAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLocation := self.Args("location").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iLocation) {
		_Watchdog, _error := model.GetWatchdogsByLocationAndTimestamp(offset, limit, iLocation,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByLocationAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByRefererAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iReferer := self.Args("referer").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iReferer) {
		_Watchdog, _error := model.GetWatchdogsByRefererAndHostname(offset, limit, iReferer,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByRefererAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByRefererAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iReferer := self.Args("referer").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iReferer) {
		_Watchdog, _error := model.GetWatchdogsByRefererAndTimestamp(offset, limit, iReferer,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByRefererAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsByHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iHostname) {
		_Watchdog, _error := model.GetWatchdogsByHostnameAndTimestamp(offset, limit, iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogsByHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Watchdog, _error := model.GetWatchdogs(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogs' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaWidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWid := self.Args("wid").MustInt64()
	if helper.IsHas(iWid) {
		_Watchdog := model.HasWatchdogViaWid(iWid)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaWid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_Watchdog := model.HasWatchdogViaUid(iUid)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_Watchdog := model.HasWatchdogViaType(iType)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMessage := self.Args("message").String()
	if helper.IsHas(iMessage) {
		_Watchdog := model.HasWatchdogViaMessage(iMessage)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVariables := self.Args("variables").Bytes()
	if helper.IsHas(iVariables) {
		_Watchdog := model.HasWatchdogViaVariables(iVariables)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSeverity := self.Args("severity").MustInt()
	if helper.IsHas(iSeverity) {
		_Watchdog := model.HasWatchdogViaSeverity(iSeverity)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink := self.Args("link").String()
	if helper.IsHas(iLink) {
		_Watchdog := model.HasWatchdogViaLink(iLink)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLocation := self.Args("location").String()
	if helper.IsHas(iLocation) {
		_Watchdog := model.HasWatchdogViaLocation(iLocation)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iReferer := self.Args("referer").String()
	if helper.IsHas(iReferer) {
		_Watchdog := model.HasWatchdogViaReferer(iReferer)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iHostname := self.Args("hostname").String()
	if helper.IsHas(iHostname) {
		_Watchdog := model.HasWatchdogViaHostname(iHostname)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasWatchdogViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_Watchdog := model.HasWatchdogViaTimestamp(iTimestamp)
		var m = map[string]interface{}{}
		m["watchdog"] = _Watchdog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasWatchdogViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaWidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iWid := self.Args("wid").MustInt64()
	if helper.IsHas(iWid) {
		_Watchdog, _error := model.GetWatchdogViaWid(iWid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaWid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_Watchdog, _error := model.GetWatchdogViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_Watchdog, _error := model.GetWatchdogViaType(iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMessage := self.Args("message").String()
	if helper.IsHas(iMessage) {
		_Watchdog, _error := model.GetWatchdogViaMessage(iMessage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iVariables := self.Args("variables").Bytes()
	if helper.IsHas(iVariables) {
		_Watchdog, _error := model.GetWatchdogViaVariables(iVariables)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSeverity := self.Args("severity").MustInt()
	if helper.IsHas(iSeverity) {
		_Watchdog, _error := model.GetWatchdogViaSeverity(iSeverity)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLink := self.Args("link").String()
	if helper.IsHas(iLink) {
		_Watchdog, _error := model.GetWatchdogViaLink(iLink)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLocation := self.Args("location").String()
	if helper.IsHas(iLocation) {
		_Watchdog, _error := model.GetWatchdogViaLocation(iLocation)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iReferer := self.Args("referer").String()
	if helper.IsHas(iReferer) {
		_Watchdog, _error := model.GetWatchdogViaReferer(iReferer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iHostname := self.Args("hostname").String()
	if helper.IsHas(iHostname) {
		_Watchdog, _error := model.GetWatchdogViaHostname(iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetWatchdogViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_Watchdog, _error := model.GetWatchdogViaTimestamp(iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the GetWatchdogViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaWidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Wid_ := self.Args("wid").MustInt64()
	if helper.IsHas(Wid_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaWid(Wid_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaWid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	if helper.IsHas(Uid_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaUid(Uid_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	if helper.IsHas(Type_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaType(Type_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Message_ := self.Args("message").String()
	if helper.IsHas(Message_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaMessage(Message_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaMessage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Variables_ := self.Args("variables").Bytes()
	if helper.IsHas(Variables_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaVariables(Variables_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaVariables's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Severity_ := self.Args("severity").MustInt()
	if helper.IsHas(Severity_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaSeverity(Severity_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaSeverity's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_ := self.Args("link").String()
	if helper.IsHas(Link_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaLink(Link_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaLink's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Location_ := self.Args("location").String()
	if helper.IsHas(Location_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaLocation(Location_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaLocation's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Referer_ := self.Args("referer").String()
	if helper.IsHas(Referer_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaReferer(Referer_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaReferer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	if helper.IsHas(Hostname_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaHostname(Hostname_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetWatchdogViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	if helper.IsHas(Timestamp_) {
		var iWatchdog model.Watchdog
		self.Bind(&iWatchdog)
		_Watchdog, _error := model.SetWatchdogViaTimestamp(Timestamp_, &iWatchdog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Watchdog)
	}
	herr.Message = "Can't get to the SetWatchdogViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddWatchdogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Wid_ := self.Args("wid").MustInt64()
	Uid_ := self.Args("uid").MustInt()
	Type_ := self.Args("type").String()
	Message_ := self.Args("message").String()
	Variables_ := self.Args("variables").Bytes()
	Severity_ := self.Args("severity").MustInt()
	Link_ := self.Args("link").String()
	Location_ := self.Args("location").String()
	Referer_ := self.Args("referer").String()
	Hostname_ := self.Args("hostname").String()
	Timestamp_ := self.Args("timestamp").MustInt()

	if helper.IsHas(Wid_) {
		_error := model.AddWatchdog(Wid_,Uid_,Type_,Message_,Variables_,Severity_,Link_,Location_,Referer_,Hostname_,Timestamp_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddWatchdog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostWatchdogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PostWatchdog(&iWatchdog)
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

func PutWatchdogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdog(&iWatchdog)
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

func PutWatchdogViaWidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Wid_ := self.Args("wid").MustInt64()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaWid(Wid_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaUid(Uid_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaType(Type_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Message_ := self.Args("message").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaMessage(Message_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Variables_ := self.Args("variables").Bytes()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaVariables(Variables_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Severity_ := self.Args("severity").MustInt()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaSeverity(Severity_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_ := self.Args("link").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaLink(Link_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Location_ := self.Args("location").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaLocation(Location_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Referer_ := self.Args("referer").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaReferer(Referer_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaHostname(Hostname_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutWatchdogViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	_int64, _error := model.PutWatchdogViaTimestamp(Timestamp_, &iWatchdog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaWidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Wid_ := self.Args("wid").MustInt64()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaWid(Wid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaType(Type_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Message_ := self.Args("message").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaMessage(Message_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Variables_ := self.Args("variables").Bytes()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaVariables(Variables_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Severity_ := self.Args("severity").MustInt()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaSeverity(Severity_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_ := self.Args("link").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaLink(Link_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Location_ := self.Args("location").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaLocation(Location_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Referer_ := self.Args("referer").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaReferer(Referer_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaHostname(Hostname_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateWatchdogViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iWatchdog model.Watchdog
	self.Bind(&iWatchdog)
	var iMap = helper.StructToMap(iWatchdog)
	_error := model.UpdateWatchdogViaTimestamp(Timestamp_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Wid_ := self.Args("wid").MustInt64()
	_int64, _error := model.DeleteWatchdog(Wid_)
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

func DeleteWatchdogViaWidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Wid_ := self.Args("wid").MustInt64()
	_error := model.DeleteWatchdogViaWid(Wid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	_error := model.DeleteWatchdogViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	_error := model.DeleteWatchdogViaType(Type_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaMessageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Message_ := self.Args("message").String()
	_error := model.DeleteWatchdogViaMessage(Message_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaVariablesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Variables_ := self.Args("variables").Bytes()
	_error := model.DeleteWatchdogViaVariables(Variables_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaSeverityHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Severity_ := self.Args("severity").MustInt()
	_error := model.DeleteWatchdogViaSeverity(Severity_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaLinkHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Link_ := self.Args("link").String()
	_error := model.DeleteWatchdogViaLink(Link_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaLocationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Location_ := self.Args("location").String()
	_error := model.DeleteWatchdogViaLocation(Location_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaRefererHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Referer_ := self.Args("referer").String()
	_error := model.DeleteWatchdogViaReferer(Referer_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	_error := model.DeleteWatchdogViaHostname(Hostname_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteWatchdogViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	_error := model.DeleteWatchdogViaTimestamp(Timestamp_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
