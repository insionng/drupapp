package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetSessionsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSessionsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["sessionssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSessionsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt64()
	_int64 := model.GetSessionsCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sessionsCount"] = 0
	}
	m["sessionsCount"] = _int64
	return self.JSON(m)
}

func GetSessionsCountViaSidHandler(self *macross.Context) error {
	Sid_ := self.Args("sid").String()
	_int64 := model.GetSessionsCountViaSid(Sid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sessionsCount"] = 0
	}
	m["sessionsCount"] = _int64
	return self.JSON(m)
}

func GetSessionsCountViaHostnameHandler(self *macross.Context) error {
	Hostname_ := self.Args("hostname").String()
	_int64 := model.GetSessionsCountViaHostname(Hostname_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sessionsCount"] = 0
	}
	m["sessionsCount"] = _int64
	return self.JSON(m)
}

func GetSessionsCountViaTimestampHandler(self *macross.Context) error {
	Timestamp_ := self.Args("timestamp").MustInt()
	_int64 := model.GetSessionsCountViaTimestamp(Timestamp_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sessionsCount"] = 0
	}
	m["sessionsCount"] = _int64
	return self.JSON(m)
}

func GetSessionsCountViaSessionHandler(self *macross.Context) error {
	Session_ := self.Args("session").Bytes()
	_int64 := model.GetSessionsCountViaSession(Session_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sessionsCount"] = 0
	}
	m["sessionsCount"] = _int64
	return self.JSON(m)
}

func GetSessionsesViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt64()
	if (offset > 0) && helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSid := self.Args("sid").String()
	if (offset > 0) && helper.IsHas(iSid) {
		_Sessions, _error := model.GetSessionsesViaSid(offset, limit, iSid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iHostname := self.Args("hostname").String()
	if (offset > 0) && helper.IsHas(iHostname) {
		_Sessions, _error := model.GetSessionsesViaHostname(offset, limit, iHostname, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTimestamp := self.Args("timestamp").MustInt()
	if (offset > 0) && helper.IsHas(iTimestamp) {
		_Sessions, _error := model.GetSessionsesViaTimestamp(offset, limit, iTimestamp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesViaSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSession := self.Args("session").Bytes()
	if (offset > 0) && helper.IsHas(iSession) {
		_Sessions, _error := model.GetSessionsesViaSession(offset, limit, iSession, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesViaSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndSidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iSid := self.Args("sid").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndSidAndHostname(offset, limit, iUid,iSid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndSidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndSidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iSid := self.Args("sid").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndSidAndTimestamp(offset, limit, iUid,iSid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndSidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndSidAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iSid := self.Args("sid").String()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndSidAndSession(offset, limit, iUid,iSid,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndSidAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndHostnameAndTimestamp(offset, limit, iUid,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndHostnameAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iHostname := self.Args("hostname").String()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndHostnameAndSession(offset, limit, iUid,iHostname,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndHostnameAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndTimestampAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimestamp := self.Args("timestamp").MustInt()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndTimestampAndSession(offset, limit, iUid,iTimestamp,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndTimestampAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesBySidAndHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").String()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iSid) {
		_Sessions, _error := model.GetSessionsesBySidAndHostnameAndTimestamp(offset, limit, iSid,iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesBySidAndHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesBySidAndHostnameAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").String()
	iHostname := self.Args("hostname").String()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iSid) {
		_Sessions, _error := model.GetSessionsesBySidAndHostnameAndSession(offset, limit, iSid,iHostname,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesBySidAndHostnameAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesBySidAndTimestampAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").String()
	iTimestamp := self.Args("timestamp").MustInt()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iSid) {
		_Sessions, _error := model.GetSessionsesBySidAndTimestampAndSession(offset, limit, iSid,iTimestamp,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesBySidAndTimestampAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByHostnameAndTimestampAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iHostname) {
		_Sessions, _error := model.GetSessionsesByHostnameAndTimestampAndSession(offset, limit, iHostname,iTimestamp,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByHostnameAndTimestampAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iSid := self.Args("sid").String()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndSid(offset, limit, iUid,iSid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndHostname(offset, limit, iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndTimestamp(offset, limit, iUid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByUidAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsesByUidAndSession(offset, limit, iUid,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByUidAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesBySidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSid) {
		_Sessions, _error := model.GetSessionsesBySidAndHostname(offset, limit, iSid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesBySidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesBySidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iSid) {
		_Sessions, _error := model.GetSessionsesBySidAndTimestamp(offset, limit, iSid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesBySidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesBySidAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").String()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iSid) {
		_Sessions, _error := model.GetSessionsesBySidAndSession(offset, limit, iSid,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesBySidAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByHostnameAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iHostname) {
		_Sessions, _error := model.GetSessionsesByHostnameAndTimestamp(offset, limit, iHostname,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByHostnameAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByHostnameAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iHostname) {
		_Sessions, _error := model.GetSessionsesByHostnameAndSession(offset, limit, iHostname,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByHostnameAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesByTimestampAndSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()
	iSession := self.Args("session").Bytes()

	if helper.IsHas(iTimestamp) {
		_Sessions, _error := model.GetSessionsesByTimestampAndSession(offset, limit, iTimestamp,iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsesByTimestampAndSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Sessions, _error := model.GetSessionses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSessionsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_Sessions := model.HasSessionsViaUid(iUid)
		var m = map[string]interface{}{}
		m["sessions"] = _Sessions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSessionsViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSessionsViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSid := self.Args("sid").String()
	if helper.IsHas(iSid) {
		_Sessions := model.HasSessionsViaSid(iSid)
		var m = map[string]interface{}{}
		m["sessions"] = _Sessions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSessionsViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSessionsViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iHostname := self.Args("hostname").String()
	if helper.IsHas(iHostname) {
		_Sessions := model.HasSessionsViaHostname(iHostname)
		var m = map[string]interface{}{}
		m["sessions"] = _Sessions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSessionsViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSessionsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_Sessions := model.HasSessionsViaTimestamp(iTimestamp)
		var m = map[string]interface{}{}
		m["sessions"] = _Sessions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSessionsViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSessionsViaSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSession := self.Args("session").Bytes()
	if helper.IsHas(iSession) {
		_Sessions := model.HasSessionsViaSession(iSession)
		var m = map[string]interface{}{}
		m["sessions"] = _Sessions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSessionsViaSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_Sessions, _error := model.GetSessionsViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSid := self.Args("sid").String()
	if helper.IsHas(iSid) {
		_Sessions, _error := model.GetSessionsViaSid(iSid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iHostname := self.Args("hostname").String()
	if helper.IsHas(iHostname) {
		_Sessions, _error := model.GetSessionsViaHostname(iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_Sessions, _error := model.GetSessionsViaTimestamp(iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSessionsViaSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSession := self.Args("session").Bytes()
	if helper.IsHas(iSession) {
		_Sessions, _error := model.GetSessionsViaSession(iSession)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the GetSessionsViaSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSessionsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	if helper.IsHas(Uid_) {
		var iSessions model.Sessions
		self.Bind(&iSessions)
		_Sessions, _error := model.SetSessionsViaUid(Uid_, &iSessions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the SetSessionsViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSessionsViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").String()
	if helper.IsHas(Sid_) {
		var iSessions model.Sessions
		self.Bind(&iSessions)
		_Sessions, _error := model.SetSessionsViaSid(Sid_, &iSessions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the SetSessionsViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSessionsViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	if helper.IsHas(Hostname_) {
		var iSessions model.Sessions
		self.Bind(&iSessions)
		_Sessions, _error := model.SetSessionsViaHostname(Hostname_, &iSessions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the SetSessionsViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSessionsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	if helper.IsHas(Timestamp_) {
		var iSessions model.Sessions
		self.Bind(&iSessions)
		_Sessions, _error := model.SetSessionsViaTimestamp(Timestamp_, &iSessions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the SetSessionsViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSessionsViaSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Session_ := self.Args("session").Bytes()
	if helper.IsHas(Session_) {
		var iSessions model.Sessions
		self.Bind(&iSessions)
		_Sessions, _error := model.SetSessionsViaSession(Session_, &iSessions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sessions)
	}
	herr.Message = "Can't get to the SetSessionsViaSession's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSessionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	Sid_ := self.Args("sid").String()
	Hostname_ := self.Args("hostname").String()
	Timestamp_ := self.Args("timestamp").MustInt()
	Session_ := self.Args("session").Bytes()

	if helper.IsHas(Uid_) {
		_error := model.AddSessions(Uid_,Sid_,Hostname_,Timestamp_,Session_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSessions's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSessionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSessions model.Sessions
	self.Bind(&iSessions)
	_int64, _error := model.PostSessions(&iSessions)
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

func PutSessionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSessions model.Sessions
	self.Bind(&iSessions)
	_int64, _error := model.PutSessions(&iSessions)
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

func PutSessionsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	_int64, _error := model.PutSessionsViaUid(Uid_, &iSessions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSessionsViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").String()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	_int64, _error := model.PutSessionsViaSid(Sid_, &iSessions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSessionsViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	_int64, _error := model.PutSessionsViaHostname(Hostname_, &iSessions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSessionsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	_int64, _error := model.PutSessionsViaTimestamp(Timestamp_, &iSessions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSessionsViaSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Session_ := self.Args("session").Bytes()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	_int64, _error := model.PutSessionsViaSession(Session_, &iSessions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSessionsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	var iMap = helper.StructToMap(iSessions)
	_error := model.UpdateSessionsViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSessionsViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").String()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	var iMap = helper.StructToMap(iSessions)
	_error := model.UpdateSessionsViaSid(Sid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSessionsViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	var iMap = helper.StructToMap(iSessions)
	_error := model.UpdateSessionsViaHostname(Hostname_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSessionsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	var iMap = helper.StructToMap(iSessions)
	_error := model.UpdateSessionsViaTimestamp(Timestamp_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSessionsViaSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Session_ := self.Args("session").Bytes()
	var iSessions model.Sessions
	self.Bind(&iSessions)
	var iMap = helper.StructToMap(iSessions)
	_error := model.UpdateSessionsViaSession(Session_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSessionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_int64, _error := model.DeleteSessions(Uid_)
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

func DeleteSessionsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_error := model.DeleteSessionsViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSessionsViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").String()
	_error := model.DeleteSessionsViaSid(Sid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSessionsViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	_error := model.DeleteSessionsViaHostname(Hostname_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSessionsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	_error := model.DeleteSessionsViaTimestamp(Timestamp_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSessionsViaSessionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Session_ := self.Args("session").Bytes()
	_error := model.DeleteSessionsViaSession(Session_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
