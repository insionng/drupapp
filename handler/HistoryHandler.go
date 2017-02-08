package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetHistoriesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetHistoriesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["historysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetHistoriesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoryCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt64()
	_int64 := model.GetHistoryCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["historyCount"] = 0
	}
	m["historyCount"] = _int64
	return self.JSON(m)
}

func GetHistoryCountViaNidHandler(self *macross.Context) error {
	Nid_ := self.Args("nid").MustInt()
	_int64 := model.GetHistoryCountViaNid(Nid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["historyCount"] = 0
	}
	m["historyCount"] = _int64
	return self.JSON(m)
}

func GetHistoryCountViaTimestampHandler(self *macross.Context) error {
	Timestamp_ := self.Args("timestamp").MustInt()
	_int64 := model.GetHistoryCountViaTimestamp(Timestamp_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["historyCount"] = 0
	}
	m["historyCount"] = _int64
	return self.JSON(m)
}

func GetHistoriesViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt64()
	if (offset > 0) && helper.IsHas(iUid) {
		_History, _error := model.GetHistoriesViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoriesViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoriesViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iNid := self.Args("nid").MustInt()
	if (offset > 0) && helper.IsHas(iNid) {
		_History, _error := model.GetHistoriesViaNid(offset, limit, iNid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoriesViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoriesViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTimestamp := self.Args("timestamp").MustInt()
	if (offset > 0) && helper.IsHas(iTimestamp) {
		_History, _error := model.GetHistoriesViaTimestamp(offset, limit, iTimestamp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoriesViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoriesByUidAndNidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iNid := self.Args("nid").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_History, _error := model.GetHistoriesByUidAndNidAndTimestamp(offset, limit, iUid,iNid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoriesByUidAndNidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoriesByUidAndNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iNid := self.Args("nid").MustInt()

	if helper.IsHas(iUid) {
		_History, _error := model.GetHistoriesByUidAndNid(offset, limit, iUid,iNid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoriesByUidAndNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoriesByUidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iUid) {
		_History, _error := model.GetHistoriesByUidAndTimestamp(offset, limit, iUid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoriesByUidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoriesByNidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iNid) {
		_History, _error := model.GetHistoriesByNidAndTimestamp(offset, limit, iNid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoriesByNidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoriesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_History, _error := model.GetHistories(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistories' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasHistoryViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_History := model.HasHistoryViaUid(iUid)
		var m = map[string]interface{}{}
		m["history"] = _History
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasHistoryViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasHistoryViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt()
	if helper.IsHas(iNid) {
		_History := model.HasHistoryViaNid(iNid)
		var m = map[string]interface{}{}
		m["history"] = _History
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasHistoryViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasHistoryViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_History := model.HasHistoryViaTimestamp(iTimestamp)
		var m = map[string]interface{}{}
		m["history"] = _History
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasHistoryViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoryViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_History, _error := model.GetHistoryViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoryViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoryViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt()
	if helper.IsHas(iNid) {
		_History, _error := model.GetHistoryViaNid(iNid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoryViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHistoryViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_History, _error := model.GetHistoryViaTimestamp(iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the GetHistoryViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetHistoryViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	if helper.IsHas(Uid_) {
		var iHistory model.History
		self.Bind(&iHistory)
		_History, _error := model.SetHistoryViaUid(Uid_, &iHistory)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the SetHistoryViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetHistoryViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt()
	if helper.IsHas(Nid_) {
		var iHistory model.History
		self.Bind(&iHistory)
		_History, _error := model.SetHistoryViaNid(Nid_, &iHistory)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the SetHistoryViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetHistoryViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	if helper.IsHas(Timestamp_) {
		var iHistory model.History
		self.Bind(&iHistory)
		_History, _error := model.SetHistoryViaTimestamp(Timestamp_, &iHistory)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_History)
	}
	herr.Message = "Can't get to the SetHistoryViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddHistoryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	Nid_ := self.Args("nid").MustInt()
	Timestamp_ := self.Args("timestamp").MustInt()

	if helper.IsHas(Uid_) {
		_error := model.AddHistory(Uid_,Nid_,Timestamp_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddHistory's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostHistoryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iHistory model.History
	self.Bind(&iHistory)
	_int64, _error := model.PostHistory(&iHistory)
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

func PutHistoryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iHistory model.History
	self.Bind(&iHistory)
	_int64, _error := model.PutHistory(&iHistory)
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

func PutHistoryViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iHistory model.History
	self.Bind(&iHistory)
	_int64, _error := model.PutHistoryViaUid(Uid_, &iHistory)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutHistoryViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt()
	var iHistory model.History
	self.Bind(&iHistory)
	_int64, _error := model.PutHistoryViaNid(Nid_, &iHistory)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutHistoryViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iHistory model.History
	self.Bind(&iHistory)
	_int64, _error := model.PutHistoryViaTimestamp(Timestamp_, &iHistory)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateHistoryViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iHistory model.History
	self.Bind(&iHistory)
	var iMap = helper.StructToMap(iHistory)
	_error := model.UpdateHistoryViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateHistoryViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt()
	var iHistory model.History
	self.Bind(&iHistory)
	var iMap = helper.StructToMap(iHistory)
	_error := model.UpdateHistoryViaNid(Nid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateHistoryViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iHistory model.History
	self.Bind(&iHistory)
	var iMap = helper.StructToMap(iHistory)
	_error := model.UpdateHistoryViaTimestamp(Timestamp_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteHistoryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_int64, _error := model.DeleteHistory(Uid_)
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

func DeleteHistoryViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_error := model.DeleteHistoryViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteHistoryViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt()
	_error := model.DeleteHistoryViaNid(Nid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteHistoryViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	_error := model.DeleteHistoryViaTimestamp(Timestamp_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
