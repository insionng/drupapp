package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetFloodsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetFloodsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["floodsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetFloodsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodCountViaFidHandler(self *macross.Context) error {
	Fid_ := self.Args("fid").MustInt64()
	_int64 := model.GetFloodCountViaFid(Fid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["floodCount"] = 0
	}
	m["floodCount"] = _int64
	return self.JSON(m)
}

func GetFloodCountViaEventHandler(self *macross.Context) error {
	Event_ := self.Args("event").String()
	_int64 := model.GetFloodCountViaEvent(Event_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["floodCount"] = 0
	}
	m["floodCount"] = _int64
	return self.JSON(m)
}

func GetFloodCountViaIdentifierHandler(self *macross.Context) error {
	Identifier_ := self.Args("identifier").String()
	_int64 := model.GetFloodCountViaIdentifier(Identifier_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["floodCount"] = 0
	}
	m["floodCount"] = _int64
	return self.JSON(m)
}

func GetFloodCountViaTimestampHandler(self *macross.Context) error {
	Timestamp_ := self.Args("timestamp").MustInt()
	_int64 := model.GetFloodCountViaTimestamp(Timestamp_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["floodCount"] = 0
	}
	m["floodCount"] = _int64
	return self.JSON(m)
}

func GetFloodCountViaExpirationHandler(self *macross.Context) error {
	Expiration_ := self.Args("expiration").MustInt()
	_int64 := model.GetFloodCountViaExpiration(Expiration_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["floodCount"] = 0
	}
	m["floodCount"] = _int64
	return self.JSON(m)
}

func GetFloodsViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFid := self.Args("fid").MustInt64()
	if (offset > 0) && helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsViaFid(offset, limit, iFid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsViaEventHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEvent := self.Args("event").String()
	if (offset > 0) && helper.IsHas(iEvent) {
		_Flood, _error := model.GetFloodsViaEvent(offset, limit, iEvent, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsViaEvent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsViaIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iIdentifier := self.Args("identifier").String()
	if (offset > 0) && helper.IsHas(iIdentifier) {
		_Flood, _error := model.GetFloodsViaIdentifier(offset, limit, iIdentifier, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsViaIdentifier's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTimestamp := self.Args("timestamp").MustInt()
	if (offset > 0) && helper.IsHas(iTimestamp) {
		_Flood, _error := model.GetFloodsViaTimestamp(offset, limit, iTimestamp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsViaExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpiration := self.Args("expiration").MustInt()
	if (offset > 0) && helper.IsHas(iExpiration) {
		_Flood, _error := model.GetFloodsViaExpiration(offset, limit, iExpiration, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsViaExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndEventAndIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iEvent := self.Args("event").String()
	iIdentifier := self.Args("identifier").String()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndEventAndIdentifier(offset, limit, iFid,iEvent,iIdentifier)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndEventAndIdentifier's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndEventAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iEvent := self.Args("event").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndEventAndTimestamp(offset, limit, iFid,iEvent,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndEventAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndEventAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iEvent := self.Args("event").String()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndEventAndExpiration(offset, limit, iFid,iEvent,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndEventAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndIdentifierAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iIdentifier := self.Args("identifier").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndIdentifierAndTimestamp(offset, limit, iFid,iIdentifier,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndIdentifierAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndIdentifierAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iIdentifier := self.Args("identifier").String()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndIdentifierAndExpiration(offset, limit, iFid,iIdentifier,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndIdentifierAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndTimestampAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iTimestamp := self.Args("timestamp").MustInt()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndTimestampAndExpiration(offset, limit, iFid,iTimestamp,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndTimestampAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByEventAndIdentifierAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEvent := self.Args("event").String()
	iIdentifier := self.Args("identifier").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iEvent) {
		_Flood, _error := model.GetFloodsByEventAndIdentifierAndTimestamp(offset, limit, iEvent,iIdentifier,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByEventAndIdentifierAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByEventAndIdentifierAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEvent := self.Args("event").String()
	iIdentifier := self.Args("identifier").String()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iEvent) {
		_Flood, _error := model.GetFloodsByEventAndIdentifierAndExpiration(offset, limit, iEvent,iIdentifier,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByEventAndIdentifierAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByEventAndTimestampAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEvent := self.Args("event").String()
	iTimestamp := self.Args("timestamp").MustInt()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iEvent) {
		_Flood, _error := model.GetFloodsByEventAndTimestampAndExpiration(offset, limit, iEvent,iTimestamp,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByEventAndTimestampAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByIdentifierAndTimestampAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iIdentifier := self.Args("identifier").String()
	iTimestamp := self.Args("timestamp").MustInt()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iIdentifier) {
		_Flood, _error := model.GetFloodsByIdentifierAndTimestampAndExpiration(offset, limit, iIdentifier,iTimestamp,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByIdentifierAndTimestampAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndEventHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iEvent := self.Args("event").String()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndEvent(offset, limit, iFid,iEvent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndEvent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iIdentifier := self.Args("identifier").String()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndIdentifier(offset, limit, iFid,iIdentifier)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndIdentifier's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndTimestamp(offset, limit, iFid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByFidAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodsByFidAndExpiration(offset, limit, iFid,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByFidAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByEventAndIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEvent := self.Args("event").String()
	iIdentifier := self.Args("identifier").String()

	if helper.IsHas(iEvent) {
		_Flood, _error := model.GetFloodsByEventAndIdentifier(offset, limit, iEvent,iIdentifier)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByEventAndIdentifier's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByEventAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEvent := self.Args("event").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iEvent) {
		_Flood, _error := model.GetFloodsByEventAndTimestamp(offset, limit, iEvent,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByEventAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByEventAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEvent := self.Args("event").String()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iEvent) {
		_Flood, _error := model.GetFloodsByEventAndExpiration(offset, limit, iEvent,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByEventAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByIdentifierAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iIdentifier := self.Args("identifier").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iIdentifier) {
		_Flood, _error := model.GetFloodsByIdentifierAndTimestamp(offset, limit, iIdentifier,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByIdentifierAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByIdentifierAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iIdentifier := self.Args("identifier").String()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iIdentifier) {
		_Flood, _error := model.GetFloodsByIdentifierAndExpiration(offset, limit, iIdentifier,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByIdentifierAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsByTimestampAndExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()
	iExpiration := self.Args("expiration").MustInt()

	if helper.IsHas(iTimestamp) {
		_Flood, _error := model.GetFloodsByTimestampAndExpiration(offset, limit, iTimestamp,iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodsByTimestampAndExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Flood, _error := model.GetFloods(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloods' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFloodViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFid := self.Args("fid").MustInt64()
	if helper.IsHas(iFid) {
		_Flood := model.HasFloodViaFid(iFid)
		var m = map[string]interface{}{}
		m["flood"] = _Flood
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFloodViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFloodViaEventHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEvent := self.Args("event").String()
	if helper.IsHas(iEvent) {
		_Flood := model.HasFloodViaEvent(iEvent)
		var m = map[string]interface{}{}
		m["flood"] = _Flood
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFloodViaEvent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFloodViaIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iIdentifier := self.Args("identifier").String()
	if helper.IsHas(iIdentifier) {
		_Flood := model.HasFloodViaIdentifier(iIdentifier)
		var m = map[string]interface{}{}
		m["flood"] = _Flood
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFloodViaIdentifier's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFloodViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_Flood := model.HasFloodViaTimestamp(iTimestamp)
		var m = map[string]interface{}{}
		m["flood"] = _Flood
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFloodViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFloodViaExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpiration := self.Args("expiration").MustInt()
	if helper.IsHas(iExpiration) {
		_Flood := model.HasFloodViaExpiration(iExpiration)
		var m = map[string]interface{}{}
		m["flood"] = _Flood
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFloodViaExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFid := self.Args("fid").MustInt64()
	if helper.IsHas(iFid) {
		_Flood, _error := model.GetFloodViaFid(iFid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodViaEventHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEvent := self.Args("event").String()
	if helper.IsHas(iEvent) {
		_Flood, _error := model.GetFloodViaEvent(iEvent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodViaEvent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodViaIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iIdentifier := self.Args("identifier").String()
	if helper.IsHas(iIdentifier) {
		_Flood, _error := model.GetFloodViaIdentifier(iIdentifier)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodViaIdentifier's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_Flood, _error := model.GetFloodViaTimestamp(iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFloodViaExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpiration := self.Args("expiration").MustInt()
	if helper.IsHas(iExpiration) {
		_Flood, _error := model.GetFloodViaExpiration(iExpiration)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the GetFloodViaExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFloodViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	if helper.IsHas(Fid_) {
		var iFlood model.Flood
		self.Bind(&iFlood)
		_Flood, _error := model.SetFloodViaFid(Fid_, &iFlood)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the SetFloodViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFloodViaEventHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Event_ := self.Args("event").String()
	if helper.IsHas(Event_) {
		var iFlood model.Flood
		self.Bind(&iFlood)
		_Flood, _error := model.SetFloodViaEvent(Event_, &iFlood)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the SetFloodViaEvent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFloodViaIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Identifier_ := self.Args("identifier").String()
	if helper.IsHas(Identifier_) {
		var iFlood model.Flood
		self.Bind(&iFlood)
		_Flood, _error := model.SetFloodViaIdentifier(Identifier_, &iFlood)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the SetFloodViaIdentifier's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFloodViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	if helper.IsHas(Timestamp_) {
		var iFlood model.Flood
		self.Bind(&iFlood)
		_Flood, _error := model.SetFloodViaTimestamp(Timestamp_, &iFlood)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the SetFloodViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFloodViaExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expiration_ := self.Args("expiration").MustInt()
	if helper.IsHas(Expiration_) {
		var iFlood model.Flood
		self.Bind(&iFlood)
		_Flood, _error := model.SetFloodViaExpiration(Expiration_, &iFlood)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Flood)
	}
	herr.Message = "Can't get to the SetFloodViaExpiration's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddFloodHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	Event_ := self.Args("event").String()
	Identifier_ := self.Args("identifier").String()
	Timestamp_ := self.Args("timestamp").MustInt()
	Expiration_ := self.Args("expiration").MustInt()

	if helper.IsHas(Fid_) {
		_error := model.AddFlood(Fid_,Event_,Identifier_,Timestamp_,Expiration_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddFlood's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostFloodHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iFlood model.Flood
	self.Bind(&iFlood)
	_int64, _error := model.PostFlood(&iFlood)
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

func PutFloodHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iFlood model.Flood
	self.Bind(&iFlood)
	_int64, _error := model.PutFlood(&iFlood)
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

func PutFloodViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	var iFlood model.Flood
	self.Bind(&iFlood)
	_int64, _error := model.PutFloodViaFid(Fid_, &iFlood)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFloodViaEventHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Event_ := self.Args("event").String()
	var iFlood model.Flood
	self.Bind(&iFlood)
	_int64, _error := model.PutFloodViaEvent(Event_, &iFlood)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFloodViaIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Identifier_ := self.Args("identifier").String()
	var iFlood model.Flood
	self.Bind(&iFlood)
	_int64, _error := model.PutFloodViaIdentifier(Identifier_, &iFlood)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFloodViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iFlood model.Flood
	self.Bind(&iFlood)
	_int64, _error := model.PutFloodViaTimestamp(Timestamp_, &iFlood)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFloodViaExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expiration_ := self.Args("expiration").MustInt()
	var iFlood model.Flood
	self.Bind(&iFlood)
	_int64, _error := model.PutFloodViaExpiration(Expiration_, &iFlood)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFloodViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	var iFlood model.Flood
	self.Bind(&iFlood)
	var iMap = helper.StructToMap(iFlood)
	_error := model.UpdateFloodViaFid(Fid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFloodViaEventHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Event_ := self.Args("event").String()
	var iFlood model.Flood
	self.Bind(&iFlood)
	var iMap = helper.StructToMap(iFlood)
	_error := model.UpdateFloodViaEvent(Event_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFloodViaIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Identifier_ := self.Args("identifier").String()
	var iFlood model.Flood
	self.Bind(&iFlood)
	var iMap = helper.StructToMap(iFlood)
	_error := model.UpdateFloodViaIdentifier(Identifier_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFloodViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iFlood model.Flood
	self.Bind(&iFlood)
	var iMap = helper.StructToMap(iFlood)
	_error := model.UpdateFloodViaTimestamp(Timestamp_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFloodViaExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expiration_ := self.Args("expiration").MustInt()
	var iFlood model.Flood
	self.Bind(&iFlood)
	var iMap = helper.StructToMap(iFlood)
	_error := model.UpdateFloodViaExpiration(Expiration_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFloodHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	_int64, _error := model.DeleteFlood(Fid_)
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

func DeleteFloodViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	_error := model.DeleteFloodViaFid(Fid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFloodViaEventHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Event_ := self.Args("event").String()
	_error := model.DeleteFloodViaEvent(Event_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFloodViaIdentifierHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Identifier_ := self.Args("identifier").String()
	_error := model.DeleteFloodViaIdentifier(Identifier_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFloodViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	_error := model.DeleteFloodViaTimestamp(Timestamp_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFloodViaExpirationHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expiration_ := self.Args("expiration").MustInt()
	_error := model.DeleteFloodViaExpiration(Expiration_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
