package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetShortcutSetUsersesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetShortcutSetUsersesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["shortcut_set_userssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetShortcutSetUsersesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutSetUsersCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt64()
	_int64 := model.GetShortcutSetUsersCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_set_usersCount"] = 0
	}
	m["shortcut_set_usersCount"] = _int64
	return self.JSON(m)
}

func GetShortcutSetUsersCountViaSetNameHandler(self *macross.Context) error {
	SetName_ := self.Args("set_name").String()
	_int64 := model.GetShortcutSetUsersCountViaSetName(SetName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["shortcut_set_usersCount"] = 0
	}
	m["shortcut_set_usersCount"] = _int64
	return self.JSON(m)
}

func GetShortcutSetUsersesViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt64()
	if (offset > 0) && helper.IsHas(iUid) {
		_ShortcutSetUsers, _error := model.GetShortcutSetUsersesViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutSetUsers)
	}
	herr.Message = "Can't get to the GetShortcutSetUsersesViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutSetUsersesViaSetNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSetName := self.Args("set_name").String()
	if (offset > 0) && helper.IsHas(iSetName) {
		_ShortcutSetUsers, _error := model.GetShortcutSetUsersesViaSetName(offset, limit, iSetName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutSetUsers)
	}
	herr.Message = "Can't get to the GetShortcutSetUsersesViaSetName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutSetUsersesByUidAndSetNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iSetName := self.Args("set_name").String()

	if helper.IsHas(iUid) {
		_ShortcutSetUsers, _error := model.GetShortcutSetUsersesByUidAndSetName(offset, limit, iUid,iSetName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutSetUsers)
	}
	herr.Message = "Can't get to the GetShortcutSetUsersesByUidAndSetName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutSetUsersesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_ShortcutSetUsers, _error := model.GetShortcutSetUserses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutSetUsers)
	}
	herr.Message = "Can't get to the GetShortcutSetUserses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutSetUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_ShortcutSetUsers := model.HasShortcutSetUsersViaUid(iUid)
		var m = map[string]interface{}{}
		m["shortcut_set_users"] = _ShortcutSetUsers
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutSetUsersViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasShortcutSetUsersViaSetNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSetName := self.Args("set_name").String()
	if helper.IsHas(iSetName) {
		_ShortcutSetUsers := model.HasShortcutSetUsersViaSetName(iSetName)
		var m = map[string]interface{}{}
		m["shortcut_set_users"] = _ShortcutSetUsers
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasShortcutSetUsersViaSetName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutSetUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_ShortcutSetUsers, _error := model.GetShortcutSetUsersViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutSetUsers)
	}
	herr.Message = "Can't get to the GetShortcutSetUsersViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetShortcutSetUsersViaSetNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSetName := self.Args("set_name").String()
	if helper.IsHas(iSetName) {
		_ShortcutSetUsers, _error := model.GetShortcutSetUsersViaSetName(iSetName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutSetUsers)
	}
	herr.Message = "Can't get to the GetShortcutSetUsersViaSetName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutSetUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	if helper.IsHas(Uid_) {
		var iShortcutSetUsers model.ShortcutSetUsers
		self.Bind(&iShortcutSetUsers)
		_ShortcutSetUsers, _error := model.SetShortcutSetUsersViaUid(Uid_, &iShortcutSetUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutSetUsers)
	}
	herr.Message = "Can't get to the SetShortcutSetUsersViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetShortcutSetUsersViaSetNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SetName_ := self.Args("set_name").String()
	if helper.IsHas(SetName_) {
		var iShortcutSetUsers model.ShortcutSetUsers
		self.Bind(&iShortcutSetUsers)
		_ShortcutSetUsers, _error := model.SetShortcutSetUsersViaSetName(SetName_, &iShortcutSetUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_ShortcutSetUsers)
	}
	herr.Message = "Can't get to the SetShortcutSetUsersViaSetName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddShortcutSetUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	SetName_ := self.Args("set_name").String()

	if helper.IsHas(Uid_) {
		_error := model.AddShortcutSetUsers(Uid_,SetName_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddShortcutSetUsers's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostShortcutSetUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iShortcutSetUsers model.ShortcutSetUsers
	self.Bind(&iShortcutSetUsers)
	_int64, _error := model.PostShortcutSetUsers(&iShortcutSetUsers)
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

func PutShortcutSetUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iShortcutSetUsers model.ShortcutSetUsers
	self.Bind(&iShortcutSetUsers)
	_int64, _error := model.PutShortcutSetUsers(&iShortcutSetUsers)
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

func PutShortcutSetUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iShortcutSetUsers model.ShortcutSetUsers
	self.Bind(&iShortcutSetUsers)
	_int64, _error := model.PutShortcutSetUsersViaUid(Uid_, &iShortcutSetUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutShortcutSetUsersViaSetNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SetName_ := self.Args("set_name").String()
	var iShortcutSetUsers model.ShortcutSetUsers
	self.Bind(&iShortcutSetUsers)
	_int64, _error := model.PutShortcutSetUsersViaSetName(SetName_, &iShortcutSetUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutSetUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iShortcutSetUsers model.ShortcutSetUsers
	self.Bind(&iShortcutSetUsers)
	var iMap = helper.StructToMap(iShortcutSetUsers)
	_error := model.UpdateShortcutSetUsersViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateShortcutSetUsersViaSetNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SetName_ := self.Args("set_name").String()
	var iShortcutSetUsers model.ShortcutSetUsers
	self.Bind(&iShortcutSetUsers)
	var iMap = helper.StructToMap(iShortcutSetUsers)
	_error := model.UpdateShortcutSetUsersViaSetName(SetName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutSetUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_int64, _error := model.DeleteShortcutSetUsers(Uid_)
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

func DeleteShortcutSetUsersViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_error := model.DeleteShortcutSetUsersViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteShortcutSetUsersViaSetNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SetName_ := self.Args("set_name").String()
	_error := model.DeleteShortcutSetUsersViaSetName(SetName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
