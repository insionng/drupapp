package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetUsersesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetUsersesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["userssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetUsersesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt64()
	_int64 := model.GetUsersCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUuidHandler(self *macross.Context) error {
	Uuid_ := self.Args("uuid").String()
	_int64 := model.GetUsersCountViaUuid(Uuid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetUsersCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersesViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt64()
	if (offset > 0) && helper.IsHas(iUid) {
		_Users, _error := model.GetUsersesViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUuid := self.Args("uuid").String()
	if (offset > 0) && helper.IsHas(iUuid) {
		_Users, _error := model.GetUsersesViaUuid(offset, limit, iUuid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_Users, _error := model.GetUsersesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUidAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUid) {
		_Users, _error := model.GetUsersesByUidAndUuidAndLangcode(offset, limit, iUid,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUidAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iUid) {
		_Users, _error := model.GetUsersesByUidAndUuid(offset, limit, iUid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUid) {
		_Users, _error := model.GetUsersesByUidAndLangcode(offset, limit, iUid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUuid) {
		_Users, _error := model.GetUsersesByUuidAndLangcode(offset, limit, iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Users, _error := model.GetUserses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUserses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_Users := model.HasUsersViaUid(iUid)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_Users := model.HasUsersViaUuid(iUuid)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Users := model.HasUsersViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_Users, _error := model.GetUsersViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_Users, _error := model.GetUsersViaUuid(iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_Users, _error := model.GetUsersViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	if helper.IsHas(Uid_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUid(Uid_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	if helper.IsHas(Uuid_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUuid(Uuid_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaLangcode(Langcode_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	Uuid_ := self.Args("uuid").String()
	Langcode_ := self.Args("langcode").String()

	if helper.IsHas(Uid_) {
		_error := model.AddUsers(Uid_,Uuid_,Langcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddUsers's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PostUsers(&iUsers)
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

func PutUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsers(&iUsers)
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

func PutUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUid(Uid_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUuid(Uuid_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaLangcode(Langcode_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUuid(Uuid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_int64, _error := model.DeleteUsers(Uid_)
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

func DeleteUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_error := model.DeleteUsersViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	_error := model.DeleteUsersViaUuid(Uuid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteUsersViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
